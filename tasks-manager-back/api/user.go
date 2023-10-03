package api

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/DavidJChavez/pkg"
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

// Constants
const (
	uri string = "/api/user"
)

// Temp DB
var UserList []UserGET

func AddUserHandlers() {
	// Uri with path params
	http.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// Uri with path params
	http.HandleFunc(uri+"/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUserById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Endpoints

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	pkg.WriteJson(w, http.StatusOK, UserList)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(params[len(params)-1])
	if err != nil {
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

func createUser(w http.ResponseWriter, r *http.Request) {
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
