package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"tast_for_AvitoAdvertising/internal/app/server"
)



var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to server conf")
}

func main() {
	fmt.Println("точка входа")
	//time.Sleep(time.Second * 5)
	flag.Parse()

	config := server.NewConfig()
	fmt.Println(configPath)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.StoreConfig)

	fmt.Println("перед стартом")
	err = server.Start(config)
	if err != nil {
		fmt.Println("fail start server")
	}
}