# Gunakan image base yang sesuai dengan aplikasi Anda
FROM golang:1.21.4-alpine AS builder

# Set default value ENV
ENV API_KEY="x-api-key"

ENV SERVER_HOST="localhost"

ENV SERVER_PORT="3030"

ENV MYSQL_PORT=3306

ENV LOG_LEVEL=INFO


# Set working directory di dalam container
WORKDIR /app

# Copy file-file yang diperlukan ke dalam container
COPY . .

# Build aplikasi
RUN go build -o schedule-api

# Expose port yang digunakan oleh aplikasi
EXPOSE 3030

# Command untuk menjalankan aplikasi ketika container dijalankan
CMD ["./schedule-api"]

