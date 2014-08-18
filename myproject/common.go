/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : common.go
 *creationDate: AUG 04 2014
 *updatedDate : 
 *updatedBy   : 
 *purpose     :  
 *****************************************************************************/
package myproject

import (
	"../dbcon"
	"../loderutil"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var Isexist bool

func Isexistdb() bool {

	db :=dbcon.Dbcreation()

	mymap := make(map[string]string)

	err := loderutil.Load("db.properties", mymap)
	if err != nil {
		fmt.Println("error in propertiesUtils---->>>", err)
	}

	rows, err := db.Query("SHOW DATABASES LIKE '" + mymap["db_name"] + "';")
	if err != nil {
		Isexist = false
		fmt.Println("Error in IsExistDataBase ", err)
		fmt.Println("Isexist---->>1", Isexist)
		return Isexist 
	}
	defer rows.Close()

	for rows.Next() {
		Isexist = true
		var name string
		if err := rows.Scan(&name); err != nil {
			Isexist = false
			fmt.Println("Error in rows1 ", err)
			fmt.Println("Isexist---->>2", Isexist)
			return Isexist
		}

		fmt.Println("%s", name)
	}
	fmt.Println("IsDataBase exist", Isexist)
	return Isexist
}
func Createdb() bool {
	var crdb bool = true

	db := dbcon.Dbcreation()
	fmt.Println("Createdb stmt null -->>", db)
	mymap := make(map[string]string)

	err := loderutil.Load("db.properties", mymap)
	if err != nil {
		fmt.Println("Error in propertiesUtils---->>>", err)
	}
	//fmt.Println("DBname-------->>>",mymap["db_name"])
	var sStmt string = "CREATE DATABASE IF NOT EXISTS "+ mymap["db_name"]+";"
	//var sStmt string = "create database tablename;"

	stmt, err := db.Prepare(sStmt)
	if err != nil {
		fmt.Println("Createdb stmt null -->>", err)
		crdb = false
		return  crdb
	}

	res, err := stmt.Exec()
	if err != nil || res == nil {
		fmt.Println("Createdb connection Exec null", err)
		crdb = false
		return  crdb
	}
	// close statement
	stmt.Close()
	// close db
	db.Close()

	return crdb
}
func Createtable() bool {
	var crtable bool = true

	db := dbcon.Dbcon()

	var sStmt string = "CREATE TABLE test (ID INT(100) NOT NULL AUTO_INCREMENT,nameu VARCHAR(255) DEFAULT NULL,organisation VARCHAR(255) DEFAULT NULL,address VARCHAR(255) DEFAULT NULL,mobilenumber VARCHAR(25) DEFAULT NULL,emailid VARCHAR(50) DEFAULT NULL,KEY ID (ID)) ENGINE=MYISAM AUTO_INCREMENT=139 DEFAULT CHARSET=latin1;"
	
	stmt, err := db.Prepare(sStmt)
	if err != nil {
		fmt.Println("Createtable stmt null -->>", err)
		crtable = false
		return crtable
	}

	res, err := stmt.Exec()
	if err != nil || res == nil {
		fmt.Println("Createtable connection Exec null", err)
		crtable = false
		return crtable
	}
	// close statement
	stmt.Close()
	// close db
	db.Close()

	return crtable
}
