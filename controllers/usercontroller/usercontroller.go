package usercontroller

import (
	"net/http"
	"note-dbs/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var user []models.User
	models.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}
func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
	}

	return c.JSON(user)
}
func Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}
func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	if models.DB.Where("id= ?", id).Updates(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat mengupdate data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
		"user":    user,
	})
}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if models.DB.Delete(&user, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data berhasil Dihapus",
	})
}
