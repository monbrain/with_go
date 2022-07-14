// cloud 테스트
package main

import (
	"fmt"
	"strings"

	// "encoding/json"
	// "encoding/csv"

	"github.com/go-gota/gota/dataframe"

	"with_go/cloud"
)

func res2csv(res [][]interface{}) string {
	csvStr := ""
	for _, row := range res {
		for _, cell := range row {
			// csvStr += `"` + cell.(string) + `"` + ","
			str := cell.(string)
			if !strings.Contains(str, " ") {
				csvStr += str + ","
			} else {
				csvStr += `"` + str + `"` + ","
			}
		}
		csvStr = csvStr[:len(csvStr)-1] + "\n"
		// csvStr += "\n"
	}
	return csvStr
}

func main() {
	srv := cloud.SrvSheets("sheets", "moonsats", "")
	spreadsheetId := "1TofihRh87iKRsNOCrjQusCgStPkIDx-wOqZa1xee-SI"
	// readRange := "sheet2!A1:C17"
	readRange := "candles"
	res := cloud.ReadSheet(srv, spreadsheetId, readRange)

	if res == nil {
		fmt.Print("Error!")
	}
	// fmt.Printf("%v", res)

	res2csv(res)

	// res -> csv
	csvStr := res2csv(res)

	// df := dataframe.LoadMaps(res)
	// csvStr -> dataframe
	df := dataframe.ReadCSV(strings.NewReader(string(csvStr)))
	fmt.Println(df)
}
