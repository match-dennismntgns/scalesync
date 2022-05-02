package main

import (
	"fmt"
	"gopkg.in/Iwark/spreadsheet.v2"
	"os"
)

func uploadtosheets(row int, weight string, date string) {
	service, err := spreadsheet.NewService()
	if err != nil {
		panic(err.Error())
	}

	spreadsheetID := os.Getenv("spreadsheetID")

	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	checkError(err)
	sheet, err := spreadsheet.SheetByIndex(0)
	checkError(err)
	for _, row := range sheet.Rows {
		for _, cell := range row {
			fmt.Println(cell.Value)
		}
	}

	columnWeight := 1
	columnDate := 2
	sheet.Update(row, columnWeight, weight)
	sheet.Update(row, columnDate, date)

	// Make sure call Synchronize to reflect the changes
	err = sheet.Synchronize()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
