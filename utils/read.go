package utils

import (
	"io"
	"os"
	"log"
	"io/fs"
	strs "strings"
)

func ReadDir(path string) []fs.DirEntry {
	dir, err := os.ReadDir(path)
	if err != nil{
		log.Printf("Error reading %s: %v", path, err)
		return nil
	}
	return dir
}

func ReadWriteText(file_name string) string {
	var content strs.Builder
	text, err := os.Open(file_name)

	if err != nil{
		log.Printf("Error opening file %s: %v", file_name, err)
		return ""
	}
	defer text.Close()
	
	buf := make([]byte, 32*1024)	
	for {
		n, err := text.Read(buf)
		
		if n > 0 {
			content.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}
	return content.String()
}
