/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : dbconnection.go
 *creationDate: AUG 04 2014
 *updatedDate : 
 *updatedBy   : 
 *purpose     :	Database Connection
 *****************************************************************************/
package dbcon

import (
	"../loderutil"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Dbcon() *sql.DB {
	mymap := make(map[string]string)
	err := loderutil.Load("db.properties", mymap)
	if err != nil {
		fmt.Println("error in propertiesUtils---->>>", err)
	}
	//fmt.Println("db_user--->>>",mymap["db_user"],"\n","db_pass--->>",mymap["db_pass"],"\n","db_host--->>",mymap["db_host"],"\n","\n","db_name--->>",mymap["db_name"])

	dsn := mymap["db_user"] + ":" + mymap["db_pass"] + "@" + mymap["db_host"] + "/" + mymap["db_name"] + "?charset=utf8"
	dbN, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("\n", err)
		return nil
	}
	fmt.Println("\n")
	fmt.Println("Connected to the database!!!")
	fmt.Println("\n")
	//defer dbN.Close()
	return dbN
}

func Dbcreation() *sql.DB {
	mymap := make(map[string]string)
	err := loderutil.Load("db.properties", mymap)
	if err != nil {
		fmt.Println("error in propertiesUtils---->>>", err)
	}
	//fmt.Println("db_user--->>>",mymap["db_user"],"\n","db_pass--->>",mymap["db_pass"],"\n","db_host--->>",mymap["db_host"],"\n","\n","db_name--->>",mymap["db_name"])

	dsn := mymap["db_user"] + ":" + mymap["db_pass"] + "@" + mymap["db_host"] + "/?charset=utf8"
	dbN, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("\n", err)
		return nil
	}
	fmt.Println("\n")
	fmt.Println("Connected to the database!!!")
	fmt.Println("\n")
	//defer dbN.Close()
	return dbN
}
