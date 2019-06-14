package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	s1 := "15:00"
	s2 := "19:00"

	d1, err := time.Parse("15:04", s1)
	if err != nil {
		log.Fatal(err)
	}
	d2, err := time.Parse("15:04", s2)
	if err != nil {
		log.Fatal(err)
	}
}
