package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync/atomic"
	"testing"

	"github.com/amagimedia/transcoders/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c             *mongo.Client
	db            *mongo.Database
	transcoderCol *mongo.Collection
	wrongCol      *mongo.Collection
	cfg           config.Properties
	jwtToken      string
	readyz        *atomic.Value
)

// Request Endpoints
const (
	RequestEndPoint = "/transcoders"
)

func getRawJSONString() string {
	return `{
		"updated_by":"me",
		"input_type":"dash",
		"output_type":"mp4",
        "asset_type":"video",
		"operation":"media_analysis",
		"template_command":"coming soon"
	}`
}
func GetDummyData(changeValue map[string]any, changeKey map[string]string) (string, error) {
	dummyJsonString := getRawJSONString()
	var dummyData map[string]any
	json.Unmarshal([]byte(dummyJsonString), &dummyData)

	for key, value := range changeValue {

		// Check if the key is present in the struct
		if _, ok := dummyData[key]; ok {
			dummyData[key] = value
		}

	}

	for key, value := range changeKey {

		// Check if the key is present in the struct
		if _, ok := dummyData[key]; ok {

			// Get the value of the key
			valueOfKey := dummyData[key]
			// delete the old key
			delete(dummyData, key)
			// add the new key
			dummyData[value] = valueOfKey
		}
	}

	if data, err := json.Marshal(dummyData); err == nil {
		return string(data), nil
	} else {
		return "", err
	}

}

func GetMongoDBCollection(connectURI string) *mongo.Collection {

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Unable to connect to database : %v", err)
	}
	db = c.Database(cfg.DBName)
	collection := db.Collection(cfg.TranscodersCollection + "_test")

	return collection
}

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read : %v", err)
	}

	connectURI := fmt.Sprintf(cfg.DBURL, cfg.DBUser, cfg.DBPass)
	transcoderCol = GetMongoDBCollection(connectURI)

	connectURI = fmt.Sprintf(cfg.DBURL, "testwrong", "testwrong")
	wrongCol = GetMongoDBCollection(connectURI)

	// STORE ONLY FOR TESTING
	jwtToken = os.Getenv("JWT_TOKEN")
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	//set up
	testCode := m.Run()

	//destory
	transcoderCol.Drop(ctx)
	db.Drop(ctx)
	os.Exit(testCode)
}
