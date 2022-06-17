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
        "termsOfService": "http://asarfi.live./terms/",
        "contact": {
            "name": "Asarfi Pay",
            "url": "http://www.asarfi.live/contact",
            "email": "asarfi.pay@gmaiil.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/customerDetails/{phone_number}": {
            "get": {
                "description": "get accounts details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CustomerDetails"
                ],
                "summary": "Give Customer Details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "phone number",
                        "name": "phone_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseCustomer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/api/v1/discount_cashback": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CashBackDiscountById"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchantDiscountCashBack"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create  CashBack/Discount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CashBack/Discount Creation"
                ],
                "summary": "Create CashBack/Discount",
                "parameters": [
                    {
                        "description": "insertData",
                        "name": "createMerchant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.requestMerchantDiscountCashBack"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchantDiscountCashBack"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/discount_cashback/approve/{id}": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApproveCashBackDiscount"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchantDiscountCashBack"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/merchants": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListMerchant"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create  merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MerchantCreation"
                ],
                "summary": "Create Merchant",
                "parameters": [
                    {
                        "description": "insertData",
                        "name": "createMerchant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.requestMerchantParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/merchants/approve/{id}": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMerchantById"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/merchants/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMerchantById"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.responseMerchant"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aerrors.Error"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aerrors.Error": {
            "type": "object",
            "properties": {
                "actualError": {
                    "type": "string"
                },
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.CardArt": {
            "type": "object",
            "properties": {
                "approved_by": {
                    "type": "integer"
                },
                "approved_date": {
                    "type": "string"
                },
                "art_name": {
                    "type": "string"
                },
                "card_height": {
                    "type": "integer"
                },
                "card_width": {
                    "type": "integer"
                },
                "correlation_id": {
                    "type": "string"
                },
                "cvv_color": {
                    "type": "string"
                },
                "cvv_font_family": {
                    "type": "string"
                },
                "cvv_font_size": {
                    "type": "integer"
                },
                "cvv_font_style": {
                    "type": "string"
                },
                "cvv_font_weight": {
                    "type": "string"
                },
                "cvv_loc_x": {
                    "type": "integer"
                },
                "cvv_loc_y": {
                    "type": "integer"
                },
                "exp_color": {
                    "type": "string"
                },
                "exp_font_family": {
                    "type": "string"
                },
                "exp_font_size": {
                    "type": "integer"
                },
                "exp_font_style": {
                    "type": "string"
                },
                "exp_font_weight": {
                    "type": "string"
                },
                "exp_loc_x": {
                    "type": "integer"
                },
                "exp_loc_y": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "name_color": {
                    "type": "string"
                },
                "name_font_family": {
                    "type": "string"
                },
                "name_font_size": {
                    "type": "integer"
                },
                "name_font_style": {
                    "type": "string"
                },
                "name_font_weight": {
                    "type": "string"
                },
                "name_loc_x": {
                    "type": "integer"
                },
                "name_loc_y": {
                    "type": "integer"
                },
                "pan_color": {
                    "type": "string"
                },
                "pan_font_family": {
                    "type": "string"
                },
                "pan_font_size": {
                    "type": "integer"
                },
                "pan_font_style": {
                    "type": "string"
                },
                "pan_font_weight": {
                    "type": "string"
                },
                "pan_loc_x": {
                    "type": "integer"
                },
                "pan_loc_y": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "token_color": {
                    "type": "string"
                },
                "token_font_family": {
                    "type": "string"
                },
                "token_font_size": {
                    "type": "integer"
                },
                "token_font_style": {
                    "type": "string"
                },
                "token_font_weight": {
                    "type": "string"
                },
                "token_loc_x": {
                    "type": "integer"
                },
                "token_loc_y": {
                    "type": "integer"
                },
                "updated_by": {
                    "type": "integer"
                },
                "updated_date": {
                    "type": "string"
                }
            }
        },
        "handlers.Message": {
            "type": "object",
            "properties": {
                "actualError": {
                    "type": "string"
                },
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.requestMerchantDiscountCashBack": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "minimum_order": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "valid_from": {
                    "type": "string"
                },
                "valid_to": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "handlers.requestMerchantParam": {
            "type": "object",
            "properties": {
                "account_no": {
                    "type": "integer"
                },
                "address_latitude": {
                    "type": "string"
                },
                "address_line": {
                    "type": "string"
                },
                "address_line2": {
                    "type": "string"
                },
                "address_longitude": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "mcc_code": {
                    "type": "integer"
                },
                "merchant_contact_citizenship": {
                    "type": "string"
                },
                "merchant_contact_name": {
                    "type": "string"
                },
                "merchant_contact_no": {
                    "type": "string"
                },
                "merchant_contact_picture": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "merchant_name": {
                    "type": "string"
                },
                "registeration_certificate": {
                    "type": "string"
                },
                "registration_year": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "tax_clearance_certificate": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "handlers.responseCard": {
            "type": "object",
            "properties": {
                "active_date": {
                    "type": "string"
                },
                "bin": {
                    "type": "string"
                },
                "card_art": {
                    "$ref": "#/definitions/handlers.CardArt"
                },
                "correlation_id": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "is_printed": {
                    "type": "boolean"
                },
                "is_virtual": {
                    "type": "boolean"
                },
                "status": {
                    "type": "integer"
                },
                "token_public": {
                    "type": "integer"
                },
                "updated_by": {
                    "type": "integer"
                },
                "updated_date": {
                    "type": "string"
                }
            }
        },
        "handlers.responseCustomer": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.responseGetAccount"
                    }
                },
                "address_line_1": {
                    "type": "string"
                },
                "address_line_2": {
                    "type": "string"
                },
                "address_line_3": {
                    "type": "string"
                },
                "bank_customer_no": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country_a3": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "credit_rating": {
                    "type": "string"
                },
                "credit_rating_source": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "email_backup": {
                    "type": "string"
                },
                "email_main": {
                    "type": "string"
                },
                "family_name": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "id_number": {
                    "type": "string"
                },
                "id_type": {
                    "type": "string"
                },
                "mobile_backup": {
                    "type": "string"
                },
                "mobile_main": {
                    "type": "string"
                },
                "phone_home": {
                    "type": "string"
                },
                "state_province": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "handlers.responseGetAccount": {
            "type": "object",
            "properties": {
                "acc_ext_ref": {
                    "type": "string"
                },
                "acc_level": {
                    "type": "integer"
                },
                "acc_type": {
                    "type": "integer"
                },
                "amount_actual": {
                    "type": "string"
                },
                "amount_blocked": {
                    "type": "string"
                },
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.responseCard"
                    }
                },
                "ccy_bill": {
                    "type": "integer"
                },
                "ccy_set": {
                    "type": "integer"
                },
                "correlation_id": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "parent_acc_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "updated_by": {
                    "type": "integer"
                },
                "updated_date": {
                    "type": "string"
                }
            }
        },
        "handlers.responseMerchant": {
            "type": "object",
            "properties": {
                "account_no": {
                    "type": "integer"
                },
                "address_latitude": {
                    "type": "string"
                },
                "address_line": {
                    "type": "string"
                },
                "address_line2": {
                    "type": "string"
                },
                "address_longitude": {
                    "type": "string"
                },
                "approved_by": {
                    "type": "integer"
                },
                "approved_date": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "mcc_code": {
                    "type": "integer"
                },
                "merchant_contact_citizenship": {
                    "type": "string"
                },
                "merchant_contact_name": {
                    "type": "string"
                },
                "merchant_contact_no": {
                    "type": "string"
                },
                "merchant_contact_picture": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "merchant_name": {
                    "type": "string"
                },
                "registeration_certificate": {
                    "type": "string"
                },
                "registration_year": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "tax_clearance_certificate": {
                    "type": "string"
                },
                "updated_date": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "handlers.responseMerchantDiscountCashBack": {
            "type": "object",
            "properties": {
                "approved_by": {
                    "type": "integer"
                },
                "approved_date": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "minimum_order": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updated_date": {
                    "type": "string"
                },
                "valid_from": {
                    "type": "string"
                },
                "valid_to": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Asarfi Acqurier API",
	Description:      "Asarfi Acqurier API is an api for card management and payment processing",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}