package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "request_total",
			Help:      "Number of request processed by this service.",
		}, []string{},
	)

	requestCountLowerHalf = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "request_lower_half",
			Help:      "Number of request processed by this service that gets a response with a number little than a half of the upper limit.",
		}, []string{},
	)

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "request_latency_seconds",
			Help:      "Time spent in this service.",
			Buckets:   []float64{1e-5,2e-5,3e-5,4e-5,5e-5,6e-5,7e-5,8e-5,9e-5,1e-4},
		}, []string{},
	)
)

// AdmissionLatency measures latency / execution time of Admission Control execution
// usual usage pattern is: timer := NewAdmissionLatency() ; compute ; timer.Observe()
type RequestLatency struct {
	histo *prometheus.HistogramVec
	start time.Time
}

func Register() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestCountLowerHalf)
	prometheus.MustRegister(requestLatency)
}


// NewAdmissionLatency provides a timer for admission latency; call Observe() on it to measure
func NewAdmissionLatency() *RequestLatency {
	return &RequestLatency{
		histo: requestLatency,
		start: time.Now(),
	}
}

// Observe measures the execution time from when the AdmissionLatency was created
func (t *RequestLatency) Observe() {
	(*t.histo).WithLabelValues().Observe(time.Now().Sub(t.start).Seconds())
}


// RequestIncrease increases the counter of request handled by this service
func RequestIncrease() {
	requestCount.WithLabelValues().Add(1)
}

// RequestIncrease increases the counter of special request handled by this service
func RequestIncreaseLowerHalf() {
	requestCountLowerHalf.WithLabelValues().Add(1)
}
