# Build the UI app (Reactjs) using node/npm
FROM node as appbuilder
WORKDIR /ui
COPY ui .
RUN ./build-app.sh

# Build the go service executable, embedding the UI app
FROM golang:alpine as gobuilder
WORKDIR /webapp
COPY --from=appbuilder /ui/app/build ui/app/build
COPY service service
COPY go.mod go.sum ./
RUN ./service/build-service.sh

# Create the final container image - just the built exe
FROM alpine
WORKDIR /app
COPY --from=gobuilder /webapp/service/webapp .
EXPOSE 8080
ENTRYPOINT [ "/app/webapp" ]
CMD []
