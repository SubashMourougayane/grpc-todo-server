package main

import (
	"fmt"
	"github.com/SubashMourougayane/grpc-todo-server/server/utils"
)

type Server struct { }
const tomlConfigPath = "../../config/"

func main() {

	config := utils.LoadToml(tomlConfigPath)

	fmt.Println(config)
}
