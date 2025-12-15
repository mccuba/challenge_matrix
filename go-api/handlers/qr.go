package handlers

//import (
    //"github.com/gofiber/fiber/v2"
    //"matrix-challenge/go-api/models"
    //"matrix-challenge/go-api/services"
//)
import (
    "github.com/gofiber/fiber/v2"
    "go-api/models"
    "go-api/services"
)
func QRHandler(c *fiber.Ctx) error {
    var req models.MatrixRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
    }

    Q, R, err := services.ComputeQR(req.Matrix)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(models.QRResponse{Q: Q, R: R})
}

