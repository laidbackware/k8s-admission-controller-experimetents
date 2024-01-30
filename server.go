
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	// "io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/klog/v2"

	controller "github.com/laidbackware/k8s-example-admission-controller/pkg/example-admission-controller"
	
)

const (
	port = "8080"
)

var (
	tlscert, tlskey string
)

func main() {
	klog.InitFlags(nil)

	flag.Parse()

	certs, err := tls.LoadX509KeyPair(tlscert, tlskey)
	if err != nil {
		klog.Errorf("Filed to load key pair: %v", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%v", port),
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
	}
	
	handler := controller.ExampleServerHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", handler.Validate)

	server.Handler = mux

	// start webhook server in new rountine
	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil {
			klog.Errorf("Failed to listen and serve webhook server: %v", err)
		}
	}()

	klog.Infof("Server running listening in port: %s", port)

	// listening shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	klog.Info("Received shutdown signal, shutting down webhook server gracefully...")
	server.Shutdown(context.Background())

}

func deserializeSvc() {
	
}

func init() {
	flag.StringVar(&tlscert, "tlsCertFile", "/etc/certs/tls.crt", "File containing the x509 Certificate for HTTPS.")
	flag.StringVar(&tlskey, "tlsKeyFile", "/etc/certs/tls.key", "File containing the x509 private key to --tlsCertFile.")
}