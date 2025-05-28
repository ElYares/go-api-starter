# Go API Starter

Starter profesional para construir APIs en Go usando el router **Chi**, estructura modular, y soporte para configuración por `.env`. Ideal para proyectos con arquitectura de microservicios.

---

## Estructura del Proyecto
<pre>
├── cmd/
│ └── server/
│       └── main.go # Punto de entrada del servidor
├── internal/
│ ├── config/
│ │     └── config.go # Carga y acceso a variables de entorno
│ └── handler/
│       └── ping.go # Handler para ruta de prueba /ping
├── Dockerfile # Imagen multietapa para desarrollo y producción
├── go.mod # Definición del módulo y dependencias
├── go.sum # Verificación de integridad de módulos
├── LICENSE # Licencia MIT
└── .env # Variables de entorno (no se sube al repo)
</pre>
---

## Ruta de prueba

http:localhost:8080/ping

---

## Ejecutar localmente

### 1. Requisitos

- Go 1.24+
- Docker (opcional)
- `.env` basado en `.env.example`

### 2. Modo local

```bash
go run cmd/server/main.go
```

### 3. Modo Docker

#### Construir la Imagen
```bash
docker build -t go-api-starter .

```

#### Correr Contenedor
```bash
docker run -p 8080:8080 --env-file .env go-api-starter
```
---

@author Arturo Yared Elizondo Regino
