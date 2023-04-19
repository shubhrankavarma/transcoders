// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Health check API",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transcoders": {
            "get": {
                "description": "Get all the transcoders",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "input_type",
                        "name": "input_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "output_type",
                        "name": "output_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Transcoder"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid limit or skip.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the transcoder",
                "summary": "Update the transcoder",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transcoder updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Please provide id in query parameter.\" example:\"Please provide id in query parameter.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Transcoder not found.\" example:\"Transcoder not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unable to pass the request payload.\" example:\"Unable to pass the request payload.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to update the transcoder.\" example:\"Unable to update the transcoder.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds the transcoder to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Transcoder",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Transcoder"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Transcoder added successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload.\" example:\"Invalid request payload.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Transcoder with the same output type and input type already exists.\" example:\"Transcoder with the same output type and input type already exists.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unable to pass the request payload.\" example:\"Unable to pass the request payload.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to process the request.\" example:\"Unable to process the request.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a transcoder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "input_type",
                        "name": "input_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "output_type",
                        "name": "output_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transcoder Delted successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Please provide output_type and input_type in query parameter.\" example:\"Please provide output_type and input_type in query parameter.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Transcoder not found.\" example:\"Transcoder not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to delete the transcoder.\" example:\"Unable to delete the transcoder.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a transcoder",
                "parameters": [
                    {
                        "type": "string",
                        "description": "input_type",
                        "name": "input_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "output_type",
                        "name": "output_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transcoder updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload.\" example:\"Invalid request payload.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Transcoder not found.\" example:\"Transcoder not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unable to pass the request payload.\" example:\"Unable to pass the request payload.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to update the transcoder.\" example:\"Unable to update the transcoder.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.StatusType": {
            "type": "string",
            "enum": [
                "active",
                "inactive"
            ],
            "x-enum-varnames": [
                "Active",
                "Inactive"
            ]
        },
        "handlers.Transcoder": {
            "type": "object",
            "required": [
                "input_type",
                "output_type",
                "status",
                "template_command",
                "updated_by"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "description": "To be used as a primary key and mandatory field",
                    "type": "string"
                },
                "input_type": {
                    "type": "string",
                    "example": "mp4"
                },
                "output_type": {
                    "description": "Types of input and output",
                    "type": "string",
                    "example": "dash"
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/handlers.StatusType"
                        }
                    ],
                    "example": "active"
                },
                "template_command": {
                    "description": "Default Value is \"Comming Soon\"",
                    "type": "string",
                    "example": "comming soon"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string",
                    "example": "me"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:51000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Transcoders API",
	Description:      "This is a transcoders API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
