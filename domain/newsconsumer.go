package domain

import "context"

type NewsConsumer interface {
	Fetch(ctx context.Context, quantity int)
}

type ConsumerTask interface {
	ConsumeNews(ctx context.Context)
}
