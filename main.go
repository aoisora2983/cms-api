package main

import (
	"cms/config"
	"cms/controller/login"
	"cms/controller/manage"
	"cms/controller/manage/article"
	"cms/controller/manage/correct"
	manageOpinion "cms/controller/manage/opinion"
	"cms/controller/manage/portfolio"
	"cms/controller/manage/systemgroup"
	"cms/controller/manage/user"
	"cms/controller/view"
	openArticle "cms/controller/view/article"
	"cms/controller/view/comment"
	"cms/controller/view/opinion"
	openPortfolio "cms/controller/view/portfolio"
	"cms/controller/view/tool"
	"cms/middleware"
	"cms/package/validation"
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	setup()
	r := gin.Default()
	setupCORS(r)
	r.Static("upload", "./upload")

	api := r.Group("/api")
	setupRoutes(api)

	port := ":" + config.ApiPort()
	r.Run(port)
}

// setup 初期化処理
func setup() {
	config.Init()
	binding.Validator = validation.NewOzzoValidator()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

// setupCORS CORS設定
func setupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"GET",
			"POST",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			config.AppDomain(),
		},
	}))
}

// setupRoutes ルーティング設定
func setupRoutes(api *gin.RouterGroup) {
	setupAuthRoutes(api)
	setupManageRoutes(api)
	setupViewRoutes(api)
}

// setupAuthRoutes 認証ルーティング
func setupAuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", login.Login)
	api.GET("/logout", login.Logout)
}

// setupManageRoutes 管理画面ルーティング
func setupManageRoutes(api *gin.RouterGroup) {
	mg := api.Group("/manage")
	mg.Use(middleware.JWT())
	{
		mg.GET("/me", manage.Me)
		mg.POST("/menu", manage.Menu)
		mg.POST("/img/upload", manage.ImgUpload)
		mg.POST("/correct", correct.CorrectSentence)

		setupBlogRoutes(mg)
		setupPortfolioRoutes(mg)
		setupUserRoutes(mg)
		setupSystemGroupRoutes(mg)
		setupOpinionRoutes(mg)
	}
}

// setupBlogRoutes ブログ関連ルーティング
func setupBlogRoutes(mg *gin.RouterGroup) {
	blog := mg.Group("/blog")
	{
		// 記事関係
		blog.GET("/article/get/", article.GetArticle)
		blog.POST("/article/get/list", article.GetArticleList)
		blog.POST("/article/register", article.RegisterArticle)
		blog.POST("/article/delete", article.DeleteArticle)

		// タグ関係
		blog.POST("/tag/register", article.RegisterTag)
		blog.POST("/tag/delete", article.DeleteTag)

		// コメント関係
		blog.GET("/comment/get/list", article.GetCommentList)
		blog.POST("/comment/approve", article.ApproveComment)
		blog.POST("/comment/delete", article.DeleteComment)

		// アクセシビリティ
		blog.GET("/accessibility/get/list", article.GetAccessibilityList)
		blog.POST("/accessibility/register", article.RegisterAccessibilityList)

		// 置換文字列
		blog.GET("/accessibility/replace/word/list", article.GetReplaceWordList)
		blog.POST("/accessibility/replace/word/register", article.RegisterReplaceWord)
	}
}

// setupPortfolioRoutes ポートフォリオ関連ルーティング
func setupPortfolioRoutes(mg *gin.RouterGroup) {
	portfolios := mg.Group("/portfolio")
	{
		portfolios.GET("/get/list", portfolio.GetPortfolioList)
		portfolios.POST("/register", portfolio.RegisterPortfolio)
		portfolios.POST("/delete", portfolio.DeletePortfolio)
	}
}

// setupUserRoutes ユーザー関連ルーティング
func setupUserRoutes(mg *gin.RouterGroup) {
	users := mg.Group("/user")
	{
		users.GET("/get/list", user.GetUserList)
		users.POST("/register", user.RegisterUser)
		users.POST("/delete", user.DeleteUser)
	}
}

// setupSystemGroupRoutes システムグループ関連ルーティング
func setupSystemGroupRoutes(mg *gin.RouterGroup) {
	sysGroup := mg.Group("/system/group")
	{
		sysGroup.GET("/get/list", systemgroup.GetSystemGroup)
		// systemGroup.POST("/register", user.RegisterUser) // システムグループ登録
		// user.POST("/delete", article.DeleteTag)     // システムグループ削除
	}
}

// setupOpinionRoutes 問合わせ関連ルーティング
func setupOpinionRoutes(mg *gin.RouterGroup) {
	opinions := mg.Group("/opinion")
	{
		opinions.POST("/get/list", manageOpinion.GetOpinionList)
	}
}

// setupViewRoutes 利用者向け画面ルーティング
func setupViewRoutes(api *gin.RouterGroup) {
	views := api.Group("/view")
	{
		views.POST("/img/upload", view.ImgUpload)

		setupViewArticleRoutes(views)
		setupViewPortfolioRoutes(views)
		setupViewOpinionRoutes(views)
		setupViewToolRoutes(views)
	}
}

// setupViewArticleRoutes 利用者向け記事関連ルーティング
func setupViewArticleRoutes(views *gin.RouterGroup) {
	articles := views.Group("/article")
	{
		articles.GET("/", openArticle.GetOpenArticle)
		articles.POST("/list", openArticle.GetOpenArticleList)
		articles.GET("/count/good", openArticle.CountUpArticleGood)
		articles.POST("/comment", comment.PostComment)
		articles.GET("/comment/count/good", comment.CountUpCommentGood)
		articles.GET("/tag", openArticle.GetTag)
		articles.GET("/tag/list", article.GetTagList)
	}
}

// setupViewPortfolioRoutes 利用者向けポートフォリオ関連ルーティング
func setupViewPortfolioRoutes(views *gin.RouterGroup) {
	portfolios := views.Group("/portfolio")
	{
		portfolios.GET("/", openPortfolio.GetPortfolio)
		portfolios.GET("/list", openPortfolio.GetOpenPortfolioList)
	}
}

// setupViewOpinionRoutes 利用者向け問合わせ関連ルーティング
func setupViewOpinionRoutes(views *gin.RouterGroup) {
	opinions := views.Group("/opinion")
	{
		opinions.POST("/post", opinion.PostOpinion)
	}
}

// setupViewToolRoutes 利用者向けツール関連ルーティング
func setupViewToolRoutes(views *gin.RouterGroup) {
	tools := views.Group("/tool")
	{
		tools.POST("/qr", tool.GetQr)
	}
}
