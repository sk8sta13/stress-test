FROM golang:1.24 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /stresstest cmd/stresstest/main.go

FROM scratch
WORKDIR /
COPY --from=build /stresstest /stresstest
ENTRYPOINT ["/stresstest"]