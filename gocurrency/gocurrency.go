package gocurrency

import (
	"fmt"
	"github.com/bojanz/currency"
	"strconv"
	"strings"
)

func FormatCurrency(amount int64, currencyCode string, location string) (string, error) {
	formatter := currency.NewFormatter(currencyCode)
	formatter.SetCurrencyCode(currencyCode)
	formatter.SetLocale(location)

	formatted := formatter.Format(amount)
	if strings.Contains(formatted, ",") {
		formatted = strings.ReplaceAll(formatted, ",", ".")
	}
	return formatted, nil
}

func ParseCurrency(currencyStr string, currencyCode string) (int64, error) {
	currencyStr = strings.ReplaceAll(currencyStr, ",", "")
	currencyStr = strings.ReplaceAll(currencyStr, ".", "")
	currencyStr = strings.ReplaceAll(currencyStr, currencyCode, "")
	currencyStr = strings.TrimSpace(currencyStr)

	amount, err := strconv.ParseInt(currencyStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency: %v", err)
	}
	return amount, nil
}
