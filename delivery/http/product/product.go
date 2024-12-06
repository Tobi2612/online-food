package product

import (
	"context"
	"net/http"
	"online-food/domain"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productRepo domain.ProductRepository
}

func New(productRouter *gin.RouterGroup, p domain.ProductRepository) {

	handler := &ProductHandler{
		productRepo: p,
	}

	productRouter.POST("/", handler.createProduct)
	productRouter.GET("/", handler.getProducts)
	productRouter.GET("/:id", handler.getProduct)
}

func (h *ProductHandler) createProduct(ctx *gin.Context) {

	var productDto *domain.ProductDto
	err := ctx.ShouldBindJSON(&productDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"type":    "error",
		})
		return
	}

	if productDto.Name == "" || productDto.Category == "" || productDto.Price == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Name, Category and Price fields are required",
			"code":    http.StatusBadRequest,
			"type":    "error",
		})
		return

	}

	product := domain.Product{
		Name:     productDto.Name,
		Price:    productDto.Price,
		Category: productDto.Category,
	}

	products, err := h.productRepo.Create(context.TODO(), &product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"type":    "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"type":    "success",
		"code":    http.StatusOK,
		"data": map[string]interface{}{
			"product": products,
		},
	})
}

func (h *ProductHandler) getProducts(ctx *gin.Context) {

	products, err := h.productRepo.Fetch(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"type":    "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"type":    "success",
		"code":    http.StatusOK,
		"data": map[string]interface{}{
			"product": products,
		},
	})
}

func (h *ProductHandler) getProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := h.productRepo.GetById(context.TODO(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"type":    "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"type":    "success",
		"code":    http.StatusOK,
		"data": map[string]interface{}{
			"product": product,
		},
	})
}
