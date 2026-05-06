package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 53, Capacity: 101, Latency: 23, Risk: 23, Weight: 10}, wantScore: 53, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 69, Capacity: 71, Latency: 8, Risk: 7, Weight: 7}, wantScore: 170, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 95, Capacity: 101, Latency: 27, Risk: 9, Weight: 7}, wantScore: 168, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
