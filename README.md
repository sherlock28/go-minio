### Pasos para ejecutar la prueba de MinIO con la API Go


1. Levantar el servicio "minio" con el comando: 

```
docker-compose up -d minio 
```

2. Ejecutar los comandos de instalación de dependencias: 

```
go mod download
npm ci
```

3. En línea 33 del archivo ```test/Go-MinIO.postman_collection.json``` definir un path hacia un archivo local
de imagen.

4. Ejecutar el servidor con el comando:

```
npm run dev
```

5. Ejecutar el test con el comando:

```
npm run test
```