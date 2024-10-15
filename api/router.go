package api

import (
	"PowerLedgerGo/api/mains"
	"PowerLedgerGo/api/middleware"
	"PowerLedgerGo/api/users"
	"PowerLedgerGo/infrastructure/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterHandler api router
type RouterHandler struct {
	AppConfig   *config.AppConfig  `inject:""`
	UserHandler *users.UserHandler `inject:""`
	MainHandler *mains.MainHandler `inject:""`
}

// Router create router handler
func (s *RouterHandler) Router() http.Handler {
	// gin mode
	s.ginMode()
	router := gin.Default()
	// Use cross-origin middleware
	router.Use(middleware.Cors)
	// When there is no matching route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "404",
		})
	})
	s.webRoute(router)
	return router
}

func (s *RouterHandler) ginMode() {
	// gin mode设置
	switch s.AppConfig.AppEnv {
	case "local", "dev":
		gin.SetMode(gin.DebugMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func (s *RouterHandler) webRoute(router *gin.Engine) {
	userGroup := router.Group("/user/")
	userGroup.POST("/register", s.UserHandler.UserRegister)
	userGroup.GET("/getBalance", s.UserHandler.GetBalance)
	userGroup.POST("/userRecharge", s.UserHandler.UserRecharge)
	userGroup.POST("/userWithdraw", s.UserHandler.UserWithdraw)

	mainGroup := router.Group("/mains/")
	mainGroup.POST("/submitBid", s.MainHandler.SubmitBid)
	mainGroup.POST("/submitOffer", s.MainHandler.SubmitOffer)
	mainGroup.GET("/queryTrade", s.MainHandler.QueryTrade)
	mainGroup.POST("/tradePayment", s.MainHandler.TradePayment)
}
