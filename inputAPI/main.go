package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const APIKeyHeader = "X-API-Key"
const ExpectedAPIKey = "your_expected_api_key_here"

type InputData struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
	Value  float64 `json:"value"`
}

type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

var metrics = make(map[string]prometheus.Gauge)

func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(APIKeyHeader)
		if apiKey != ExpectedAPIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func InputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var inputData InputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if metric, ok := metrics[inputData.Name]; !ok {

		metric = promauto.NewGauge(prometheus.GaugeOpts{
			Name: inputData.Name,
			Help: "_______",
		})
		metric.Set(inputData.Value)
		metrics[inputData.Name] = metric
	} else {
		metric.Set(inputData.Value)
	}

	// Here you would handle the inputData, e.g., save it to a database or process it

	response := ResponseMessage{Message: "Data successfully received."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/input", APIKeyMiddleware(http.HandlerFunc(InputHandler)))
	mux.Handle("/metrics", APIKeyMiddleware(promhttp.Handler()))

	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

var (
	MPeopleNearSensore = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "people_near_sesore",
		Help: "total nummber of people near a Sensore",
	})
)
