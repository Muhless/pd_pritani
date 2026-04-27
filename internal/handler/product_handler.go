package handler

import (
	"net/http"
	"pd_pritani/internal/helper"
	"pd_pritani/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

// @Summary      List semua produk
// @Description  Mengambil semua data produk dengan pagination
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        page   query  int  false  "Halaman"
// @Param        limit  query  int  false  "Limit data"
// @Success      200  {object}  helper.PaginationResponse
// @Failure      500  {object}  helper.Response
// @Security     BearerAuth
// @Router       /products/ [get]
func (h *ProductHandler) GetAll(ctx *gin.Context) {
	page, limit := helper.GetPagination(ctx)
	products, total, err := h.productService.GetAll(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	meta := helper.NewPagination(page, limit, total)
	helper.SuccessWithPagination(ctx, http.StatusOK, "successfully getting product data", products, meta)
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	product, err := h.productService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

// @Summary      Buat produk baru
// @Description  Membuat produk baru
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request  body  service.ProductRequest  true  "Product data"
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Security     BearerAuth
// @Router       /products/ [post]
func (h *ProductHandler) Create(ctx *gin.Context) {
	var req service.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.productService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "product successfully created"})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req service.ProductUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.productService.Update(uint(id), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product successfully updated"})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.productService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product deleted"})

}
