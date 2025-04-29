# Go tabanlı bir builder image'ı kullanıyoruz
FROM golang:1.23 AS builder

# Çalışma dizinini ayarla
WORKDIR /app

# Go mod dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./ 
RUN go mod download

# Uygulama kaynak kodunu kopyala
COPY . .

# Go uygulamanı derle
RUN go build -o app .

# Ubuntu 22.04 (veya başka bir uyumlu sürüm) kullanarak base image'ı değiştiriyoruz
FROM ubuntu:22.04

# Çalışma dizinini ayarla
WORKDIR /root/

# Gerekli bağımlılıkları yükle
RUN apt-get update && apt-get install -y \
    curl \
    openssh-client \
    libc6 \
    git \
    && apt-get clean 

# Builder aşamasından derlenmiş uygulamayı kopyala
COPY --from=builder /app/app . 
COPY path.txt . 

# Uygulamayı çalıştır
CMD ["./app"]
