FROM golang:1.23

RUN apt update && apt install libasound2-dev pkg-config -y

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /vogoder

CMD ["/vogoder"]
