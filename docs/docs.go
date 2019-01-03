// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-03 22:46:44.098845222 +0800 CST m=+0.064122648

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "云盘的 Api 服务.",
        "title": "云盘 Api 服务",
        "termsOfService": "https://github.com/zm-dev",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/wq1019",
            "email": "2013855675@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api",
    "paths": {
        "/file/rename": {
            "put": {
                "description": "通过文件 ID 重命名文件",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "multipart/form-data"
                ],
                "tags": [
                    "文件"
                ],
                "summary": "重命名文件",
                "operationId": "rename-file",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "文件 ID",
                        "name": "file_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "文件所属的目录 ID",
                        "name": "folder_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "新的文件名",
                        "name": "new_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "404": {
                        "description": "文件不存在 | 目录不存在",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            }
        },
        "/folder": {
            "get": {
                "description": "加载指定的目录及子目录和文件列表",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "multipart/form-data"
                ],
                "tags": [
                    "目录"
                ],
                "summary": "加载指定的目录及子目录和文件列表",
                "operationId": "load-folder",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "目录 ID",
                        "name": "folder_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Folder"
                        }
                    },
                    "404": {
                        "description": "目录不存在 | 没有访问权限 | id 格式不正确",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            },
            "post": {
                "description": "创建一个目录",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "multipart/form-data"
                ],
                "tags": [
                    "目录"
                ],
                "summary": "创建一个目录",
                "operationId": "create-folder",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "父级目录的 ID",
                        "name": "parent_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "新目录的名称",
                        "name": "folder_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Folder"
                        }
                    },
                    "401": {
                        "description": "请先登录",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "404": {
                        "description": "目录名称不能为空 | (父)目录不存在 | 目录已经存在",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            },
            "delete": {
                "description": "批量删除资源(文件/目录)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "资源"
                ],
                "summary": "批量删除资源(文件/目录)",
                "operationId": "delete-source",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "当前目录的 ID",
                        "name": "current_folder_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "要删除的文件 ids",
                        "name": "file_ids",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "description": "要删除的目录 ids",
                        "name": "folder_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "204": {},
                    "401": {
                        "description": "请先登录",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "404": {
                        "description": "请指定要删除的文件或者目录ID | 当前目录不存在",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            }
        },
        "/folder/rename": {
            "put": {
                "description": "通过目录 ID 重命名目录",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "multipart/form-data"
                ],
                "tags": [
                    "目录"
                ],
                "summary": "重命名目录",
                "operationId": "rename-folder",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "所属的目录 ID",
                        "name": "folder_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "新的目录名",
                        "name": "new_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "404": {
                        "description": "目录不存在",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.GlobalError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10001
                },
                "inner_err": {
                    "type": "error"
                },
                "message": {
                    "type": "string",
                    "example": "error message"
                },
                "service_name": {
                    "type": "string",
                    "example": "cloud_disk"
                },
                "status_code": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "model.File": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "extra": {
                    "type": "string"
                },
                "filename": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Folder": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.File"
                    }
                },
                "folder_name": {
                    "type": "string"
                },
                "folders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Folder"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
