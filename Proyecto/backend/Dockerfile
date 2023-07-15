FROM golang:1.18

ADD . /backend

WORKDIR /backend
RUN go mod init backend
RUN go mod tidy
RUN go build -o backend .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/backend"]
