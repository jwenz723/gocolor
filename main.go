package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Account string
	Color   string
	Region  string
}

func getConfig() Config {
	account := getEnvOrDefault("GOCOLOR_ACCOUNT", "unknown")
	color := getEnvOrDefault("GOCOLOR_COLOR", "aquamarine")
	region := getEnvOrDefault("GOCOLOR_REGION", "unknown")
	return Config{
		Account: account,
		Color:   color,
		Region:  region,
	}
}

func main() {
	addr := flag.String("addr", ":8080", "the address to run the http server on")
	flag.Parse()

	config := getConfig()

	http.Handle("/", handle(config))
	fmt.Printf("starting http on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
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
