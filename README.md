# webapp

## Build Steps

### Build the app

```bash
./build-app-docker.sh
```

### Build the service

```bash
./build-service-docker.sh
```

### Build everything in one go

```bash
./build-docker.sh
```

## Run the built service

```bash
./webapp
```

## Run the app development server

```
./start-app-docker.sh
```

## Recreate react app from fresh

```bash
rm -rf app
./create-app-docker.sh
```

Edit the file app/package.json to set the `/app` path prefix to the application.<br>
Add the setting...

```json
  "homepage": "/app",
```
