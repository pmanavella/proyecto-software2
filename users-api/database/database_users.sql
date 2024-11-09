-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS users_api;

-- Usar la base de datos
USE users_api;

-- Crear la tabla de usuarios
CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Crear índices para optimizar las consultas
CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_users_email ON users (email);

-- Insertar datos iniciales
INSERT INTO users (username, password, nombre, apellido, email, admin) VALUES
('admin', '$2a$10$7EqJtq98hPqEX7fNZaFWoOa8z8lF.5e4xK4j6h5h5h5h5h5h5h5h5', 'Admin', 'User', 'admin@example.com', TRUE),
('jdoe', '$2a$10$7EqJtq98hPqEX7fNZaFWoOa8z8lF.5e4xK4j6h5h5h5h5h5h5h5h5', 'John', 'Doe', 'jdoe@example.com', FALSE),
('asmith', '$2a$10$7EqJtq98hPqEX7fNZaFWoOa8z8lF.5e4xK4j6h5h5h5h5h5h5h5h5', 'Alice', 'Smith', 'asmith@example.com', FALSE);