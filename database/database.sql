/* Script para crear la base de datos en PostgreSQL */
CREATE DATABASE dioPlanner;
USE dioPlanner;

-- Tabla events
CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  date TIMESTAMP NOT NULL,
  description TEXT,
  img VARCHAR(255),
  location VARCHAR(255),
  accessibility VARCHAR(255),
  topic VARCHAR(255),
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla grades
CREATE TABLE grades (
  id SERIAL PRIMARY KEY,
  opinion TEXT,
  grade INTEGER,
  event_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla guests
CREATE TABLE guests (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  event_id INTEGER NOT NULL,
  status VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla memories
CREATE TABLE memories (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  picture VARCHAR(255),
  event_id INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla purchases
CREATE TABLE purchases (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  price FLOAT,
  event_id INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla todos
CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  text VARCHAR(255),
  date TIMESTAMP,
  complete BOOLEAN,
  event_id INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

-- Tabla users
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  description TEXT,
  picture VARCHAR(255),
  social_login BOOLEAN,
  social_provider VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
