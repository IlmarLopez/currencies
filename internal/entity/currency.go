package entity

// Currency represents a currency record.
type Currency struct {
	ID    int     `json:"id"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
