BINARY_NAME=k8s_dashboard
OUTDIR=bin
UNAME_S := $(shell uname)

.PHONY:fmt
fmt:
	go fmt ./...

.PHONY:lint
lint: fmt
	golint ./...

.PHONY:vet
vet: lint
	go vet ./...

.PHONY:build
build: vet
ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
	@echo "Building for $(OS) platform"
	GOOS=windows GOARCH=amd64 go build -o ${OUTDIR}/${BINARY_NAME}_win main.go
else ifeq ($(UNAME_S),Linux)
	@echo "Building for $(UNAME_S) platform"
	GOOS=linux GOARCH=amd64 go build -o ${OUTDIR}/${BINARY_NAME} main.go
else ifeq ($(UNAME_S),Darwin)
	@echo "Building for $(UNAME_S) platform"
	GOOS=darwin GOARCH=amd64 go build -o ${OUTDIR}/${BINARY_NAME}_darwin main.go
else
	@echo "Platform $(UNAME_S) not supported"
endif

.PHONY:test
test: build
	go test -v ./...

.PHONY:test-cover
test-cover: test
	go test -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html 

.PHONY:run
run:
	./bin/${BINARY_NAME}

.PHONY:clean
clean:
	go clean
	if [ -d "${OUTDIR}" ]; then \
        rm -rf ${OUTDIR}; \
    fi \

all: clean fmt lint vet build run

go_run:	
	go run main.go