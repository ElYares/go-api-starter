# syntax=docker/dockerfile:1

# Etapa 1: build de la aplicación
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copiar go.mod y go.sum y descargar dependencias
COPY go.mod ./
RUN go mod download

# Copiar todo el proyecto
COPY . .

# Compilar el binario
RUN go build -o main ./cmd/server

# Etapa 2: contenedor final, más pequeño
FROM alpine:latest

WORKDIR /app

# Copiar binario compilado
COPY --from=builder /app/main .

# Copiar archivos de configuración si es necesario
COPY .env .env

# Exponer el puerto que usa tu app
EXPOSE 8080

# Comando por defecto
CMD ["./main"]

