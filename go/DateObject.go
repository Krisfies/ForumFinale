package main

import (
	"fmt"
	"time"
)

// Fonction enregistrant la date du jour dans un string puis le renvoie
func DateObject() string {
	var Date string
	DateNow := time.Now()
	Test := fmt.Sprintln(DateNow)
Loop:
	for _, letter := range Test {
		if letter == ' ' {
			break Loop
		}
		Date += string(letter)
	}
	return (Date)
}
