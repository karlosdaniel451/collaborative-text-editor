FROM golang:1.21rc2-alpine

WORKDIR /app

# # Copy and download dependency using go mod.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ## Install GCC to allow CGO to work in order to enable data race detector.
RUN apk add build-base

# Set envorinment variables.
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# Build go application with data race detector enabled.
RUN go build --race -o /backend

CMD [ "/backend" ]
