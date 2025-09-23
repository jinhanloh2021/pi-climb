package main

import (
	"github.com/jinhanloh2021/pi-climb/scripts/internal/config"
	"github.com/jinhanloh2021/pi-climb/scripts/internal/seed"
)

func main() {
	config.LoadSeedConfig()
	seed.SeedUsers()
	seed.SeedFollows()
	seed.SeedPosts()
}
