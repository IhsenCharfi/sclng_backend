module github.com/IhsenCharfi/sclng_backend

go 1.21.5

replace (
	github.com/IhsenCharfi/sclng_backend/configuration => ./configuration
	github.com/IhsenCharfi/sclng_backend/internalHandlers => ./internalHandlers
	github.com/IhsenCharfi/sclng_backend/models => ../models
	github.com/IhsenCharfi/sclng_backend/utils => ../utils

)

require (
	github.com/Scalingo/go-handlers v1.8.1
	github.com/Scalingo/go-utils/logger v1.2.0
	github.com/google/go-github/v32 v32.1.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.17.0
)

require (
	github.com/Scalingo/go-utils/crypto v1.0.0 // indirect
	github.com/Scalingo/go-utils/errors/v2 v2.3.0 // indirect
	github.com/Scalingo/go-utils/security v1.0.0 // indirect
	github.com/gofrs/uuid/v5 v5.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/urfave/negroni v1.0.0 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/errgo.v1 v1.0.1 // indirect
)
