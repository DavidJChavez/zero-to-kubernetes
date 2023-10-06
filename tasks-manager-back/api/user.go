package api

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/DavidJChavez/pkg"
	"github.com/julienschmidt/httprouter"
)

// Structs
type UserGET struct {
	Id       uint      `json:"_id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastName"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

type UserPOST struct {
	Name     string    `json:"name" required:"true"`
	LastName string    `json:"lastName"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

// Temp DB
var UserList []UserGET

func AddUserHandlers() {
	Router.GET("/api/user", getAllUsers)
	Router.GET("/api/user/:id", getUserById)
	Router.POST("/api/user", createUser)
}

// Endpoints

func getAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pkg.WriteJson(w, http.StatusOK, UserList)
}

func getUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != err {
		pkg.PrintErr(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userIdFound := sort.Search(len(UserList), func(i int) bool {
		return UserList[i].Id == uint(id)
	})
	if userIdFound == len(UserList) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	user := UserList[userIdFound]
	pkg.WriteJson[UserGET](w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userDto, err := pkg.GetStructFromJson[UserPOST](r.Body)
	if err != nil {
		pkg.PrintErr(err)
		pkg.WriteJson(w, http.StatusBadRequest, http.StatusBadRequest)
		return
	}
	user := newUser(userDto)
	pkg.PrintStruct(user)
	UserList = append(UserList, user)
	pkg.WriteJson(w, http.StatusOK, user)
}

// Helpers

func newUser(nu UserPOST) UserGET {
	var newId uint
	if len(UserList) == 0 {
		newId = 1
	} else {
		newId = UserList[len(UserList)-1].Id + 1
	}

	return UserGET{
		newId,
		nu.Name,
		nu.LastName,
		nu.Email,
		nu.Birthday,
	}
}
