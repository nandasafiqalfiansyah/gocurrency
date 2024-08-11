# 🌍 gocurrency v1.1.0

`gocurrency` is a Golang library for formatting float values ​​into currency strings in various worldwide currency formats, as well as converting formatted currency strings back into float values.
<hr>

[![Go](https://skillicons.dev/icons?i=go)](https://skillicons.dev)

## 📦 Installation
To use the latest stable version:
```bash
go get github.com/nandasafiqalfiansyah/gocurrency@v1.1.0
```

## 🌐 CurrencyCode and Locale Usage Table

| 🌏 Currency         | 💲 Currency Code | 🗺️ Locale     | 📊 Example Format                                |
|--------------------|----------------|------------|-----------------------------------------------|
| 🇺🇸 Dolar Amerika   | USD            | en_US      | $1,234,567.89                                 |
| 🇪🇺 Euro            | EUR            | de_DE      | 1.234.567,89 €                                |
| 🇬🇧 Pound Sterling  | GBP            | en_GB      | £1,234,567.89                                 |
| 🇯🇵 Yen Jepang      | JPY            | ja_JP      | ￥1,234,567                                    |
| 🇮🇩 Rupiah Indonesia| IDR            | id_ID      | Rp1.234.567,89                                |
| 🇮🇳 Rupee India     | INR            | hi_IN      | ₹1,23,45,678.89                               |
| 🇦🇺 Dolar Australia | AUD            | en_AU      | $1,234,567.89                                 |
| 🇨🇦 Dolar Kanada    | CAD            | en_CA      | $1,234,567.89                                 |
| 🇨🇭 Franc Swiss     | CHF            | fr_CH      | 1'234'567.89 CHF                              |
| 🇨🇳 Yuan Tiongkok   | CNY            | zh_CN      | ￥1,234,567.89                                 |
| 🇷🇺 Ruble Rusia     | RUB            | ru_RU      | 1 234 567,89 ₽                                |
| 🇰🇷 Won Korea Selatan| KRW            | ko_KR      | ₩1,234,567                                    |
| 🇲🇽 Peso Meksiko    | MXN            | es_MX      | $1,234,567.89                                 |
| 🇿🇦 Rand Afrika Selatan| ZAR       | en_ZA      | R1,234,567.89                                 |
| 🇧🇷 Real Brasil     | BRL            | pt_BR      | R$1.234.567,89                                |
| 🇲🇾 Ringgit Malaysia| MYR            | ms_MY      | RM1,234,567.89                                |
| 🇹🇭 Baht Thailand   | THB            | th_TH      | ฿1,234,567.89                                 |
| 🇵🇭 Peso Filipina   | PHP            | en_PH      | ₱1,234,567.89                                 |
| 🇰🇼 Dinar Kuwait    | KWD            | ar_KW      | د.ك1,234,567.890                               |
| 🇸🇦 Riyal Saudi     | SAR            | ar_SA      | ﷼1,234,567.89                                 |

### 🛠️ Use

To format or convert back currency with `currencyCode` and `locale`, you can use the `FormatCurrency` and `ParseCurrency` functions with the appropriate parameters as shown in the table above.

#### Example
```go
formattedUsd, err := gocurrency.FormatCurrency(123456789, "USD", "en_US")
decodedUsd, err := gocurrency.ParseCurrency(formattedUsd, "USD")
```
