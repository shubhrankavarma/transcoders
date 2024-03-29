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

var acceptedInputAndOutputTypes = map[string]bool{
	"mp4":  true,
	"hls":  true,
	"mov":  true,
	"ts":   true,
	"dash": true,
	"mxf":  true,
	"drm":  true,
}

type Transcoder struct {

	// To be used as a primary key and mandatory field
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedBy string             `json:"updated_by,omitempty" bson:"updated_by,omitempty" validate:"required" example:"me"`

	// Types of input and output
	OutputType string     `json:"output_type,omitempty" bson:"output_type,omitempty" validate:"required" example:"dash"`
	InputType  string     `json:"input_type,omitempty" bson:"input_type,omitempty" validate:"required" example:"mp4"`
	Status     StatusType `json:"status,omitempty" bson:"status,omitempty" example:"active"`

	// * Fields to added
	// Codec Field String - Will be used with Input and Output type
	Codec string `json:"codec,omitempty" bson:"codec,omitempty" validate:"required" example:"h264"`

	MultiAudio   *bool `json:"multi_audio,omitempty" bson:"multi_audio,omitempty" validate:"required" example:"false"`
	MultiCaption *bool `json:"multi_caption,omitempty" bson:"multi_caption,omitempty" validate:"required" example:"false"`

	// TrackConfiguration interface{}

	// Descriptor String - required | examples - Media Analysis, Encoding, Packaging, Processing, Extraction, QC, Sticking, Spliting
	Descriptor string `json:"descriptor,omitempty" bson:"descriptor,omitempty" validate:"required" example:"media_analysis"`

	// Default Value is "Comming Soon"
	TemplateCommand string `json:"template_command,omitempty" bson:"template_command,omitempty" validate:"required" example:"comming soon"`
}
