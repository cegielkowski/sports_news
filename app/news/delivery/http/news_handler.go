package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sports_news/domain"
	"sports_news/transport/request"
)

// NewsHandler Represent the http-handler for news.
type newsHandler struct {
	NUsecase domain.NewsUsecase
	Logger   *zap.Logger
}

// NewNewsHandler Will initialize the news / resources endpoint.
func NewNewsHandler(g *gin.RouterGroup, us domain.NewsUsecase, logger *zap.Logger) {
	handler := &newsHandler{
		NUsecase: us,
		Logger:   logger,
	}

	g.GET("/teams/:teamId/news", handler.FetchByTeam)

	g.GET("/teams/:teamId/news/:articleId", handler.GetByIDAndTeam)

	g.GET("/teams/news", handler.Fetch)

	g.GET("/teams/news/:articleId", handler.GetByID)
}

// Fetch Get all news.
func (n *newsHandler) Fetch(c *gin.Context) {
	news, err := n.NUsecase.Fetch(c)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrNotFound.Error()})
			return
		}
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ErrInternalServerError.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// FetchByTeam Get news by team id.
func (n *newsHandler) FetchByTeam(c *gin.Context) {
	var byTeam request.ByTeamId
	if err := c.ShouldBindUri(&byTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrBadParamInput.Error()})
		return
	}

	news, err := n.NUsecase.FetchByTeam(c, byTeam.TeamId)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrNotFound.Error()})
			return
		}
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})

}

// GetByID Get news by article id.
func (n *newsHandler) GetByID(c *gin.Context) {
	var byArticle request.ByArticleId
	if err := c.ShouldBindUri(&byArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrBadParamInput.Error()})
		return
	}

	news, err := n.NUsecase.GetByID(c, byArticle.ArticleId)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrNotFound.Error()})
			return
		}
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

// GetByIDAndTeam Get news by id and team id.
func (n *newsHandler) GetByIDAndTeam(c *gin.Context) {
	var byTeamAndArticle request.ByTeamAndArticleId
	if err := c.ShouldBindUri(&byTeamAndArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrBadParamInput.Error()})
		return
	}

	news, err := n.NUsecase.GetByIDAndTeam(c, byTeamAndArticle.ArticleId, byTeamAndArticle.TeamId)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrNotFound.Error()})
			return
		}
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": domain.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}
