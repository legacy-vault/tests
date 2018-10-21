// main/main.go

package main

// External Imports.
import (
	"github.com/go-redis/redis"
)

// Local Imports (Wrappers)
import (
	redisLocal "../redis"
)

// Built-in Imports.
import (
	"fmt"
)

func main() {

	var err error
	var redisClient *redis.Client

	fmt.Print("Connecting to Redis...")

	// 3. Redis Client.
	redisClient, err = redisLocal.Connect()
	if err != nil {
		fmt.Println("Fail.")
		fmt.Println("Connection Error.", err)
	}
	defer redisClient.Close()

	fmt.Println("Done.")
	fmt.Printf("redisClient is:\r\n%+v.\r\n", *redisClient)
}
