package main

import "strconv"

// Fonction qui prend un ID et retourne l'url d'une catégorie
func TakeIDReturnUrlCategorie(ID int64) string {
	url := "D" + strconv.Itoa(int(ID))
	return url
}
