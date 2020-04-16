// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "taxirestapi",
    "version": "1.0.0"
  },
  "paths": {
    "/admin/requests": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "requests",
        "responses": {
          "200": {
            "description": "list of requests",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Fatal",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/request": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "request",
        "responses": {
          "200": {
            "description": "get request",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Fatal",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "Error",
      "type": "object",
      "properties": {
        "attributes": {
          "description": "values for error code placeholders",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "code": {
          "description": "error code",
          "type": "string"
        },
        "detail": {
          "description": "a human-readable explanation specific to this occurrence of the problem.",
          "type": "string",
          "example": "Value of ID must be an integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "taxirestapi",
    "version": "1.0.0"
  },
  "paths": {
    "/admin/requests": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "requests",
        "responses": {
          "200": {
            "description": "list of requests",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Fatal",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/request": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "request",
        "responses": {
          "200": {
            "description": "get request",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Fatal",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "Error",
      "type": "object",
      "properties": {
        "attributes": {
          "description": "values for error code placeholders",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "code": {
          "description": "error code",
          "type": "string"
        },
        "detail": {
          "description": "a human-readable explanation specific to this occurrence of the problem.",
          "type": "string",
          "example": "Value of ID must be an integer"
        }
      }
    }
  }
}`))
}
