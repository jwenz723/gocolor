package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Account string
	Address string
	Color string
	Region string
}

func getConfig() Config {
	account := getEnvOrDefault("GOCOLOR_ACCOUNT", "unknown")
	address := getEnvOrDefault("GOCOLOR_ADDRESS", ":8080")
	color := getEnvOrDefault("GOCOLOR_COLOR", "aquamarine")
	region := getEnvOrDefault("GOCOLOR_REGION", "unknown")
	return Config{
		Account: account,
		Address: address,
		Color:   color,
		Region:  region,
	}
}

func main() {
	config := getConfig()

	http.Handle("/", handle(config))
	fmt.Printf("starting http on %s\n", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

func handle(c Config) http.HandlerFunc {
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