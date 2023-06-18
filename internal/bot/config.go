package bot

import (
	"fmt"
	"os"
)

type BotConfig struct {
	token string
}

func getBotConfig() (*BotConfig, error) {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &BotConfig{token: string(token)}, nil
}
