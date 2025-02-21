package controllers

import (
	"go-admin/db"
	"go-admin/models"
	"go-admin/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	var u models.User

	checkEmail := db.DB.Where("email = ?", data["email"]).First(&u).Error

	if checkEmail == nil {
		return c.JSON(fiber.Map{
			"message": "Email aleady exists!",
		})
	}

	user := &models.User{
		FullName: data["fullName"],
		Username: data["username"],
		Email:    data["email"],
		RoleId:   1,
	}

	user.SetPassword(data["password"])

	db.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email      string `json:"email"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		RememberMe bool   `json:"rememberMe"`
	}

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User

	if req.Email != "" {
		if err := db.DB.Where("LOWER(email) = ?", strings.ToLower(req.Email)).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Invalid credentials!",
			})
		}
	} else if req.Username != "" {
		if err := db.DB.Where("LOWER(username) = ?", strings.ToLower(req.Username)).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Invalid credentials!",
			})
		}
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid credentials!",
		})
	}

	if !user.ComparePassword(req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid credentials!",
		})
	}

	var cookieDuration time.Duration
	if req.RememberMe {
		cookieDuration = 365 * 24 * time.Hour // 1 year
	} else {
		cookieDuration = 24 * time.Hour // 1 day
	}

	tokenString, err := utils.GenerateJwt(user.Id.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "user_session",
		Value:    tokenString,
		HTTPOnly: true,
		Expires:  time.Now().Add(cookieDuration),
	})

	return c.JSON(fiber.Map{
		"message": "Successfully Logged In!",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("user_session")

	id, _ := utils.ParseJwt(cookie)

	uid, _ := uuid.Parse(id)

	user := models.User{
		Id: uid,
	}

	if err := db.DB.Model(&user).Preload("Role").Preload("Role.Permissions").Find(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Database error",
		})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "user_session",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	cookie := c.Cookies("user_session")

	id, _ := utils.ParseJwt(cookie)

	// * store the data parsed from the request body
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	// * store the existing user data fetched from the database
	var existingUser models.User

	db.DB.Where("id = ?", id).First(&existingUser)

	if input.FullName != "" {
		existingUser.FullName = input.FullName
	}

	if input.Email != "" && input.Email != existingUser.Email {
		var existingUserByEmail models.User
		if err := db.DB.Where("email = ?", input.Email).First(&existingUserByEmail).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email already exists",
			})
		}
		existingUser.Email = input.Email
	}

	if input.Username != "" && input.Username != existingUser.Username {
		var existingUserByUsername models.User
		if err := db.DB.Where("username = ?", input.Username).First(&existingUserByUsername).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Username already exists",
			})
		}
		existingUser.Username = input.Username
	}

	db.DB.Save(&existingUser)

	return c.JSON(existingUser)
}

func UpdatePassword(c *fiber.Ctx) error {
	cookie := c.Cookies("user_session")

	id, _ := utils.ParseJwt(cookie)

	uid, _ := uuid.Parse(id)

	var req map[string]string

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	if req["password"] != req["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	user := models.User{
		Id: uid,
	}

	user.SetPassword(req["password"])

	if err := db.DB.Model(&user).Updates(user).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Cannot update password!",
		})
	}

	return c.JSON(user)
}
