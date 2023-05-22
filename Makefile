GO ?= go 

TARGET ?= main

GOFLAGS ?= -v 

.PHONY: build
build:
	$(GO) build $(GOFLAGS) -o $(TARGET) .


.PHONY: run
run:
	./$(TARGET)


.PHONY: lint
lint:
	gofmt -s -w .


.PHONY: docker-build 
docker-build:
	docker build -t go-logs-app . 


.PHONY: docker-run 
docker-run: 
	docker run --name logs-monitoring --rm -p 80:80 go-logs-app 
	
.PHONY: docker-compose
docker-compose:
	docker-compose up 

.PHONY: clean 
clean:
	$(GO) clean $(GOFLAGS); rm -f $(TARGET)