package shop

import (
	"coffee_api/modules/shop/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type Business interface {
	GetListShop(context.Context)
	GetShopById(context.Context) (entity.Shop, error)
	CreateShop(context.Context)
	DeleteShop(context.Context)
}

type Repository interface {
	GetListShop(context.Context)
	GetShopById(context.Context) (entity.Shop, error)
	CreateShop(context.Context)
	DeleteShop(context.Context)
}

type API interface {
	ListShopHandler() gin.HandlerFunc
	GetShopHandler() gin.HandlerFunc
	CreateShopHandler() gin.HandlerFunc
	UpdateShopHandler() gin.HandlerFunc
	DeleteShopHandler() gin.HandlerFunc
}
