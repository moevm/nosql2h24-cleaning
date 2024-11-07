package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moevm/nosql2h24-cleaning/cleaning/docs"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/config"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/auth"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/controllers/http/middlewares"
	mongorepo "github.com/moevm/nosql2h24-cleaning/cleaning/internal/repository/mongo"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/services/jwt"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/hasher"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/httpserver"
	"github.com/moevm/nosql2h24-cleaning/cleaning/pkg/mongo"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	docs.SwaggerInfo.Version = "v0.5.0"
	// init logger
	logger := zap.Must(zap.NewDevelopment())

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

	// create services
	jwt := jwt.New(cfg.Auth.Secret)
	authService := services.NewAuthService(
		logger,
		jwt,
		hasher.New(10),
		userRepo,
	)

	// setup router
	router := chi.NewRouter()
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
