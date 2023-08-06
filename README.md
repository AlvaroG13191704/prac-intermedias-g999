# prac-intermedias-g999

## Para crear un workspace
```bash
# Correr en la carpeta del proyecto
yarn init 

# crear directorios
mkdir server
mkdir client

```

## Iniciar backend - golang
```bash
# Correr en la carpeta del proyecto
cd server
go mod init <nombre>
# instalar dependencias
go get github.com/gofiber/fiber/v2
# correr 
go run server.go
```

## Iniciar frontend - sveltekit
```bash
# Correr en la carpeta del proyecto
cd client
# instalar sveltekit
npm create svelte@latest
# instalar dependencias
npm install
# correr
npm run dev -- --open
```

## unir todo en el workspace
```json
  "scripts": {
    "dev:server":" cd ./server && go run server.go"
  }
```
