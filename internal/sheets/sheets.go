package sheets

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"go-pay-reminder/internal/logger"
)

type Row struct {
	Name     string
	Url      string
	IP       string
	LLC      string
	Login    string
	PayDate  *time.Time
	PayValue string
	DaysLeft int
}

const (
	sheetRange = "%s!A1:H"
	dateLayout = "2006-01-02" // YYYY-MM-DD
)

func Get(ctx context.Context, c *http.Client, sheetID string) []*Row {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(c))
	if err != nil {
		logger.Error("Unable to retrieve Sheets client", "error", err)
		return nil
	}

	spreadsheet, err := srv.Spreadsheets.Get(sheetID).Do()
	if err != nil {
		logger.Error("Unable to retrieve spreadsheet data", "error", err)
		return nil
	}

	logger.Info("Retrieved spreadsheet data successfully")

	sheet := spreadsheet.Sheets[0]
	resp, err := srv.Spreadsheets.Values.Get(sheetID, fmt.Sprintf(sheetRange, sheet.Properties.Title)).Do()
	if err != nil {
		logger.Error("Unable to retrieve sheet values", "error", err)
		return nil
	}

	logger.Info("Retrieved sheet values successfully")

	if len(resp.Values) < 2 {
		logger.Info("No data found in the sheet")
		return nil
	}

	rows := make([]*Row, len(resp.Values))
	for _, value := range resp.Values[1:] {
		if len(value) < 8 {
			log.Printf("Skipping a row due to not enough columns\n")
			continue
		}

		var payDate *time.Time

		if value[6] != nil && value[6] != "" {
			parsedDate, e := time.Parse(dateLayout, value[6].(string))
			if e != nil {
				logger.Error("Unable to parse date", "error", err, "date", value[6].(string))
			} else {
				payDate = &parsedDate
			}
		}

		row := &Row{
			Name:     value[0].(string),
			Url:      value[1].(string),
			IP:       value[2].(string),
			LLC:      value[3].(string),
			Login:    value[4].(string),
			PayDate:  payDate,
			PayValue: value[7].(string),
		}

		rows = append(rows, row)
	}

	return rows
}
