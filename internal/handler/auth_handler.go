package handler

import (
	"net/http"
	"pd_pritani/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin employee"`
}

// @Summary      Login
// @Description  Login user dengan username dan password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  handler.LoginRequest  true  "Login data"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	// accept the request
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and Password must be filled",
		})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// @Summary      Register
// @Description  Register user baru
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  handler.RegisterRequest  true  "Register data"
// @Success      201  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Security     BearerAuth
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	// receive request
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// call service
	err := h.authService.Register(req.Username, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return
	c.JSON(http.StatusCreated, gin.H{
		"message": "User data successfully created",
	})
}

// @Summary      Get Profile
// @Description  Ambil data profile user yang sedang login
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  helper.Response
// @Failure      401  {object}  helper.Response
// @Security     BearerAuth
// @Router       /profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	id := uint(userID.(float64))

	user, err := h.authService.GetProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

// @Summary      Update Profile
// @Description  Update data profile user yang sedang login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  service.UpdateProfileRequest  true  "Update data"
// @Success      200  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Security     BearerAuth
// @Router       /profile [patch]
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	// get user profile from jwt
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}

	id := uint(userID.(float64))

	// bind request
	var req service.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 3. call service
	err := h.authService.UpdateProfile(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "profile updated",
	})
}

func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := h.authService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func (h *AuthHandler) GetUserByID(c *gin.Context) {
	// get id from url
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id doesn't valid",
		})
		return
	}

	// call service
	user, err := h.authService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id doesn't valid",
		})
		return
	}

	var req service.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.authService.UpdateUser(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "profile updated successfully",
	})

}

func (h *AuthHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id not valid",
		})
		return
	}

	err = h.authService.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
