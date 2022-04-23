package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// room struct
type room struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Gamemode    string  `json:"gamemode"`
	Players     int     `json:"players"`
	MaxPlayers  int     `json:"max_players"`
	VulnMachine string  `json:"vuln_machine"`
	King 	    string  `json:"king"`
	Status 	    string  `json:"status"`
}

	// user struct
type user struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Level     int  `json:"level"`
}

// This is temporary data to be used while developing the API.
// users is a slice of user structs
var users = []user{
	{ID: "1", Username: "Troan", Level: 5},
	{ID: "2", Username: "Kent", Level: 1},
	{ID: "3", Username: "Sarah", Level: 4},
}
var rooms = []room{
	{ID: "1", Name: "Lobby 1", Gamemode: "King of the Hill", Players: 5, MaxPlayers: 10, VulnMachine: "Linux RootMe", King: "Lastest-king", Status: "Finished"},
	{ID: "2", Name: "Troan's Private Game", Gamemode: "Attack & Defense", Players: 1, MaxPlayers: 10, VulnMachine: "Multiple", King: "Troan", Status: "Running"},
	{ID: "3", Name: "Public Koth", Gamemode: "King of the Hill", Players: 4, MaxPlayers: 10, VulnMachine: "Windows Blue", King: "Not-Set", Status: "Waiting to Start"},
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of users, looking for
	// a user whose ID value matches the parameter.
	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}

	// If no user with the specified ID was found,
	// c.Abort with an error 404 (Not Found).
	c.AbortWithStatus(http.StatusNotFound)
}

// postUsers adds an user from JSON received in the request body.
func postUsers(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

// putUser updates an user based on the specified ID.
func putUser(c *gin.Context) {
	// First get the ID parameter from the request.
	id := c.Param("id")

	// Find and update the user based on the ID.
	for index, u := range users {
		if u.ID == id {
			users[index] = user{ID: u.ID, Username: u.Username, Level: u.Level}
			break
		}
	}

	c.String(http.StatusOK, "User updated!")
}


// deleteUser removes an user based on the specified ID.
func deleteUser(c *gin.Context) {
	// First get the ID parameter from the request.
	id := c.Param("id")

	// Find and remove the user based on the ID.
	for index, u := range users {
		if u.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}

	c.String(http.StatusOK, "User deleted!")
}

// getRooms responds with the list of all rooms as JSON.
func getRooms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rooms)
}

// getRoomByID locates the room whose ID value matches the id
// parameter sent by the client, then returns that room as a response.
func getRoomByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of rooms, looking for
	// a room whose ID value matches the parameter.
	for _, r := range rooms {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}

	// If no room with the specified ID was found,
	// c.Abort with an error 404 (Not Found).
	c.AbortWithStatus(http.StatusNotFound)
}

// postRooms adds an room from JSON received in the request body.
func postRooms(c *gin.Context) {
	var newRoom room

	// Call BindJSON to bind the received JSON to newRoom.
	if err := c.BindJSON(&newRoom); err != nil {
		return
	}

	// Add the new room to the slice.
	rooms = append(rooms, newRoom)

	c.IndentedJSON(http.StatusCreated, newRoom)
}

// putRoom updates an room based on the specified ID.
func putRoom(c *gin.Context) {
	// First get the ID parameter from the request.
	id := c.Param("id")

	// Find and update the room based on the ID.
	for index, r := range rooms {
		if r.ID == id {
			rooms[index] = room{ID: r.ID, Name: r.Name, Gamemode: r.Gamemode, Players: r.Players, MaxPlayers: r.MaxPlayers, VulnMachine: r.VulnMachine, King: r.King, Status: r.Status}
			break
		}
	}

	c.String(http.StatusOK, "Room updated!")
}

// deleteRoom removes an room based on the specified ID.
func deleteRoom(c *gin.Context) {
	// First get the ID parameter from the request.
	id := c.Param("id")

	// Find and remove the room based on the ID.
	for index, r := range rooms {
		if r.ID == id {
			rooms = append(rooms[:index], rooms[index+1:]...)
			break
		}
	}

	c.String(http.StatusOK, "Room deleted!")
}

// addVulnMachine adds vulnerable machines to the room based on the specified ID.
func addVulnMachine(c *gin.Context) {
	// First get the ID parameter from the request.
	id := c.Param("id")

	// Find and update the room based on the ID.
	for index, r := range rooms {
		if r.ID == id {
			// append the vulnerable machine to the room
			rooms[index].VulnMachine = rooms[index].VulnMachine + c.PostForm("vulnMachine")

			break
		}
	}

	c.String(http.StatusOK, "Vulnerable Machine added!")
}


// getKing responds with the current king of the room as JSON.
func getKing(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of rooms, looking for
	// a room whose ID value matches the parameter.
	for _, r := range rooms {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r.King)
			return
		}
	}

	// If no room with the specified ID was found,
	// c.Abort with an error 404 (Not Found).
	c.AbortWithStatus(http.StatusNotFound)
}

// curl 