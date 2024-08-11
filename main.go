package main

import (
	"fmt"
	"github.com/nandasafiqalfiansyah/gocurrency/gocurrency"
	"log"
)

func main() {
	// Format USD
	formattedUsd, err := gocurrency.FormatCurrency(13, "USD", "en_US")
	if err != nil {
		log.Fatalf("Error formatting currency: %v", err)
	}
	fmt.Printf("Formatted USD: %s\n", formattedUsd)

	// Decode USD
	decodedUsd, err := gocurrency.ParseCurrency(formattedUsd, "USD")
	if err != nil {
		log.Fatalf("Error parsing currency: %v", err)
	}
	fmt.Printf("Decoded USD: %f\n", decodedUsd)

	// Format IDR
	formattedIdr, err := gocurrency.FormatCurrency(123456789, "IDR", "id_ID")
	if err != nil {
		log.Fatalf("Error formatting currency: %v", err)
	}
	fmt.Printf("Formatted IDR: %s\n", formattedIdr)

	// Decode IDR
	decodedIdr, err := gocurrency.ParseCurrency(formattedIdr, "IDR")
	if err != nil {
		log.Fatalf("Error parsing currency: %v", err)
	}
	fmt.Printf("Decoded IDR: %f\n", decodedIdr)
}
