package sheets

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func createSheet(srv *sheets.Service, spreadsheetID, sheetTitle string) error {
    requests := []*sheets.Request{
        {
            AddSheet: &sheets.AddSheetRequest{
                Properties: &sheets.SheetProperties{
                    Title: sheetTitle,
                },
            },
        },
    }

    batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
        Requests: requests,
    }

    _, err := srv.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
    if err != nil {
        return fmt.Errorf("unable to create sheet: %v", err)
    }

    return nil
}