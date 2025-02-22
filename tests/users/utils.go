package users

import (
	"SparkGuardBackend/cmd/rest/controllers/users"
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/tests/utils"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
	"net/http"
	"net/http/httptest"
	"strings"
)

func CreateUser(router *gin.Engine, username, email string, accessLevel int) (res *db.User, password string, err error) {
	var salt string

	w := httptest.NewRecorder()

	if salt, err = utils.GenerateRandomString(20); err != nil {
		return
	}

	if password, err = utils.GenerateRandomString(45); err != nil {
		return
	}

	hash := sha3.New256()
	hash.Write([]byte(password + salt))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	user := users.CreateUserRequest{
		Username:    username,
		Email:       email,
		AccessLevel: accessLevel,

		Salt: salt,
		Hash: hashedPassword,
	}
	userJson, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/", strings.NewReader(string(userJson)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		err = fmt.Errorf("Expected 201, got %d", w.Code)
		return
	}

	var response users.GetUserResponse
	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		return
	}

	res = &response.User

	return
}

func EditUser(router *gin.Engine, user *db.User) (err error) {
	w := httptest.NewRecorder()

	userJson, _ := json.Marshal(user)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/users/%d", user.ID), strings.NewReader(string(userJson)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = fmt.Errorf("Expected 200, got %d", w.Code)
		return
	}

	if err = json.Unmarshal(w.Body.Bytes(), user); err != nil {
		return
	}

	return
}

func GetUser(router *gin.Engine, id uint) (user *db.User, err error) {
	user = new(db.User)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = fmt.Errorf("Expected 200, got %d", w.Code)
		return
	}

	if err = json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		return
	}

	return
}
