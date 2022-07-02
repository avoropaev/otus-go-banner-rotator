package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"

	"github.com/avoropaev/otus-go-banner-rotator/cmd"
)

func main() {
	cmd.Execute(context.Background())
}
