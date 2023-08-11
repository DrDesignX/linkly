package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	Initializers "github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/middlewares"
	"github.com/drdesignx/linkly/models"
	"github.com/drdesignx/linkly/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Routes(port string) error {
	router := setupRoutes()
	log.Printf("Server running on %v port...", port)
	if err := router.Listen("localhost:3000"); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func setupRoutes() *fiber.App {
	router := fiber.New()
	router.Static("/", "./view")
	router.Get("/", index)
	router.All("/login", login)
	router.Post("/logout", logout)
	router.Post("/register", register)
	router.Get("/validate", middlewares.RequireAuth, validate)
	return router
}

func logout(ctx *fiber.Ctx) error {
	// clear cookies
	ctx.ClearCookie("Authentication")
	return ctx.Redirect("/login", 200)
}

func login(ctx *fiber.Ctx) error {
	if ctx.Method() == fiber.MethodGet {
		return ctx.SendFile("./view/login.html")
	} else if ctx.Method() == fiber.MethodPost {
		var data map[string]string
		if err := ctx.BodyParser(&data); err != nil {
			ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to Read Body",
			})
			return err
		}

		user, err := findUserByUsername(data["username"])
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg":  "Username or Password is incorrect",
				"date": time.Now(),
			})
		}

		err = utils.ComparePassword(user.Password, data["password"])
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg":  "Username or Password is incorrect",
				"date": time.Now(),
			})
		}

		tokenString, err := generateToken(user.Username)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg":  "Failed to Create Token",
				"date": time.Now(),
			})
		}

		ctx.Cookie(&fiber.Cookie{
			Name:     "Authentication",
			Value:    tokenString,
			Expires:  time.Now().Add(time.Second * 60),
			HTTPOnly: true,
			SameSite: "lax",
		})

		return ctx.Status(fiber.StatusOK).Send([]byte(ctx.BaseURL() + "\n / -  "))
	} else {
		return ctx.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}
}

func findUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := Initializers.DB.Select("id, username, email, password").First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Second * 60).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func validate(ctx *fiber.Ctx) error {
	data := ctx.Get("user")
	fmt.Print(data)
	fmt.Println(" validation ")
	return ctx.JSON(fiber.Map{
		"Status":   ctx.SendStatus(fiber.StatusOK),
		"Location": "I am in the Validate",
	})
}

func register(ctx *fiber.Ctx) error {
	if ctx.Method() == fiber.MethodPost {
		var data map[string]string
		if err := ctx.BodyParser(&data); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to Read Body",
			})
		}
		// if user exist return alreay exit
		if userExists(data["username"]) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User Already Exists",
			})
		}

		hash, err := utils.HashPassword(data["password"])
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		user := models.User{
			Username: data["username"],
			Email:    data["email"],
			Password: hash,
		}

		if err := Initializers.DB.Create(&user).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	} else {
		return ctx.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}
}

func index(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).SendString(ctx.BaseURL() + "\n / -  ")
}

func userExists(username string) bool {
	var count int64
	err := Initializers.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		fmt.Printf("Error checking user existence: %s\n", err)
		return false
	}
	return count > 0
}
