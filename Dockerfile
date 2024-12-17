FROM golang:1.23

COPY go.mod go.sum ./

COPY . .

RUN go build -o secretly

ENTRYPOINT ["./secretly"]
