package internal

import (
	"fmt"

	"github.com/robin-andreasson/yt-aitool/internal/controllers"

	_ "github.com/joho/godotenv/autoload"
	"github.com/robin-andreasson/fox"
)

func New() {
	fox.CORS(fox.CorsOptions{
		Origins: []string{"http://localhost:5500"},
		Methods: []string{"GET", "POST", "OPTIONS"},
		Headers: []string{"content-type"},
	})

	r := fox.Init()

	r.Static("public", "../")

	r.Get("/", func(c *fox.Context) error {
		return c.File(fox.Status.Ok, "html/index.html")
	})

	r.Post("/subtitles", controllers.Subtitles)

	ai := r.Group("ai")

	ai.Post("/summarize", controllers.Summarize)
	ai.Post("/explain", controllers.Explain)

	fmt.Println("Server running at port 4000")
	r.Listen(4000)
}
