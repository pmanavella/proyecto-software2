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

-- Crear la tabla intermedia inscripción 
CREATE TABLE IF NOT EXISTS inscripción (
    user_id BIGINT NOT NULL,
    course_id BIGINT NOT NULL,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, course_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE);

-- Crear índices para optimizar las consultas
CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_users_email ON users (email);

-- Insertar datos iniciales
INSERT INTO users (username, password, nombre, apellido, email, admin) VALUES
('admin', '21232f297a57a5a743894a0e4a801fc3', 'Admin', 'User', 'admin@example.com', TRUE),
('pmanavella', 'dc9f4d858cff6dc1960d5292eecbf9df', 'Pilar', 'Manavella', 'Pmanavella@gmail.com', FALSE),
('vsponton', '77bfd8ff0493ccaf9f598975b4748d34', 'Victoria', 'Sponton', 'Vsponton@gmail.com', FALSE),
('valentinaC', 'd70932b58cc9bc3eb9f97ffb24dfa485', 'Valentina', 'Cervellini', 'valentinacervellini@gmail.com', FALSE),
('JuanG', '152cc4bc44aadf123c0f35449aff336e', 'Juan', 'Gutierrez', 'juangutierrez@gmail.com', FALSE),
('LucasM', 'df9dc9b984b278ec6880df2c0ab731aa', 'Lucas', 'Mendez', 'lucasmendez@gmail.com', FALSE); 

-- armar el insert into de inscripción para que me muestre por ejemplo, user 1 se inscribió en curso 1, user 2 se inscribió en curso 2, etc.
 INSERT INTO inscripción (user_id, course_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4); 
