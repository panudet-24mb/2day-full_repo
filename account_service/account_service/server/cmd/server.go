package cmd

import (
	"account_service/config"
	"account_service/events"
	"account_service/repositories"
	"account_service/services"
	"context"

	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func Execute() {
	config.InitConfig()
	config.InitTimeZone()
	db := config.InitDatabase()

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to kafka group: %s\n", viper.GetStringSlice("kafka.group"))
	defer consumer.Close()

	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("Account consumer started...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}
}
