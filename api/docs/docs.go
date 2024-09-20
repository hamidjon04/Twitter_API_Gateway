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
        "/twits/addLike": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan tanlangan twitga like qo'shadi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Twitga like qo'shish",
                "parameters": [
                    {
                        "description": "Twit ID ma'lumotlari",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/twit.AddLikeReq1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Like qo'shish muvaffaqiyatli bo'ldi",
                        "schema": {
                            "$ref": "#/definitions/twit.AddLikeResp"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "AddLike request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/createTwit": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan yangi twit yaratadi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Twits"
                ],
                "summary": "Yangi twit yaratish",
                "parameters": [
                    {
                        "description": "Twit ma'lumotlari",
                        "name": "CreateTwitReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/twit.CreateTwitReq1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/twit.CreateTwitResp"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "CreateTwit request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/deleteTwit/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan yaratilgan twitni o'chiradi",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Twits"
                ],
                "summary": "Twitni o'chirish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Twit ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/twit.Message"
                        }
                    },
                    "500": {
                        "description": "DeleteTwit request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/getFollowerTwit": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchining kuzatuvchilari tomonidan yaratilgan barcha twitlarni qaytaradi",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Twits"
                ],
                "summary": "Kuzatuvchilar twitlarini olish",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/twit.GetTwitsResp"
                        }
                    },
                    "500": {
                        "description": "GetFollowerTwit request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/getUserTwit": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchining barcha twitlarini qaytaradi",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Twits"
                ],
                "summary": "Foydalanuvchi twitlarini olish",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/twit.GetTwitsResp"
                        }
                    },
                    "500": {
                        "description": "GetTwits request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/unlike": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan tanlangan twitdan like olib tashlanadi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Twitdan like olib tashlash",
                "parameters": [
                    {
                        "description": "Twit ID ma'lumotlari",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/twit.DeleteTwitReq1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Muvaffaqiyatli olib tashlandi",
                        "schema": {
                            "$ref": "#/definitions/twit.Message"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "RemoveLike request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/twits/updateTwit": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan mavjud twitni yangilaydi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Twits"
                ],
                "summary": "Mavjud twitni yangilash",
                "parameters": [
                    {
                        "description": "Yangilash uchun twit ma'lumotlari",
                        "name": "UpdateTwitReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/twit.UpdateReq1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/twit.UpdateTwitResp"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "UpdateTwit request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/deleteFollower": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi kuzatuvchisini o'chirish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchi kuzatuvchisini o'chirish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Kuzatuvchi ID'si",
                        "name": "followerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Massage"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri kuzatuvchi ID'si",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "DeleteFollower request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/deleteFollowing": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi kuzatayotgan boshqa foydalanuvchini kuzatishni bekor qiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchini kuzatishni bekor qilish",
                "parameters": [
                    {
                        "description": "Bekor qilinadigan following foydalanuvchi IDsi",
                        "name": "followingId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Massage"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "DeleteFollowing request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchini ID orqali o'chirish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchini o'chirish",
                "parameters": [
                    {
                        "description": "Foydalanuvchi ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Massage"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "DeleteUsers so'rov xatosi",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/getByIdFollower": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Kuzatuvchini ID bo'yicha olish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Kuzatuvchini ID bo'yicha olish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Kuzatuvchi ID'si",
                        "name": "followerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Follow"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri kuzatuvchi ID'si",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetByIdFollower request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/getFollowers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Limit va offset query parametrlari orqali foydalanuvchi kuzatuvchilarini olish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchilar ro'yxatidagi kuzatuvchilarni olish",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.GetaFollowersRes"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri limit yoki offset qiymatlari",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetFollowers so'rov xatosi",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/getFollowing": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi kuzatayotgan odamlarning ro'yxatini oladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchi kuzatayotgan odamlarni olish",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sahifada ko'rsatiladigan elementlar soni",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Qaysi sahifadan boshlab olish kerakligi",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.GetaFollowingRes"
                        }
                    },
                    "400": {
                        "description": "Limit yoki offset noto'g'ri",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetFollowing request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/getUsers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Limit va offset query parametrlari orqali foydalanuvchilar ro'yxatini olish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchilar ro'yxatini olish",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.GetUserRes"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri limit yoki offset qiymatlari",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetUsers so'rov xatosi",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/subscribe": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi boshqa bir foydalanuvchini kuzatadi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Kuzatuvchi qo'shish (subscribe)",
                "parameters": [
                    {
                        "description": "Following ID ma'lumotlari",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.FollowingReq1"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Massage"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri kuzatilayotgan foydalanuvchi ID'si",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "AddFollowing request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/users/getByIdFollowing": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi tomonidan kuzatib borilayotgan ma'lum bir foydalanuvchi haqida ma'lumot olish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Kuzatib boruvchi foydalanuvchi haqida ma'lumot olish",
                "parameters": [
                    {
                        "description": "Kuzatilayotgan foydalanuvchi IDsi",
                        "name": "followingId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.Follow"
                        }
                    },
                    "400": {
                        "description": "Ma'lumotlarni olishda xatolik",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetByIdFollowing request error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/users/getUserById": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Foydalanuvchi ma'lumotlarini ID orqali olish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Foydalanuvchini ID orqali olish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Foydalanuvchi ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.User"
                        }
                    },
                    "400": {
                        "description": "Noto'g'ri ID",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "GetByIdUsers so'rov xatosi",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "twit.AddLikeReq1": {
            "type": "object",
            "properties": {
                "twit_id": {
                    "type": "string"
                }
            }
        },
        "twit.AddLikeResp": {
            "type": "object",
            "properties": {
                "massage": {
                    "type": "string"
                }
            }
        },
        "twit.CreateTwitReq1": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "media": {
                    "type": "string"
                }
            }
        },
        "twit.CreateTwitResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "twit.DeleteTwitReq1": {
            "type": "object",
            "properties": {
                "twit_id": {
                    "type": "string"
                }
            }
        },
        "twit.GetTwitsResp": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "twits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/twit.Twit"
                    }
                }
            }
        },
        "twit.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "twit.Twit": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "media": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "twit.UpdateReq1": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "media": {
                    "type": "string"
                }
            }
        },
        "twit.UpdateTwitResp": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "media": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "users.Follow": {
            "type": "object",
            "properties": {
                "birth_day": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "users.Followers": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "follower_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "users.Following": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "following_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "users.FollowingReq1": {
            "type": "object",
            "properties": {
                "following_id": {
                    "type": "string"
                }
            }
        },
        "users.GetUserRes": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.User"
                    }
                }
            }
        },
        "users.GetaFollowersRes": {
            "type": "object",
            "properties": {
                "followers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.Followers"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "users.GetaFollowingRes": {
            "type": "object",
            "properties": {
                "following": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.Following"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "users.Massage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "birth_day": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Auth Service API3",
	Description:      "This is a sample server for Auth Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
