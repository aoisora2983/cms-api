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
	// 環境変数初期化
	config.Init()
	binding.Validator = validation.NewOzzoValidator()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r := gin.Default()

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

	// 静的ファイル
	r.Static("upload", "./upload")

	// API
	api := r.Group("/api")
	{
		// 認証
		api.POST("/login", login.Login)
		api.GET("/logout", login.Logout)

		// 管理画面
		mg := api.Group("/manage")
		mg.Use(middleware.JWT()) // 認証ミドルウェアを通す
		{
			mg.GET("/me", manage.Me)      // ログイン確認
			mg.POST("/menu", manage.Menu) // ログイン後メニュー画面情報

			mg.POST("/img/upload", manage.ImgUpload) // 画像アップロード

			mg.POST("/correct", correct.CorrectSentence) // 記述チェック

			// blog関係
			blog := mg.Group("/blog")
			{
				// 記事関係
				blog.GET("/article/get/", article.GetArticle)           // 記事情報取得
				blog.POST("/article/get/list", article.GetArticleList)  // 記事一覧取得
				blog.POST("/article/register", article.RegisterArticle) // 記事登録
				blog.POST("/article/delete", article.DeleteArticle)     // 記事削除

				// タグ関係
				blog.POST("/tag/register", article.RegisterTag) // タグ登録
				blog.POST("/tag/delete", article.DeleteTag)     // タグ削除

				// コメント関係
				blog.GET("/comment/get/list", article.GetCommentList) // コメント一覧取得
				blog.POST("/comment/approve", article.ApproveComment) // コメント承認
				blog.POST("/comment/delete", article.DeleteComment)   // コメント削除

				// アクセシビリティ
				blog.GET("/accessibility/get/list", article.GetAccessibilityList)       // 一覧取得
				blog.POST("/accessibility/register", article.RegisterAccessibilityList) // 一覧登録

				// 置換文字列
				blog.GET("/accessibility/replace/word/list", article.GetReplaceWordList)       // 一覧取得
				blog.POST("/accessibility/replace/word/register", article.RegisterReplaceWord) // 登録
			}

			// 固有ページ
			portfolios := mg.Group("/portfolio")
			{
				portfolios.GET("/get/list", portfolio.GetPortfolioList)   // ポートフォリオ一覧取得
				portfolios.POST("/register", portfolio.RegisterPortfolio) // ポートフォリオ登録
				portfolios.POST("/delete", portfolio.DeletePortfolio)     // ポートフォリオ削除
			}

			// ユーザー関係
			users := mg.Group("/user")
			{
				users.GET("/get/list", user.GetUserList)   // ユーザー一覧取得
				users.POST("/register", user.RegisterUser) // ユーザー登録
				users.POST("/delete", user.DeleteUser)     // ユーザー削除
			}

			// システムグループ関係
			sysGroup := mg.Group("/system/group")
			{
				sysGroup.GET("/get/list", systemgroup.GetSystemGroup) // システムグループ一覧取得
				// systemGroup.POST("/register", user.RegisterUser) // システムグループ登録
				// user.POST("/delete", article.DeleteTag)     // システムグループ削除
			}

			// 問合わせ関係
			opinions := mg.Group("/opinion")
			{
				opinions.POST("/get/list", manageOpinion.GetOpinionList)
			}
		}

		// 利用者向け画面
		views := api.Group("/view")
		{
			views.POST("/img/upload", view.ImgUpload) // 画像アップロード

			articles := views.Group("/article")
			{
				articles.GET("/", openArticle.GetOpenArticle)          // ページ情報
				articles.POST("/list", openArticle.GetOpenArticleList) // ページ一覧情報

				articles.GET("/count/good", openArticle.CountUpArticleGood) // goodボタン押下

				articles.POST("/comment", comment.PostComment)                  // コメント投稿
				articles.GET("/comment/count/good", comment.CountUpCommentGood) // コメントgoodボタン押下

				articles.GET("/tag", openArticle.GetTag)      // タグ取得
				articles.GET("/tag/list", article.GetTagList) // タグ一覧取得
			}

			portfolios := views.Group("/portfolio")
			{
				portfolios.GET("/")                                         // ポートフォリオ情報
				portfolios.GET("/list", openPortfolio.GetOpenPortfolioList) // ポートフォリオ情報一覧
			}

			opinions := views.Group("/opinion")
			{
				opinions.POST("/post", opinion.PostOpinion) // お問合せ送信
			}

			tools := views.Group("/tool")
			{
				tools.POST("/qr", tool.GetQr) // QRコード生成
			}
		}
	}

	var port = ":" + config.ApiPort()

	r.Run(port)
}
