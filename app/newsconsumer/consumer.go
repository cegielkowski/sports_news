package newsconsumer

import (
	"context"
	"sports_news/domain"
)

type consumerTask struct {
	consumers []domain.NewsConsumer
	quantity  int
}

// NewConsumerTask Will create new an consumerTask object representation of domain.ConsumerTask interface.
func NewConsumerTask(consumers []domain.NewsConsumer, quantity int) domain.ConsumerTask {
	return consumerTask{
		consumers: consumers,
		quantity:  quantity,
	}
}

// ConsumeNews Handle all consumers and fetch data from all of them.
func (c consumerTask) ConsumeNews(ctx context.Context) {
	for _, consumer := range c.consumers {
		consumer.Fetch(ctx, c.quantity)
	}
}
