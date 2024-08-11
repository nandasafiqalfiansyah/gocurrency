package gocurrency

import (
	"fmt"
	"github.com/bojanz/currency"
	"regexp"
	"strconv"
	"strings"
)

// sampel code formating
func FormatCurrency(amount int64, currencyCode string, locale string) (string, error) {
	amountValue, err := currency.NewAmount(fmt.Sprintf("%d", amount), currencyCode)
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

func ParseCurrency(currencyStr string, currencyCode string) (int64, error) {
	re := regexp.MustCompile(`[^\d]`)
	cleanedStr := re.ReplaceAllString(currencyStr, "")
	amount, err := strconv.ParseInt(cleanedStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency: %v", err)
	}
	return amount, nil
}
