package repository

import (
	"testing"
)

func TestMysqlClient_Connect(t *testing.T) {
	mysql := MysqlClient{DBName:"scheduler", Ip:"127.0.0.1:3306", Password:"miguelmikael", Username:"root"}
	testConnection(mysql, t)
}

func TestMongoClient_Connect(t *testing.T) {
	mongo := MongoClient{Ip:"127.0.0.1", DBName:"Schedules"}
	testConnection(mongo, t)
}

func TestEsClient_Connect(t *testing.T) {
	//TODO not yet implemented
	elastic_search := EsClient{Index:"_schedule", Ip:"127.0.0.1", port:1092}
	testConnection(elastic_search, t)
}

func testConnection (connection Connection, t *testing.T) {
	//TODO make test more realistic. Esclient does not have real implementation
	t.Logf("Session: %#v", connection)
	session, err := connection.Connect()
	t.Logf("Session: %#v", session)
	if err != nil {
		t.Errorf("Error: %#v", err)
	}

	if len(session.Name) == 0 && len(session.Type) == 0 {
		t.Errorf("Error: %#v", session)
	}
	if session.IsClose {
		t.Errorf("Session is already closed. %#v", session)
	}
	// lets call it twice to check if the session is really close
	session.Close()
	if !session.IsClose {
		t.Errorf("Session is not close, while calling DbSession.Close(). %#v", session)
	}
	session.Close()
	if !session.IsClose {
		t.Errorf("Session is not close, while calling DbSession.Close(). %#v", session)
	}
}
