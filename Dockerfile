FROM golang:1.19 AS build

WORKDIR /app

RUN go mod init sonarqube-ci

COPY . .

RUN go build -o main && \
    go clean -cache

FROM scratch

COPY --from=build /app/main /main

USER nonroot

CMD ["/main"]

HEALTHCHECK --interval=5s --timeout=3s \
  CMD wget -qO- localhost:8080 || exit 1
