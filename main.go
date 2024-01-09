package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/create", func(c *fiber.Ctx) error {
		var person Person
		if err := c.BodyParser(&person); err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
		}

		// Set QR Code
		qrCodeData := fmt.Sprintf("Name: %s\nAge: %d\nMarital Status: %s", person.Name, person.Age, person.MaritalStatus)
		person.QRCode = GenerateQRCode(qrCodeData)

		// Simpan data diri ke MongoDB
		insertedID, err := InsertPerson(person)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(insertedID)
	})

	app.Post("/scan", func(c *fiber.Ctx) error {
		var scanData struct {
			QRCode string `json:"qrCode,omitempty"`
		}
		if err := c.BodyParser(&scanData); err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
		}

		// Temukan data diri berdasarkan QR Code
		person, err := FindPersonByQRCode(scanData.QRCode)
		if err != nil {
			return c.Status(http.StatusNotFound).SendString("QR Code not found")
		}

		// Perbarui status menjadi "sudah menikah"
		err = UpdateMaritalStatus(person.ID)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(person)
	})

	app.Listen(":8080")
}
