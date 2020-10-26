package main

import (
	"flag"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Config collects all configuration for gocolor
type Config struct {
	Account               string
	Color                 string
	InfoDurationHistogram *prometheus.Histogram
	Region                string
}

// getConfig will instantiate a Config instance from environment variables
func getConfig() Config {
	account := getEnvOrDefault("GOCOLOR_ACCOUNT", "unknown")
	color := getEnvOrDefault("GOCOLOR_COLOR", "aquamarine")
	region := getEnvOrDefault("GOCOLOR_REGION", "unknown")

	infoDurationHistogram := prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
		Name: "info_http_request_durations_seconds",
		Help: "HTTP latency distributions to the /info endpoint",
		ConstLabels: map[string]string{
			labelAccount: account,
			labelRegion:  region,
		},
	}, nil)

	return Config{
		Account:               account,
		Color:                 color,
		InfoDurationHistogram: infoDurationHistogram,
		Region:                region,
	}
}

const (
	labelAccount = "account"
	labelRegion  = "region"
)

func main() {
	addr := flag.String("addr", ":8080", "the address to run the http server on")
	flag.Parse()

	config := getConfig()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	http.Handle("/info", handleInfo(config))
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	l := logger.With(zap.String("addr", *addr))
	l.Info("starting http listener...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleInfo(c Config) http.HandlerFunc {
	t, err := template.New("t").Parse(`
<html>
	<body style="background-color: {{ .Color }};">
    	<div>Account: {{ .Account }}</div>
		<div>Region: {{ .Region }}</div>
	</body>
</html>`)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		defer func(begin time.Time) {
			c.InfoDurationHistogram.Observe(time.Since(begin).Seconds())
		}(time.Now())
		log.Printf("handled request")
		t.Execute(w, c)
	}
}

func getEnvOrDefault(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		v = def
	}
	return v
}
