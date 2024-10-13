SHELL=/bin/bash

requests=48

.PHONY: docker-build
docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial-go
	docker rmi -f trial-java
	cd ./go ; docker build -t trial-go .
	cd ./java ; docker build -t trial-java .

.PHONY: compose-up
compose-up:
	docker compose up -d

.PHONY: compose-down
compose-down:
	docker compose down

.PHONY: build
build:
	$(MAKE) clean
	cd ./go ; go mod tidy ; go build -v .
	cd ./java ; mvn clean install

.PHONY: go-rate-limit-test
go-rate-limit-test:
	go run utils/rateLimit.go http://localhost:8080 $(requests)

.PHONY: java-rate-limit-test
java-rate-limit-test:
	go run utils/rateLimit.go http://localhost:8081 $(requests)

.PHONY: clean
clean:
	rm -rf ./go/trial
	cd ./java ; mvn clean
