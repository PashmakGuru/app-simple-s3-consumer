// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/blobs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "List All Blobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/blobs.SuccessResponse-array_blobs_Blob"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/blobs.FailureResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a New Blob",
                "parameters": [
                    {
                        "description": "Details of the blob",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/blobs.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Blob is created successfully",
                        "schema": {
                            "$ref": "#/definitions/blobs.SuccessResponse-blobs_Blob"
                        }
                    },
                    "400": {
                        "description": "There was an error creating the blob",
                        "schema": {
                            "$ref": "#/definitions/blobs.FailureResponse"
                        }
                    }
                }
            }
        },
        "/blobs/{blob}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get a Blob",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Blob Name",
                        "name": "blob",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blob is fetched",
                        "schema": {
                            "$ref": "#/definitions/blobs.SuccessResponse-blobs_ShowResponse"
                        }
                    },
                    "400": {
                        "description": "There was an error fetching the blob",
                        "schema": {
                            "$ref": "#/definitions/blobs.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "There was an error fetching the blob",
                        "schema": {
                            "$ref": "#/definitions/blobs.FailureResponse"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a Blob",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Blob Name",
                        "name": "blob",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blob is deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/blobs.SuccessResponse-blobs_Blob"
                        }
                    },
                    "400": {
                        "description": "There was an error deleting the blob",
                        "schema": {
                            "$ref": "#/definitions/blobs.FailureResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "blobs.Blob": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "version_id": {
                    "type": "string"
                }
            }
        },
        "blobs.CreateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "blobs.FailureResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "blobs.ShowResponse": {
            "type": "object",
            "properties": {
                "blob": {
                    "$ref": "#/definitions/blobs.Blob"
                },
                "content": {
                    "type": "string"
                }
            }
        },
        "blobs.SuccessResponse-array_blobs_Blob": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/blobs.Blob"
                    }
                }
            }
        },
        "blobs.SuccessResponse-blobs_Blob": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/blobs.Blob"
                }
            }
        },
        "blobs.SuccessResponse-blobs_ShowResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/blobs.ShowResponse"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Simple Blob Storage Consumer API",
	Description:      "A simple implementation example for Azure blob storage",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
