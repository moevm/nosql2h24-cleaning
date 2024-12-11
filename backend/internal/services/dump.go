package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/models"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type (
	Droper interface {
		Drop(context.Context) error
	}
	UserRepoDumper interface {
		ImportUsers(context.Context, []models.User) error
		ExportUsers(context.Context) ([]models.User, error)
		Droper
	}
	OrderRepoDumper interface {
		ImportOrders(context.Context, []models.Order) error
		ExportOrders(context.Context) ([]models.Order, error)
		Droper
	}
	ServiceRepoDumper interface {
		ImportServices(context.Context, []models.Service) error
		ExportServices(context.Context) ([]models.Service, error)
		Droper
	}
)

type DumpService struct {
	log     *zap.Logger
	user    UserRepoDumper
	order   OrderRepoDumper
	service ServiceRepoDumper
}

func NewDumpService(
	logger *zap.Logger,
	user UserRepoDumper,
	order OrderRepoDumper,
	service ServiceRepoDumper,
) *DumpService {
	return &DumpService{
		log:     logger,
		user:    user,
		order:   order,
		service: service,
	}
}

func (d *DumpService) Import(ctx context.Context, data []byte, force bool) error {
	l := d.log.With(
		zap.Any("operation", "DumpService.Import"),
	)

	var dump models.Dump
	if err := json.Unmarshal(data, &dump); err != nil {
		l.Error("unmarshal dump failed", zap.Any("error", err))
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	if force {
		if err := d.drop(ctx); err != nil {
			l.Error("dropping data failed", zap.Any("error", err))
			return err
		}
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		return d.user.ImportUsers(ctx, dump.Users)
	})

	eg.Go(func() error {
		return d.service.ImportServices(ctx, dump.Services)
	})

	eg.Go(func() error {
		return d.order.ImportOrders(ctx, dump.Orders)
	})

	if err := eg.Wait(); err != nil {
		l.Error("loading dump failed", zap.Any("error", err))
		return err
	}

	return nil
}

func (d *DumpService) Export(ctx context.Context) (*models.Dump, error) {
	l := d.log.With(
		zap.Any("operation", "DumpService.Export"),
	)

	dump := &models.Dump{}
	eg := errgroup.Group{}

	eg.Go(func() error {
		users, err := d.user.ExportUsers(ctx)
		if err != nil {
			l.Error("export users failed", zap.Any("error", err))
			return err
		}
		dump.Users = users
		return nil
	})

	eg.Go(func() error {
		services, err := d.service.ExportServices(ctx)
		if err != nil {
			l.Error("export services failed", zap.Any("error", err))
			return err
		}
		dump.Services = services
		return nil
	})

	eg.Go(func() error {
		orders, err := d.order.ExportOrders(ctx)
		if err != nil {
			l.Error("export orders failed", zap.Any("error", err))
			return err
		}
		dump.Orders = orders
		return nil
	})

	if err := eg.Wait(); err != nil {
		l.Error("export dump failed", zap.Any("error", err))
		return nil, err
	}

	return dump, nil
}

func (d *DumpService) drop(ctx context.Context) error {
	l := d.log.With(
		zap.Any("operation", "DumpService.drop"),
	)

	eg := errgroup.Group{}

	eg.Go(func() error {
		return d.user.Drop(ctx)
	})

	eg.Go(func() error {
		return d.service.Drop(ctx)
	})

	eg.Go(func() error {
		return d.order.Drop(ctx)
	})

	if err := eg.Wait(); err != nil {
		l.Error("dropping data failed", zap.Any("error", err))
		return err
	}

	return nil
}
