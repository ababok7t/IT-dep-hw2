package file

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	inputFile, inputErr := os.Open(path)
	if inputErr != nil {
		fmt.Println("input error: ", inputErr)
		return "input error: " + inputErr.Error()
	}
	defer inputFile.Close()
	data := make([]byte, 1024)
	readingByte, readingErr := inputFile.Read(data)
	if readingErr != nil {
		fmt.Println("reading error: ", readingByte)
		return "reading error: " + readingErr.Error()
	}
	text := string(data[:readingByte])
	return text
}
