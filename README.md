Candidate Name: Janarthan

### Start server

cd .\server\interview-accountapi\ (Navigate to api directory)
docker-compose up (Start server)

##### Test

cd client (Navigate to operation directory)

# Run unit tests

cd src
go test -v .\operation\



# Run main
set baseUrl
cd src
go run .



##### Run as service using docker

Navigate form3-app\client
docker build -t form3-client .
Remove container with name as form3-client usiing docker rm <containerId>
docker run --name form3-client -p 8085:8085 form3-client

# Create network
docker network inspect form3-network
docker network create form3-network

# Add containers to network

docker network connect form3-network interview-accountapi
docker network connect form3-network form3-client
