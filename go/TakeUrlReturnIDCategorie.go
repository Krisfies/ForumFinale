package main

import (
	"fmt"
	"strconv"
)

// Fonction qui prend un url et retourne l'ID de la catégorie correspondante
func TakeUrlReturnIDCategorie(url string) int64 {
	var IDstr string
	var Idint int64
	for _, letter := range url {
		if letter >= 48 && letter <= 57 {
			IDstr += string(letter)
		}
	}
	Idint, err := strconv.ParseInt(IDstr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return Idint
}
