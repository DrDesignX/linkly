package middlewares

import (
	"fmt"
	"os"
	"time"

	Initializers "github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("Authentication")
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println(err)
		return ctx.Redirect("/login", fiber.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Claims: ", claims)
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			return ctx.Redirect("/login")
		}
		var user models.User
		res := Initializers.DB.Where("username = ?", claims["username"]).First(&user)
		if res.Error != nil {
			fmt.Println(res.Error)
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Set("User", string(user.Username))
		ctx.Next()
	} else {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	return nil
}
