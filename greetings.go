package main

import (
	"math/rand"
	"time"
)

func GetGreetings(custom string) string {
	greetings := []string{
		"Hi there!",
		"Hey people.",
		"Hello team.",
		"Hey! May I ask your attention please!",
	}

	if custom != "" {
		return custom
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	randomIndex := r.Intn(len(greetings))
	return greetings[randomIndex]
}
