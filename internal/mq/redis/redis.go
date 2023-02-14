package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type ChannelMQ struct {
	client      redis.UniversalClient
	channelName string
	pubSub      *redis.PubSub
}

func NewChannelMQ(client redis.UniversalClient) *ChannelMQ {
	return &ChannelMQ{client: client, channelName: "test-channel"}
}

func (c *ChannelMQ) Name() string {
	return fmt.Sprintf("redis channel:%s", c.channelName)
}

func (c *ChannelMQ) Start(ctx context.Context) error {
	c.pubSub = c.client.Subscribe(ctx, c.channelName)
	ch := c.pubSub.Channel()
	go func() {
		defer func() {
			log.Printf("channel %s exiting...", c.channelName)
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-ch:
				if !ok {
					return
				}
				log.Printf("received msg from channel %s:%s", c.channelName, data.Payload)
			}
		}
	}()
	return nil
}

func (c *ChannelMQ) Stop(ctx context.Context) error {
	return c.pubSub.Close()
}
