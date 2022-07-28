package main

import (
	action "main/actions"
	dao "main/dao"
)

func main() {
	db, err := dao.Initialize()
	if err != nil {
		panic(err)
	}
	action.Query(db)
}
