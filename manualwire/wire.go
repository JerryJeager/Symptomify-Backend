package manualwire

import (
	"github.com/JerryJeager/Symptomify-Backend/config"
	"github.com/JerryJeager/Symptomify-Backend/internal/http"
	"github.com/JerryJeager/Symptomify-Backend/internal/service/users"
	"github.com/JerryJeager/Symptomify-Backend/internal/service/tabs"
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

