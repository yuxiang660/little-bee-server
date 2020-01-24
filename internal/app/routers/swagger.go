/* 
Usage:
	go get -u github.com/swaggo/swag/cmd/swag
	swag init -g ./internal/app/routers/swagger.go -o ./docs
	http://localhost:8181/swagger
*/

package routers

// @title Little Bee Server
// @version 0.1.0
// @description Restful API description about little bee server
// @schemes http https
// @host 127.0.0.1:8181
// @basePath /
// @contact.name Little Bee
// @contact.email yuxiang660@gmail.com