# webapp

## Build Steps

### Build the app

```bash
./ui/build-app-docker.sh
```

### Build the service

```bash
./service/build-service-docker.sh
```

### Build everything in one go

```bash
./build-docker.sh
```

## Run the built service

```bash
./service/webapp
```

The service runs on [http://localhost:8080/](http://localhost:8080/).

## Run the app development server

```bash
./ui/start-app-docker.sh
```

The development server runs on [http://localhost:3000/app](http://localhost:3000/app).

## Recreate react app from fresh

```bash
rm -rf ./ui/app/
./ui/create-app-docker.sh
```

Edit the file ./ui/app/package.json to set the `/app` path prefix to the application.<br>
Add the setting...

```json
  "homepage": "/app",
```

## Build Automation

### GitHub Action - `Build and Publish Container`

The GitHub action [`build-and-publish-container.yml`](.github/workflows/build-and-publish-container.yml) will build the container image and publish to DockerHub.

The image will be tagged as...
* branch name
* `vX.Y.Z` for git tags conforming to semver notation
* `latest` for commits to the default branch of the repo (`main`)
