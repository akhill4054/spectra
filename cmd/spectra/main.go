package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := resolvePort()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("starting HTTP server on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

// resolvePort determines the port precedence:
//   1. --p flag
//   2. PORT environment variable
//   3. 8080 (hardcoded default)
func resolvePort() int {
	defaultPort := 8080

	envPort := os.Getenv("PORT")
	if envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			defaultPort = p
		} else {
			log.Printf("invalid PORT value %q, falling back to default %d", envPort, defaultPort)
		}
	}

	flagPort := flag.Int("p", 0, "port to listen on")
	flag.Parse()

	if *flagPort != 0 {
		return *flagPort
	}

	return defaultPort
}

