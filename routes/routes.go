package routes

import (
	"aseprayana-skeleton-go/configs"
	"aseprayana-skeleton-go/docs"
	h_auth "aseprayana-skeleton-go/handlers/auth"
	h_transaction "aseprayana-skeleton-go/handlers/transaction"
	"aseprayana-skeleton-go/middlewares"
	r_auth "aseprayana-skeleton-go/repositories/auth"
	r_transaction "aseprayana-skeleton-go/repositories/transaction"
	s_auth "aseprayana-skeleton-go/services/auth"
	s_transaction "aseprayana-skeleton-go/services/transaction"

	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	DB = configs.InitDB()

	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)

	transactionR = r_transaction.NewTransactionRepository(DB)
	transactionS = s_transaction.NewTransactionService(transactionR, JWT)
	transactionH = h_transaction.NewTransactionHandler(transactionS, JWT)
)

func New() *echo.Echo {

	e := echo.New()

	middlewares.LoggerMiddleware(e)

	docs.SwaggerInfo.Title = os.Getenv("APP")
	docs.SwaggerInfo.Version = os.Getenv("VERSION")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SCHEME")}

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	v1 := e.Group("api/v1", middlewares.AuthorizeJWT(JWT))
	{
		authRoutes := e.Group("api/auth")
		{
			authRoutes.POST("/register", authH.Signup)
			authRoutes.POST("/login", authH.Signin)
			authRoutes.GET("/profile/:id", authH.Profile)
		}

		n := v1.Group("/transaction", middlewares.AuthorizeJWT(JWT))
		{
			n.POST("", transactionH.Create)
			n.GET("", transactionH.GetAll)
			n.GET("/:id", transactionH.GetById)
			n.PUT("/:id", transactionH.Update)
			n.DELETE("/:id", transactionH.Delete)
		}

	}
	return e
}
