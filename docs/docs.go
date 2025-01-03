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
        "/api/formulario": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retornar todos os formulários ativos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "formulario"
                ],
                "summary": "Listar Formulário",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Formulario"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retornar todos os formulários ativos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "formulario"
                ],
                "summary": "Criar Formulário",
                "parameters": [
                    {
                        "description": "Formulário",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Formulario"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.createResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Atualiza um usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "formulario"
                ],
                "summary": "Atualizar Formulário",
                "parameters": [
                    {
                        "description": "Formulário",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Formulario"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.createResponse"
                        }
                    }
                }
            }
        },
        "/api/formulario/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deleta o usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "formulario"
                ],
                "summary": "Deletar Formulário",
                "parameters": [
                    {
                        "description": "Formulário",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Formulario"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.createResponse"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retornar os dados do usuário logado",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Usuário",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/api/user/create": {
            "put": {
                "description": "Criar um novo usuário",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Criar Usuário",
                "parameters": [
                    {
                        "description": "Usuário",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/api/user/grupamento": {
            "post": {
                "description": "Fazer ligação do usuário com o grupamento",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Adicionar Grupamento",
                "parameters": [
                    {
                        "description": "Dados de Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.loginResponse"
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "Pegar token JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Dados de Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.loginResponse"
                        }
                    }
                }
            }
        },
        "/api/variaveis": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Listar todos as Variaveis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "variaveis"
                ],
                "summary": "Listar Variaveis",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Variavel"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Criar somente a variável",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "variaveis"
                ],
                "summary": "Criar Variáveis",
                "parameters": [
                    {
                        "description": "Variavel",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Variavel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.createResponse"
                        }
                    }
                }
            }
        },
        "/api/variaveis/grupamento": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Listar todos os Grupamentos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "variaveis"
                ],
                "summary": "Listar Grupamentos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Grupamento"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Criar um grupamento",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "variaveis"
                ],
                "summary": "Criar grupamento",
                "parameters": [
                    {
                        "description": "Grupamento",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Grupamento"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.createResponse"
                        }
                    }
                }
            }
        },
        "/api/variaveis/tipo": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Listar todos tipos de variáveis",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "variaveis"
                ],
                "summary": "Listar tipos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TipoVariavel"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.createResponse": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.loginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controller.loginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.Formulario": {
            "type": "object",
            "required": [
                "ativo",
                "nome"
            ],
            "properties": {
                "ativo": {
                    "type": "boolean"
                },
                "descricao": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "model.Grupamento": {
            "type": "object",
            "required": [
                "nome"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "model.TipoVariavel": {
            "type": "object",
            "required": [
                "nome"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "senha",
                "usuario"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "empresa": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "manager_id": {
                    "type": "integer"
                },
                "senha": {
                    "type": "string"
                },
                "usuario": {
                    "type": "string"
                }
            }
        },
        "model.Variavel": {
            "type": "object",
            "required": [
                "texto",
                "tipo_variavel_id"
            ],
            "properties": {
                "grupamento_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "obrigatorio": {
                    "type": "boolean"
                },
                "pergunta_id": {
                    "type": "integer"
                },
                "possui_item": {
                    "type": "boolean"
                },
                "texto": {
                    "type": "string"
                },
                "tipo_variavel_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "http://victor.controllab.com:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Controllab :: REST API",
	Description:      "API for Proficiency Testing and Internal Control integration",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
