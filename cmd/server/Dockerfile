FROM golang:alpine AS build

WORKDIR /app

# the module files are copied first so there is no need to re-download them if there are no changes.
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY . .
RUN make build

FROM alpine:latest
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/server .
COPY --from=build /app/cmd/server/entrypoint.sh .
COPY --from=build /app/config/local.yml ./config/
RUN ls -la
ENTRYPOINT ["/entrypoint.sh"]
