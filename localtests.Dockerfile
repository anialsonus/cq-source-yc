# syntax=docker/dockerfile:1

FROM golang:1.16
WORKDIR /app

# Copy YC provider
COPY . .

WORKDIR resources
RUN go test -c -o test
CMD ./test -test.v