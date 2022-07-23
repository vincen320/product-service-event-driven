package helper

import (
	"github.com/vincen320/product-service-event-driven/event"
	"github.com/vincen320/product-service-event-driven/model/domain"
)

func ToProductCreatedEvent(product domain.Product) event.ProductCreatedEvent {
	return event.ProductCreatedEvent{
		Id:           product.Id,
		IdUser:       product.IdUser,
		NamaProduk:   product.NamaProduk,
		Harga:        product.Harga,
		Kategori:     product.Kategori,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		LastModified: product.LastModified,
	}
}
