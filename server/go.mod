module server

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-openapi/errors v0.19.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	gorm.io/driver/mysql v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.11
	helm.sh/helm/v3 v3.6.0 // indirect
	sigs.k8s.io/controller-runtime v0.9.0 // indirect
)
