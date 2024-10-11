package connection

import (
	"database/sql"
	"fmt"
	"inventory-service/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func Database() *sql.DB {
	c := config.Configuration()

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBname))
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil
	}
	return db
}
