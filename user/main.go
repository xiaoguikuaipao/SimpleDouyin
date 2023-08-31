package main

import (
	api "SimpleDouyin/user/kitex_gen/api/register"
	"log"
)

func main() {
	svr := api.NewServer(new(RegisterImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
