package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateQueue(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kategori harus A atau B"})
		return
	}

	queue, err := h.publicSvc.CreateQueue(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, queue)
}

func (h *Handler) GetPublicCalledQueues(c *gin.Context) {
	queues, err := h.publicSvc.GetQueuesForAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get queue data"})
		return
	}

	c.JSON(http.StatusOK, queues)
}
