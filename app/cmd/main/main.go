package main

import (
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/Vityalimbaev/Example-Backend/internal/adapter"
	"github.com/Vityalimbaev/Example-Backend/internal/logger"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	config.InitConfig()
	logger.SetupLogger()
	adapter.InitApp()
}
