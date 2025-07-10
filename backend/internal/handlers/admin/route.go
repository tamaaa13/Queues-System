package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/Queue/internal/middleware"
	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

type adminService interface {
	Login(username, password string) (string, *admin.User, error)
	HashPassword(password string) (string, error)
	CreateUser(user *admin.User) error
	DeleteUser(userID string) error
	GetAllUser() ([]*admin.User, error)
	FindUserByUsername(username string) (*admin.User, error)
	UpdateAdmin(username, newUsername, newPassword, role, category string) error
	GetNextQueue(category string, calledBy string) (*public.QueueModel, error)
	GetTodayQueues() ([]*public.QueueModel, error)
	GetTodayQueuesFiltered(category string) ([]*public.QueueModel, error)
}

type Handler struct {
	*gin.Engine
	adminSvc adminService
}

func NewHandler(api *gin.Engine, adminSvc adminService) *Handler {
	return &Handler{
		Engine:   api,
		adminSvc: adminSvc,
	}
}

func (h *Handler) RegisterRoute() {
	h.POST("/login", h.Login)

	// Protected routes (admin only)
	route := h.Group("/admin")
	route.Use(middleware.AuthMiddleware())
	{
		// Routes yang hanya dapat diakses oleh superadmin
		superadminRoute := route.Group("/")
		superadminRoute.Use(middleware.SuperAdminMiddleware()) // Middleware untuk superadmin
		{
			superadminRoute.POST("/add-admin", h.AddAdminLoket)
			superadminRoute.PUT("/update-admin", h.UpdateAdminLoket)
			superadminRoute.GET("/all-admin", h.GetAllAdmin)
			superadminRoute.DELETE("/delete-admin/:id", h.DeleteUser)
		}

		// Routes yang dapat diakses oleh admin dan superadmin
		route.GET("/call-next-queue/:category", h.CallNextQueue) // Bisa di akses sesuai loginnya pake akun loket category yang sama
		route.GET("/called-queues", h.GetCalledQueues)           // Bisa nampilin sesuai category tergantung loginnya pake akun loket category apa
	}
}
