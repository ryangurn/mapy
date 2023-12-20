package mongo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Mongo struct {
	Uri               string
	Client            *mongo.Client
	CurrentDatabase   string
	Database          *mongo.Database
	CurrentCollection string
	Collection        *mongo.Collection
}

func Connect() Mongo {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("You must set your 'MONGODB_URI' environment variable in the .env file.")
	}

	// check if database is specified
	db := os.Getenv("MONGODB_DB")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if db != "" {
		return Mongo{Uri: uri, Client: client, CurrentDatabase: db, Database: client.Database(db)}
	}

	return Mongo{Uri: uri, Client: client}
}

// SetDatabase
// This method will specify the database that should be connected to.
func (cnf *Mongo) SetDatabase(database string) {
	cnf.CurrentDatabase = database
	cnf.Database = cnf.Client.Database(database)
}

// SetCollection
// This method will set the collection that we wish to connect to
func (cnf *Mongo) SetCollection(collection string) {
	if cnf.CurrentDatabase == "" {
		panic("Please set the database before setting the collection")
	}
	if cnf.Database == nil {
		panic("Database connection not established")
	}

	cnf.CurrentCollection = collection
	cnf.Collection = cnf.Database.Collection(collection)
}

// AddCollection
// This method will create a collection
func (cnf *Mongo) AddCollection(collection string) {
	if !cnf.CheckSetup() {
		panic("Environment database connection is not properly set-up")
	}

	cmd := bson.D{{"create", collection}}
	var result bson.M
	fmt.Println(cnf)
	if err := cnf.Database.RunCommand(context.TODO(), cmd).Decode(&result); err != nil {
		panic(err)
	}
	
	// get a list of existing collections
//	collections, _ := cnf.Database.ListCollectionNames(context.TODO(), bson.D{})
//	for _, item := range collections {
//		fmt.Println(item)
//	}
}

// CheckSetup
// This method will check if the database and collection are set to
// ensure that all methods that rely on these objects are set.
func (cnf *Mongo) CheckSetup() bool {
	if cnf.Database == nil {
		fmt.Println("Database is empty")
		return false
	}

	if cnf.Collection == nil {
		fmt.Println("Collection is empty")
		return false
	}

	if cnf.CurrentDatabase == "" {
		fmt.Println("CurrentDatabase is empty")
		return false
	}

	if cnf.CurrentCollection == "" {
		fmt.Println("CurrentCollection is empty")
		return false
	}

	return true
}

func (cnf *Mongo) CheckSetupWithoutCollection() bool {
	if cnf.Database == nil {
		fmt.Println("Database is empty")
		return false
	}

	if cnf.CurrentDatabase == "" {
		fmt.Println("CurrentDatabase is empty")
		return false
	}

	return true
}
