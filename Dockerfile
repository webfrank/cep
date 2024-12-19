FROM golang:1.23

WORKDIR /src

COPY cmd ./cmd
COPY internal ./internal

COPY go.mod .
COPY go.sum .

RUN go mod download 

RUN CGO_ENABLED=0 go build -o cep ./cmd/main.go

FROM scratch

WORKDIR /app

COPY plugins ./plugins
COPY --from=0 /src/cep /app/cep

EXPOSE 50051
EXPOSE 50052
EXPOSE 3322
EXPOSE 3320
EXPOSE 9092

CMD ["/app/cep"]