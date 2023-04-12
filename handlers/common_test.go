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
	db *mongo.Database
	// col           *mongo.Collection
	transcoderCol *mongo.Collection
	cfg           config.Properties
	jwtToken      string
)

func getRawJSONString() string {
	return `{
		"updated_by":"me",
		"output_type":"mp4",
		"input_type":"dash",
		"codec":"h264",
		"multi_audio":true,
		"multi_caption":false,
		"descriptor":"media_analysis",
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

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read : %v", err)
	}

	connectURI := fmt.Sprintf(cfg.DBURL, cfg.DBUser, cfg.DBPass)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Unable to connect to database : %v", err)
	}
	db = c.Database(cfg.DBName)
	transcoderCol = db.Collection(cfg.TranscodersCollection + "_test")

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
