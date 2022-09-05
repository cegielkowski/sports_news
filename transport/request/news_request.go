package request

type ByTeamAndArticleId struct {
	TeamId    string `uri:"teamId" binding:"required"`
	ArticleId string `uri:"articleId" binding:"required"`
}

type ByTeamId struct {
	TeamId string `uri:"teamId" binding:"required"`
}

type ByArticleId struct {
	ArticleId string `uri:"articleId" binding:"required"`
}
