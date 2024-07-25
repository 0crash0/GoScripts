package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tadvi/winc"
	"github.com/xuri/excelize/v2"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"time"
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

	//var opFlName11 string

	var tagIds [][]string
	var items []Item
	log.Println(tagIds)

	mainWindow := winc.NewForm(nil)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	//none := winc.Shortcut{}

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	openMn := fileMn.AddItem("Open", winc.Shortcut{winc.ModControl, winc.KeyO})
	saveMn := fileMn.AddItem("Save", winc.Shortcut{winc.ModControl, winc.KeyS})
	menu.Show()

	edtAthUrl := winc.NewEdit(mainWindow)
	edtAthUrl.SetPos(10, 20)
	edtAthUrl.SetSize(200, 20)
	edtAthUrl.SetText("Enter here Oauth2 url")

	edtAthUsr := winc.NewEdit(mainWindow)
	edtAthUsr.SetPos(10, 20)
	edtAthUsr.SetSize(200, 20)
	edtAthUsr.SetText("Enter here Oauth2 client user")

	edtAthPwd := winc.NewEdit(mainWindow)
	edtAthPwd.SetPos(10, 20)
	edtAthPwd.SetSize(200, 20)
	edtAthPwd.SetText("Enter here Oauth2 client password")

	btnAuth := winc.NewPushButton(mainWindow)
	btnAuth.SetText("Auth")
	btnAuth.SetPos(0, 0)
	btnAuth.SetSize(98, 38)
	btnAuth.OnClick().Bind(func(arg *winc.Event) {
		ctx := context.Background()

		conf := &oauth2.Config{
			ClientID:     "YOUR_CLIENT_ID",
			ClientSecret: "YOUR_CLIENT_SECRET",
			Scopes:       []string{"SCOPE1", "SCOPE2"},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://provider.com/o/oauth2/token",
				AuthURL:  "https://provider.com/o/oauth2/auth",
			},
		}

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		// Use the authorization code that is pushed to the redirect
		// URL. Exchange will do the handshake to retrieve the
		// initial access token. The HTTP Client returned by
		// conf.Client will refresh the token as necessary.
		var code string
		if _, err := fmt.Scan(&code); err != nil {
			log.Fatal(err)
		}

		// Use the custom HTTP client when requesting a token.
		httpClient := &http.Client{Timeout: 2 * time.Second}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

		tok, err := conf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}

		client := conf.Client(ctx, tok)
		_ = client
	})

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
	edt := winc.NewEdit(mainWindow)
	edt.SetPos(10, 20)
	edt.SetSize(200, 20)
	edt.SetText("edit text")
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
	dock.Dock(edtAthUrl, winc.Top)
	dock.Dock(edtAthUsr, winc.Top)
	dock.Dock(edtAthPwd, winc.Top)
	dock.Dock(btnAuth, winc.Top)
	dock.Dock(btnDelAll, winc.Top)
	dock.Dock(btnSelAll, winc.Top)
	dock.Dock(ls, winc.Left)
	dock.Dock(split, winc.Left)
	dock.Dock(edt, winc.Left)
	dock.Dock(chk, winc.Left)

	// if err := dock.LoadStateFile("layout.json"); err != nil {
	// 	log.Println(err)
	// }
	openMn.OnClick().Bind(func(arg *winc.Event) {
		if filePath, ok := winc.ShowOpenFileDlg(mainWindow,
			"Открыть 11 прил", "Excel files (*.xls;*.xlsx;*.xlsm)|*.xls;*.xlsx;*.xlsm|All files (*.*)|*.*", 0, ""); ok {

			//opFlName11 = OpFile("Открыть 11 прил")
			tagIds = getExcel11(filePath, chk.Checked())

			for _, s := range tagIds {
				items = append(items, Item{[]string{s[0]}, false})

			}
			for _, s := range items {
				ls.AddItem(&s)
			}
		}

	})
	saveMn.OnClick().Bind(func(arg *winc.Event) {
		/*if filePath, ok := winc.ShowSaveFileDlg(mainWindow,
			"Открыть 11 прил", "Excel files (*.xls;*.xlsx;*.xlsm)|*.xls;*.xlsx;*.xlsm|All files (*.*)|*.*", 0, ""); ok {

			//opFlName11 = OpFile("Открыть 11 прил")

		}*/

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
