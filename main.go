package main

import (
	"fmt"
	"time"

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
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)
	router.PUT("/users/:id", putUser)
	router.DELETE("/users/:id", deleteUser)

	// Rooms
	router.GET("/rooms", getRooms)
	router.GET("/rooms/:id", getRoomByID)
	router.POST("/rooms", postRooms)
	router.PUT("/rooms/:id", putRoom)
	router.DELETE("/rooms/:id", deleteRoom)
	router.POST("/rooms/:id/vulnMachine", addVulnMachine)
	router.GET("/rooms/:id/king", getKing)
	router.PUT("/rooms/:id/king", putKing)

	// By default it serves on :8080 unless a PORT environment variable was defined.
	go router.Run("localhost:8080") // Initialize a goroutine

	i := 0
	for i < 3 {

		fmt.Print("Test")
		// close the database connection
		//defer db.Close()

		netID := createNetwork("vulnNet2")
		time.Sleep(10 * time.Second)
		inspectNetwork("vulnNet2")

		// create a new container and return the container ID
		container, err := CreateNewContainer("nginx")
		if err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)

		ListAllContainers()

		LogsSpecificContainer(container.ID)

		// StopAllContainers()

		// ListAllImages()

		// Pull an image from DockerHub
		// PullImage("hello-world")

		// Get the states of all containers in the docker environment
		// GetAllContainerState()

		GetContainerState(container.ID)

		connectContainerToNetwork(container.ID, netID)
		time.Sleep(10 * time.Second)

		getContainerIPAddress(container.ID)

		StopContainer(container.ID)
		time.Sleep(10 * time.Second)

		removeNetwork(netID)
		i++
	}
}
