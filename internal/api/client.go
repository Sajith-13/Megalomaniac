package api

import (
	"net/http"
	"time"
)

type TradeClient struct {
	httpClient *http.Client
	userAgent  string
}

func NewTradeClient(contactEmail string) *TradeClient {
	return &TradeClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		userAgent:  "MegaloMarketAnalyzer/1.0 (Contact: " + contactEmail + ")",
	}
}

// TODO: Add functions for Search() and Fetch() using rate-limiters