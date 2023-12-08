package main

import (
	"fmt"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent"
	"github.com/KasumiMercury/uo-patradb-dogtrot/ent/proto/entpb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func loadEnv() {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("failed to load .env file: %v", err)
		}
	}
}

func main() {
	loadEnv()

	usr := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	hst := os.Getenv("MYSQL_HOST")
	prt := os.Getenv("MYSQL_PORT")
	db := os.Getenv("MYSQL_DATABASE")
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", usr, pwd, hst, prt, db))
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed to close mysql connection: %v", err)
		}
	}(client)

	server := grpc.NewServer()

	// Initialize ent services
	vs := entpb.NewVideoService(client)
	ds := entpb.NewDescriptionService(client)
	pds := entpb.NewPeriodicDescriptionTemplateService(client)
	cds := entpb.NewCategoryDescriptionTemplateService(client)
	cd := entpb.NewChannelService(client)
	pcs := entpb.NewPatChatService(client)

	// Register ent services
	entpb.RegisterVideoServiceServer(server, vs)
	entpb.RegisterDescriptionServiceServer(server, ds)
	entpb.RegisterPeriodicDescriptionTemplateServiceServer(server, pds)
	entpb.RegisterCategoryDescriptionTemplateServiceServer(server, cds)
	entpb.RegisterChannelServiceServer(server, cd)
	entpb.RegisterPatChatServiceServer(server, pcs)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
