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
	opFlName := OpFile()

	getExcel(opFlName)

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
	wr.Write(csvHeaders)
	wr.Flush()
}

func getExcel(string2 string) {

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

	columnName := "A"
	sheetName := f.GetSheetList()[0]
	totalNumberOfRows := 2

	for i := 3; i < totalNumberOfRows; i++ {
		cellName := fmt.Sprintf("%s%d", columnName, i)
		// fmt.Println(cellName)
		cellValue, _ := f.GetCellValue(sheetName, cellName)
		fmt.Printf("%s\t", cellValue)
	}

	rows, err := f.Rows(f.GetSheetList()[0])
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

	}

	/*cell, err1 := f.GetCellValue(f.GetSheetList()[0], "B2")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(cell)*/
	f.Close()

	return
}

func OpFile() string {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
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
