<p align="center">
  <img
  src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg"
		width="200" alt="Nest Logo"
	/>
</p>

# Dio Planner Back-End

Repositorio con el código del back-end del proyecto dio-planner de la materia de Aplicaciones móviles, este proyecto esta hecho con Go usando el Framework / Librería de Iris.

## Requisitos del proyecto

Para poder ejecutar el proyecto primero debes tener una instancia de base de datos con PostgreSQL, crear una base de datos llamada `dioPlanner` después copias el contenido del archivo `.env.example` en un archivo `.env` y agregas la cadena de conexión correspondiente a tu base de datos.

Después instalas las dependencias del proyecto con el comando:

```bash
go download
```

Posteriormente ejecutas el proyecto, cuando es la primera vez Gorm ejecutara las migraciones correspondientes para crear las tablas en la base de datos.

```bash
go run main.go
```
