package main

import (
	"context"
	"log"
	"net/rpc"

	"github.com/dkmelnik/GophKeeper/configs"
	rpcrouter "github.com/dkmelnik/GophKeeper/internal/delivery/rpc"
	userHR "github.com/dkmelnik/GophKeeper/internal/delivery/rpc/user"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/database"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/jwt"
	mongorp "github.com/dkmelnik/GophKeeper/internal/infrastructure/persistence/mongo"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/server"
	"github.com/dkmelnik/GophKeeper/internal/usecase/binary"
	"github.com/dkmelnik/GophKeeper/internal/usecase/card"
	"github.com/dkmelnik/GophKeeper/internal/usecase/text"
	"github.com/dkmelnik/GophKeeper/internal/usecase/user"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// CONFIG
	conf, err := configs.NewServer()
	if err != nil {
		return err
	}
	// CONFIG

	// MONGO
	mongoClient, err := database.NewMongoConnection(conf.MongoDSN)
	if err != nil {
		return err
	}
	defer mongoClient.Disconnect(context.Background())
	dbname := mongoClient.Database(conf.MongoDB)
	// MONGO

	// REPOS -----------------------
	userRepo := mongorp.NewUserRepository(dbname)
	textRepo := mongorp.NewTextRepository(dbname)
	cardRepo := mongorp.NewCardRepository(dbname)
	binaryRepo := mongorp.NewBinaryRepository(dbname)
	// REPOS -----------------------

	// INFRASTRUCTURE -----------------------
	jwtSVC := jwt.NewJwt(conf.Secret, conf.ExpDuration)
	userMW := userHR.NewMiddleware(jwtSVC)
	// INFRASTRUCTURE -----------------------

	// UC -----------------------
	userUC := user.NewUseCase(userRepo, jwtSVC)
	textUC := text.NewUseCase(textRepo)
	cardUC := card.NewUseCase(cardRepo)
	binaryUC := binary.NewUseCase(binaryRepo)
	// UC -----------------------

	// RPC
	rpcServer := rpc.NewServer()
	if err = rpcrouter.Register(rpcServer, userMW, userUC, textUC, cardUC, binaryUC); err != nil {
		return err
	}
	// RPC

	// TCP
	tcp, err := server.NewTCPServer(conf.ADDR)
	if err != nil {
		return err
	}
	// TCP
	log.Println("listener run on port: ", conf.ADDR)
	rpcServer.Accept(tcp.Listener())

	return nil
}
