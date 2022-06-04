package dto

import (
	"time"
)

type OrderDto struct {
	Id         int       `json:"id_order"`
	MontoFinal float32   `json:"monto_final"`
	Fecha      time.Time `json:"fecha"`
	IdUsuario int `json:"id_user"`
}

type OrdersDto []OrderDto
