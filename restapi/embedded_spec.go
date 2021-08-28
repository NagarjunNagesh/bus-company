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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a pack-and-go bus company Swagger API to manage the trips that PackAndGo offers",
    "title": "PackAndGo Application",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@packandgo.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "packandgo.com",
  "basePath": "/v1",
  "paths": {
    "/trip": {
      "get": {
        "description": "Returns all the trips from PackAndGo",
        "produces": [
          "application/json"
        ],
        "tags": [
          "trip"
        ],
        "summary": "Get all trips from PackAndGo",
        "operationId": "getAllTrips",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Trips"
            }
          },
          "404": {
            "description": "No trips found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
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
        "tags": [
          "trip"
        ],
        "summary": "Add a new trip to the packandgo",
        "operationId": "addNewTrip",
        "parameters": [
          {
            "description": "Trip object that needs to be added to the packandgo",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AddTrip"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Successfully created the trip",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/trip/{tripId}": {
      "get": {
        "description": "Returns a single trip",
        "produces": [
          "application/json"
        ],
        "tags": [
          "trip"
        ],
        "summary": "Find trip by ID",
        "operationId": "getTripById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of trip to return",
            "name": "tripId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Trip"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Trip not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AddTrip": {
      "type": "object",
      "required": [
        "originId",
        "destinationId",
        "dates",
        "price"
      ],
      "properties": {
        "dates": {
          "type": "string"
        },
        "destinationId": {
          "type": "number",
          "format": "int32"
        },
        "originId": {
          "type": "number",
          "format": "int32"
        },
        "price": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Trip": {
      "type": "object",
      "required": [
        "origin",
        "destination",
        "dates",
        "price"
      ],
      "properties": {
        "dates": {
          "type": "string"
        },
        "destination": {
          "type": "string"
        },
        "origin": {
          "type": "string"
        },
        "price": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "Trips": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Trip"
      }
    }
  },
  "tags": [
    {
      "description": "Manage your trips",
      "name": "packandgo",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://packandgo.com"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a pack-and-go bus company Swagger API to manage the trips that PackAndGo offers",
    "title": "PackAndGo Application",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@packandgo.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "packandgo.com",
  "basePath": "/v1",
  "paths": {
    "/trip": {
      "get": {
        "description": "Returns all the trips from PackAndGo",
        "produces": [
          "application/json"
        ],
        "tags": [
          "trip"
        ],
        "summary": "Get all trips from PackAndGo",
        "operationId": "getAllTrips",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Trips"
            }
          },
          "404": {
            "description": "No trips found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
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
        "tags": [
          "trip"
        ],
        "summary": "Add a new trip to the packandgo",
        "operationId": "addNewTrip",
        "parameters": [
          {
            "description": "Trip object that needs to be added to the packandgo",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AddTrip"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Successfully created the trip",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "405": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/trip/{tripId}": {
      "get": {
        "description": "Returns a single trip",
        "produces": [
          "application/json"
        ],
        "tags": [
          "trip"
        ],
        "summary": "Find trip by ID",
        "operationId": "getTripById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of trip to return",
            "name": "tripId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Trip"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "404": {
            "description": "Trip not found",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AddTrip": {
      "type": "object",
      "required": [
        "originId",
        "destinationId",
        "dates",
        "price"
      ],
      "properties": {
        "dates": {
          "type": "string"
        },
        "destinationId": {
          "type": "number",
          "format": "int32"
        },
        "originId": {
          "type": "number",
          "format": "int32"
        },
        "price": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Trip": {
      "type": "object",
      "required": [
        "origin",
        "destination",
        "dates",
        "price"
      ],
      "properties": {
        "dates": {
          "type": "string"
        },
        "destination": {
          "type": "string"
        },
        "origin": {
          "type": "string"
        },
        "price": {
          "type": "number",
          "format": "float64"
        }
      }
    },
    "Trips": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Trip"
      }
    }
  },
  "tags": [
    {
      "description": "Manage your trips",
      "name": "packandgo",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://packandgo.com"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
