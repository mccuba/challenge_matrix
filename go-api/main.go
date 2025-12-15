//package main

//import "github.com/gofiber/fiber/v2"

//func main() {
    //app := fiber.New()

    //app.Post("/qr", func(c *fiber.Ctx) error {
        //return c.JSON(fiber.Map{
            //"message": "QR endpoint working",
        //})
    //})

    //app.Listen(":8080")
//}


package main

import (
    "github.com/gofiber/fiber/v2"
    "go-api/handlers"
)

func main() {
    app := fiber.New()

    app.Post("/qr", handlers.QRHandler)

    app.Listen(":8080")
}
