package main

import (
	"context"

	"go-pay-reminder/internal/config"
	"go-pay-reminder/internal/controller"
	"go-pay-reminder/internal/logger"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Handler(ctx context.Context) (*Response, error) {
	cfg, err := config.MustLoad()
	if err != nil {
		return nil, err
	}

	logger.Init(cfg.LogLevel)

	client := cfg.Service.Client(ctx)

	ctrl, err := controller.New(client, cfg)
	if err != nil {
		return nil, err
	}

	ctrl.GetSheetsData(ctx)
	ctrl.SendReminder(ctx)

	return &Response{
		StatusCode: 200,
		Body:       "Done",
	}, nil
}
