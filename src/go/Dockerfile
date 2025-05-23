FROM golang:1.22.5

WORKDIR /src/go

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 1234

CMD ["./main"]