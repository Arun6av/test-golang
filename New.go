package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeTables(url string) [][]string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error loading document: %s", err)
	}

	var tableData [][]string
	doc.Find("table").First().Find("tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			cellText := cell.Text()
			rowData = append(rowData, cellText)
		})
		tableData = append(tableData, rowData)
	})

	return tableData
}

func main() {
	url := "https://www.investing.com/equities/52-week-high?country=india"
	data := ScrapeTables(url)

	file, err := os.Create("New.xlsx")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			log.Fatalf("Error writing row: %s", err)
		}
	}
	fmt.Println("Data saved to New.xlsx")
}
