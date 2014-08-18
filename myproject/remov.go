/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : remove.go
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

func Removerecord(r *http.Request) error {
	v := r.URL.Query()
	var sStmt string = "DELETE FROM test WHERE ID=" + v.Get("codeid") + ";"
	fmt.Println("====", sStmt)
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