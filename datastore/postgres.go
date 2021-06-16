package datastore

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"gitlab.com/TheChertila/REST_API_TV_Market/config"
)

func NewDBConnection() *pgx.Conn {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
