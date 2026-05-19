FROM golang:1.26

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

ENV PATH="/root/go/bin:${PATH}"

EXPOSE 8080

CMD ["air"]