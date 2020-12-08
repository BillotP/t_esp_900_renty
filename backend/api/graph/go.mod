module github.com/BillotP/t_esp_900_renty/v2/backend/api/graph

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/getsentry/sentry-go v0.8.0
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/rs/xid v1.2.1
	github.com/stretchr/testify v1.5.1
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)

replace github.com/99designs/gqlgen => github.com/j75689/gqlgen v0.11.4-0.20200428101547-3390d0df18d8
