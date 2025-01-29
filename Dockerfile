# Use a imagem oficial do Golang
FROM golang:1.23

# Defina o diretório de trabalho
WORKDIR /app

# Copie os arquivos do projeto
COPY . .

# Instale as dependências
RUN go mod tidy

# Compile a aplicação
RUN go build -o golabs cmd/main.go

# Comando para rodar a API
CMD ["./golabs"]