FROM golang:alpine AS build_base
RUN apk add --no-cache git
WORKDIR /build
COPY . /build
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -x -installsuffix cgo -o stitch .

FROM alpine:latest
WORKDIR /app
COPY --from=build_base /build/stitch .
EXPOSE 8000
ENTRYPOINT ["./stitch"]

