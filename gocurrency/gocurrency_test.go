package gocurrency

import (
	"testing"
)

func TestFormatCurrency(t *testing.T) {
	tests := []struct {
		name         string
		amount       float64
		currencyCode string
		locale       string
		want         string
	}{
		{"USD-en_US", 1234.56, "USD", "en_US", "$1,234.56"},
		{"EUR-de_DE", 1234.56, "EUR", "de_DE", "1.234,56 €"},
		{"JPY-ja_JP", 1234, "JPY", "ja_JP", "￥1,234"},
		{"IDR-id_ID", 1234.56, "IDR", "id_ID", "Rp1.234,56"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatCurrency(tt.amount, tt.currencyCode, tt.locale)
			if err != nil {
				t.Errorf("FormatCurrency() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("FormatCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCurrency(t *testing.T) {
	tests := []struct {
		name         string
		currencyStr  string
		currencyCode string
		want         float64
	}{
		{"USD", "$1,234.56", "USD", 1234.56},
		{"EUR", "1.234,56 €", "EUR", 1234.56},
		{"JPY", "￥1,234", "JPY", 1234},
		{"IDR", "Rp1.234,56", "IDR", 1234.56},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCurrency(tt.currencyStr, tt.currencyCode)
			if err != nil {
				t.Errorf("ParseCurrency() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("ParseCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
