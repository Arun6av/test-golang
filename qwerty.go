package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func scrapeTables(url string) [][][]string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error loading document: %s", err)
	}

	var tableData [][][]string
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		var tableDataRow [][]string
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var rowData []string
			row.Find("td").Each(func(k int, cell *goquery.Selection) {
				rowData = append(rowData, cell.Text())
			})
			tableDataRow = append(tableDataRow, rowData)
		})
		tableData = append(tableData, tableDataRow)
	})

	return tableData
}

func main() {
	baseURL := "https://www.screener.in/screens/874888/dwm/?page="
	file, err := os.Create("qwerty.xlsx")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 156; i++ {
		url := baseURL + strconv.Itoa(i)
		data := scrapeTables(url)

		for _, tables := range data {
			for _, row := range tables {
				if err := writer.Write(row); err != nil {
					log.Fatalf("Error writing row: %s", err)
				}
			}
		}
		fmt.Printf("Page %d done\n", i)
	}
	fmt.Println("Data saved to qwerty.xlsx")
}
