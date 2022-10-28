package controllers

import (
	"github.com/fabriciojlm/api-golang-users/src/app/controllers/dto"
	"github.com/fabriciojlm/api-golang-users/src/app/database"
	"github.com/fabriciojlm/api-golang-users/src/app/entity"
	"github.com/gin-gonic/gin"
)

// CreateUser    api-golang-users
// @Summary      Criar user
// @Tags         Users
// @Produce      json
// @Param        User  body     models.User true "User JSON"
// @Success      200  {object}  models.User
// @failure      400  {string}  string    "request não possui body"
// @Router       / [post]
func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var userCreateDTO dto.UserCreateDTO

	if err := c.ShouldBindJSON(&userCreateDTO); err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user := entity.User{
		UserName:   userCreateDTO.UserName,
		FirstName:  userCreateDTO.FirstName,
		LastName:   userCreateDTO.LastName,
		Email:      userCreateDTO.Email,
		Password:   userCreateDTO.Password,
		Phone:      userCreateDTO.Phone,
		UserStatus: userCreateDTO.UserStatus}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "usuário não criado: " + err.Error(),
		})
		return
	}

	c.JSON(201, user)
}

// CreateUsers   api-golang-users
// @Summary      Cria uma lista de usuários de acordo com o array
// @Tags         Users
// @Produce      json
// @Param        User  body     []models.User true "User JSON"
// @Success      200  {object}  []models.User
// @Router       /createWithArray [post]
func CreateUserArray(c *gin.Context) {
	db := database.GetDatabase()

	var p []entity.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "usuários não criados" + err.Error(),
		})
		return
	}

	c.JSON(200, "Todos os usuários criados")
}

// ShowAllUsers  api-golang-users
// @Summary      Lista de usuários cadastrados
// @Tags         Users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       / [get]
func ShowAllUsers(c *gin.Context) {
	db := database.GetDatabase()
	var p []entity.User
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user by name: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// GetUserById   api-golang-users
// @Summary      Trazer um usuário pelo username
// @Tags         Users
// @Produce      json
// @Param        id  path      string  true  "Id do usuario a ser acessado"
// @Success      200
// @Router       /{id} [get]
func GetUserById(c *gin.Context) {
	db := database.GetDatabase()

	id := c.Param("id")

	var p entity.User

	err := db.Where("id = ?", id).First(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "usuário não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func Health(c *gin.Context) {
	c.JSON(200, "API is healthy")
}

// GetUsersByUsername      api-golang-users
// @Summary      Trazer um usuário pelo username
// @Tags         Users
// @Produce      json
// @Param        name  path      string  true  "username do usuario a ser acessado"
// @Success      200
// @Router       ?user_name [get]
func ShowUserByName(c *gin.Context) {
	db := database.GetDatabase()

	name := c.Query("user_name")

	var p entity.User

	err := db.Where("user_name = ?", name).First(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "usuário não encontrado: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// DeleteUsersByUsername      api-golang-users
// @Summary      Deletar user
// @Tags         Users
// @Produce      json
// @Param        name  path      string  true  "username do usuário a ser apagado"
// @Success      200
// @Router       /{id} [delete]
func DeleteUser(c *gin.Context) {

	db := database.GetDatabase()
	id := c.Param("id")

	var p entity.User

	err := db.Where("id = ?", id).First(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"id": "invalido" + err.Error(),
		})
		return
	}

	err = db.Delete(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"usuário": "não encontrado " + err.Error(),
		})
		return
	}

	c.JSON(204, "")
}

// GetUsersByUsername      api-golang-users
// @Summary      Atualizar user
// @Tags         Users
// @Produce      json
// @Param        name  path      string  true  "username do usuário a ser atualizado"
// @Success      200
// @Router       /{name} [put]
func EditUser(c *gin.Context) {

	//TODO: Fix the update operation

	db := database.GetDatabase()

	id := c.Param("id")

	var user entity.User
	var userUpdateDto dto.UserUpdateDTO

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "usuário não encontrado: " + err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&userUpdateDto); err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.UserName = userUpdateDto.UserName
	user.FirstName = userUpdateDto.FirstName
	user.LastName = userUpdateDto.LastName
	user.Email = userUpdateDto.Email
	user.Password = userUpdateDto.Password
	user.Phone = userUpdateDto.Phone
	user.UserStatus = userUpdateDto.UserStatus

	db.Save(&user)

	c.JSON(200, user)

}
