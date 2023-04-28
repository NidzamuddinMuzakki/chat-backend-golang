package controllers

import (
	"github.com/NidzamuddinMuzakki/chat-golang-backend/configs"
	"github.com/NidzamuddinMuzakki/chat-golang-backend/models"
	"github.com/NidzamuddinMuzakki/chat-golang-backend/responses"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	configs.Database.Find(&users)
	data := responses.UserResponse{
		Status:  200,
		Message: "success",
		Data:    users,
	}
	return c.Status(200).JSON(data)
}

func GetChats(c *fiber.Ctx) error {
	var users []models.Chat

	configs.Database.Order("created_at DESC, id DESC").Find(&users)
	data := responses.UserResponse{
		Status:  200,
		Message: "success",
		Data:    users,
	}
	return c.Status(200).JSON(data)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.User

	result := configs.Database.First(&dog, "name=?", id)

	data := responses.UserResponse{
		Status:  404,
		Message: "not found",
		Data:    nil,
	}
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(data)
	}
	data.Status = 200
	data.Message = "success"
	data.Data = &dog

	return c.Status(200).JSON(data)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	data := responses.UserResponse{
		Status: fiber.ErrBadRequest.Code,

		Data: nil,
	}
	if err := c.BodyParser(user); err != nil {
		data.Message = err.Error()
		return c.Status(fiber.ErrBadRequest.Code).JSON(data)
	}
	datas := configs.Database.First(&user, "name=?", user.Name)
	if datas.RowsAffected != 0 {
		data.Message = "duplicate name"
		return c.Status(fiber.ErrBadRequest.Code).JSON(data)
	}
	configs.Database.Create(&user)
	return c.Status(201).JSON(user)
}
