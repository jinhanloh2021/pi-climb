package main

import (
	"github.com/jinhanloh2021/beta-blocker/scripts/internal/config"
	"github.com/jinhanloh2021/beta-blocker/scripts/internal/seed"
)

func main() {
	config.LoadSeedConfig()
	// SeedUsers()
	seed.SeedFollows()
}
