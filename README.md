# Purpose

Coinbase Pro doesn't provide very good historical reporting capabilities, and the web-based report generation seems to only provide transaction data on a per-month basis.  This is unacceptable if you need a larger order history.

This tool will fetch all of your orders from Coinbase Pro.  In the spirit of unix simplicity, the orders are output line-by-line to stdout in JSON format.  You can easily analyze them later with whatever tools you prefer, like [Pandas](https://pandas.pydata.org).  By default, only `done` orders that are also `filled` orders will be returned.  You can get all orders by passing the `-all` command line argument.

# Installation, setup, usage

* Installation

    `go get -u github.com/adamdrake/coinbase-order-export/...`

* Setup/API Keys

    You'll need your API keys from Coinbase Pro, and the tool assumes that `COINBASE_SECRET`, `COINBASE_KEY`, and `COINBASE_PASSPHRASE` are all present as environment variables.

* Usage

    Once the environment variables are set, you can simply run the tool:

    `coinbase-order-export`

    If you would like all orders, run the tool with the `-all` flag:

    `coinbase-order-export -all`

    If you'd like to save the output to a file for further analysis, just redirect the output to a file
    
    `coinbase-order-export > orders.json`.

    The orders will be retrieved in chunks from the Coinbase Pro API, and there will be a delay of 1 second between retrievals so as not to run afoul of any rate limiting.

# Despair

If you do a lot of market orders and want to know how much money you've given to Coinbase Pro in fees, just pipe the `coinbase-order-export` output to `jq` and pipe the `fill_fees` to awk:

`coinbase-order-export | jq -r '.fill_fees' | awk '{s+=$1} END {print s}'`