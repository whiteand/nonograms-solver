package boardfromsite

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse(input io.Reader) (rowBlocks, colBlocks [][]int, err error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, nil, fmt.Errorf("read error: %v", err)
	}

	rowBlocks, err = parseRows(bytes.NewReader(data))
	if err != nil {
		return nil, nil, fmt.Errorf("row parse error: %v", err)
	}
	colBlocks, err = parseCols(bytes.NewReader(data))
	if err != nil {
		return nil, nil, fmt.Errorf("col parse error: %v", err)
	}
	return rowBlocks, colBlocks, nil
}
func ParseFromFile(fileName string) (rowBlocks, colBlocks [][]int, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	return Parse(file)

}
func parseRows(r io.Reader) (rowBlocks [][]int, err error) {

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	var parseErr error = nil

	doc.Find("#nonogram_table td.nmtl table tbody tr").Each(func(i int, tr *goquery.Selection) {
		currentRow := make([]int, 0)
		tr.Find("td div").Each(func(k int, div *goquery.Selection) {
			html, err := div.Html()
			if err != nil {
				parseErr = err
				return
			}
			html = strings.TrimSpace(html)
			if html == "" {
				return
			}
			num, err := strconv.Atoi(strings.TrimSpace(html))
			if err != nil {
				trHtml, _ := tr.Html()
				fmt.Println(trHtml)
				parseErr = err
				return
			}
			currentRow = append(currentRow, num)
		})
		rowBlocks = append(rowBlocks, currentRow)
	})
	if parseErr != nil {
		return nil, parseErr
	}
	return rowBlocks, nil
}

func parseCols(r io.Reader) (colBlocks [][]int, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	textTable := make([][]string, 0)
	var parseErr error = nil
	doc.Find("#nonogram_table td.nmtt table tr").Each(func(i int, tr *goquery.Selection) {
		currentTextRow := make([]string, 0)
		tr.Find("td div").Each(func(j int, div *goquery.Selection) {
			numText, err := div.Html()
			if err != nil {
				parseErr = err
				return
			}
			currentTextRow = append(currentTextRow, strings.TrimSpace(numText))
		})
		textTable = append(textTable, currentTextRow)
	})
	if parseErr != nil {
		return nil, parseErr
	}

	transposedTable := make([][]string, len(textTable[0]))
	for rowInd := 0; rowInd < len(textTable[0]); rowInd++ {
		newRow := make([]string, len(textTable))
		for colInd := 0; colInd < len(textTable); colInd++ {
			newRow[colInd] = textTable[colInd][rowInd]
		}
		transposedTable[rowInd] = newRow
	}
	colBlocks = make([][]int, len(transposedTable))
	for i := range colBlocks {
		colBlocks[i] = make([]int, 0)
		for _, numText := range transposedTable[i] {
			if numText == "" {
				continue
			}
			num, err := strconv.Atoi(numText)
			if err != nil {
				return nil, fmt.Errorf("strconv error: %v", err)
			}
			colBlocks[i] = append(colBlocks[i], num)
		}
	}

	return colBlocks, nil
}
