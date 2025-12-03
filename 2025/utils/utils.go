package utils

import (
	"bufio"
	"log"
	"os"
)

func ForEachLineInFile(path string, f func(s string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		f(scanner.Text())
	}
}
