# rastwit
Rest App with Golang

# Init a Golang API

`$ go mod init github.com/rasjsus/rastwit`

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

- Get token \ 
`$ curl -X POST -H "Content-Type: application/json" -d '{"email": "jesuslz@gmail.com", "password": "1234567"}' http://localhost:8080/login` 
- Export token as environment variable \
`$ export TOKEN=value` 
- Creating a User \
`$ curl -X POST -H "Content-Type: application/json" -d '{"nam": "linuxize", "email": "linuxize@example.com", "password":"12asdfsdfassdf"}' -H "Authorization: {$TOKEN}" http://localhost:8080/register` 
- Creating a Tweet \
`$ curl -v -X POST -H "Content-Type: application/json" -d '{"message": "first tweet using curl"}' -H "Authorization: Bearer$TOKEN" http://localhost:8080/tweet`
- Get Tweets \
`$ curl -v -X GET -H "Content-Type: application/json" -H "Authorization: Bearer$TOKEN" -G http://localhost:8080/tweets -d 'id=61a7da9bf730db9ef9f0f7d1' -d 'page=1'`


# Connecting Mongo DB Locally

`$ docker exec -it mongojlz bash` \
`$ mongo -u "mongoadmin" -p "secret" --authenticationDatabase "admin"`

# Creating a user on admin database
`$ use admin`
`$ db.createUser({user:"mongoadmin", pwd:"secret", roles:[{role:"root", db:"rastwitdb"}]`

# Swagger Documentation

Pendiente por desarrollar