package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func main() {

	//open 11 prilozhenie
	opFlName11 := OpFile("Открыть 11 прил")

	csvData11 := getExcel11(opFlName11)
	//fmt.Printf("%s\t", csvData11)

	//open 5 prilozhenie
	opFlName5 := OpFile("Открыть 5 прил")

	csvData5 := getExcel5(opFlName5)
	//fmt.Printf("%s\t", csvData5)

	var csvData [][]string

	//combine two arrays from two files
	fmt.Printf("%s\t", "combine")
	for _, element11 := range csvData11 {
		for _, element5 := range csvData5 {
			if element11[0] == element5[0] {
				//fmt.Printf("%s\t", "found")
				appBool := true
				for _, itemTest := range csvData {
					if element11[0] == itemTest[0] {
						appBool = false
						//fmt.Println("Duplicate")
					}
				}
				if appBool {
					csvLine := []string{
						element11[0],
						"x",
						element11[1],
						"",
						"",
						"ms",
						"1",
						"0",
						"",
						"0",
						"0.5",
						"0:00:00",
						"0:00:10",
						"swingingdoor",
						"0",
						element5[1],
						"",
						"local",
						element11[2],
						"",
						"0",
						"0",
						"",
						"1",
						"1",
						"1",
						"1",
						"",
						"[\"common\"]",
						"0",
						"0",
						"0:00:00",
						"0:00:00",
						"deadband",
						"",
						"0",
						"0",
						"-1",
						"0",
						"0",
					}
					csvData = append(csvData, csvLine)
				}

			} /*else {
				appBool := true
				for _, itemTest := range csvData {
					if element11[0] == itemTest[0] {
						appBool = false
						//fmt.Println("Duplicate")
					}
				}
				if appBool {
					csvLine := []string{
						element11[0],
						"TEST",
						element11[1],
						"",
						"",
						"ms",
						"1",
						"0",
						"",
						"0",
						"0.5",
						"0:00:00",
						"0:00:10",
						"swingingdoor",
						"0",
						element5[1],
						"",
						"local",
						element11[2],
						"",
						"0",
						"0",
						"",
						"1",
						"1",
						"1",
						"1",
						"",
						"[\"common\"]",
						"0",
						"0",
						"0:00:00",
						"0:00:00",
						"deadband",
						"",
						"0",
						"0",
						"-1",
						"0",
						"0",
					}
					csvData = append(csvData, csvLine)
				}
			}*/
		}
	}
	//fmt.Printf("%s\t", csvData)
	//save csv result
	svFlName := SvFile()

	file2, err := os.Create(svFlName)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	csvHeaders := []string{
		"_Name",
		"_Select (x)",
		"_ValueType",
		"_Delete (x)",
		"_NewName",
		"Accuracy",
		"Archive",
		"CalcAggregates",
		"CalculationAlgorithmExpression",
		"Compression",
		"CompressionDeviation",
		"CompressionTimeDeadBand",
		"CompressionTimespan",
		"CompressionType",
		"Convers",
		"Description",
		"DictId",
		"DictSource",
		"EngUnits",
		"InstrumentTag",
		"Interpolation",
		"Output",
		"PointSource",
		"Reception",
		"SaveAddTs",
		"SaveQuality",
		"Scan",
		"ScanClass",
		"SecurityGroups",
		"SourceCompression",
		"SourceCompressionDeviation",
		"SourceCompressionTimeDeadBand",
		"SourceCompressionTimespan",
		"SourceCompressionType",
		"SourceTag",
		"Span",
		"SquareRoot",
		"TTL",
		"TotalCode",
		"Zero",
	}
	wr := csv.NewWriter(file2)
	wr.Comma = ';'
	wr.Write(csvHeaders)
	for _, csvLine := range csvData {
		wr.Write(csvLine)
	}
	wr.Flush()

}

