package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"crypto/sha256"	
	
)
func myFunc(c *gin.Context) {
	fmt.Println(c)
} 
var db *sql.DB
var err error

type userData struct {
	Id string `form:"id" binding:"required"`
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Role string `form:"role" binding:"required"`
	FullName string `form:"fullname" binding:"required"`
	Status string `form:"status" binding:"required"`
}
//Hash implements root.Hash
type Hash struct{}

type LoginJSON struct {
    UserName string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}
func auth(c *gin.Context) {
	var json LoginJSON
    c.Bind(&json)
	
	var verif userData = DBVerif(json.UserName, json.Password)
	fmt.Println(verif)
	if (verif.Status == "true"){
		fmt.Println(true)
		c.IndentedJSON(200, verif)
	} else{
		fmt.Println(false)
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

func DBVerif(username string, pass string) userData{
	hashedPass := hashPass(pass)
	fmt.Println(hashedPass)
	// ____________ INIT DB ___________________

	db, err := sql.Open("mysql", "root@root")
	fmt.Println(db)
	if err != nil {
        fmt.Println(err.Error())
    } else {
		fmt.Println("connnect success")
	}
	// ____________ / INIT DB ___________________

	query := "SELECT * FROM admins_admin_panel WHERE username ='"+ username+"' AND password ='"+hashedPass+"'"
	fmt.Println(query)

	rows, err := db.Query(query)
	//
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range values {
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string)
	i := 0

	var currentUser userData 
	
	for rows.Next() { 
		if err := rows.Scan(scans...); err != nil { 
			fmt.Println(err)
		}

		row := make(map[string]string) 

		for k, v := range values { 
			key := cols[k]
			row[key] = string(v)
		}
		results[i] = row // Загружаем набор результатов
		i++
	}

	if(len(results) > 0) {
		currentUser.Id = results[0]["id"]
		currentUser.UserName = results[0]["username"]
		currentUser.Password = results[0]["password"]
		currentUser.Role = results[0]["role"]
		currentUser.FullName = results[0]["name"]
		currentUser.Status = "true"
	} else {
		currentUser.Status = results[0]["false"]
	}

	
	return currentUser
	
}

// ___________________________________ Hash ___________________________________

func hashPass(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	res := fmt.Sprintf("%x", h.Sum(nil))
	return res
}
// ___________________________________ / Hash ___________________________________

func main() {
    router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", myFunc)
	router.POST("/auth", auth)

    router.Run("localhost:8080")
	
}