package user

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wael-boudissaa/marquinoBackend/services/auth"
	"github.com/wael-boudissaa/marquinoBackend/types"
	"github.com/wael-boudissaa/marquinoBackend/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// customer
	router.HandleFunc("/login", h.loginUser).Methods("POST")
	router.HandleFunc("/signup", h.signUpUser).Methods("POST")
	//admin
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	// verifiy if the user exists
	var user types.UserLogin

	if err := utils.ParseJson(r, &user); err != nil {
		// if not return 404
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}


	body, err := io.ReadAll(r.Body)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
	fmt.Println("Received JSON:", string(body))

	u, err := h.store.GetUserByEmail(user)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
    
    fmt.Println("user found ")

	if !auth.ComparePasswords([]byte(user.Password), []byte(u.Password)) {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("Invalid Password"))
		return
	}

	token, err := auth.CreateRefreshToken(*u)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) signUpUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	token, err := auth.CreateRefreshToken(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	idUser, err := auth.CreateAnId()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	hashedPassword, err := auth.HashedPassword(user.Password)

	if err := h.store.CreateUser(user, idUser, token, string(hashedPassword)); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token, "message": "User created successfully"})

	// get the user data from the request
	// validate the data
	// if the data is valid
	// create the user
	// return the user data
	// else return the error

}
