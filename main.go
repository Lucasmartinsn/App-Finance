package main

import (
	conf "development/application/fiance/conf"
	"fmt"
)

func main() {
	t := conf.Conn()
	te, err := t.Conn.GetUser(t.Cxt, "lucas")
	if err != nil {
		return
	}
	fmt.Println(te)
}
