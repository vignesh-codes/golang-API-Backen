package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")
	db_user, ok := os.LookupEnv(("DB_USER"))

	if !ok {
		logger.Log(errors.New("DB_USER not found in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	db_password, ok := os.LookupEnv("DB_PASS")

	if !ok {
		logger.Log(errors.New("DB_PASS not found in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	db_host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(errors.New("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	connect_db := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime-True", db_user, db_password, db_host)

	database, err := sql.Open("mysql", connect_db)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Connect to DB")
	}

	_, err = database.Exec(`CREATE DATABASE new_db`)

	if err != nil {
		fmt.Println(err, "IGNORE IF EXISTS")
	}

	_, err = database.Exec(`USE new_db`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE todo_table (
		    id INT AUTO_INCREMENT,
		    task TEXT NOT NULL,
		    description TEXT NOT NULL,
			status TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err, "IGNORE IF EXISTS")
	}

	_, err = database.Exec(`
		SET AUTOCOMMIT = 1;
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database

}
