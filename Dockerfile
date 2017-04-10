# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/gcristofol/connect-mock

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get gopkg.in/gin-gonic/gin.v1
RUN go install github.com/gcristofol/connect-mock

COPY CheckoutResponse /tmp
COPY OrderCreatedResponse /tmp
COPY TicketTypesResponse /tmp

RUN ls -lisa /go/bin

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/connect-mock

# Document that the service listens on port 8080.
EXPOSE 8080
