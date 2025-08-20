// main.go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// This program converts an XLSX file to one or more CSV files.
// It requires an input XLSX file path and an output directory path as command-line arguments.
func main() {
	// Check if the correct number of command-line arguments are provided.
	if len(os.Args) < 3 {
		fmt.Println("Usage: convertxls2csv <input.xlsx> <output_directory>")
		os.Exit(1)
	}

	xlsxPath := os.Args[1]
	outputDir := os.Args[2]

	// Create the output directory if it doesn't already exist.
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	// Open the XLSX file for reading.
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		log.Fatalf("Error opening XLSX file: %v", err)
	}
	defer f.Close()

	// Get a list of all sheet names in the XLSX file.
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		log.Println("No sheets found in the XLSX file.")
		return
	}

	// Iterate through each sheet and convert it to a separate CSV file.
	for _, sheetName := range sheets {
		// Get all rows from the current sheet.
		rows, err := f.GetRows(sheetName)
		if err != nil {
			log.Printf("Warning: Could not get rows from sheet '%s', skipping: %v\n", sheetName, err)
			continue
		}

		// Skip if the sheet is empty.
		if len(rows) == 0 {
			log.Printf("Sheet '%s' is empty, skipping.\n", sheetName)
			continue
		}

		// Construct the output CSV file path based on the sheet name.
		outputCSVPath := filepath.Join(outputDir, fmt.Sprintf("%s.csv", sheetName))

		// Create the new CSV file.
		csvFile, err := os.Create(outputCSVPath)
		if err != nil {
			log.Printf("Warning: Could not create CSV file for sheet '%s', skipping: %v\n", sheetName, err)
			continue
		}
		defer csvFile.Close()

		// Create a new CSV writer that writes to our new file.
		csvWriter := csv.NewWriter(csvFile)

		// Write all the rows from the XLSX sheet to the CSV file.
		if err := csvWriter.WriteAll(rows); err != nil {
			log.Printf("Warning: Could not write data to CSV for sheet '%s', skipping: %v\n", sheetName, err)
			continue
		}

		// Flush the writer to ensure all buffered data is written to the file.
		csvWriter.Flush()

		fmt.Printf("Successfully converted sheet '%s' to '%s'\n", sheetName, outputCSVPath)
	}
}
