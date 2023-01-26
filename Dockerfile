FROM golang:1.19-alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dest/server

FROM golang:1.19-alpine
WORKDIR /app
COPY --from=build /app/dest/server /app/server
ADD .env .env
LABEL maintainer="Julian<PrinzJuliano@users.noreply.github.com>"
CMD ["./server"]