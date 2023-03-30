package handlers

import (
	"context"
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
