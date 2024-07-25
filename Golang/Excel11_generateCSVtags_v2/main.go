package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/tadvi/winc"
	"github.com/xuri/excelize/v2"
	"os"
)

func btnOnClick(arg *winc.Event) {
	//edt.SetCaption("Got you !!!")
	fmt.Println("Button clicked")
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

type Item struct {
	T       []string
	checked bool
}

func (item Item) Text() []string    { return item.T }
func (item *Item) SetText(s string) { item.T[0] = s }

func (item Item) Checked() bool            { return item.checked }
func (item *Item) SetChecked(checked bool) { item.checked = checked }
func (item Item) ImageIndex() int          { return 0 }

func main() {
	csvHeaders := []string{
		"_Name",
		"_Select(x)",
		"_ValueType",
		"_Delete(x)",
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

	//var opFlName11 string

	var csvData11 [][]string
	var items []Item
	var csvData [][]string

	mainWindow := winc.NewForm(nil)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	//none := winc.Shortcut{}

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	openMn := fileMn.AddItem("Open", winc.Shortcut{winc.ModControl, winc.KeyO})
	saveMn := fileMn.AddItem("Save", winc.Shortcut{winc.ModControl, winc.KeyS})
	menu.Show()

	ls := winc.NewListView(mainWindow)
	ls.EnableEditLabels(true)
	ls.SetCheckBoxes(true)
	//ls.EnableFullRowSelect(true)
	//ls.EnableHotTrack(true)
	//ls.EnableSortHeader(true)
	//ls.EnableSortAscending(true)
	ls.OnEndLabelEdit().Bind(func(e *winc.Event) {
		println("edited", e.Data)
		// acccept label edit event!
		//d := e.Data.(*winc.LabelEditEventData)
		//d.Item.SetText(d.Text)
		//fmt.Println(d.Item.Text())
	})
	ls.AddColumn("tag name", 120)
	ls.SetPos(0, 0)
	ls.SetSize(300, 100)

	ls.EnableSortHeader(true)

	btnDelAll := winc.NewPushButton(mainWindow)
	btnDelAll.SetText("Delete Selected")
	btnDelAll.SetPos(0, 0)
	btnDelAll.SetSize(98, 38)
	btnDelAll.OnClick().Bind(func(arg *winc.Event) {
		fmt.Println()
		ls.DeleteAllItems()

	})

	chk := winc.NewCheckBox(mainWindow)
	chk.SetText("с фильтром")

	btnSelAll := winc.NewPushButton(mainWindow)
	btnSelAll.SetText("Select All")
	btnSelAll.SetPos(0, 0)
	btnSelAll.SetSize(98, 38)
	btnSelAll.OnClick().Bind(func(arg *winc.Event) {
		ls.DeleteAllItems()
		for _, s := range items {
			//fmt.Println(i, s)
			s.checked = true
			ls.AddItem(&s)
		}
	})

	split := winc.NewVResizer(mainWindow)
	mainWindow.Center()
	mainWindow.Show()

	dock := winc.NewSimpleDock(mainWindow)
	//mainWindow.SetLayout(dock)
	dock.Dock(btnDelAll, winc.Top)
	dock.Dock(btnSelAll, winc.Top)
	dock.Dock(ls, winc.Left)
	dock.Dock(split, winc.Left)
	dock.Dock(chk, winc.Left)

	// if err := dock.LoadStateFile("layout.json"); err != nil {
	// 	log.Println(err)
	// }
	openMn.OnClick().Bind(func(arg *winc.Event) {
		if filePath, ok := winc.ShowOpenFileDlg(mainWindow,
			"Открыть 11 прил", "Excel files (*.xls;*.xlsx;*.xlsm)|*.xls;*.xlsx;*.xlsm|All files (*.*)|*.*", 0, ""); ok {

			//opFlName11 = OpFile("Открыть 11 прил")
			csvData11 = getExcel11(filePath, chk.Checked())

			for _, s := range csvData11 {
				items = append(items, Item{[]string{s[0]}, false})

			}
			for _, s := range items {
				ls.AddItem(&s)
			}

		}

	})
	saveMn.OnClick().Bind(func(arg *winc.Event) {
		if filePathSv, ok := winc.ShowSaveFileDlg(mainWindow,
			"Открыть 11 прил", "CSV files (*.csv)|*.csv|All files (*.*)|*.*", 0, ""); ok {

			fmt.Printf("%s\t", "combine")
			for _, element11 := range csvData11 {
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
					element11[3],
					"",
					"local",
					element11[2],
					element11[0],
					"0",
					"0",
					"[\"csv-import\"]",
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
			file2, err := os.Create(filePathSv + ".csv")
			if err != nil {
				panic(err)
			}
			defer file2.Close()
			wr := csv.NewWriter(file2)
			wr.Comma = ';'
			wr.Write(csvHeaders)
			for _, csvLine := range csvData {
				wr.Write(csvLine)
			}
			wr.Flush()
		}

	})

	mainWindow.OnClose().Bind(func(e *winc.Event) {
		dock.SaveStateFile("layout.json") // error gets ignored
		winc.Exit()
	})

	dock.Update()
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
	// --- end of Dock and main window management

}
func getExcel11(string2 string, onlyVisible bool) [][]string {

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

		cellNameDscr := fmt.Sprintf("%s%d", "B", i)
		cellNameDscrs, _ := f.GetCellValue(sheetName, cellNameDscr)

		csvLine := []string{cellValueName, cellValueType, cellValueUnits, cellNameDscrs}

		//GET ONLY VISIBLE ROWS (FILTERED IN EXCEL)
		if onlyVisible == true {
			include, _ := f.GetRowVisible(sheetName, i)
			if include {
				csvData = append(csvData, csvLine)
				//fmt.Printf("%s\t", cellValueName)
				//fmt.Printf("%s\t", csvLine)
			}
		} else {
			csvData = append(csvData, csvLine)
		}

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
