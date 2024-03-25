FROM golang:1.21.0

# Set destination for COPY
WORKDIR /app

ADD . /app/
# Download Go modules
COPY go.mod go.sum /app/

RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux go build .

EXPOSE 8000

CMD ["/app/apiService"]