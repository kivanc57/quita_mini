package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)


func WriteExcel(freqMap map[string]int) {
	f := excelize.NewFile()

	style, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "left"},
	})

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Word")
	
	id := 2
	for word, freq := range freqMap {
		if freq == 1 {
			cellA := fmt.Sprintf("A%d", id)
			cellB := fmt.Sprintf("B%d", id)

			f.SetCellValue("Sheet1", cellA, id-1)
			f.SetCellValue("Sheet1", cellB, word)

			f.SetCellStyle("Sheet1", cellA, cellA, style)
			id++
		}
	}
	f.SaveAs("hapax_list.xlsx")
}
