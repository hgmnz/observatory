package main

import (
	"time"
	"fmt"
)

type observation struct {
	databaseUrl string
	connections  int
	updatedAt   time.Time
}

func (observation observation) display() {
	fmt.Println("=== ", observation.databaseUrl)
	fmt.Println("Connections: ", observation.connections)
	fmt.Printf("(%s)\n\n", observation.updatedAt)
}

func (observation observation) persist() {
	fmt.Println("Persisting observation:")
	observation.display()
	time.Sleep(1 * time.Second)
}

