package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nikola-enter21/wms/backend/cmd/grpcservice"
	"github.com/nikola-enter21/wms/backend/cmd/httpgateway"
	"github.com/nikola-enter21/wms/backend/database/model"
	"github.com/nikola-enter21/wms/backend/database/postgres"
	fakesender "github.com/nikola-enter21/wms/backend/integrations/sender/fake"
	"github.com/nikola-enter21/wms/backend/logging"
)

const (
	httpPort = "8090"
	grpcPort = "8080"
)

var (
	log = logging.MustNewLogger()
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	postgresDSN := fmt.Sprintln("host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	db := postgres.MustNew(postgresDSN)
	err := db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderLine{}, &model.Invoice{})
	if err != nil {
		panic(err)
	}

	grpcServer := &grpcservice.Server{
		DB:          db,
		EmailSender: fakesender.NewFakeSender("test@gmail.com"),
	}

	log.Infow("gRPC server", "listening at", grpcPort)
	log.Infow("http server", "listening at", httpPort)

	go httpgateway.Serve(ctx, fmt.Sprintf(":%s", httpPort), fmt.Sprintf(":%s", grpcPort))
	grpcServer.ServeForever(fmt.Sprintf(":%s", grpcPort))
}
