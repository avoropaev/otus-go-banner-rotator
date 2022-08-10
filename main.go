package main

import (
	"context"

	"github.com/avoropaev/otus-go-banner-rotator/cmd"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cmd.Execute(context.Background())
}
