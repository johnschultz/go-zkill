package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := &websocket.Dialer{HandshakeTimeout: 1 * time.Second}
	ws, _, err := dialer.Dial("wss://zkillboard.com/websocket/", map[string][]string{"referrer": []string{"https://goonfleet.com/index.php/user/38964-ajaxify/"}})
	if err != nil {
		fmt.Printf("error dialing websocket: %s", err)
		os.Exit(1)
	}

	if err = ws.WriteMessage(websocket.TextMessage, []byte("{\"action\":\"sub\",\"channel\":\"killstream\"}")); err != nil {
	// if err = ws.WriteMessage(websocket.TextMessage, []byte("{\"action\":\"sub\",\"channel\":\"public\"}")); err != nil {
		fmt.Printf("error writing subscription request to websocket: %s", err)
		os.Exit(2)
	}

	fmt.Fprintln(os.Stderr, "Connected and subscribed")

	for true {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			panic(fmt.Sprintf("encountered error reading from websocket: %+v", err))
		} else if msgType != websocket.TextMessage {
			fmt.Fprintf(os.Stderr, "encountered non-text data frame, skipping...")
		}
		fmt.Printf("Message received: %s\n", msg)
	}
}
