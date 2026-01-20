package main

import (
	"bufio"
	"context"
	"control_users/config"
	"control_users/db"
	"control_users/repository/postgres"
	"control_users/service"
	"control_users/ui"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("erroc config, %v", err)
	}
	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("db connection err: %v", err)
	}
	ctx := context.Background()
	err = postgres.CreateTableForUsers(ctx, db)
	if err != nil {
		log.Fatal("error created users table: %v", err)
	}
	repo := postgres.NewDBUserRepositore(db)
	scanner := bufio.NewScanner(os.Stdin)
	service := service.NewUserService(repo)
	//arrays users
	ui.Menu(service, scanner)
}
