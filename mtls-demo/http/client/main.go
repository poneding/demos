package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	caCert     = "../../certs/ca.crt"
	clientCert = "../../certs/client.crt"
	clientKey  = "../../certs/client.key"
)

func main() {
	pool := x509.NewCertPool()

	caCertBytes, err := os.ReadFile(caCert)
	if err != nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(caCertBytes)

	clientCertBytes, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			Certificates:       []tls.Certificate{clientCertBytes},
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: tr,
	}

	r, err := client.Get("https://127.0.0.1:8443") // server
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
