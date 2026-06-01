package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourorg/ratls/pkg/ratls"
)

func main() {
	audience := getEnv("RATLS_AUDIENCE", "ratls-buffer-tee")
	listenAddr := getEnv("LISTEN_ADDR", ":443")

	log.Printf("gpu-cs: starting (audience=%s)", audience)

	srv, err := ratls.NewServer(audience)
	if err != nil {
		log.Fatalf("gpu-cs: init server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle(ratls.ConnectPath, srv.ConnectHandler())
	mux.HandleFunc("/api/load-model", handleLoadModel)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	httpSrv := &http.Server{
		Addr:      listenAddr,
		Handler:   mux,
		TLSConfig: srv.TLSConfig(),
	}

	log.Printf("gpu-cs: listening on %s", listenAddr)
	if err := httpSrv.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("gpu-cs: server error: %v", err)
	}
}

func handleLoadModel(w http.ResponseWriter, r *http.Request) {
	log.Printf("gpu-cs: load model request from %s", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
