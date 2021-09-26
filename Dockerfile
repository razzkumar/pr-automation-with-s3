FROM golang:alpine as builder
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o /s3 .

## copy only build file
FROM node:14-alpine

LABEL maintainer="razzkumar <razzkumar.dev@gmail.com>"
LABEL version="1.0.1"
LABEL repository="https://github.com/razzkumar/pr-automation-s3-utils"

LABEL "com.github.actions.name"="PR Automation"
LABEL "com.github.actions.description"="Deploy each PR to s3 bucket by create \
  new s3 bucket and comment url to the PR and delete s3 after merge"
LABEL "com.github.actions.icon"="upload-cloud"
LABEL "com.github.actions.color"="green"

COPY --from=builder /s3 /
COPY ./entrypoint.sh /

# Command to run when starting the container
CMD ["/entrypoint.sh"]
