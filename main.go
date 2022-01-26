package main // only main package can be compiled

import (
	"os"

	"github.com/estellechoi/go/scrapper"
	"github.com/labstack/echo"
)

const FILE_NAME = "jobs.csv"

func handleHome(c echo.Context) error {
	// c.String(http.StatusOK, "Hello World!")
	return c.File("home.html")
}

func handleScrap(c echo.Context) error {
	defer os.Remove(FILE_NAME)
	searchText := c.FormValue("searchText")
	scrapper.Scrap(searchText)
	return c.Attachment(FILE_NAME, FILE_NAME)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrap)
	e.Logger.Fatal(e.Start(":1323"))
}
