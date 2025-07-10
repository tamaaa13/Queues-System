package admin

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tamaaa13/Queue/internal/model/admin"
	"github.com/tamaaa13/Queue/internal/model/public"
)

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, user, err := h.adminSvc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
			"category": user.Category,
		},
	})
}

func (h *Handler) AddAdminLoket(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Category string `json:"category"` // A or B
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hanya superadmin yang bisa menambahkan admin
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Lakukan type assertion ke jwt.MapClaims
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Periksa role dari claims
	if claims["role"] != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
		return
	}

	// Hash password
	passwordHash, err := h.adminSvc.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create user
	admin := &admin.User{
		Username:     req.Username,
		PasswordHash: passwordHash,
		Role:         "admin",
		Category:     sql.NullString{String: req.Category, Valid: true},
	}

	if err := h.adminSvc.CreateUser(admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create admin loket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin loket added successfully"})
}

func (h *Handler) GetAllAdmin(c *gin.Context) {
	// Ambil data user dari context (di-set oleh middleware auth)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi user ke JWT claims
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
		return
	}

	// Cek role user, harus superadmin
	if role, ok := claims["role"].(string); !ok || role != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
		return
	}

	// Ambil semua user dari service
	users, err := h.adminSvc.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data user"})
		return
	}

	// Kirim response
	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateAdminLoket(c *gin.Context) {
	var req struct {
		Username    string `json:"username"`
		Password    string `json:"password,omitempty"`
		NewUsername string `json:"new_username,omitempty"`
		Role        string `json:"role,omitempty"`
		Category    string `json:"category,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid token claims"})
		return
	}

	if role, ok := claims["role"].(string); !ok || role != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
		return
	}

	var passwordHash string
	var err error
	if req.Password != "" {
		passwordHash, err = h.adminSvc.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}
	}

	err = h.adminSvc.UpdateAdmin(req.Username, req.NewUsername, passwordHash, req.Role, req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update admin loket"})
		return
	}

	updatedUser, err := h.adminSvc.FindUserByUsername(req.NewUsername)
	if err != nil {
		// Jika new_username kosong, pakai username lama
		updatedUser, err = h.adminSvc.FindUserByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User updated but failed to fetch updated data"})
			return
		}
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	err := h.adminSvc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *Handler) CallNextQueue(c *gin.Context) {
	category := c.Param("category")

	// Validasi kategori
	if category != "A" && category != "B" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	// Ambil user dari context dan cast ke jwt.MapClaims
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
		return
	}

	claims := user.(jwt.MapClaims) // ✅ fix utama ada di sini

	if claims["role"] != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
		return
	}

	if claims["category"] != category {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to call this category"})
		return
	}

	// Dapatkan antrian berikutnya
	queue, err := h.adminSvc.GetNextQueue(category, claims["username"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"queue_number": queue.Number, "status": "called"})
}

func (h *Handler) GetCalledQueues(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You do not have permission"})
		return
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	role := claims["role"].(string)

	var categoryFilter string
	if role == "admin" {
		// Admin hanya boleh lihat kategori-nya sendiri
		category, ok := claims["category"].(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Category is missing for admin"})
			return
		}
		categoryFilter = category
	}

	// Dapatkan semua antrian hari ini
	queues, err := h.adminSvc.GetTodayQueues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve called queues"})
		return
	}

	// Filter berdasarkan kategori jika user adalah admin
	if role == "admin" {
		filtered := []*public.QueueModel{}
		for _, q := range queues {
			if q.Category == categoryFilter {
				filtered = append(filtered, q)
			}
		}
		queues = filtered
	}

	c.JSON(http.StatusOK, queues)
}
