package main

import "192.168.205.151/vq2-go/go-template/cmd"

//	@title			Swagger Go Template Project API
//	@version		1.0
//	@description	This is a sample Go Template server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		petstore.swagger.io
// @BasePath	/v2
func main() {
	cmd.Execute()
}
