package main

import (
	"fmt"
	"path/filepath"
	"strings"
	utils"quita_mini/pkg/utils"
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
	
	tokenFreqMap := utils.GetFreqMap(corpus, false)
	fmt.Println(tokenFreqMap)
	
	
	var token_sum int = 0
	for _, freq := range tokenFreqMap{
		token_sum += freq
	}
	fmt.Println(token_sum)

}
