package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/amagimedia/transcoders/config"
	"github.com/amagimedia/transcoders/handlers"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo-contrib/prometheus"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/amagimedia/transcoders/docs"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

const (
	CorrelationID = "x-Request-ID"
)

var (
	c                     *mongo.Client
	db                    *mongo.Database
	transcodersCollection *mongo.Collection
	cfg                   config.Properties
)

// addCorrelationID is a custom middleware function.
// This method will generate a 20 digit requestID and
// added to header of both request and response for traceability.
// Instead of X-Request-ID a custom id X-Correlation-ID is generated.
func addCorrelationID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		//Generate the Correlation ID
		id := ctx.Request().Header.Get(CorrelationID)
		var cID string
		if id == "" {
			cID = random.String(20)
		} else {
			cID = id
		}
		ctx.Request().Header.Set(CorrelationID, cID)
		ctx.Response().Header().Set(CorrelationID, cID)
		return next(ctx)
	}
}

func init() {

	// Reading configurations from environment variables
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configurations cannot be read: %v", err)
	}

	// Connecting to MongoDB
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf(cfg.DBURL, cfg.DBUser, cfg.DBPass)).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	c, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	// Getting a handle for database and collection
	db = c.Database(cfg.DBName)

	// Getting a handle for collection
	transcodersCollection = db.Collection(cfg.TranscodersCollection)

}

func adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hToken := c.Request().Header.Get("Authorization")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(hToken, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(cfg.JwtTokenSecret), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse token")
		}
		if !claims["authorized"].(bool) {
			return echo.NewHTTPError(http.StatusForbidden, "Not Authorized")
		}
		return next(c)
	}
}

//@title Transcoders API
//@version 1.0
//@description This is a transcoders API
//@host localhost:51000
//@BasePath /
//@schemes http

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Logger.SetLevel(log.INFO)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(addCorrelationID)
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.JwtTokenSecret),
		TokenLookup: "header:Authorization"})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
		`"request_ID":"${header:x-Request-ID}"+"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n"}))
	ch := &handlers.TranscoderHandler{Col: transcodersCollection}

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", ch.Healthz)
	e.POST("/transcoders", ch.AddTranscoder, middleware.BodyLimit("1M"), jwtMiddleware)
	e.PUT("/transcoders", ch.PutTranscoder, middleware.BodyLimit("1M"), jwtMiddleware)
	e.PATCH("/transcoders", ch.PatchTranscoder, middleware.BodyLimit("1M"), jwtMiddleware)
	e.GET("/transcoders", ch.GetTranscoders, jwtMiddleware)
	e.DELETE("/transcoders", ch.DeleteTranscoder, jwtMiddleware, adminMiddleware)

	e.Logger.Infof("listening for requests on %s:%s", cfg.Host, cfg.Port)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
