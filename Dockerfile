FROM golang:1.16-buster AS build

WORKDIR /app

COPY *.go /app

RUN go build -o /http main.go


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /http /http 

ENTRYPOINT ["/http"]
