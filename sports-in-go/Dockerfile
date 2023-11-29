# Usa una imagen de Go como base
FROM golang

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos locales al contenedor
COPY . .

# Compila el programa Go
RUN go build -o app

# Ejecuta el programa al iniciar el contenedor
CMD ["./app"]