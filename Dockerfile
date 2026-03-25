From golang:1.26.1 as base

WORKDIR /app

Copy go.mod .

Run go mod download

Copy . .

Run go build -o main .

#final stage - Distroless image
From gcr.io/distroless/base

Copy --from=base /app/main .

Copy --from=base /app/static ./static

Expose 8080

CMD ["./main"]