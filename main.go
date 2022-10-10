package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	
)
func myFunc(c *gin.Context) {
	fmt.Println(c)
} 
var db *sql.DB
var err error

type LoginJSON struct {
    UserName string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}
func auth(c *gin.Context) {
	var json LoginJSON
    c.Bind(&json)
	
	var verif bool = DBVerif(json.UserName, json.Password)
	if (verif){
		c.IndentedJSON(200, true)
	} else{
		c.IndentedJSON(200, false)
	}
}
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        } else{
			c.Next()
		}
    }
}
func DBVerif(username string, pass string) bool{
	// ____________ INIT DB ___________________

	db, err := sql.Open("mysql", "root:root@")
	fmt.Println(db)
	if err != nil {
        fmt.Println(err.Error())
    } else {
		fmt.Println("connnect success")
	}
	// ____________ / INIT DB ___________________

	query := "SELECT * FROM admins_admin_panel WHERE username ='"+ username+"' AND password ='"+pass+"'"
	fmt.Println(query)
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}
	
    rows.Close()
	
	return true
}
func main() {
    router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", myFunc)
	router.POST("/auth", auth)

    router.Run("localhost:8080")
	
}