package order

import (
	"bufio"
	"context"
	"net/http"
	"online-food/domain"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderRepo   domain.OrderRepository
	ProductRepo domain.ProductRepository
}

func New(orderRouter *gin.RouterGroup, o domain.OrderRepository, p domain.ProductRepository) {

	handler := &OrderHandler{
		OrderRepo:   o,
		ProductRepo: p,
	}

	orderRouter.POST("/", handler.createOrder)
}

func checkCouponInFile(filename, couponCode string, results chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(filename)
	if err != nil {
		results <- false
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == couponCode {
			results <- true
			return
		}
	}
	results <- false
}

func (h *OrderHandler) createOrder(ctx *gin.Context) {
	var orderDto *domain.OrderReqDTO
	err := ctx.ShouldBindJSON(&orderDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if orderDto.CouponCode != "" {
		if len(orderDto.CouponCode) < 8 || len(orderDto.CouponCode) > 10 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"type":    "error",
				"message": "invalid coupon code",
			})
			return
		}

		results := make(chan bool, 3)
		var wg sync.WaitGroup
		couponFiles := []string{"coupon1.txt", "coupon2.txt", "coupon3.txt"}

		wg.Add(len(couponFiles))
		for _, file := range couponFiles {
			go checkCouponInFile(file, orderDto.CouponCode, results, &wg)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		validCount := 0
		for result := range results {
			if result {
				validCount++
			}
		}

		if validCount < 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"type":    "error",
				"message": "invalid coupon code",
			})
			return
		}
	}

	order := &domain.Order{
		CouponCode: orderDto.CouponCode,
	}

	for _, item := range orderDto.Items {
		if item.ProductID == "" || item.Quantity == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Product ID and Quantity are required",
				"type":    "error",
				"code":    http.StatusBadRequest,
			})
			return
		}

		product, err := h.ProductRepo.GetById(context.TODO(), item.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"type":    "error",
				"code":    http.StatusInternalServerError,
			})
			return
		}

		if product == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
				"type":    "error",
				"code":    http.StatusNotFound,
			})
			return
		}

		order.Items = append(order.Items, domain.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	new_order, err := h.OrderRepo.Create(context.TODO(), order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": err.Error(),
			"type":     "error",
			"code":     http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":  http.StatusCreated,
		"msg":   "Order created successfully",
		"type":  "success",
		"order": new_order,
	})
}
