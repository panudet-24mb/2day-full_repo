package cmd

import (
	"account_gateway/internal/config"
	"account_gateway/internal/controllers"
	"account_gateway/internal/services"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Execute() {
	config.InitTimeZone()
	config.InitConfig()
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "2day-gateway",
		AppName:       "AccountGateway API V1",
	})
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ hello")
		return c.SendString(msg)
	})
	app.Post("/open-account", accountController.OpenAccount)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
