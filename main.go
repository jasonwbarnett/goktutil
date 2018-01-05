package main

import (
	"fmt"

	"github.com/jasonwbarnett/goktutil/cmd"

	"gopkg.in/jasonwbarnett/gokrb5.v555/keytab"
)

func main() {
	fmt.Println("Jason")
	kt, err := keytab.Load("/etc/krb5.keytab")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	fmt.Printf("%+v", kt)
	cmd.KtRemove()
}
