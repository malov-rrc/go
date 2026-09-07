package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseLine(line string) (command string, amount float64, err error) {
	splittedLine := strings.Split(line, " ")
	if len(splittedLine) != 2 {
		return "", 0, errors.New("Invalid line format")
	}
	amount, err = strconv.ParseFloat(splittedLine[1], 64)
	if err != nil {
		return "", 0, errors.New("cant parse amount")
	}
	return splittedLine[0], amount, nil
}
