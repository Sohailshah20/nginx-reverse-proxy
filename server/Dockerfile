FROM golang:1.22-bullseye as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-server

FROM scratch

COPY --from=builder /go-server .

EXPOSE 8080

CMD ["./go-server"]
