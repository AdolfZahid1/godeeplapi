package models

// UsageAndLimitResponse represents API response from usage endpoint.
type UsageAndLimitResponse struct {
	//Characters translated so far in the current billing period.
	CharCount int64 `json:"character_count,omitempty"`
	//Current maximum number of characters that can be translated per billing period. If cost control is set, the cost control limit will be returned in this field.
	CharLimit int64 `json:"character_limit,omitempty"`
}
