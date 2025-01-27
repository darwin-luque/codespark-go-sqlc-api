FROM golang:1.21.3-alpine AS builder

WORKDIR /go/app
COPY go.* ./
ENV GOPROXY https://proxy.golang.org,direct
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN GOOS=linux go build -o /app.bin

FROM scratch

WORKDIR /app
COPY --from=builder /app.bin .
EXPOSE 9635

CMD [ "/app/app.bin" ]
