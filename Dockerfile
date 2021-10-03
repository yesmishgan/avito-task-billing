FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito-billing ./cmd/main.go

CMD ["./avito-billing"]