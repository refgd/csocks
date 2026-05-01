FROM golang:1.25-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/csocks .


FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /out/csocks /app/csocks

EXPOSE 1080

ENTRYPOINT ["/app/csocks"]
CMD ["--listen","1080"]
