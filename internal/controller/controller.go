package controller

import (
	"context"
	"go-pay-reminder/internal/config"
	"go-pay-reminder/internal/sheets"
	"go-pay-reminder/internal/telegram"
	"net/http"
	"time"
)

type controller struct {
	sheetsClient *http.Client
	config       *config.Config
	rows         []*sheets.Row
}

type Controller interface {
	GetSheetsData(ctx context.Context)
	SendReminder(ctx context.Context)
}

func New(client *http.Client, cfg *config.Config) (Controller, error) {
	return &controller{
		sheetsClient: client,
		config:       cfg,
	}, nil
}

func (c *controller) GetSheetsData(ctx context.Context) {
	rows := sheets.Get(ctx, c.sheetsClient, c.config.SheetID)

	c.rows = []*sheets.Row{}

	for _, row := range rows {
		if row == nil {
			continue
		}

		if row.PayDate == nil {
			continue
		}

		days := int(time.Until(*row.PayDate).Hours() / 24)
		switch {
		case days == 12, days == 7, days == 5, days == 3, days <= 0:
			row.DaysLeft = days
			c.rows = append(c.rows, row)
		}
	}
}

func (c *controller) SendReminder(ctx context.Context) {
	for _, row := range c.rows {
		telegram.SendToBot(c.config.TGToken, c.config.ChatID, row)
	}
}
