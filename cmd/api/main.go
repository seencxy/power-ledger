package main

import (
	"PowerLedgerGo/api"
	"PowerLedgerGo/providers"
)

func main() {
	var apiHandle api.NewsHandler

	if err := providers.Inject(&apiHandle); err != nil {
		panic("Failed to provider: " + err.Error())
	}

	apiHandle.Run()
}
