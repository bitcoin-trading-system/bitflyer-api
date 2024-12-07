# ベースイメージを指定
FROM golang:1.23.2-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールと依存関係をコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .
