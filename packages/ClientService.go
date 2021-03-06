package services

import (
	"Golang/onlineBank/database/postgres"
	"Golang/onlineBank/models"
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/pgxpool"
)

func AddClient(name string, surname string, login string, password string, age int, gender string, phone string, db *pgxpool.Pool) (err error) {
	status := true
	password = MakeHash(password)
	//fmt.Println(name, surname, login, password,age,gender,phoneNumber,status)
	_, err = db.Exec(context.Background(), postgres.AddClient, name, surname, login, password, age, gender, phone, status)
	fmt.Println(postgres.AddClient)
	if err != nil {
		log.Fatalf("Пользователь недобавлен: %s", err)
		return err
	}
	return nil
}

func MakeHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func QueryError(text string) (err error) {
	return fmt.Errorf(text)
}

func GetAllClients(db *sql.DB) (clients []models.Client, err error) {
	rows, err := db.Query(postgres.GetAllClients)
	if err != nil {
		log.Fatalf("1 wrong")
		return nil, err
	}

	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			clients = nil
		}
	}()

	for rows.Next() {
		client := models.Client{}
		err = rows.Scan(&client.ID, &client.Name, &client.Surname, &client.Phone, &client.Login, &client.Password)
		if err != nil {
			log.Fatalf("2 wrong")
			return nil, err
		}
		clients = append(clients, client)
	}
	if rows.Err() != nil {
		log.Fatalf("3 wrong")
		return nil, rows.Err()
	}
	return clients, nil
}

func Login(login string, password string, db *pgxpool.Pool) (loginPredicate bool, err error) {
	var dbLogin, dbPassword string
	fmt.Println(login, password)
	err = db.QueryRow(context.Background(), postgres.LoginSQL, login).Scan(&dbLogin, &dbPassword)
	fmt.Println(dbLogin, dbPassword)
	fmt.Println(err)
	if err != nil {
		//		fmt.Printf("%s, %e\n", loginSQL, err)
		return false, err
	}
	err = QueryError("Несовпадение пароля")
	if MakeHash(password) != dbPassword {
		//fmt.Println(makeHash(password), " ", dbPassword)
		return true, err
	}
	//fmt.Println(makeHash(password), " ", dbPassword)
	return true, nil
}

func SearchByLogin(login string, db *pgxpool.Pool) (id int64, surname string) {
	err := db.QueryRow(context.Background(), postgres.SearchClientByLogin, login).Scan(&id, &surname)
	if err != nil {
		log.Fatalf("Ошибка в %s", postgres.SearchClientByLogin)
	}
	return id, surname
}
