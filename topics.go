// Package topics implements topics for resty services.
// Should be used in pub/sub distributed systems.
package topics

// Topic represents Topic for pub/sub distributed systems.
type Topic string

func (t Topic) String() string { return string(t) }

// pub/sub topics.

const (
	OrderCreated      Topic = "order.created"
	Canceled          Topic = "order.canceled"
	Paid              Topic = "order.paid"
	Cooking           Topic = "order.cooking"
	CookingFinished   Topic = "order.cooking.finished"
	WaitingForCourier Topic = "order.waiting"
	CourierTook       Topic = "order.taken"
	Delivering        Topic = "order.delivering"
	Delivered         Topic = "order.delivered"
	Closed            Topic = "order.closed"
)
