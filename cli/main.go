/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/cmd"
	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//set up database global variable
	var err error
	data.DB, err = gorm.Open(sqlite.Open("../python/course.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	cmd.Execute()
	sqlDb, _ := (*data.DB).DB()
	sqlDb.Close()
}
