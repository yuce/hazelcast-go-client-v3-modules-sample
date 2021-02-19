package main

import (
	"fmt"
	"log"

	"github.com/hazelcast/hazelcast-go-client/v3"
)

func main() {
	hz, err := hazelcast.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hz.Name())
	myMap, err := hz.GetMap("my-map")
	if err != nil {
		log.Fatal("error getting map:", err)
	}
	err = myMap.Set("foo", "bar")
	if err != nil {
		log.Fatal("error setting value:", err)
	}
	val, err := myMap.Get("foo")
	if err != nil {
		log.Fatal("error getting value:", err)
	}
	fmt.Println("Val:", val)
}
