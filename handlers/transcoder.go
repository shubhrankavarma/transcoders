package handlers

import (
	"github.com/amagimedia/transcoders/config"
	"sync/atomic"
	"time"

	"github.com/amagimedia/transcoders/dbiface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TranscoderHandler struct {
	Col     dbiface.MongoCollectionAPI
	Cfg     config.Properties
	IsReady *atomic.Value
}

// StatusType is a type for status
type StatusType string

const (
	Active   StatusType = "active"
	Inactive StatusType = "inactive"
)

type Transcoder struct {

	// To be used as a primary key and mandatory field
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedBy string             `json:"updated_by,omitempty" bson:"updated_by,omitempty" validate:"required" example:"me"`

	// Types of input and output
	OutputType string     `json:"output_type,omitempty" bson:"output_type,omitempty"  example:"dash"`
	InputType  string     `json:"input_type,omitempty" bson:"input_type,omitempty"  example:"mp4"`
	Status     StatusType `json:"status,omitempty" bson:"status,omitempty" example:"active"`

	Description string `json:"description,omitempty" bson:"description,omitempty" example:"media_analysis"`
	AssetType   string `json:"asset_type,omitempty" bson:"asset_type,omitempty" validate:"required"  example:"media_analysis"`

	// Operation String - required | examples - Media Analysis, Encoding, Packaging, Processing, Extraction, QC, Sticking, Spliting
	Operation string `json:"operation,omitempty" bson:"operation,omitempty" validate:"required" example:"media_analysis"`

	// Default Value is "Coming Soon"
	TemplateCommand string `json:"template_command,omitempty" bson:"template_command,omitempty" validate:"required" example:"coming soon"`
}
