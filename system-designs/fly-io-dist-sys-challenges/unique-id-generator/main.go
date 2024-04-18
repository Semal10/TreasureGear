package main

import (
	"encoding/json"
	"log"
	"fmt"
	"os"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	// Register a handler for the "echo" message that responds with an "echo_ok".
	n.Handle("unique-ids", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]interface{}
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type.
		body["type"] = "generate_ok"
		log.Printf("%s", fmt.Sprintf("%v", msg))
		body["id"] = msg.Src + "-" + msg.Dest + body["msg_id"].(string)

		resp := map[string]interface{}{
			"type": "generate_ok",
			"id": body["id"],
		}

		// Echo the original message back with the updated message type.
		return n.Reply(msg, resp)
	})

	// Execute the node's message loop. This will run until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}