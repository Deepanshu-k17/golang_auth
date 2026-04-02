package app

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config      config.Config
	MongoCLient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	mongoCLi, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &App{
		Config:      cfg,
		MongoCLient: mongoCLi.Client,
		DB:          mongoCLi.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoCLient == nil {
		return nil
	}
	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := a.MongoCLient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect MongoDB client: %w", err)
	}
	return nil
}
