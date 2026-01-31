package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
	log.Printf("Tracker listening on port %v\n", port)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	clientIP := r.Header.Get("CF-Connecting-IP")
	if clientIP == "" {
		clientIP = r.RemoteAddr
	}

	logger.Info("request", "ip", clientIP, "agent", r.UserAgent())
	fmt.Fprintf(w, `<html>
		<body style="background:#111; color:#0f0; font-family:monospace; display:flex; flex-direction:column; align-items:center; justify-content:center; height:100vh; text-align:center;">
			<h1>NODE</h1>
			<p>Connection Secure.</p>
			<p>Your IP: %s</p>
			<small>Hosted on Proxmox @ Home-lab</small>
		</body>
		</html>
	`, clientIP)
}
