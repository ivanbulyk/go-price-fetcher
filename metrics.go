package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	fmt.Println("pushing to prometheus")
	//your metrics storage. Push to Prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
