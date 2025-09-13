package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// IncidentReport holds details of the report
type IncidentReport struct {
	Complainant string
	Respondent  string
	Location    string
	Date        time.Time
	Description string
	Officer     string
}

// Create a local random generator (no deprecated rand.Seed)
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for {
		fmt.Println("\n=== Barangay Incident Report System ===")
		fmt.Println("1. Create Incident Report")
		fmt.Println("2. Check Reports")
		fmt.Println("Type exit to close the program.")

		choice := strings.ToLower(readInput("Choose an option: "))

		if choice == "exit" {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		switch choice {
		case "1":
			// Collect report details
			complainant := readInput("Enter Complainant Name : ")
			respondent := readInput("Enter Respondent Name  : ")
			location := readInput("Enter Location         : ")
			description := readInput("Enter Incident Details : ")
			officer := readInput("Enter Officer's Name   : ")

			// Create report struct
			report := IncidentReport{
				Complainant: complainant,
				Respondent:  respondent,
				Location:    location,
				Date:        time.Now(),
				Description: description,
				Officer:     officer,
			}

			// Print and save report
			fmt.Println("\nGenerated Report:")
			fmt.Println(report.GenerateReport())
			saveReport(report.GenerateReport())

			// Notify officers concurrently (with spinner)
			notifyOfficersConcurrently(report)

		case "2":
			// Admin login before viewing reports
			username := readInput("Enter admin username: ")
			password := readInput("Enter admin password: ")

			if username == "brenty" && password == "brenty" {
				fmt.Println("Admin login successful.")
				viewReports()
			} else {
				fmt.Println("Invalid admin credentials.")
			}

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// Format the report for saving/printing
func (ir IncidentReport) GenerateReport() string {
	return fmt.Sprintf(`

      BARANGAY INCIDENT REPORT

Date       : %s
Location   : %s

Complainant: %s
Respondent : %s

Incident Details:
%s

Prepared by: %s

`, ir.Date.Format("January 02, 2006 03:04 PM"), ir.Location, ir.Complainant, ir.Respondent, ir.Description, ir.Officer)
}

// Read input safely
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Save reports to file
func saveReport(report string) {
	file, err := os.OpenFile("reports.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving report:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(report + "\n\n"); err != nil {
		fmt.Println("Error writing report:", err)
	}
	fmt.Println("Report saved to reports.txt")
}

// View saved reports
func viewReports() {
	data, err := os.ReadFile("reports.txt")
	if err != nil {
		fmt.Println("No reports found or error reading file.")
		return
	}
	fmt.Println("\n=== ALL SAVED REPORTS ===")
	fmt.Println(string(data))
}

// Progress spinner goroutine
// reference https://stackoverflow.com/questions/4995733/how-to-create-a-spinning-command-line-cursor
func spinner(done chan bool) {
	chars := `|/-\`
	for {
		for _, r := range chars {
			select {
			case <-done:
				fmt.Print("\r") // clear line when done
				return
			default:
				fmt.Printf("\r%c Notifying officers...", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

// Notify officers concurrently (with spinner feedback)
func notifyOfficersConcurrently(report IncidentReport) {
	officers := []string{"Officer Alpha", "Officer Bravo", "Officer Charlie"}
	results := make(chan string, len(officers))

	// Channel to stop spinner
	done := make(chan bool)

	// Start spinner animation in background
	go spinner(done)

	// Start concurrent goroutines to notify officers
	for _, officer := range officers {
		go func(officer string) {
			// Random delay simulating response time
			delay := time.Duration(r.Intn(3)+1) * time.Second
			time.Sleep(delay)

			// Send result back
			results <- fmt.Sprintf("[Done] %s received the report about %s vs %s at %s (after %v)",
				officer, report.Complainant, report.Respondent, report.Location, delay)
		}(officer)
	}

	// Collect results
	for i := 0; i < len(officers); i++ {
		fmt.Println("\n" + <-results)
	}

	// Stop spinner
	done <- true
}
