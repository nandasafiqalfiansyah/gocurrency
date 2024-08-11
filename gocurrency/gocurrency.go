package gocurrency

import (
	"fmt"
	"github.com/bojanz/currency"
	"regexp"
	"strconv"
	"strings"
)

func FormatCurrency(amount float64, currencyCode string, locale string) (string, error) {
	amountValue, err := currency.NewAmount(fmt.Sprintf("%f", amount), currencyCode)
	if err != nil {
		return "", fmt.Errorf("error creating currency amount: %v", err)
	}
	formatter := currency.NewFormatter(currency.NewLocale(locale))
	formatted := formatter.Format(amountValue)
	if strings.Contains(formatted, ",") {
		formatted = strings.ReplaceAll(formatted, ",", ".")
	}
	return formatted, nil
}

func ParseCurrency(currencyStr string, currencyCode string) (float64, error) {
	re := regexp.MustCompile(`[^\d.,]`)
	cleanedStr := re.ReplaceAllString(currencyStr, "")
	if strings.Contains(cleanedStr, ".") && strings.Contains(cleanedStr, ",") {
		cleanedStr = strings.ReplaceAll(cleanedStr, ",", "")
	} else if strings.Count(cleanedStr, ",") > 1 {
		cleanedStr = strings.ReplaceAll(cleanedStr, ",", "")
	} else if strings.Count(cleanedStr, ".") > 1 {
		cleanedStr = strings.ReplaceAll(cleanedStr, ".", "")
		cleanedStr = strings.ReplaceAll(cleanedStr, ",", ".")
	} else {
		if strings.Contains(cleanedStr, ",") {
			cleanedStr = strings.ReplaceAll(cleanedStr, ",", ".")
		}
	}
	amount, err := strconv.ParseFloat(cleanedStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency: %v", err)
	}
	return amount, nil
}
