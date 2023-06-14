package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
	"github.com/enzofoucaud/exrond-notifier/config"
	"github.com/enzofoucaud/exrond-notifier/jexchange"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// LOG
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	log.Info().Msg("Starting bot")

	c, err := config.GetConfig()
	if err != nil {
		log.Err(err).Msg("error getting config")
		return
	}

	switch c.LogLevel {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	for {
		tokens, err := jexchange.GetTokens()
		if err != nil {
			log.Err(err).Msg("error getting tokens")
			return
		}

		for _, token := range tokens {
			log.Debug().Msg("--------------")
			log.Debug().Msg("Checking token " + token.Name)
			price, _ := jexchange.GetPricesTokensIdentifier(token.Identifier)
			log.Debug().Float64("price.Rate", price.Rate).Msg("Token price")
			// Check if pair is WAGMI
			for _, cToken := range c.Tokens {
				if cToken.Token == token.Identifier {
					if cToken.IsBelow {
						if price.Rate <= cToken.Price {
							message := token.Name + " alert: token is below " + fmt.Sprintf("%f", price.Rate) + " " + price.Unit
							err := Discord(message, c.DiscordID, c.DiscordToken)
							if err != nil {
								log.Err(err).Msg("error sending discord notification")
							}
							log.Info().Msg("Notifier sent for " + token.Name)
						}
					}
					if cToken.IsAbove {
						if price.Rate >= cToken.Price {
							message := token.Name + " alert: token is above " + fmt.Sprintf("%f", price.Rate) + " " + price.Unit
							err := Discord(message, c.DiscordID, c.DiscordToken)
							if err != nil {
								log.Err(err).Msg("error sending discord notification")
							}
							log.Info().Msg("Notifier sent for " + token.Name)
						}
					}
				}
			}
		}
		log.Debug().Msg("--------------")

		time.Sleep(time.Duration(c.SleepTime) * time.Second)
	}
}

func Discord(message, discordID, discordToken string) error {
	client := webhook.New(snowflake.MustParse(discordID), discordToken)
	defer client.Close(context.TODO())

	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().
		SetContent(message).
		Build(),
	); err != nil {
		return err
	}

	return nil
}
