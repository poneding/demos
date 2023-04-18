package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"mtls-demo/grpc/server/rpc"
	"net"
	"os"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {

	certificate, err := tls.LoadX509KeyPair("../../certs/server.crt", "../../certs/server.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("../../certs/ca.crt")
	if err != nil {
		panic(err)

	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		panic("AppendCertsFromPEM failed")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	rpc.RegisterServerServer(grpcServer, Server{})

	log.Println("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}
