setup:
	docker pull envoyproxy/nighthawk-dev; cd cmd; go mod tidy;
run:
	go run cmd/main.go
