package storage

import (
	"adminka/models"
	"crypto/sha256"
	"database/sql"
	"fmt"
)

func DBVerif(username string, pass string) models.UserData {
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

	query := "SELECT * FROM admins_admin_panel WHERE username ='" + username + "' AND password ='" + hashedPass + "'"
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

	var currentUser models.UserData

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

	if len(results) > 0 {
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

func hashPass(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	res := fmt.Sprintf("%x", h.Sum(nil))
	return res
}
