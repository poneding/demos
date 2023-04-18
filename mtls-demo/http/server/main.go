package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	caCert     = "../../certs/ca.crt"
	serverCert = "../../certs/server.crt"
	serverKey  = "../../certs/server.key"
)

type mtlsHandler struct {
}

func (m *mtlsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! ", time.Now())
}

func main() {
	pool := x509.NewCertPool()

	caCertBytes, err := os.ReadFile(caCert)
	if err != nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(caCertBytes)

	server := &http.Server{
		Addr:    ":8443",
		Handler: &mtlsHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert, // 需要客户端证书
		},
	}

	log.Println("server started...")
	log.Fatalln(server.ListenAndServeTLS(serverCert, serverKey))
}
