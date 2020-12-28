FROM golang:1.15.6-alpine3.12 AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/main ./cmd/api/

FROM scratch
COPY --from=build /bin/main /bin/main

EXPOSE 8080

ENTRYPOINT ["/bin/main"]