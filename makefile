BINARY_NAME=k8s_dashboard
OUTDIR=bin

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
	echo "Building for linux and windows"
	GOOS=linux GOARCH=386 go build -o ${OUTDIR}/${BINARY_NAME} main.go
	GOOS=windows GOARCH=386 go build -o ${OUTDIR}/${BINARY_NAME}_win main.go

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