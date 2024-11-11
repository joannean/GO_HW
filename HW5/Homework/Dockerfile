FROM golang:1.23

WORKDIR /class04

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY handlers.go handlers.go
COPY errors.go errors.go
COPY models.go models.go
COPY config.yaml config.yaml

RUN go mod tidy  # 下載 viper 和其他依賴
RUN go build -v -o GO_HW

ENTRYPOINT ["./GO_HW"]
