package handler

import (
	"net/http"
	"pd_pritani/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService}
}

// @Summary      Get customer data
// @Description  get all customer data
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        page   query  int  false  "Halaman"
// @Param        limit  query  int  false  "Limit data"
// @Success      200  {object}  helper.PaginationResponse
// @Failure      500  {object}  helper.Response
// @Security     BearerAuth
// @Router       /customers/ [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.customerService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

func (h *CustomerHandler) GetByID(C *gin.Context) {
	id, err := strconv.Atoi(C.Param("id"))
	if err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customer, err := h.customerService.GetByID(uint(id))
	if err != nil {
		C.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	C.JSON(http.StatusOK, gin.H{"data": customer})

}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req service.CustomerRequest
	if err := c.ShouldBindJSON(&req); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.customerService.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "data successfully created",
	})
}

func (h *CustomerHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req service.CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.customerService.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated customer data",
	})
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.customerService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}
