#! /bin/bash
for i in {a,b,c}; do
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./service-$i/service-$i ./service-$i/main.go
    docker build -t poneding/telepresence-demo-service-$i ./service-$i -f ./service-$i/Dockerfile
    docker push poneding/telepresence-demo-service-$i
done