package db

import (
	"database/sql"
	"log"
)

type User struct {
	Id     int
	Name   string
	Salary int
}

func getEmployeeByID(userID int) (*User, error) {
	query := "SELECT Id, Name, Salary FROM employee WHERE id = ?"

	var user User
	err := DB.QueryRow(query, userID).Scan(&user.Id, &user.Name, &user.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 用户不存在
		}
		log.Println("Failed to get user:", err)
		return nil, err
	}
	return &user, nil
}
