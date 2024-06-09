package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	// здесь мы испольщуем имя контейнера bd, в кастомной сети bridge, обращаясь к контейнеру по имени,
	// которое при помощи DNS преобразуется в ip адрес контейнера
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@db:5432/postgres")
	for i := 0; i < 3; i++ {
		conn, err = pgx.Connect(context.Background(), "postgres://postgres:1234@db:5432/postgres")
		if err == nil {
			break
		}
		fmt.Println("could not connect to database, restarting...")
		time.Sleep(time.Second)
	}
	if err != nil {
		panic(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("error while pinging database")
	} else {
		fmt.Println("database is running successfully!")
	}

}
