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

Levantar base de datos con Docker:

```bash
docker compose up -d
```

Posteriormente ejecutas el proyecto, cuando es la primera vez Gorm ejecutara las migraciones correspondientes para crear las tablas en la base de datos.

```bash
go run main.go
```

(Opcional) Si quieres compilar el proyecto a una imagen usando Docker puedes usar el siguiente comando:

```bash
docker build -t dio-planner:1.0.0 .
```

Para ejecutar el contenedor de la imagen compilada puedes usar el siguiente comando:

```bash
docker container run -dp 4000:4000 `  
> --name dio-planner `
> --network dio-planner-back_dio_planner_network `
> --env-file .env `
> dio-planner:1.0.0
```
