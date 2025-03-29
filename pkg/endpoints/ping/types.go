package ping

// PingResponse represents the response from the Ping API
type PingResponse struct {
	// Gecko_says is the response message from the API
	GeckoSays string `json:"gecko_says"`
}
