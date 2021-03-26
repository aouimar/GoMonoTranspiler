package main

import (
	"os"

	"aouimar.com/translater"
)

func main() {
	inputFilePath := os.Args[1]
	outputFile := os.Args[2]
	translater.TranslateFile(inputFilePath, outputFile)
}
