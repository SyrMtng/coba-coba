package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"ilmudata/task1/controllers"
	"ilmudata/task1/models"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 200 * 1024 * 1024,
	})

	// static
	app.Static("/public", "./public")
	models.InitDbModels()

	userController := controllers.InitUserController(store)
	videoController := controllers.InitVideoController(store)
	advertController := controllers.InitAdsController(store)

	ads := app.Group("/ads")
	ads.Get("/", advertController.Ads)

	user := app.Group("")
	user.Get("/login", userController.Login)
	user.Post("/login", userController.LoginPosted)
	user.Get("/logout", userController.Logout)
	user.Get("/dashboarduser", userController.DashboardUser)

	//Untuk testing harus ada user register agar foreignKey relasi user to video dan advert bisa tersambung
	user.Post("/register", userController.AddRegisteredUser)

	// video := app.Group("/videos")
	video := app.Group("/videos")
	video.Get("/", videoController.IndexVideo)
	video.Get("/create", videoController.AddVideo)
	video.Post("/create", videoController.AddPostedVideo)
	// advert := app.Group("/advert")

	app.Listen(":3001")
}
