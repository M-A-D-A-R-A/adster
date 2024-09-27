package binding

import (
	"microservices/src/handler"
	"microservices/src/repository"
	"microservices/src/service"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
)

type handlerBinding struct {
	supabase *supabase.Client
	group    *echo.Group
}

func NewHandlerBinding(supabase *supabase.Client, group  *echo.Group) *handlerBinding {
	return &handlerBinding{supabase, group}
}

func (b *handlerBinding) Run() {
	repo := repository.NewTargetRepository(b.supabase)
	service := service.NewTargetService(repo)
	handler := handler.NewTargetHandler(service)

	b.group.GET("/files", handler.GetAllFiles)

}