package utils

import "time"

type LoginResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type StockResponse struct {
	ID           uint        `json:"id"`
	ProductID    uint        `json:"product_id"`
	Type         string      `json:"description"`
	ProductName  string      `json:"product_name"`
	ProductStock int         `json:"product_stock"`
	Variant      interface{} `json:"variant"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type ResponseOK struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
type PaginationResponse struct {
	Message    string      `json:"message,omitempty"`
	Page       int         `json:"page,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	TotalItems int         `json:"total_items,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
type ErrorResponse struct {
	ErrorMsg string `json:"error_msg,omitempty"`
	Message  string `json:"message"`
}
