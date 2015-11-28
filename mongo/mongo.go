package mongo

import (
	"log"
	"gopkg.in/mgo.v2"
	"sync"
)

const (
	MongoDbHost = "mongodb://192.168.99.100:27017"
	MongoDbSchema = "go-test"
	MongoDbUserTable = "user"
)

var session *mgo.Session
var err error

type MongoConnection struct {
	session *mgo.Session
}

var instance *MongoConnection
var once sync.Once

func GetConnection() *MongoConnection {

	once.Do(func() {
		log.Println("Connecting to Mongo...")
		session := connect()
		instance = &MongoConnection{session}
		log.Println("MongoConnection has been initialised.")
	})
	return instance
}

func connect() (*mgo.Session) {

	session, err = mgo.Dial(MongoDbHost)
	if err != nil {
		log.Fatal("mongo connection not found! ", err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

func (MongoConnection) Session() *mgo.Session {
	return session
}


