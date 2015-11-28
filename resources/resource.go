package resources

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/juanjoqmelian/go-rest/users/entities"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

const (
	MongoDbHost = "mongodb://192.168.99.100:27017"
	MongoDbSchema = "go-test"
	MongoDbUserTable = "user"
)

type UserWebService interface {
	NewUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type DefaultUserWebService struct {
	UserWebService
}

func (DefaultUserWebService) NewUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	user := assembleUserFromRequest(request)

	session, err := mgo.Dial(MongoDbHost)
	if err != nil {
		log.Fatal("mongo connection not found! ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(MongoDbSchema).C(MongoDbUserTable)
	err = collection.Insert(user)
	if err != nil {
		log.Fatal("couldn't insert users into mongo! ", err)
	}

	log.Println("User created: ", user)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(204)
}

func (DefaultUserWebService) GetUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	session, err := mgo.Dial(MongoDbHost)
	if err != nil {
		log.Fatal("mongo connection not found! ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(MongoDbSchema).C(MongoDbUserTable)

	result := []entities.User{}
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal("couldn't find user in mongo! ", err)
	}

	log.Println("Retrieving users ", result)

	json.NewEncoder(writer).Encode(result)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
}

func (DefaultUserWebService) GetUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	email := params.ByName("email")

	session, err := mgo.Dial(MongoDbHost)
	if err != nil {
		log.Fatal("mongo connection not found! ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(MongoDbSchema).C(MongoDbUserTable)

	result := entities.User{}
	err = collection.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		log.Fatal("couldn't find user in mongo! ", err)
	}

	log.Println("Retrieving user ", result)

	json.NewEncoder(writer).Encode(result)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
}


func assembleUserFromRequest(request *http.Request) entities.User {

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("Unable to parse request body!", err)
	}

	user := entities.User{}
	json.Unmarshal(requestBody, &user)
	return user
}
