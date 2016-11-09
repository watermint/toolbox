package main

import (
	"github.com/watermint/toolbox/infra"
)

func main() {
	infra.InfraStartup()
	defer infra.InfraShutdown()
}
