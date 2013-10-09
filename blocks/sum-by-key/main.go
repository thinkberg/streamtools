package main

import (
	"flag"
	"github.com/bitly/go-simplejson"
	"github.com/nytlabs/streamtools/streamtools"
	"log"
)

var (
	topic    = flag.String("topic", "", "topic to read from")
	window   = flag.Float64("window", 10, "size of window in seconds")
	rulePort = flag.String("port", "8080", "port to listen for new rules on")
	name     = flag.String("name", "sum", "name of block")
	key      = flag.String("key", "", "key against the value you'd like to sum")
)

func main() {

	flag.Parse()
	streamtools.SetupLogger(name)
	block := streamtools.NewStateBlock(streamtools.Sum, "count")
	// make sure the block has a key waiting for it
	rule, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		log.Fatal(err.Error())
	}
	rule.Set("window", *window)
	rule.Set("key", *key)
	block.RuleChan <- rule
	// set it going
	block.Run(*topic, *rulePort)
}