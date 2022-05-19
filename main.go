package main

import (
    "github.com/gin-gonic/gin"

)

// main function
func main() {
	// Create a router without any middleware by default

	
	// connect to mysql
	db, err := connect_db()
	if err != nil {
		panic(err)
	}
	
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
	router.Run("localhost:8080")

	// close the database connection
	defer db.Close()
	
}

