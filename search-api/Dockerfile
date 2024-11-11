# Usa la imagen oficial de Golang
FROM golang:1.22-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el go.mod y go.sum (si existe) para aprovechar la caché de Docker
COPY go.mod ./
COPY go.sum ./

# Instala las dependencias
RUN go mod tidy

# Copia el resto del código fuente
COPY . .

# Compila la aplicación
RUN go build -o main .

# Expone el puerto que utilizará la aplicación (ajusta según tu configuración)
EXPOSE 8084

# Comando para ejecutar la aplicación
CMD ["./main"]
