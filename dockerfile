# FROM golang:1.18 as build

# WORKDIR /app

# COPY . .

# RUN go mod download

# RUN go build -o account-bff ./cmd/main.go

# FROM alpine:latest

# COPY --from=build /app/account-bff /app/account-bff
# COPY --from=build /app/.env /app/.env

# RUN chmod +x app/account-bff

# CMD ["./app/account-bff"]
FROM golang:1.18 as build

WORKDIR /app
COPY . .
# downloads the app dependencies and prints the progress to standard out.
RUN go mod download 
RUN go build ./cmd/main.go
# Expose necessary port
EXPOSE 4000
EXPOSE 5000

CMD [ "./app" ]