package main

import (
	"github.com/jntullo/task/cmd"
)

func main(){
	cmd.Open()
	defer cmd.Close()
	cmd.RootCmd.Execute()
}

