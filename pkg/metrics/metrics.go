package metrics

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Only allow the metrics to be instantiated and registered once
var metricsInit sync.Once

var (
	totalRequests  *prometheus.CounterVec
	responseStatus *prometheus.CounterVec
	httpDuration   *prometheus.HistogramVec
)

type OpenAzureEmissionsEmitter struct {
	TotalRequests  *prometheus.CounterVec
	ResponseStatus *prometheus.CounterVec
	HttpDuration   *prometheus.HistogramVec
}

func initOpenAzurEmissionsMetrics() {
	metricsInit.Do(func() {
		totalRequests = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Number of get requests.",
			},
			[]string{"path"},
		)
		prometheus.MustRegister(totalRequests)

		responseStatus = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "response_status",
				Help: "Status of HTTP response",
			},
			[]string{"status"},
		)
		prometheus.MustRegister(responseStatus)

		httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_response_time_seconds",
			Help: "Duration of HTTP requests.",
		}, []string{"path"})
		prometheus.Register(httpDuration)
	})
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}

func NewOpenAzureEmissionsEmitter() {
	initOpenAzurEmissionsMetrics()
}
