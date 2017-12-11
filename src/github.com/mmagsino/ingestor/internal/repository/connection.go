package repository
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

type MysqlClient struct {
	Username string
	Password string
	DBName string
	Ip string
}

type MongoClient struct {
	Ip string
	DBName string
}

type EsClient struct {
	Ip string
	port int16
	Index string
}

type DbSession struct {
	Name string
	Type string
	IsClose bool
	Resource interface{}
}

type Connection interface {
	 Connect() (*DbSession, error)
}

type Resource interface {
	Close()
}

func (client MysqlClient) Connect() (*DbSession, error) {
	var dataSourceName string
	if len(client.Password) > 0 && len(client.Username) > 0 {
		dataSourceName = client.Username + ":" + client.Password + "@" + "tcp(" + client.Ip + ")/" + client.DBName + "?parseTime=true&charset=utf8"
	} else if len(client.Username) > 0 {
		dataSourceName = client.Username + "@" + "tcp(" + client.Ip + ")/" + client.DBName + "?parseTime=true&charset=utf8"
	}
	db, err := sql.Open("mysql", dataSourceName)
	return &DbSession{Name: client.DBName,  Type:"mysql", Resource: db, IsClose:false}, err
}

func (client MongoClient) Connect() (*DbSession, error) {
	connection, err := mgo.Dial(client.Ip)
	return &DbSession{Name:client.DBName, Type:"mongo", Resource:connection, IsClose:false}, err
}

func (client EsClient) Connect() (*DbSession, error) {
	return &DbSession{Name:client.Index, Type:"es", Resource:nil, IsClose:false}, nil
}

func (session *DbSession) Close() {
	if !session.IsClose {
		if session.Type == "mysql" {
			session.Resource.(*sql.DB).Close()
		} else if session.Type == "mongo" {
			session.Resource.(*mgo.Session).Close()
		}
		session.IsClose = true
	}
}
