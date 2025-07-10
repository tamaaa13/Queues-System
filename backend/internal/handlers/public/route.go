package public

import (
	"github.com/gin-gonic/gin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

type publicService interface {
	CreateQueue(category string) (*public.QueueModel, error)
	GetQueuesForAdmins() ([]public.QueueResponse, error)
}

type Handler struct {
	*gin.Engine
	publicSvc publicService
}

func NewHandler(api *gin.Engine, publicSvc publicService) *Handler {
	return &Handler{
		Engine:    api,
		publicSvc: publicSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("queue")
	route.POST("/create", h.CreateQueue)

	h.GET("/public/called-queues", h.GetPublicCalledQueues)
	// route.GET("/all", h.GetTodayQueues) // GET
	// route.GET("/next", h.GetNextQueue)  // GET /queues/next?category=A
}
