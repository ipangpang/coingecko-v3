package key

// KeyResponse represents the response from the Key API
type KeyResponse struct {
	// Plan is the subscription plan name
	Plan string `json:"plan"`
	// RateLimitRequestPerMinute is the rate limit for requests per minute
	RateLimitRequestPerMinute int `json:"rate_limit_request_per_minute"`
	// MonthlyCallCredit is the total monthly call credit
	MonthlyCallCredit int `json:"monthly_call_credit"`
	// CurrentTotalMonthlyCalls is the current total monthly calls made
	CurrentTotalMonthlyCalls int `json:"current_total_monthly_calls"`
	// CurrentRemainingMonthlyCalls is the current remaining monthly calls
	CurrentRemainingMonthlyCalls int `json:"current_remaining_monthly_calls"`
}
