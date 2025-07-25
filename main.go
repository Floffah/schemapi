package main

import (
	"github.com/floffah/schemapi/internal/logger"
	"os"
	"strings"
)

func main() {
	os.Exit(prog())
}

func prog() int {
	logger.PrintHeader()

	workingDir, err := os.Getwd()
	logger.HandleError(err)

	files, err := os.ReadDir(workingDir)
	logger.HandleError(err)

	count := 0

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sapi") {
			count++
		}
	}

	if count == 0 {
		logger.PrintError("No Schemapi files found in the current directory", nil)

		return 1
	}

	return 0
}
