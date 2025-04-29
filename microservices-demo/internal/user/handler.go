package user

import "github.com/gin-gonic/gin"

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler { return &Handler{repo} }

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/users/:id", h.GetUser)
	r.POST("/users", h.CreateUser)
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	// 实际应查询数据库
	c.JSON(200, gin.H{"id": id, "name": "Alice"})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(&u); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, u)
}
