/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : fetch.go
 *creationDate: AUG 04 2014
 *updatedDate : 
 *updatedBy   : 
 *purpose     :
 *****************************************************************************/
package myproject

import (
	"../dbcon"
	"fmt"
	"net/http"
	"strings"
)

type Vertex struct {
	value string
}

func Fetch(w http.ResponseWriter) {
	db := dbcon.Dbcon()

	rows, err := db.Query("SELECT  CONCAT(CONCAT('{','\"Name\":','\"',nameu,'\"', ',')  ,CONCAT('\"Organisation\":','\"',organisation,'\"',','),CONCAT('\"ID\":','\"',ID,'\"',',') ,CONCAT('\"Address\":','\"',address,'\"',',') ,CONCAT('\"Mobilenumber\":','\"',mobilenumber,'\"',','),CONCAT('\"Emailid\":','\"',emailid,'\"','}')  ) json FROM test ")
	if err != nil {
		fmt.Println("Error in Fetch ", err)
	}
	defer rows.Close()

	var Str, Strend string
	Str = "{\"rows\":["
	Strend = ",]}"

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println("Error in rows1 ", err)
		}
		Str = Str + name + ","
	}
	Str = Str + Strend
	Str = strings.Replace(Str, ",,", "", 1)

	if strings.Contains(Str, "{\"rows\":[,]}") {
		Str = "{\"rows\":[]}"
	}

	w.Write([]byte(Str))
	if err := rows.Err(); err != nil {
		fmt.Println("Error in rows2 ", err)
	}
}
