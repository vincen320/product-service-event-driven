package service

import (
	"context"

	"github.com/vincen320/product-service-event-driven/event"
	"github.com/vincen320/product-service-event-driven/model/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	Create(ctx context.Context, createRequest web.ProductCreateRequest) (web.ProductResponse, error)
	Update(ctx context.Context, updateRequest web.ProductUpdateRequest) (web.ProductResponse, error)
	Delete(ctx context.Context, idProduct primitive.ObjectID, idUser primitive.ObjectID) error
	FindById(ctx context.Context, idProduct primitive.ObjectID) (web.ProductResponse, error)
	FindAll(ctx context.Context) (web.ProductResponses, error)
	FindByIdCache(ctx context.Context, idProduct primitive.ObjectID) (web.ProductResponse, error)
	FindAllCache(ctx context.Context) (web.ProductResponses, error)
	PublishProductCreated(productCreateEvent event.ProductCreatedEvent) error
}
