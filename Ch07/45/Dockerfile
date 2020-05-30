# first stage: build module image
FROM golang:1.14.2 as builder

# setting project direactory
WORKDIR /app

# copy module file first
COPY go.mod .

# download dependencies
RUN go mod download

# copy all files under porject directory
COPY . .

# build go files (using Go Constraints)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webservice

# second stage: Run built image
FROM scratch

# copy built files from first
COPY --from=builder /app/webservice /app/

# set port and entry point
EXPOSE 3000

ENTRYPOINT ["/app/webservice"]