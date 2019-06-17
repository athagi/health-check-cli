package delete

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Delete(args []string) {
	filePath := viper.GetString("recordFileName")

	lines, err := readLines(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range args {
		lines = removeItem(lines, arg)
	}

	write(filePath, lines)

}

func removeItem(lists []string, word string) []string {
	var res []string
	for _, line := range lists {
		if line == word {
			continue
		}
		res = append(res, line)
	}
	return res
}

func write(filePath string, lists []string) {
	f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, line := range lists {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readLines(filePath string) ([]string, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Open file error")
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
