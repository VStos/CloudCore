package main

import (
	"time"

	"github.com/VStos/CloudCore/grpc/server"
	"github.com/VStos/CloudCore/tools"
)

func main() {
	// TODO: инициализировать объект конфига //

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения

	var grpcServerOption *server.Options
	grpcServerOption.Init()

	tools.Rest(time.Minute)
}
