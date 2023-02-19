package utils

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
)

// isSupportedCurrency returns if currency param is supported else false
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR:
		return true
	}
	return false
}
