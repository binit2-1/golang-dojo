package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var http_requests_total = promauto.NewCounter(prometheus.CounterOpts{
	Name: "ticketblitz_requests_total",
	Help: "The total number of processed requests",
})

var http_request_duration_seconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "ticketblitz_request_duration_seconds",
	Help: "The duration of processed requests in seconds",
}, []string{"method", "path"})