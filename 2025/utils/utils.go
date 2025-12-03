package utils

import (
	"bufio"
	"os"
)

func ForEachLineInFile(path string, f func(s string)) {
	file, err := os.Open(path)
	if err != nil {
		panic(err) // I'm here for a good time, not for a long time
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		f(scanner.Text())
	}

}
