package main

import (
	"SIEMSEasy/internal/database"
	"SIEMSEasy/internal/realtime"
	"context"
	"fmt"
	"log"
)

func main() {

	var queue realtime.RingBuffer
	queue.Head = 0
	queue.Tail = 0

	ctx := context.Background()

	conn, err := database.DatabaseConnect()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query(ctx, "SELECT @@version")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
	}

}
