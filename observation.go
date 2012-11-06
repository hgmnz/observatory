package main

import (
	"fmt"
	"time"
)

type observation struct {
	databaseUrl      string
	connections      uint16
	bytes            uint64
	serviceAvailable bool
	pgVersion        string
	updatedAt        time.Time
}

func (observation observation) display() {
	fmt.Println("=== ", observation.databaseUrl)
	fmt.Println("Connections: ", observation.connections)
	fmt.Println("Bytes: ", observation.bytes)
	fmt.Println("Service Available: ", observation.serviceAvailable)
	fmt.Printf("(%s)\n\n", observation.updatedAt)
}

func (observation observation) persist() {
	fmt.Println("Persisting observation:")
	observation.display()
	time.Sleep(1 * time.Second)
}
