package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"visitors/handler"
)

func main() {
	log := handler.NewVisitorLog()
	scanner := bufio.NewScanner(os.Stdin)

	running := true

	for running {
		fmt.Println("Choose the command (Save | GetHistory | GetLastVisit | Exit)")
		if !scanner.Scan() {
			break
		}
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))
		switch command {
		case "save":
			handler.Save(scanner, log)
		case "gethistory":
			handler.GetHistory(scanner, log)
		case "getlastvisit":
			handler.GetLastVisit(scanner, log)
		case "exit":
			running = false
		default:
			fmt.Println("Unknown command")
		}
	}
}
