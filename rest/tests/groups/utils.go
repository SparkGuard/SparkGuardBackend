package groups

import (
	"SparkGuardBackend/controllers/groups"
	"SparkGuardBackend/db"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func CreateGroup(router *gin.Engine, name string) (res *db.Group, err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/groups", nil)
	req.Header.Set("Content-Type", "application/json")

	group := groups.CreateGroupRequest{
		Name: name,
	}
	groupJson, err := json.Marshal(group)

	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(groupJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		err = errors.New(w.Body.String())
		return
	}

	res = new(db.Group)
	err = json.Unmarshal(w.Body.Bytes(), res)

	return
}

func GetGroup(router *gin.Engine, id uint) (group *db.Group, err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("GET", fmt.Sprintf("/groups/%d", id), nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	group = new(db.Group)
	err = json.Unmarshal(w.Body.Bytes(), &group)

	return
}

func EditGroup(router *gin.Engine, group *db.Group) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PATCH", fmt.Sprintf("/groups/%d", group.ID), nil)
	req.Header.Set("Content-Type", "application/json")

	editGroup := groups.EditGroupRequest{
		Group: *group,
	}
	editGroupJson, err := json.Marshal(editGroup)

	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(editGroupJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}

func AddUserToGroup(router *gin.Engine, groupID, userID uint) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", fmt.Sprintf("/groups/%d/users", groupID), nil)
	req.Header.Set("Content-Type", "application/json")

	addUserToGroup := groups.AddUserToGroupRequest{
		UserID: userID,
	}

	addUserToGroupJson, err := json.Marshal(addUserToGroup)
	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(addUserToGroupJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}

func RemoveUserFromGroup(router *gin.Engine, groupID, userID uint) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/groups/%d/users", groupID), nil)
	req.Header.Set("Content-Type", "application/json")

	removeUserFromGroup := groups.RemoveUserFromGroupRequest{
		UserID: userID,
	}

	removeUserFromGroupJson, err := json.Marshal(removeUserFromGroup)
	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(removeUserFromGroupJson))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}

func AddStudentToGroup(router *gin.Engine, groupID, studentID uint) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", fmt.Sprintf("/groups/%d/students", groupID), nil)
	req.Header.Set("Content-Type", "application/json")

	addStudentToGroup := groups.AddStudentToGroupRequest{
		StudentID: studentID,
	}

	addStudentToGroupJson, err := json.Marshal(addStudentToGroup)
	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(addStudentToGroupJson))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}

func RemoveStudentFromGroup(router *gin.Engine, groupID, studentID uint) (err error) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/groups/%d/students", groupID), nil)
	req.Header.Set("Content-Type", "application/json")

	removeStudentFromGroup := groups.RemoveStudentFromGroupRequest{
		StudentID: studentID,
	}
	removeStudentFromGroupJson, err := json.Marshal(removeStudentFromGroup)

	if err != nil {
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(removeStudentFromGroupJson))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		err = errors.New(w.Body.String())
		return
	}

	return
}
