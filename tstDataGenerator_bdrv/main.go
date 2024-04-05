package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"github.com/tadvi/winc"
	"github.com/xuri/excelize/v2"
	"log"
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

	var opFlName11 string

	var tagIds []string
	var items []Item
	log.Println(tagIds)

	mainWindow := winc.NewForm(nil)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	//none := winc.Shortcut{}

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	openMn := fileMn.AddItem("Open", winc.Shortcut{winc.ModControl, winc.KeyO})

	menu.Show()
	/*
		openMn.OnClick().Bind(func(e *winc.Event) {
			dlg := winc.NewDialog(mainWindow)
			dlg.Center()
			dlg.Show()
			dlg.OnClose().Bind(func(arg *winc.Event) {
				dlg.Close()
			})
		})*/
	//tipRun := winc.NewToolTip(mainWindow)
	//tipRun.AddTool(btnRun, "Run project")

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
	ls.SetPos(10, 180)

	openMn.OnClick().Bind(func(arg *winc.Event) {
		/*if filePath, ok := winc.ShowOpenFileDlg(mainWindow,
			"Select EDI X12 file", "All files (*.*)|*.*", 0, ""); ok {

			if err := imgv.DrawImageFile(filePath); err != nil {
				winc.Errorf(mainWindow, "Error: %s", err)
			}
		}*/
		opFlName11 = OpFile("Открыть 11 прил")
		tagIds = getExcel11(opFlName11)

		for _, s := range tagIds {
			items = append(items, Item{[]string{s}, false})

		}
		for _, s := range items {
			ls.AddItem(&s)
		}
	})

	btnEdit := winc.NewPushButton(mainWindow)
	btnEdit.SetText(" Edit")
	btnEdit.SetPos(0, 0)
	btnEdit.SetSize(98, 38)
	btnEdit.OnClick().Bind(func(arg *winc.Event) {
		ls.DeleteAllItems()
		for i, s := range items {
			fmt.Println(i, s)
			s.checked = true
			ls.AddItem(&s)
		}
	})

	left := winc.NewMultiEdit(mainWindow)
	left.SetPos(5, 5)
	left.SetSize(300, 38)
	split := winc.NewVResizer(mainWindow)

	mainWindow.Center()
	mainWindow.Show()

	dock := winc.NewSimpleDock(mainWindow)
	//mainWindow.SetLayout(dock)
	dock.Dock(btnEdit, winc.Top)
	dock.Dock(left, winc.Left)
	dock.Dock(split, winc.Left)
	dock.Dock(ls, winc.Left)

	// if err := dock.LoadStateFile("layout.json"); err != nil {
	// 	log.Println(err)
	// }

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
func getExcel11(string2 string) []string {

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
	var csvData []string
	sheetName := f.GetSheetList()[0]
	rosws, _ := f.GetSheetDimension(sheetName)
	fmt.Sprintf("%d", len(rosws))

	rre, _ := f.GetRows(sheetName)
	fmt.Println(len(rre))
	totalNumberOfRows := len(rre)
	for i := 3; i < totalNumberOfRows; i++ {
		cellNameName := fmt.Sprintf("%s%d", "A", i)
		cellValueName, _ := f.GetCellValue(sheetName, cellNameName)

		csvData = append(csvData, cellValueName)
		//fmt.Printf("%s\t", csvLine)
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
