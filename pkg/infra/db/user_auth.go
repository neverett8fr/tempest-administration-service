package db

import (
	"fmt"
	"log"
	"tempest-administration-service/pkg/entities"

	"golang.org/x/crypto/bcrypt"
)

func (conn *DBConn) CreateUser(user entities.User) error {

	if !user.CheckValid() {
		err := fmt.Errorf("error creating user, entity not valid")
		log.Println(err)
		return err
	}

	id, err := conn.ReadUserID(user.Username)
	if err == nil {
		err := fmt.Errorf("error, expected no results / error - username must already exist")
		log.Println(err)
		return err
	}
	if id != "" {
		err := fmt.Errorf("error username exists")
		log.Println(err)
		return err
	}

	_, err = conn.Conn.Exec(fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2)", userTableUsers, userColumnUsername, userColumnPasswordHash), user.Username, user.PasswordHashed)
	if err != nil {
		err := fmt.Errorf("error inserting new user, err %v", err)
		log.Println(err)
		return err
	}

	return nil
}

func (conn *DBConn) ReadUserID(username string) (string, error) {

	row := conn.Conn.QueryRow(fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", userColumnID, userTableUsers, userColumnUsername), username)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error scanning id, err %v", err)
	}

	return id, nil
}

func (conn *DBConn) VerifyPassword(username string, password string) error {

	row := conn.Conn.QueryRow(fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", userColumnPasswordHash, userTableUsers, userColumnUsername), username)

	var passwordHash string
	err := row.Scan(&passwordHash)
	if err != nil {
		return fmt.Errorf("error scanning hash, err %v", err)
	}

	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
