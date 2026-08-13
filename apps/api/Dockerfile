# Production image for the OpenSwim API.
# Multi-stage build; runtime is intentionally minimal.
#
# HEALTHCHECK: do not probe stub /livez until E12-F5-T3 implements real
# liveness semantics (see docs/adr/0002-api-observability.md).

FROM golang:1.22-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/api /app/api
COPY migrations /app/migrations
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/api"]
