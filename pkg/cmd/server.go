package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tkennes/open-azure-emissions/pkg/azure"
	"github.com/tkennes/open-azure-emissions/pkg/env"
	"github.com/tkennes/open-azure-emissions/pkg/log"
	"github.com/tkennes/open-azure-emissions/pkg/metrics"
	"github.com/tkennes/open-azure-emissions/pkg/version"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// AgentOpts contain configuration options that can be passed to the Execute() method
type ServerOpts struct {
	// Stubbed for future configuration
}

type Response struct {
	Status     string `json:"status"`
	StatucCode int    `json:"statusCode"`
	Reason     string `json:"reason"`
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Length", "0")
	w.Header().Set("Content-Type", "text/plain")
	log.Infof("Received healthz request at %s%s from %s", r.Host, r.RequestURI, r.RemoteAddr)
}

func FootPrintRoute(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received footprint request at %s%s from %s", r.Host, r.RequestURI, r.RemoteAddr)

	if r.Header.Get("Content-Type") == "application/json" {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error reading request body: %v", err)
		}
		footprint, err := azure.GetFootprint(reqBody)
		if err != nil {
			log.Infof("Error getting footprint: %v", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&Response{"Fail", 400, fmt.Sprintf("Field Validation error %s", err)})
		} else {
			w.Header().Set("Content-Type", "application/json")
			log.Infof("Result: %s", footprint.(string))
			json.NewEncoder(w).Encode(&Response{"Success", 200, footprint.(string)})
		}
	} else {
		log.Infof("Received footprint request with invalid Content-Type header at %s%s from %s", r.Host, r.RequestURI, r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&Response{"Fail", 400, "Content-Type header must be application/json"})
	}
}

func RunServer(opts *ServerOpts) error {
	log.Infof("Starting Emissions Agent version %s", version.FriendlyVersion())

	// ServeMuxes are also known as routers
	router := mux.NewRouter()

	// Middleware
	router.Use(metrics.PrometheusMiddleware)

	// Routes
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/footprint", FootPrintRoute).Methods("POST")

	// set a HTTP request handle function for path /greeting and registrate it
	router.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Received greeting request from %s", r.RemoteAddr)
		// when receive the request, print the greeting meassage
		fmt.Fprint(w, "Hello World")
	}).Methods("GET")

	// Metrics
	metrics.NewOpenAzureEmissionsEmitter()
	router.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(fmt.Sprintf(":%d", env.GetOpenAzureEmissionsPort()), router)
}
