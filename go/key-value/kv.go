package main

import (
	"fmt"
	"time"
)

type admin struct {
	username string
	hash     string
	created  time.Time
	business string
	role     string
	group    string
}

func main() {
	fmt.Printf("KV store")
	user := admin{
		username: "ramesh",
		hash:     "",
		business: "Spark",
		role:     "super",
		group:    "chootiya",
	}
	fmt.Printf("User = %v", user)
}

func addElem(list map[string]admin, itm admin) bool {
	list[itm.username] = itm
	return true
}

func delElem(list map[string]admin, itm admin) map[string]admin {
	delete(list, itm.username)
	return list

}

func lookup(list map[string]admin, username string) *admin {
	if itm, ok := list[username]; !ok {
		return nil
	} else {
		return &itm
	}
	return nil

}

func updateElem(list map[string]admin, itm admin, username string) map[string]admin {
	list[username] = itm
	return list
}
