package manualwire

import (
	"github.com/JerryJeager/Symptomify-Backend/config"
	"github.com/JerryJeager/Symptomify-Backend/internal/http"
	"github.com/JerryJeager/Symptomify-Backend/internal/service/chats"
	"github.com/JerryJeager/Symptomify-Backend/internal/service/tabs"
	"github.com/JerryJeager/Symptomify-Backend/internal/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}

func GetTabRepository() *tabs.TabRepo {
	repo := config.GetSession()
	return tabs.NewTabRepo(repo)
}

func GetTabService(repo tabs.TabStore) *tabs.TabServ {
	return tabs.NewTabService(repo)
}

func GetTabController() *http.TabController {
	repo := GetTabRepository()
	service := GetTabService(repo)
	return http.NewTabController(service)
}

func GetChatRepository() *chats.ChatRepo {
	repo := config.GetSession()
	return chats.NewChatRepo(repo)
}

func GetChatService(repo chats.ChatStore) *chats.ChatServ {
	return chats.NewChatService(repo)
}

func GetChatController() *http.ChatController {
	repo := GetChatRepository()
	service := GetChatService(repo)
	return http.NewChatController(service)
}
