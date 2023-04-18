package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"mtls-demo/grpc/client/rpc"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certificate, err := tls.LoadX509KeyPair("../../certs/client.crt", "../../certs/client.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("../../certs/ca.crt")
	if err != nil {
		panic(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	})

	conn, err := grpc.Dial("127.0.0.1:8000",
		grpc.WithTransportCredentials(creds))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := rpc.NewServerClient(conn)

	helloRespone, err := ServerClient.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloRespone, err)

	registerResponse, err := ServerClient.Register(context.Background(), &rpc.RegisterRequest{Name: "chihuo", Password: "123456"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(registerResponse, err)

}
