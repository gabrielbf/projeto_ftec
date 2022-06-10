// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/samples": {
            "post": {
                "description": "Sample Description.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Request an Sample creation",
                "parameters": [
                    {
                        "description": "Requested body for Sample Creation",
                        "name": "CreateSample",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateSample"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Sample request successfully created",
                        "schema": {
                            "$ref": "#/definitions/response.Sample"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "The endpoint location for the created ressource. E.g /samples/da74025d-07cf-4347-8e97-3073e83adc82"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/v1/samples/:referenceUUID": {
            "get": {
                "description": "In this endpoint you can request an sample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sample"
                ],
                "summary": "Request an Sample",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reference",
                        "name": "referenceUUID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample requested",
                        "schema": {
                            "$ref": "#/definitions/response.Sample"
                        }
                    },
                    "404": {
                        "description": "Sample not found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateSample": {
            "type": "object",
            "required": [
                "reference_uuid"
            ],
            "properties": {
                "reference_uuid": {
                    "description": "UUID sent by client and used as a client generated id. It will be the external reference for the requested sample.",
                    "type": "string",
                    "example": "6fe2d9d5-687d-47da-aaf2-a31d88aa70ca"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "invalid_params": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidParam"
                    }
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "response.InvalidParam": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "response.Sample": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Date time when the sample ressource was saved",
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                },
                "deleted_at": {
                    "description": "Date time when the sample ressource was deleted",
                    "type": "string",
                    "x-nullable": true,
                    "example": "null"
                },
                "reference_uuid": {
                    "description": "UUID sent by client and used as a client generated id. It will be the external reference for the requested sample.",
                    "type": "string",
                    "example": "2e9948d6-5b0c-4984-a963-21e3ff0b2e25"
                },
                "updated_at": {
                    "description": "Date time when the sample ressource was updated",
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Sample Project",
	Description:      "This doc presents the http endpoints that could be used for manage samples.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
