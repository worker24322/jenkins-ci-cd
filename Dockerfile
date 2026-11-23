# Bước 1: biên dịch
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /api .

# Bước 2: image chạy thật, chỉ chứa file đã biên dịch
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /api /api
EXPOSE 8081
ENTRYPOINT ["/api"]
