module github.com/devminnu/weatherapp/location

go 1.15

require (
	github.com/akath19/gin-zap v0.0.0-20180806194049-1d9e5171c53e
	github.com/devminnu/weatherapp/pb v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/urfave/negroni v1.0.0
	go.mongodb.org/mongo-driver v1.4.4
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.34.0
)

replace (
	github.com/devminnu/weatherapp/location/router 	=> /Users/minhaj/go/src/apps/weatherapp/location/router
	github.com/devminnu/weatherapp/location/config => /Users/minhaj/go/src/apps/weatherapp/location/config
	github.com/devminnu/weatherapp/location/db => /Users/minhaj/go/src/apps/weatherapp/location/db
	github.com/devminnu/weatherapp/location/logger => /Users/minhaj/go/src/apps/weatherapp/location/logger
	github.com/devminnu/weatherapp/location/service => /Users/minhaj/go/src/apps/weatherapp/location/service
	github.com/devminnu/weatherapp/pb => /Users/minhaj/go/src/apps/weatherapp/pb
)
