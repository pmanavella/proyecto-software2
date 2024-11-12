# Usa la imagen oficial de Node.js para construir el frontend
FROM node:18-alpine AS build

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos package.json y package-lock.json
COPY package*.json ./

# Instala las dependencias del proyecto
RUN npm install

# Copia el resto del código fuente
COPY . .

# Construye la aplicación de React para producción
RUN npm run build

# Usa una imagen ligera de Nginx para servir la aplicación construida
FROM nginx:alpine

# Copia los archivos estáticos del build de React al directorio que Nginx servirá
COPY --from=build /app/build /usr/share/nginx/html

# Expone el puerto 80
EXPOSE 80

# Comando para ejecutar Nginx
CMD ["nginx", "-g", "daemon off;"]
