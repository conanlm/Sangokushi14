package api

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/go-redis/redis/v7"
)

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// client.RPush("个性", "胆小", "功名")

	fi, err := os.Open("战法.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		// fmt.Println(string(a))
		client.RPush("战法", string(a))
	}

}
func CreatClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func GetdataClient(client *redis.Client) []string {

	str, _ := client.LRange("个性", 0, -1).Result()
	return str

}

func GetdataClient1(client *redis.Client) []string {

	str, _ := client.LRange("战法", 0, -1).Result()
	return str

}
