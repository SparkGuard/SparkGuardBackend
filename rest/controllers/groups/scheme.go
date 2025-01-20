package groups

import "SparkGuardBackend/db"

type GetGroupsResponse struct {
	Groups []db.Group `json:"groups"`
}

type GetGroupRequest struct {
	ID uint `uri:"id"`
}

type GetGroupResponse struct {
	db.Group
}

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateGroupResponse struct {
	db.Group
}

type EditGroupRequest struct {
	ID uint `uri:"id"`
	db.Group
}

type DeleteGroupRequest struct {
	ID uint `uri:"id"`
}

type AddUserToGroupRequest struct {
	GroupID uint `uri:"id"`
	UserID  uint `json:"user_id" binding:"required"`
}

type RemoveUserFromGroupRequest struct {
	GroupID uint `uri:"id"`
	UserID  uint `json:"user_id" binding:"required"`
}

type AddStudentToGroupRequest struct {
	GroupID   uint `uri:"id"`
	StudentID uint `json:"student_id" binding:"required"`
}

type RemoveStudentFromGroupRequest struct {
	GroupID   uint `uri:"id"`
	StudentID uint `json:"student_id" binding:"required"`
}
