FROM golang AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 go build -o app ./cmd/trie-maintainer/main.go



FROM alpine

WORKDIR /app

COPY --from=builder /build/app /app/

EXPOSE 4321

CMD [ "./app" ]