/*******************************************************************************
 *Copyright (c) COMPANYNAME 2014 . All Rights Reserved.
 *author      : BALAKRISHNAN.K
 *version     : 0.2
 *FileName    : server.go
 *creationDate: AUG 04 2014
 *updatedDate : 
 *updatedBy   : 
 *purpose     :
 *****************************************************************************/
package main

import (
	"./myproject"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	var isexistdb bool = myproject.Isexistdb()
	if isexistdb == true {
		fmt.Println("Yes DataBase is there")
	} else {
		fmt.Println("DataBase is Not There ...")
		var createdb bool = myproject.Createdb()
		if createdb == true {
			fmt.Println("Yes DataDase Created Successfully...")
			var createtable bool = myproject.Createtable()
			if createtable == true {
				fmt.Println("table created successfully....")
			} else {
				fmt.Println("table not created....")
			}
		} else {
			fmt.Println("Data Base Not created")
		}
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	http.HandleFunc("/select", fetch)
	http.HandleFunc("/update", update)
	http.HandleFunc("/remove", removerec)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err in ListenAndServe check")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var compen string
	if strings.TrimSpace(v.Get("org")) == "" {
		compen = "Please enter Organisation "
	} else if strings.TrimSpace(v.Get("name")) == "" {
		compen = "Please enter Nmae"
	} else if strings.TrimSpace(v.Get("address")) == "" {
		compen = "Please enter Address"
	} else if strings.TrimSpace(v.Get("mob")) == "" {
		compen = "Please enter MobileNo"
	} else if strings.TrimSpace(v.Get("email")) == "" {
		compen = "Please enter Email"
	} else {
		err := myproject.Insert(r)
		if err != nil {
			compen = "Error occured Please Reinsert!!!"
		} else {
			compen = "Success"
		}
	}
	w.Header().Set("Server", "A Go Web Server")
	w.Write([]byte(compen))
}
func removerec(w http.ResponseWriter, r *http.Request) {
	err := myproject.Removerecord(r)
	var remov string
	if err != nil {
		remov = "Error occured Please Reinsert!!!"
	} else {
		remov = "Successfully Deleted"
	}
	w.Write([]byte(remov))
}

func fetch(w http.ResponseWriter, r *http.Request) {
	myproject.Fetch(w)
	//fmt.Fprintf(w, Mymap)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r.Referer())
	render(w, "templates/index.html")
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf(tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("template parsing error:", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		fmt.Println("template executing error:", err)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	var compen string
	if strings.TrimSpace(v.Get("org")) == "" {
		compen = "Please enter Organisation"
	}else if strings.TrimSpace(v.Get("name")) == "" {
		compen = "Please enter Nmae"
	} else if strings.TrimSpace(v.Get("address")) == "" {
		compen = "Please enter Address"
	} else if strings.TrimSpace(v.Get("mob")) == "" {
		compen = "Please enter MobileNo"
	} else if strings.TrimSpace(v.Get("email")) == "" {
		compen = "Please enter Email"
	} else {
		err := myproject.Update(r)
		if err != nil {
			compen = "Error occured Please Reinsert!!!"
		} else {
			compen = "Success"
		}
	}
	w.Header().Set("Server", "A Go Web Server")
	w.Write([]byte(compen))
}
