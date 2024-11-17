package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/moevm/nosql2h24-cleaning/cleaning/docs"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/config"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/auth"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	v1orders "github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/v1/orders"
	v1services "github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/v1/services"
	v1users "github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/v1/users"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	mongorepo "github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository/mongo"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services/jwt"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/hasher"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httpserver"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/mongo"
	_ "github.com/moevm/nosql2h24-cleaning/cleaning/pkg/validate"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const Version = "v0.5.0"

func Run(cfg *config.Config) {
	docs.SwaggerInfo.Version = Version
	// init logger
	logger := zap.Must(zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel)))

	// connect to database
	client, err := mongo.Connect(
		cfg.MongoDB.User,
		cfg.MongoDB.Password,
		cfg.MongoDB.Hostname,
		cfg.MongoDB.Port,
	)
	if err != nil {
		logger.Panic("mongo connection failed", zap.Any("err", err))
	}

	// create repositories
	db := client.Database(cfg.MongoDB.DB)
	userRepo := mongorepo.NewUserRepo(db)
	serviceRepo := mongorepo.NewServiceRepo(db)
	orderRepo := mongorepo.NewOrderRepo(db)

	// load dump
	// bad practice, but it's just a prototype
	// my last brain cell is dying
	if file, err := os.OpenFile("/dump/dump.json", os.O_RDONLY, 0666); err == nil {
		defer file.Close()
		logger.Info("app - Run - loading dump")

		data, _ := io.ReadAll(file)

		var dump models.Dump
		if err := json.Unmarshal(data, &dump); err != nil {
			logger.Error("app - Run - unmarshal dump failed", zap.Any("error", err))
		}
		logger.Info("dump", zap.Any("dump.Users", dump.Users))
		logger.Info("dump", zap.Any("dump.Orders", dump.Orders))
		logger.Info("dump", zap.Any("dump.Services", dump.Services))
		// insert dump
		eg := errgroup.Group{}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		eg.Go(func() error {
			return userRepo.InsertMany(ctx, dump.Users)
		})
		eg.Go(func() error {
			return serviceRepo.InsertMany(ctx, dump.Services)
		})
		eg.Go(func() error {
			return orderRepo.InsertMany(ctx, dump.Orders)
		})

		if err := eg.Wait(); err != nil {
			logger.Error("app - Run - loading dump failed", zap.Any("error", err))
		}
		cancel()
	}

	// create services
	jwt := jwt.New(cfg.Auth.Secret)
	hasher := hasher.New(10)
	authService := services.NewAuthService(
		logger,
		jwt,
		hasher,
		userRepo,
	)
	userService := services.NewUserService(
		logger,
		hasher,
		userRepo,
	)
	addressService := services.NewAddressService(
		logger,
		userRepo,
	)
	orderService := services.NewOrderService(
		logger,
		orderRepo,
	)
	service := services.NewService(
		logger,
		serviceRepo,
	)
	// setup router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(r chi.Router) {
		r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
		})
		r.Get("/swagger*", httpSwagger.Handler())

		// route auth
		r.Mount("/auth", auth.New(authService).Routes())

		// route v1
		r.Route("/v1", func(r chi.Router) {
			r.Use(middlewares.NewAuthMiddleware(jwt).JWT)
			// TODO: Add v1 routes
			r.Mount("/users", v1users.NewHandler(userService, addressService).Routes())
			r.Mount("/orders", v1orders.NewHandler(orderService).Routes())
			r.Mount("/services", v1services.NewHandler(service).Routes())
			// hardcoded to test
			r.Get("/secret", func(w http.ResponseWriter, r *http.Request) {
				id := r.Context().Value(middlewares.UserID).(string)
				w.Write([]byte(fmt.Sprintf("Hello %s", id)))
			})

		})
	})

	// run http-server
	server := httpserver.New(
		router,
		httpserver.Port(cfg.HTTP.Port),
		httpserver.WriteTimeout(cfg.HTTP.Timeout),
	)

	// wait signal to shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sig:
		logger.Info("app - Run - get signal: " + s.String())
	case err := <-server.Notify():
		logger.Error(fmt.Sprintf("app - Run - server.Notify: %s", err.Error()))
	}

	// shutdown
	if err := server.Shutdown(); err != nil {
		logger.Error(fmt.Sprintf("app - Run - get signal: %s", err.Error()))
	}
}
