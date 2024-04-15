FROM golang:1.20-alpine
WORKDIR /app
COPY . .

RUN go mod download
RUN go install github.com/google/wire/cmd/wire@latest

RUN cd internal/ && wire
RUN cd ../
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/main ./cmd/main.go

EXPOSE 5020
ENTRYPOINT ["/build/main"]