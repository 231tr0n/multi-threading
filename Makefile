SHELL=/bin/bash

requests=45

.PHONY: docker-build
docker-build:
	$(MAKE) build
	cd ./go ; docker build -t trial-go .
	cd ./java-platform ; docker build -t trial-java-platform .
	cd ./java-virtual ; docker build -t trial-java-virtual .
	cd ./java-dropwizard ; docker build -t trial-java-dropwizard .

.PHONY: compose-up
compose-up:
	docker compose up -d

.PHONY: compose-down
compose-down:
	docker compose down

.PHONY: build
build:
	$(MAKE) clean
	cd ./visualization ; go mod tidy ; go build -v .
	cd ./go ; go mod tidy ; go build -v .
	cd ./java-platform ; mvn clean install
	cd ./java-virtual ; mvn clean install
	cd ./java-dropwizard ; mvn clean install

.PHONY: go-rate-limit-test
go-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8080 -requests=$(requests)

.PHONY: java-platform-rate-limit-test
java-platform-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8081 -requests=$(requests)

.PHONY: java-virtual-rate-limit-test
java-virtual-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8082 -requests=$(requests)

.PHONY: java-dropwizard-virtual-rate-limit-test
java-dropwizard-virtual-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8083/virtual -requests=$(requests)

.PHONY: java-dropwizard-platform-rate-limit-test
java-dropwizard-platform-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8083/platform -requests=$(requests)

.PHONY: pkgsite
pkgsite:
	cd ./visualization ; go run golang.org/x/pkgsite/cmd/pkgsite@latest

.PHONY: clean
clean:
	rm -rf ./visualization/visualization.io
	rm -rf ./go/trial
	cd ./java-platform ; mvn clean
	cd ./java-virtual ; mvn clean
	cd ./java-dropwizard ; mvn clean
