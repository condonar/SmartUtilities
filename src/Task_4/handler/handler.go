package handler

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

var (
	ErrNotFoundVisit     = errors.New("not found visit")
	ErrInvalidDateFormat = errors.New("invalid date format (need YYYY-MM-DD)")
	ErrInput             = errors.New("input error")
)

type Visit struct {
	spec      string
	visitDate time.Time
}

type VisitorLog struct {
	storedData map[string][]Visit
}

func NewVisitorLog() *VisitorLog {
	return &VisitorLog{
		storedData: make(map[string][]Visit),
	}
}

func readLine(scanner *bufio.Scanner, prompt string) string {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return ""
		}
		result := strings.TrimSpace(scanner.Text())
		if result != "" {
			return result
		}
	}
}

func (data *VisitorLog) AddVisit(fullName, specInput, dateStr string) error {
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return ErrInvalidDateFormat
	}
	data.storedData[fullName] = append(data.storedData[fullName], Visit{specInput, date})
	return nil
}

func Save(scanner *bufio.Scanner, log *VisitorLog) {
	fullName := readLine(scanner, "Full Name: ")
	specInput := readLine(scanner, "Doctor: ")
	dateStr := readLine(scanner, "Date: ")

	if err := log.AddVisit(fullName, specInput, dateStr); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Visit successfully saved")
	}
}

func GetHistory(scanner *bufio.Scanner, log *VisitorLog) {
	fullName := readLine(scanner, "Full Name: ")
	result, ok := log.storedData[fullName]
	if !ok {
		fmt.Println(ErrNotFoundVisit)
	} else {
		for _, visit := range result {
			fmt.Printf("%s %s\n", visit.spec, visit.visitDate.Format(dateFormat))
		}
	}
}

func GetLastVisit(scanner *bufio.Scanner, log *VisitorLog) {
	fullName := readLine(scanner, "Full Name: ")
	specInput := readLine(scanner, "Doctor: ")

	result, ok := log.storedData[fullName]
	if !ok {
		fmt.Println(ErrNotFoundVisit)
	} else {
		var lastVisit Visit
		found := false
		for _, visit := range result {
			if visit.spec == specInput {
				if visit.visitDate.After(lastVisit.visitDate) {
					lastVisit = visit
					found = true
				}
			}
		}
		if !found {
			fmt.Println(ErrNotFoundVisit)
		}
		fmt.Println(lastVisit.spec, lastVisit.visitDate.Format(dateFormat))
	}
}
