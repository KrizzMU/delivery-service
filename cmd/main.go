package main

import (
	"fmt"
	"time"

	"github.com/KrizzMU/delivery-service/internal/transport/nats/sender"
	"github.com/KrizzMU/delivery-service/internal/transport/nats/subscriber"
)

func main() {
	natsURL := "stan://localhost:4222"
	clusterID := "alancid"
	channel := "Krizzis"
	subID := "subs"
	sendID := "send"

	sender := sender.NewSender(natsURL, clusterID, sendID)
	sub := subscriber.NewSub(natsURL, clusterID, subID)

	go sender.SendFake(channel)
	time.Sleep(time.Second * 5)
	fmt.Println("Started listening to the channel")
	go sub.SubToChannel(channel)

	select {}
}