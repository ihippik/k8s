FROM golang:1.16-alpine AS build
WORKDIR /build
COPY . .

#Build back-end
WORKDIR /build
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app


FROM scratch as final
COPY --from=build /build/app /app
WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/app"]

