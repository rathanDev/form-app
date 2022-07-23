Candidate Name: Janarthan




### Start server
cd .\server\interview-accountapi\             (Navigate to api directory)
docker-compose up                             (Start server)  




### Test 
cd client                                     (Navigate to operation directory)
# Run unit tests
cd src
go test -v .\operation\                       

# Run main
cd src 
go run .