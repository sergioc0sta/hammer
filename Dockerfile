FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /out/hammer ./cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/hammer /usr/local/bin/hammer

ENTRYPOINT ["/usr/local/bin/hammer"]
