

# Multi stage build

# Build stage I : Go lang and Alpine Linux is only needed to build the program
FROM golang:1.11-alpine AS build
WORKDIR /src/
COPY hello-http.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/hello-http-server


# Build stage II : Go binaries are self-contained executables.
FROM scratch
COPY --from=build /bin/hello-http-server /bin/hello-http-server

ENTRYPOINT ["/bin/hello-http-server"]