SHELL=/bin/bash

requests=47

.PHONY: docker-build
docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial-go
	docker rmi -f trial-java-platform
	docker rmi -f trial-java-virtual
	cd ./go ; docker build -t trial-go .
	cd ./java-platform ; docker build -t trial-java-platform .
	cd ./java-virtual ; docker build -t trial-java-virtual .

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

.PHONY: go-rate-limit-test
go-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8080 -requests=$(requests)

.PHONY: java-platform-rate-limit-test
java-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8081 -requests=$(requests)

.PHONY: java-virtual-rate-limit-test
java-rate-limit-test:
	go run utils/rateLimit.go -debug -host=http://localhost:8081 -requests=$(requests)

.PHONY: clean
clean:
	rm -rf ./visualization/visualization.io
	rm -rf ./go/trial
	cd ./java-platform ; mvn clean
	cd ./java-virtual ; mvn clean
