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
        "/auth/login": {
            "post": {
                "description": "Authenticates the user with valid credentials and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login User Details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "This endpoint logs out the currently logged-in user. The user must be authenticated,",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Logs out the user by invalidating their JWT token.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Processes user registration by validating the input and creating a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "User Details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/contacts/action": {
            "post": {
                "description": "Handles the blocking or removal of a contact based on user input.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "parameters": [
                    {
                        "description": "Contact action request payload",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContactActionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/contacts/add": {
            "post": {
                "description": "This endpoint processes a request to add or update a contact request between two users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Adds or updates a contact request between users.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Contact request details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ContactRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/messages/get": {
            "get": {
                "description": "Fetches all messages exchanged with a specific recipient.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipient ID",
                        "name": "reciever",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/messages/sent": {
            "post": {
                "description": "Sends a new message from the authorized user to the recipient.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "parameters": [
                    {
                        "description": "Message payload",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "delete": {
                "description": "Deletes a user account by the specified username.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to delete",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/user/fetchUser": {
            "get": {
                "description": "Handles requests to retrieve user details using a username provided in the query parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Fetch user details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to fetch",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/user/updateUserAndProfile": {
            "post": {
                "description": "Updates the details of a user based on the provided username and JSON body payload.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user and profile details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to update",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "JSON payload containing updated user details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserAndProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ContactActionRequest": {
            "type": "object",
            "properties": {
                "action": {
                    "description": "remove, block",
                    "type": "string"
                },
                "contact_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.ContactRequest": {
            "type": "object",
            "properties": {
                "from_user_id": {
                    "type": "string"
                },
                "status": {
                    "description": "pending, accepted, blocked",
                    "type": "string"
                },
                "to_user_id": {
                    "type": "string"
                }
            }
        },
        "models.GenericResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "models.LoginUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "chat_id": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "media_url": {
                    "type": "string"
                },
                "message_id": {
                    "type": "string"
                },
                "recipient_id": {
                    "type": "string"
                },
                "sender_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "models.Profile": {
            "type": "object",
            "properties": {
                "achievements": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "contact_preferences": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "boolean"
                    }
                },
                "cover_photo_url": {
                    "type": "string"
                },
                "education": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "interests": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_profile_public": {
                    "type": "boolean"
                },
                "occupation": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "profile_completeness": {
                    "type": "integer"
                },
                "social_links": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Role": {
            "type": "string",
            "enum": [
                "ADMIN",
                "CLIENT"
            ],
            "x-enum-varnames": [
                "Admin",
                "Client"
            ]
        },
        "models.UpdateUserAndProfile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar_url": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "last_seen": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/models.Profile"
                },
                "role": {
                    "$ref": "#/definitions/models.Role"
                },
                "status_message": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar_url": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "last_seen": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/models.Profile"
                },
                "role": {
                    "$ref": "#/definitions/models.Role"
                },
                "status_message": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