func getExcel11(string2 string) [][]string {

	//f, err := excelize.OpenFile(string2)
	reportBytes, _ := os.ReadFile(string2)
	reader := bytes.NewReader(reportBytes)
	f, err := excelize.OpenReader(reader)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	var csvData [][]string
	sheetName := f.GetSheetList()[0]
	rosws, _ := f.GetSheetDimension(sheetName)
	fmt.Sprintf("%d", len(rosws))

	/*
		for _, row := range rosws {

			csvLine := []string{row[0], row[12], row[13]}
			csvData = append(csvData, csvLine)
			fmt.Print(row[12], "\n")

			fmt.Println()
		}*/
	rre, _ := f.GetRows(sheetName)
	fmt.Println(len(rre))
	totalNumberOfRows := len(rre)
	for i := 3; i < totalNumberOfRows; i++ {
		cellNameName := fmt.Sprintf("%s%d", "A", i)
		cellValueName, _ := f.GetCellValue(sheetName, cellNameName)

		cellNameType := fmt.Sprintf("%s%d", "M", i)
		cellValueType, _ := f.GetCellValue(sheetName, cellNameType)

		cellNameUnits := fmt.Sprintf("%s%d", "N", i)
		cellValueUnits, _ := f.GetCellValue(sheetName, cellNameUnits)

		csvLine := []string{cellValueName, cellValueType, cellValueUnits}
		csvData = append(csvData, csvLine)
		//fmt.Printf("%s\t", csvLine)
	}

	/*rows, err := f.Rows(f.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}

	rows.Next()
	for rows.Next() {
		row, err1 := rows.Columns()

		if err1 != nil {
			log.Fatal(err)
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Printf(row[20]) // Print values in columns B and D
		return

	}*/

	/*cell, err1 := f.GetCellValue(f.GetSheetList()[0], "B2")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(cell)*/
	f.Close()

	return csvData
}

func getExcel5(string2 string) [][]string {

	reportBytes, _ := os.ReadFile(string2)
	reader := bytes.NewReader(reportBytes)
	f, err := excelize.OpenReader(reader)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := f.GetSheetList()[0]
	var csvData [][]string
	/*rosws, _ := f.GetRows(sheetName)
	fmt.Sprintf("%d", len(rosws))
	//totalNumberOfRows := 35

	for _, row := range rosws {

		csvLine := []string{row[21], row[22]}
		csvData = append(csvData, csvLine)
		fmt.Print(row[12], "\n")

		fmt.Println()
	}
	*/
	rre, _ := f.GetRows(sheetName)
	fmt.Println(len(rre))
	totalNumberOfRows := len(rre)

	for i := 3; i < totalNumberOfRows; i++ {
		cellNameName := fmt.Sprintf("%s%d", "W", i)
		cellValueName, _ := f.GetCellValue(sheetName, cellNameName)

		cellNameDesc := fmt.Sprintf("%s%d", "X", i)
		cellValueDesc, _ := f.GetCellValue(sheetName, cellNameDesc)

		csvLine := []string{cellValueName, cellValueDesc}
		csvData = append(csvData, csvLine)
	}

	f.Close()

	return csvData
}

func OpFile(wTitle string) string {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: wTitle,
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Excel Files",
				Pattern:     "*.xls;*.xlsx;*.xlsm",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 0,
		FileName:                "file.xlsx",
		DefaultExtension:        "xlsx",
	})
	if errors.Is(err, cfd.ErrorCancelled) {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file: %s\n", result)
	return result
}

func SvFile() string {
	fCsV, err2 := cfdutil.ShowSaveFileDialog(cfd.DialogConfig{
		Title: "Save A File",
		Role:  "SaveFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.csv)",
				Pattern:     "*.csv",
			},
		},
		SelectedFileFilterIndex: 1,
		FileName:                "export.csv",
		DefaultExtension:        "csv",
	})
	if errors.Is(err2, cfd.ErrorCancelled) {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("Chosen file: %s\n", fCsV)
	return fCsV
}
