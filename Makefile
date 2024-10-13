SHELL=/bin/bash

.PHONY: docker-build
docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial-go
	docker rmi -f trial-java
	cd ./go ; docker build -t trial-go .
	cd ./java ; docker build -t trial-java .

.PHONY: build
build:
	$(MAKE) clean
	cd ./go ; go mod tidy ; go build -v .
	cd ./java ; mvn clean install

.PHONY: rate-limit-test
rate-limit-test:
	go run utils/rateLimit.go http://localhost:8080 1000

.PHONY: clean
clean:
	rm -rf ./go/trial
	cd ./java ; mvn clean
