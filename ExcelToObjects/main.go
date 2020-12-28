package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type DBValues struct {
	Npi        string `json:"npi"`
	ZipCode    string `json:"zipcode"`
	DivPanelId string `json:"divpanelid"`
}

func main() {
	fmt.Println("Demo")
	xlsx, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// // Get value from cell by given worksheet name and axis.
	// cell, _ := xlsx.GetCellValue("Sheet1", "B2")
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, _ := xlsx.GetRows("Sheet1")
	var dbvalues []DBValues
	for j, row := range rows {
		var npi, divpanelid, zipcode string
		if j > 0 {
			for i, colCell := range row {
				if i == 0 {
					divpanelid = colCell
				} else if i == 1 {
					zipcode = colCell
				} else if i == 2 {
					npi = colCell
				}
				// fmt.Print(i, "colCell  ", colCell, "\t")
			}
			// dbvalues.append(dbvalues,)
			// fmt.Println(npi, divpanelid, zipcode)
			dbValue := DBValues{Npi: npi, DivPanelId: divpanelid, ZipCode: zipcode}
			// fmt.Println(dbValue)
			dbvalues = append(dbvalues, dbValue)
			fmt.Println()
		}

	}
	fmt.Println(dbvalues)
}
