package middlewares

import (
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/database"
	model "github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/models"
	"github.com/VieiraGabrielAlexandre/ms-vendaonline/infraestruture/structs"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var identityKey = "id"

func Auth() *jwt.GinJWTMiddleware {

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     Payload(),
		IdentityHandler: IdentityHandler(identityKey),
		Authenticator:   Authenticator(),
		Unauthorized:    Unathorized(),
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}

func IdentityHandler(identityKey string) func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &structs.User{
			Name: claims[identityKey].(string),
		}
	}
}

func Payload() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*structs.User); ok {
			return jwt.MapClaims{
				identityKey: v.Name,
			}
		}
		return jwt.MapClaims{}
	}
}

func Unathorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func Authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals structs.Login

		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		userID := loginVals.Username
		password := loginVals.Password

		userData, _ := getLoginData(userID, password)

		if userData.Active == true {

			if userData != nil {
				return &structs.User{
					UserEmail: userData.Email,
					Name:      userData.Name,
					Active:    userData.Active,
				}, nil
			}
		}

		return nil, jwt.ErrForbidden
	}
}

func getLoginData(email string, password string) (user *model.Users, err string) {
	var loginVals model.Users

	email = email
	password = password

	database.DB.Model(&loginVals).Where("email = ?", email).Where("password = ?", password).Find(&loginVals)

	if loginVals.ID > 0 {
		return &loginVals, ""
	}
	return nil, "User not found"
}
