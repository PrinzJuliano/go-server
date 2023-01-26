package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"prinzjuliano.com/server/types"
)

var founders []types.Founder

func greetingsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func foundersHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&founders)
}

func formHandler(c echo.Context) (err error) {
	u := new(types.Founder)
	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	founders = append(founders, *u)
	return foundersHandler(c)
}
func foundersEmailHandler(c echo.Context) error {
	email := c.Param("email")
	for _, founder := range founders {
		if founder.Email == email {
			return c.JSON(http.StatusOK, founder)
		}
	}
	return c.String(http.StatusNotFound, "No user found with email: '"+email+"'.")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())

	founders = append(founders, types.Founder{Name: "Julian", Age: 24, Email: "PrinzJuliano@users.noreply.github.com", Company: "amicaldo"})

	e.GET("/", greetingsHandler)
	e.POST("/form", formHandler)
	e.GET("/founders", foundersHandler)
	e.GET("/founders/:email", foundersEmailHandler)

	fmt.Printf("Starting server at port 8000\n")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	log.Fatal(e.Start("0.0.0.0:" + port))
}
