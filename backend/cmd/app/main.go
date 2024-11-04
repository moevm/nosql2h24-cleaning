package main

import (
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/app"
	"github.com/moevm/nosql2h24-cleaning/cleaning/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.Run(cfg)
}
