package main

import (
	"fmt"
	"strconv"
)

// Fonction qui prend un url et retourne l'ID d'un topic
func TakeUrlReturnIDTopic(url string) int64 {
	var LeIDstr string
	var LeIDint int64
	for _, letter := range url {
		if letter != '/' && letter != 't' && letter != 'o' && letter != 'p' && letter != 'i' && letter != 'c' {
			LeIDstr += string(letter)
		}
	}
	LeIDint, err := strconv.ParseInt(LeIDstr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return LeIDint
}
