Candidate Name: Janarthan

### Build form3-client as docker image

cd client 
docker build -t form3tech-client .


### Update given docker-compose.yml
# Add clint to given docker compose 

  form3techclient:
    image: form3tech-client:latest
    container_name: form3tech-client
    depends_on:
      - accountapi
    restart: on-failure

# Add container_name to account api 

container_name: interview-accountapi

# Run docker-compose 

cd server/interview-accountapi
docker-compose config             
docker-compose up               



