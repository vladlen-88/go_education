package main

import (
	"conf_reader"
	"fmt"
)

func main() {
	c := conf_reader.ReadFromFile()
	fmt.Println(*c.Port)
	fmt.Println(*c.Dburl)
	fmt.Println(*c.JaegerUrl)
	fmt.Println(*c.SentryUrl)
	fmt.Println(*c.KafkaBrokerPort)
	fmt.Println(*c.AppID)
	fmt.Println(*c.AppKey)
}
