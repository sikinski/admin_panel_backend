package models

type UserData struct {
	Id       string `form:"id" binding:"required"`
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Role     string `form:"role" binding:"required"`
	FullName string `form:"fullname" binding:"required"`
	Status   string `form:"status" binding:"required"`
}

type LoginJSON struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type TaskIT struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	DateStart   string `form:"dateStart" `
	DateEnd     string `form:"dateEnd" `
}

type PointsProject struct {
	ProjectID   int    `form:"projectID" binding:"required"`
	ActivePoint string `form:"active_points" binding:"required"`
	EndedPoint  string `form:"ended_points" binding:"required"`
}
