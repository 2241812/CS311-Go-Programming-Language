package demo

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type IncidentReport struct {
	Complainant string
	Respondent  string
	Location    string
	Date        time.Time
	Description string
	Officer     string
}

func main() {
	rand.Seed(time.Now().UnixNano())

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
			var complainant string = readInput("Enter Complainant Name : ")
			var respondent string = readInput("Enter Respondent Name  : ")
			var location string = readInput("Enter Location         : ")
			description := readInput("Enter Incident Details : ")
			officer := readInput("Enter Officer's Name   : ")

			report := IncidentReport{
				Complainant: complainant,
				Respondent:  respondent,
				Location:    location,
				Date:        time.Now(),
				Description: description,
				Officer:     officer,
			}

			fmt.Println("\nGenerated Report:")
			fmt.Println(report.GenerateReport())

			saveReport(report.GenerateReport())

		
			notifyOfficersConcurrently(report)

		case "2":
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

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

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

func viewReports() {
	data, err := os.ReadFile("reports.txt")
	if err != nil {
		fmt.Println("No reports found or error reading file.")
		return
	}
	fmt.Println("\n=== ALL SAVED REPORTS ===")
	fmt.Println(string(data))
}


func notifyOfficersConcurrently(report IncidentReport) {
	officers := []string{"Officer A", "Officer B", "Officer C"}
	results := make(chan string, len(officers))

	fmt.Println("\n📢 Notifying barangay officers...")

	for _, officer := range officers {
		go func(officer string) {
			delay := time.Duration(rand.Intn(3)+1) * time.Second
			time.Sleep(delay)
			results <- fmt.Sprintf("%s received the report about %s vs %s at %s (after %v)",
				officer, report.Complainant, report.Respondent, report.Location, delay)
		}(officer)
	}

	
	for i := 0; i < len(officers); i++ {
		fmt.Println(<-results)
	}
}
