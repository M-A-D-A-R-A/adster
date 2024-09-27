package routes

import (
	"microservices/core/routes/bindings"


	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
	"github.com/labstack/echo/v4/middleware"
)

type routes struct {
	supabaseClient *supabase.Client
}

func NewRoutes(supabaseClient *supabase.Client) *routes {
	return &routes{supabaseClient}
}


// func (r *routes) RunAppRouter() {
// 	router := echo.New()
// 	api := router.Group("/api/v1")

// 	binding.NewFilesHandlerBinding(r.supabaseClient, api).Run()

// 	router.Run()
// }

func (r *routes) RunAppRouter() {
	router:= echo.New()
	 router.Use(middleware.CORS())
	// router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//   }))  
	api := router.Group("/api/v1")

	binding.NewHandlerBinding(r.supabaseClient, api).Run()

	router.Logger.Fatal(router.Start(":8080"))
}