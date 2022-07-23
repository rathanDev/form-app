Candidate Name: Janarthan

# Start server
cd cd .\server\interview-accountapi\          (Navigate to api directory)
docker-compose up                             (Start server)  

# Test 
cd client                                     (Navigate to operation directory)
go test -v .\operation\                       (Run unit tests)    
go run .                                      (Run main)