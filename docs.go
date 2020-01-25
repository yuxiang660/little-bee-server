// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-01-25 11:21:27.7671941 +0800 CST m=+0.122169101

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Little Bee",
            "email": "yuxiang660@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/pub/login": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "Login with username and password.",
                "parameters": [
                    {
                        "description": "JSON format username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token information",
                        "schema": {
                            "$ref": "#/definitions/schema.LoginTokenInfo"
                        }
                    },
                    "400": {
                        "description": "Bad request parameters or invalid username/password",
                        "schema": {
                            "$ref": "#/definitions/errors.impl"
                        }
                    }
                }
            }
        },
        "/api/v1/pub/login/exit": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "Logout with a token.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/errors.impl"
                        }
                    }
                }
            }
        },
        "/api/v1/pub/users": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "Query users with username.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username to query string: ...?user_name=xxx",
                        "name": "user_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "All matched users",
                        "schema": {
                            "$ref": "#/definitions/schema.UserQueryResults"
                        }
                    },
                    "400": {
                        "description": "Bad request parameters",
                        "schema": {
                            "$ref": "#/definitions/errors.impl"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "Create a user with username and password.",
                "parameters": [
                    {
                        "description": "Create a user with username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/errors.impl"
                        }
                    },
                    "400": {
                        "description": "Bad request parameters",
                        "schema": {
                            "$ref": "#/definitions/errors.impl"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.impl": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "schema.LoginParam": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schema.LoginTokenInfo": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "required": [
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "record_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schema.UserQueryResults": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.User"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "127.0.0.1:8181",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Little Bee Server",
	Description: "Restful API description about little bee server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}