/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : insert.go
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
)

func Insert(r *http.Request) error {
	v := r.URL.Query()
	var sStmt string = "INSERT INTO test (nameu,organisation,address,mobilenumber,emailid) VALUES ('" + v.Get("name") + "','" + v.Get("org") + "','" + v.Get("address") + "','" + v.Get("mob") + "','" + v.Get("email") + "');"
	db := dbcon.Dbcon()

	stmt, err := db.Prepare(sStmt)
	if err != nil {
		fmt.Println("DB stmt null -->>", err)
		return err
	}

	res, err := stmt.Exec()
	if err != nil || res == nil {
		fmt.Println("DB connection Exec null", err)
		return err
	}
	// close statement
	stmt.Close()
	// close db
	db.Close()
	return err
}