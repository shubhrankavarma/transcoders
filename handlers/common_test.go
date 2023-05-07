package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/amagimedia/transcoders/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// c             *mongo.Client
	db            *mongo.Database
	transcoderCol *mongo.Collection
	wrongCol      *mongo.Collection
	cfg           config.Properties
	jwtToken      string
	// readyz        *atomic.Value
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
		"template_command":"coming soon",
		"status":"active"
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

func seedDataInDB(changeValue map[string]any, changeKey map[string]string) {

	// Dummy data
	dummyData, err := GetDummyData(changeValue, changeKey)
	if err != nil {
		log.Fatalf("Unable to get dummy data : %v", err)
	}
	var transcoder Transcoder

	// Unmarshal the dummy data
	err = json.Unmarshal([]byte(dummyData), &transcoder)
	if err != nil {
		log.Fatalf("Unable to unmarshal dummy data : %v", err)
	}

	// Insert dummy data
	_, err = transcoderCol.InsertOne(context.Background(), transcoder)
	if err != nil {
		log.Fatalf("Unable to insert dummy data : %v", err)
	}
}

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read : %v", err)
	}

	// Read from hashicorp vault
	// utils.GetHashiCorpVaultValues(&cfg)
	connectURI := fmt.Sprintf(cfg.DBURL, cfg.DBUser, cfg.DBPass)
	transcoderCol = GetMongoDBCollection(connectURI)

	connectURI = fmt.Sprintf(cfg.DBURL, "testwrong", "testwrong")
	wrongCol = GetMongoDBCollection(connectURI)

	// STORE ONLY FOR TESTING
	jwtToken = os.Getenv("JWT_TOKEN")
}

func BeforeEach() {

	// Clear collection
	transcoderCol.Drop(context.Background())

	// Seed data
	seedDataInDB(nil, nil)
	seedDataInDB(map[string]any{"asset_type": "audio"}, nil)
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
