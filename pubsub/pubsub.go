package main

import (
	"fmt"

	"github.com/google/uuid"
)

type SubscriptionID uuid.UUID
type AggregationID uuid.UUID

type Aggregation struct {
	city string
	days int // #days for moving average computation
}

type Aggregator struct {
}

var aggregations map[AggregationID]*Aggregation
var subscriptions map[SubscriptionID]AggregationID
var m map[Aggregation]int

func subscribe(city string, movingAverageDays int) (id SubscriptionID) {
	id = SubscriptionID(uuid.New())
	return
}

func unsubscribe(id SubscriptionID) {
	delete(subscriptions, id)
}

func main() {
	aggregations = make(map[AggregationID]*Aggregation)
	subscriptions = make(map[SubscriptionID]AggregationID)
	m = make(map[Aggregation]int)
	fmt.Printf("Hi\n")
	sid := subscribe("Boston", 7)
	fmt.Printf("%v\n", uuid.UUID(sid))

	var a Aggregation
	a = Aggregation{city: "Boston", days: 3}
	m[a]++
	a = Aggregation{city: "San Jose", days: 3}
	m[a]++
	a = Aggregation{city: "Boston", days: 7}
	m[a]++
	a = Aggregation{city: "Boston", days: 3}
	m[a]++
	fmt.Printf("%v\n", m)
}
