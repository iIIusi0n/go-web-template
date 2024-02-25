package user

import (
	"log"
	"strconv"

	"api-server/config"
	"api-server/controllers/auth"
	"api-server/data"
	"github.com/gin-gonic/gin"
)

func TemporaryTokenRouter(c *gin.Context) {
	var user data.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	uid, err := data.Manager.CreateUser(&user)
	if err != nil {
		log.Println("Failed to create user", err)

		c.JSON(500, gin.H{"error": "Failed to create user"})
		return

	}

	user.ID = uid
	token, err := auth.BuildNewToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", token, 60*60*24, "/", config.ServerName, false, true)

	c.JSON(200, gin.H{"uid": uid, "token": token})
}

func GetLoggedInUserRouter(c *gin.Context) {
	uid, _ := c.Get("uid")

	user, err := data.Manager.ReadUser(uid.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read user"})
		return
	}

	c.JSON(200, user)
}

func GetUserRouter(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	user, err := data.Manager.ReadUser(idNum)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read user"})
		return
	}

	c.JSON(200, user)
}

func UpdateUserRouter(c *gin.Context) {
	uid, _ := c.Get("uid")

	var u data.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	u.ID = uid.(int)
	err := data.Manager.UpdateUser(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{"message": "User updated"})
}

func CreateUserRouter(c *gin.Context) {
	var u data.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	uid, err := data.Manager.CreateUser(&u)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"uid": uid})
}
