package handlers

import (
	"time"

	"github.com/amagimedia/transcoders/dbiface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TranscoderHandler struct {
	Col dbiface.MongoCollectionAPI
}

// StatusType is a type for status
type StatusType string

const (
	Active   StatusType = "active"
	Inactive StatusType = "inactive"
)

type InputAndOutputType string

const (
	MP4  InputAndOutputType = "mp4"
	HLS  InputAndOutputType = "hls"
	MOV  InputAndOutputType = "mov"
	TS   InputAndOutputType = "ts"
	DASH InputAndOutputType = "dash"
)

type Transcoder struct {

	// To be used as a primary key and mandatory field
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by" validate:"required"`

	// Types of input and output
	OutputType string     `json:"output_type" bson:"output_type" validate:"required"`
	InputType  string     `json:"input_type" bson:"input_type" validate:"required"`
	Status     StatusType `json:"status" bson:"status"`

	// Default Value is "Comming Soon"
	TemplateCommand string `json:"template_command" bson:"template_command" validate:"required"`
}
