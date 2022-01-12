package main

import (
	"github.com/andreani-publico/go-platform/web"
	"github.com/customer-experience/go-template-repository/api"
)

// @title Andreani Go Api
// @version 1.0
// @description API Restful to register shipments
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v2

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	server := web.NewServer(web.ReadConfigFromEnv())
	server.AddMetrics()
	server.AddCorsAllOrigins()
	server.AddHealth(web.HealthAlwaysUp)
	server.AddApiDocs()
	// go get -u github.com/swaggo/swag/cmd/swag
	// swag init
	api.SetupRouter(server.GetRouter())
	server.ListenAndServe()
}

// export GO111MODULE=on
