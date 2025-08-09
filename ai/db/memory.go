package db

import (
	_"database/sql"
)

func GetChatHistory() ([]string, error) {
	rows, err := DB.Query("SELECT message FROM chat_history ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []string
	for rows.Next() {
		var msg string
		if err := rows.Scan(&msg); err != nil {
			return nil, err
		}
		history = append(history, msg)
	}
	return history, nil
}

func SaveMessage(msg string) error {
	_, err := DB.Exec("INSERT INTO chat_history (message) VALUES (?)", msg)
	return err
}