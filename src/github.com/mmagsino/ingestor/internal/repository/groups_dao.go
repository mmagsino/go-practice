package repository

import (
	"time"
	"database/sql"
)

type Group struct {
	Id int32
	Name string
	Description string
	DateCreated time.Time
	DateUpdated time.Time
}

func FindById(id int32) (group Group) {
	conn, _ := getConnection()

	defer conn.Close()
	session := conn.Resource.(*sql.DB)
	rows, _ := session.Query(`SELECT ID, NAME, DESCRIPTION, DATE_CREATED, DATE_UPDATED FROM scheduler.SYSTEM_GROUP WHERE ID = ?`, id)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&group.Id,
			&group.Name,
			&group.Description,
			&group.DateCreated,
			&group.DateUpdated)
	}
	return group
}

func FindAll() (groups[] Group){
	conn, _ := getConnection()
	
	defer conn.Close()
	session := conn.Resource.(*sql.DB)
	rows, _ := session.Query(`SELECT ID, NAME, DESCRIPTION, DATE_CREATED, DATE_UPDATED FROM scheduler.SYSTEM_GROUP`)
	defer rows.Close()
	for rows.Next() {
		group := Group{};
		rows.Scan(
			&group.Id,
			&group.Name,
			&group.Description,
			&group.DateCreated,
			&group.DateUpdated)
		groups = append(groups, group)
	}
	return groups
}

func getConnection() (*DbSession, error) {
	mysql := MysqlClient{DBName:"scheduler", Ip:"127.0.0.1:3306", Password:"miguelmikael", Username:"root"}
	return mysql.Connect()
}
