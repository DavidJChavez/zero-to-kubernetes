package api

import (
	"log"
	"net/http"
	"time"

	"github.com/DavidJChavez/internal"
	api_utils "github.com/DavidJChavez/internal/api"
)

type User struct {
	Id       uint      `json:"_id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastName"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

type NewUser struct {
	Name     string    `json:"name" required:"true"`
	LastName string    `json:"lastName"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

var UserList []User

func AddUserHandlers() {
	// g := Router.Group("/user")
	// {
	// 	g.GET("/", getAllUsers)
	// 	g.POST("/", createUser)
	// }
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Endpoints

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	api_utils.WriteJson(w, http.StatusOK, UserList)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var newUser NewUser
	// err := decoder.Decode(&newUser)
	newUser, err := api_utils.GetStructFromJson[NewUser](r.Body)
	if err != nil {
		internal.PrintStruct(newUser)
		log.Println(err)
		api_utils.WriteJson(w, http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	internal.PrintStruct(newUser)

	UserList = append(UserList, buildUser(newUser))

	api_utils.WriteJson(w, http.StatusOK, newUser)
}

// Helpers

func buildUser(nu NewUser) User {
	var newId uint
	if len(UserList) == 0 {
		newId = 1
	} else {
		newId = UserList[len(UserList)-1].Id + 1
	}

	return User{
		newId,
		nu.Name,
		nu.LastName,
		nu.Email,
		nu.Birthday,
	}
}
