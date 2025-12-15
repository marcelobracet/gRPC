package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/marcelobracet/grpc/internal/database"
	"github.com/marcelobracet/grpc/internal/pb"
	"github.com/marcelobracet/grpc/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	// added new change
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := services.NewCategory(*categoryDB)

	gRPCserver := grpc.NewServer()
	pb.RegisterCategoryServiceServer(gRPCserver, categoryService)
	reflection.Register(gRPCserver)

	addr := ":50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen on %s: %v", addr, err)
	}

	log.Printf("gRPC server listening on %s", addr)
	if err := gRPCserver.Serve(listener); err != nil {
		log.Fatalf("serve gRPC: %v", err)
	}
}
