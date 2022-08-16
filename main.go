package main

import (
	"fmt"
	"time"

	"github.com/akseltroan/openkoth/api"
	"github.com/akseltroan/openkoth/dockerapi"
	"github.com/gin-gonic/gin"
)

// main function
func main() {
	// Create a router without any middleware by default

	// connect to mysql
	//db, err := connect_db()
	//if err != nil {
	//	panic(err)
	//}

	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Users
	router.GET("/users", api.GetUsers)
	router.GET("/users/:id", api.GetUserByID)
	router.POST("/users", api.PostUsers)
	router.PUT("/users/:id", api.PutUser)
	router.DELETE("/users/:id", api.DeleteUser)

	// Rooms
	router.GET("/rooms", api.GetRooms)
	router.GET("/rooms/:id", api.GetRoomByID)
	router.POST("/rooms", api.PostRooms)
	router.PUT("/rooms/:id", api.PutRoom)
	router.DELETE("/rooms/:id", api.DeleteRoom)
	router.POST("/rooms/:id/vulnMachine", api.AddVulnMachine)
	router.GET("/rooms/:id/king", api.GetKing)
	router.PUT("/rooms/:id/king", api.PutKing)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	go router.Run("localhost:8080") // Initialize a goroutine

	i := 0
	for i < 3 {

		fmt.Print("Test")
		// close the database connection
		//defer db.Close()

		netID := dockerapi.CreateNetwork("vulnNet2")
		time.Sleep(10 * time.Second)
		dockerapi.InspectNetwork("vulnNet2")

		// create a new container and return the container ID
		container, err := dockerapi.CreateNewContainer("nginx")
		if err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)

		dockerapi.ListAllContainers()

		dockerapi.LogsSpecificContainer(container.ID)

		// StopAllContainers()

		// ListAllImages()

		// Pull an image from DockerHub
		// PullImage("hello-world")

		// Get the states of all containers in the docker environment
		// GetAllContainerState()

		dockerapi.GetContainerState(container.ID)

		dockerapi.ConnectContainerToNetwork(container.ID, netID)
		time.Sleep(10 * time.Second)

		dockerapi.GetContainerIPAddress(container.ID)

		dockerapi.StopContainer(container.ID)
		time.Sleep(10 * time.Second)

		dockerapi.RemoveNetwork(netID)
		i++
	}
}
