FROM golang:1.22-alpine

WORKDIR /app

# Instala dependências necessárias
RUN apk add --no-cache gcc musl-dev

# Copia os arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Compila a aplicação
RUN go build -o main .

# Expõe a porta da API
EXPOSE 8000

# Roda a aplicação
CMD ["./main"]