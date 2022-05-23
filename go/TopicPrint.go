package main

import "fmt"

// Fonction affichant l'ensemble des topic d'une catégorie correspondante en les assignants à un tableau
func TopicPrint(WichCategorie int64, trier string) ([]Data_Topic, error) {

	// Gestion du trie en ordre croissant
	if trier == "Old" {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? ORDER BY `data_topic`.`ID` ASC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil

		// Gestion du trie de tous les topics contentant le tag Aide
	} else if trier == "Aide" {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? AND `IsAide` = 1  ORDER BY `data_topic`.`ID` DESC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil

		// Gestion du trie de tous les topics contentant le tag Bug
	} else if trier == "Bug" {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? AND `IsBug` = 1 ORDER BY `data_topic`.`ID` DESC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil

		// Gestion du trie de tous les topics contentant le tag Boss
	} else if trier == "Boss" {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? AND `IsBoss` = 1 ORDER BY `data_topic`.`ID` DESC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil

		// Gestion du trie de tous les topics contentant le tag Lore
	} else if trier == "Lore" {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? AND `IsLore` = 1 ORDER BY `data_topic`.`ID` DESC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil

		// Gestion du trie par ordre décroissant
	} else {
		var topics []Data_Topic

		rows, err := db.Query("SELECT * FROM `data_topic` WHERE `CategorieID` = ? ORDER BY `data_topic`.`ID` DESC ", WichCategorie)
		if err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var top Data_Topic
			if err := rows.Scan(&top.ID, &top.Name, &top.IsAide, &top.IsBug, &top.IsBoss, &top.IsLore, &top.CategorieID); err != nil {
				return nil, fmt.Errorf("TopicPrint error: %v", err)
			}
			top.AddTagsToTopic()
			topics = append(topics, top)
		}
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("TopicPrint error: %v", err)
		}
		return topics, nil
	}
}
