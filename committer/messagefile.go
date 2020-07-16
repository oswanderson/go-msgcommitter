package committer

import (
	"log"
	"os"
)

// WriteMessageToFile writes the commit formated commit message to the file
// passed as an argumetn by the command git commit.
func WriteMessageToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0664)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = file.WriteString(content); err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}
