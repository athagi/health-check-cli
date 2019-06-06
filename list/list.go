package list

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func ListTargets() {
	filePath := viper.GetString("recordFileName")

	lines, err := readLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range lines {
		fmt.Fprintln(os.Stdout, l)
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
