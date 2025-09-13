# Barangay Incident Report System (GoLang)

## Overview
The Barangay Incident Report System is a simple Go-based console application designed to help barangay offices (or schools) record and manage incident reports.  
This project demonstrates a practical barangay-level solution while showcasing Go’s concurrency features (goroutines and channels).  

When an incident report is created, the system:  
- Saves it locally to `reports.txt`.  
- Notifies multiple barangay officers concurrently, simulating real-time acknowledgment.  

---

## Features
- **Create Incident Reports** – Collect details such as complainant, respondent, location, date, description, and officer-in-charge.  
- **View Saved Reports** – Admins can log in and view all stored reports from `reports.txt`.  
- **Concurrent Officer Notification (Concurrency Demo)** – Multiple officers receive the report simultaneously using goroutines and channels.  

---

## Technologies Used
- **Language:** Go (Golang)  
- **Standard Libraries:** `fmt`, `os`, `bufio`, `strings`, `time`, `math/rand`  
- **Concurrency:** Goroutines and Channels  

---

## Project Structure
```
.
├── main.go         # Main program logic
├── reports.txt     # Auto-generated file where reports are saved
└── README.md       # Documentation
```

---

## How to Build
1. Ensure Go is installed (version 1.19+ recommended).  
2. Clone or download the repository.  
3. Run the build command in the project folder:  
```bash
go build -o incident-report-system main.go
```  
The executable `incident-report-system` will be generated in your directory.  

---

## How to Run
Run directly with Go using terminal:  
```bash
go run Demo/demo.go
```

**Menu Options:**  
- `1` → Create Incident Report  
- `2` → Check Reports (admin login required: username: `brenty`, password: `brenty`)  
- `exit` → Close program  

---

## Sample Output (Concurrency Demo)
```
Notifying barangay officers...
Officer Bravo received the report about Juan Dela Cruz vs Pedro Santos at Session Road (after 1s)
Officer Alpha received the report about Juan Dela Cruz vs Pedro Santos at Session Road (after 2s)
Officer Charlie received the report about Juan Dela Cruz vs Pedro Santos at Session Road (after 3s)
```
