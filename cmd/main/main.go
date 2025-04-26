package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"github.com/xuri/excelize/v2"
	utils "quita_mini/pkg/utils"
	calc "quita_mini/pkg/calc"
)

func main() {
	var sb strings.Builder
	absDirPath := "/home/kivanc57/Desktop/quita_mini/data"
	
	dir := utils.ReadDir(absDirPath)
	for _, f := range dir {
		file_name := f.Name()
		
		if strings.HasSuffix(file_name, ".txt") {
			absFilePath := filepath.Join(absDirPath, file_name)
			sb.WriteString(utils.ReadWriteText(absFilePath))
		}
	}
	corpus := sb.String()
	
	lemmaFeqMap := utils.GetFreqMap(corpus, true)
	tokenFreqMap := utils.GetFreqMap(corpus, false)

	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet in the Excel file
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// Write header row
	headers := []string{
		"TTR", "Entropy", "NormEntropy", "AvgTokenLen", "AvgTypeLen", 
		"TokenLenSD", "TypeLenSD", "YulesK", "AdjustedModulus", 
		"PercOfTextHapax", "PercOfTypeHapax", "R1", "RIndex", "RrIndex", "LIndex",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		f.SetCellValue(sheet, cell, header)
	}

	// Write data row
	data := []float64{
		calc.CalcTTR(lemmaFeqMap, tokenFreqMap),
		calc.CalcEntropy(tokenFreqMap),
		calc.CalcNormEntropy(tokenFreqMap),
		calc.CalcAvgTokenLen(tokenFreqMap),
		calc.CalcAvgTypeLen(lemmaFeqMap),
		calc.CalcTokenLenSD(tokenFreqMap),
		calc.CalcTypeLenSD(lemmaFeqMap),
		calc.CalcYulesK(tokenFreqMap),
		calc.CalcAdjustedModulus(lemmaFeqMap, tokenFreqMap),
		calc.CalcPercOfTextHapax(tokenFreqMap),
		calc.CalcPercOfTextHapax(lemmaFeqMap),
		calc.CalcR1(tokenFreqMap),
		calc.CalcRIndex(lemmaFeqMap, tokenFreqMap),
		calc.CalcRrIndex(tokenFreqMap),
		calc.CalcLIndex(lemmaFeqMap, tokenFreqMap),
	}

	// Fill data into the row
	for i, value := range data {
		cell := fmt.Sprintf("%s2", string('A'+i))
		f.SetCellValue(sheet, cell, calc.RoundToDecimal(value))
	}

	// Save the file
	if err := f.SaveAs("results.xlsx"); err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Results written to 'results.xlsx'")
}
