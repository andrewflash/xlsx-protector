package main

import (
    "fmt"
    "log"
    "os"

    "github.com/xuri/excelize/v2"
)

func main() {
    if len(os.Args) != 5 {
        fmt.Println("Usage: excel-protector <-e|-d> <input.xlsx> <output.xlsx> <password>")
        os.Exit(1)
    }

    operation := os.Args[1]
    inputFilePath := os.Args[2]
    outputFilePath := os.Args[3]
    password := os.Args[4]

    if operation != "-e" && operation != "-d" {
        fmt.Println("Operation must be either -e (encrypt) or -d (decrypt)")
        os.Exit(1)
    }

    if operation == "-e" {
		f, err := excelize.OpenFile(inputFilePath)
		if err != nil {
			log.Fatalf("Failed to open the input file: %v", err)
		}
		defer f.Close()

		if err := f.SaveAs(outputFilePath, excelize.Options{Password: password}); err != nil {
			log.Fatalf("Failed to save the file: %v", err)
		}
		fmt.Println("Successfully encrypted the workbook")
    } else {
		f, err := excelize.OpenFile(inputFilePath, excelize.Options{Password: password})
		if err != nil {
			log.Fatalf("Failed to open the input file: %v", err)
		}
		defer f.Close()

		if err := f.SaveAs(outputFilePath, excelize.Options{}); err != nil {
			log.Fatalf("Failed to save the file: %v", err)
		}
		fmt.Println("Successfully decrypted the workbook")
    }
}