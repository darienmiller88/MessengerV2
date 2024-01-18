package main

import (
	"context"
	"fmt"
	"net/http"

	// "net/http"
	"os"

	// "time"

	"MessengerV2/api/database"
	"MessengerV2/api/pusherclient"
	"MessengerV2/api/routes"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"


	// "github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
)

const MAX_SIZE int64 = 5 * 1024 * 1024 //Max number of bytes, which is 5mb or 5,248,000 bytes

func main(){
	err := godotenv.Load()
	engine := html.New(".", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	index := routes.Index{}
	successful := false
	filename := ""
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"), 
		os.Getenv("CLOUDINARY_API_KEY"), 
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		fmt.Println("err", err)
	}
	
	if err != nil {
		fmt.Println("err:", err)
	}

	if err != nil{
		fmt.Println("err:", err)
	}
	
	database.Init()
	pusherclient.Init()
	index.Init()
	
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, https://messengerv2.netlify.app",
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Static("/static", "./static")
	// app.Config()

	successful = false
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
            "Successful": successful,
			"Image": filename,
        })
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")

		if err != nil{
			fmt.Println("err:", err)
			successful = false
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		if file.Size > MAX_SIZE{
			return c.Status(http.StatusBadRequest).SendString("File too fucking big.")
		}

		filename =  file.Filename
		if err := c.SaveFile(file, fmt.Sprintf("./static/img/%s", file.Filename)); err != nil{
			return c.Status(http.StatusInternalServerError).SendString("Could not save file")
		}

		res, err := cld.Upload.Upload(
			context.Background(), 
			file, 
			uploader.UploadParams{},
		)

		fmt.Println("res:", res.URL, "err:", err)
		fmt.Println("file name:", file.Filename, "file size: ", (float64(file.Size) / float64(1024)) / float64(1024), "mb")
		
		successful = true
		return c.Redirect("/")
		// return c.SendString("success")
	})

	app.Mount("/api/v1", index.Router)

	fmt.Println("running on port:", os.Getenv("PORT"))
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}