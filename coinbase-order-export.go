package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	gdax "github.com/preichenberger/go-gdax"
)

func orders(c *gdax.Client, allOrders bool) error {
	w := json.NewEncoder(os.Stdout)
	cursor := c.ListOrders(gdax.ListOrdersParams{Status: "done"})
	var ts []gdax.Order
	for cursor.HasMore {
		if err := cursor.NextPage(&ts); err != nil {
			panic(err.Error())
		}
		for _, t := range ts {
			if allOrders {
				w.Encode(t)
			} else {
				if t.DoneReason != "filled" {
					// Only orders that are done because they are filled should printed, so
					continue
				}
				w.Encode(t)
			}
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	secret := os.Getenv("COINBASE_SECRET")
	key := os.Getenv("COINBASE_KEY")
	passphrase := os.Getenv("COINBASE_PASSPHRASE")
	client := gdax.NewClient(secret, key, passphrase)
	allOrders := flag.Bool("all", false, "Set this flag if you want ALL orders, filled or not")
	flag.Parse()

	// Due to the fact that Rob Pike doesn't want the time lib
	// (and therefore flag lib, which uses ParseDuration) to be aware of the concept
	// of a day, there are no command-line flags for the duration in which to retrieve
	// orders.  This may change in a future version.

	orders(client, *allOrders)
}
