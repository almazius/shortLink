FROM golang:latest
LABEL authors="sigy"

WORKDIR /links

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o runner api/cmd/main.go

CMD [". /runner"]
#CMD ["go", "run", "api/cmd/main.go"]
