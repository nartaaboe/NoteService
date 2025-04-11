# Project Setup
## Prerequisites
Before running the project, ensure you have the following installed:

`Go (version 1.18 or later)`

`Docker`

`Docker Compose`

## Run the following command to install the necessary Go dependencies:

`go mod tidy`

# Docker Setup
## Dockerfile
This project comes with a Dockerfile to build and run the application in a Docker container.

## Build and Run the Docker Container
#### Build the Docker image:

`docker build -t notes-microservice.`
#### Run the Docker container:

`docker run -p 8080:8080 notes-microservice`

This will start the Notes microservice on http://localhost:8080.

## API Endpoints
The API exposes the following endpoints:

#### 1. GET /notes
Retrieves all notes stored in memory.

Response:

Status: `200` OK

Body: List of notes in JSON format.

#### 2. POST /notes/create
Creates a new note. The request body must include the note details.

Request:

```
{
  "title": "Note title",
  "content": "Note content"
}
```

Response:

Status: `201` Created

Body: The created note in JSON format, including the ID.

