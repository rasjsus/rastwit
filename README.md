# rastwit
Rest App with Golang

# Init a Golang API

`$ go mod init github.com/rasjsus/rastwit.git`

# Dependencies

`$ go get go.mongodb.org/mongo-driver/mongo` \
`$ go get go.mongodb.org/mongo-driver/mongo/options` \
`$ go get go.mongodb.org/mongo-driver/bson` \
`$ go get golang.org/x/crypto/bcrypt` \
`$ go get github.com/gorilla/mux` \
`$ go get github.com/rs/cors` \
`$ go get github.com/dgrijalva/jwt-go` \
`$ go mod vendor`

# Config Mongo DB Locally

`$ docker pull mongo` \
`$ docker run -d --name mongojlz -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo`  \
`$ docker stop mongojlz` \
`$ docker rm mongojlz` \
`$ docker exec -it mongojlz bash` \
`$ docker ps`

# Test endpoints
`export TOKEN=value` \
`curl -X POST -H "Content-Type: application/json" -d '{"nam": "linuxize", "email": "linuxize@example.com", "password":"12asdfsdfassdf"}' -H "Authorization: {$TOKEN}" http://localhost:8080/register`