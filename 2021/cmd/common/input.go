package common

import (
	"bytes"
	"log"
	"os"
)

func ReadInputFile(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Error opening "+fileName, err)
	}

	defer file.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(file)
	return buffer.String()
}
