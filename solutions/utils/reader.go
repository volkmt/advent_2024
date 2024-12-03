package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertToInt(s string) int {
	intValue, err := strconv.Atoi(s)
	CheckError(err)

	return intValue
}

func ReadFileIntoSlices(fp string, d string) ([]int, []int, error) {
	file, err := os.Open(fp)
	CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sliceOne := []int{}
	sliceTwo := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, d)

		if len(parts) == 2 {
			sliceOne = append(sliceOne, ConvertToInt(parts[0]))
			sliceTwo = append(sliceTwo, ConvertToInt(parts[1]))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return sliceOne, sliceTwo, nil
}
