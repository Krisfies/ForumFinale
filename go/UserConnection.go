package main

import "fmt"

// Fonction assignant un utilisateur de la base de donnée à l'utilisateur du site
func UserConnection(usr Data_User) {
	rows, err := db.Query("SELECT * FROM `data_user` WHERE `Name` = ?", usr.Name)
	if err != nil {
		fmt.Println("UserConnection error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.Admin, &usr.Date, &usr.Biography); err != nil {
			fmt.Println("UserConnection error: ", err)
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Println("UserConnection error: ", err)
	}
}
