package dto

type OrderInsertDto struct {
	UserId       int                   `json:"user_id"`
	OrderDetails OrderDetailsInsertDto `json:"details"`
}
