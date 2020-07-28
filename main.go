package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "./")

	// imgs
	// proxy := httputil.NewSingleHostReverseProxy(&url.URL{
	// 	Scheme: "https",
	// 	Host:   "faro-hig-store-dev.s3.eu-central-1.amazonaws.com",
	// 	Path:   "/",
	// })

	// url1, err := url.Parse("https://faro-hig-store-dev.s3-eu-central-1.amazonaws.com/faro-hig-store-dev/")
	url1, err := url.Parse("http://localhost:1235/")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// "https://s3.eu-central-1.amazonaws.com/faro-hig-store-dev/file-folder-n_l-dark.svg"

	fmt.Printf("Url: %+v", url1)

	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	g := e.Group("/img")
	// g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
	g.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(targets),
		Rewrite: map[string]string{
			"/img/*": "/$1",
		},
	}), middleware.Logger())

	// e.GET("/img/*", echo.WrapHandler(proxy))

	second := echo.New()

	second.Use(middleware.Logger())
	second.Use(middleware.Recover())

	// second.GET("/:sth", func(c echo.Context) error {
	// 	fmt.Println(c.Param("sth"))
	// 	return c.String(http.StatusOK, "Hello")
	// })
	second.Static("/", "./img")

	go second.Start(":1235")

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Handler
func hello(c echo.Context) error {
	// fmt.Printf("names: %+v", c.ParamNames())
	// fmt.Printf("values: %+v", c.ParamValues())
	img := c.Param("image")
	fmt.Printf("Image: %s", img)
	return c.Redirect(http.StatusOK, "https://faro-hig-store-dev.s3.eu-central-1.amazonaws.com/file-folder-n_l-dark.svg")
	// return c.File("https://faro-hig-store-dev.s3.eu-central-1.amazonaws.com/file-folder-n_l-dark.svg")
	// return c.String(http.StatusOK, "Hello, World!")
}
