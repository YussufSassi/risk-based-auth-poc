# Build backend
FROM golang:1.20.5-alpine as backend-builder

WORKDIR /usr/app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /usr/app/build/server .

EXPOSE 8080
ENTRYPOINT [ "./build/server" ]
