package main

import (
	"golang_web_programming/internal"
	"golang_web_programming/internal/controller"
	"golang_web_programming/internal/repository"
	"golang_web_programming/internal/service"
)

func main() {
	repo := repository.NewRepository(map[string]repository.Membership{})
	serv := service.NewApplication(repo)
	cont := controller.NewMemberController(serv)
	server := internal.NewServer(cont)

	server.Run("8080")

}
