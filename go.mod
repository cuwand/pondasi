module github.com/cuwand/pondasi

go 1.15

replace github.com/cariapo/cService => ../../@cariapo/backend/cservice

require (
	github.com/MicahParks/keyfunc v1.9.0 // indirect
	github.com/bpdlampung/banklampung-core-backend-go v1.0.18
	github.com/cariapo/cservice v0.0.1 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/nats-io/nats.go v1.25.0 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/redis/go-redis/v9 v9.0.2
	github.com/rs/zerolog v1.26.1
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.6.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
