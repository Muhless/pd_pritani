package handler

import (
	"net/http"
	"pd_pritani/internal/helper"
	"pd_pritani/internal/service"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService service.DashboardService
}

func NewDashboardHandler(dashboarService service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboarService}
}

func (h *DashboardHandler) GetDashboard(ctx *gin.Context) {
	data, err := h.dashboardService.GetDashboard()
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(ctx, http.StatusOK, "successfully getting dashboard data", data)

}
