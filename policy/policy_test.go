package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 53, Capacity: 101, Latency: 23, Risk: 23, Weight: 10}
	if got := Score(signal); got != 53 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 69, Capacity: 71, Latency: 8, Risk: 7, Weight: 7}
	if got := Score(signal); got != 170 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 95, Capacity: 101, Latency: 27, Risk: 9, Weight: 7}
	if got := Score(signal); got != 168 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
