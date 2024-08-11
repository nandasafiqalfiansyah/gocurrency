package gocurrency

import (
	"fmt"
	"github.com/bojanz/currency"
	"strconv"
	"strings"
)

// FormatCurrency converts an integer to a formatted currency string.
func FormatCurrency(amount int64, currencyCode string, location string) (string, error) {
	// Konversi amount dari int64 ke currency.Amount
	amountValue, err := currency.NewAmount(fmt.Sprintf("%d", amount), currencyCode)
	if err != nil {
		return "", fmt.Errorf("error creating currency amount: %v", err)
	}

	// Buat formatter dengan locale yang diberikan
	formatter := currency.NewFormatter(currency.WithLocale(location))

	// Format nilai menjadi string mata uang
	formatted := formatter.Format(amountValue)
	return formatted, nil
}

// ParseCurrency converts a formatted currency string back to an integer.
func ParseCurrency(currencyStr string, currencyCode string) (int64, error) {
	// Buang simbol mata uang dan pemisah ribuan
	currencyStr = strings.ReplaceAll(currencyStr, ",", "")
	currencyStr = strings.ReplaceAll(currencyStr, ".", "")
	currencyStr = strings.ReplaceAll(currencyStr, currencyCode, "")
	currencyStr = strings.TrimSpace(currencyStr)

	// Konversi string ke int64
	amount, err := strconv.ParseInt(currencyStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency: %v", err)
	}
	return amount, nil
}
