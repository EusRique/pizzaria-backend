# Use a imagem oficial do Go
FROM golang:1.21-alpine

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o código-fonte
COPY . .

# Baixar as dependências Go
RUN go mod tidy

# Expor a porta (caso o Go faça algo que precise expor)
EXPOSE 8080

# Comando para rodar o app Go
CMD ["tail", "-f", "/dev/null"]
