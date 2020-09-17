package ut_compare

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"math/rand"
	"testing"
	"time"
)

func TestEarliest(t *testing.T) {
	earliest := time.Now()
	entries := []time.Time{
		earliest,
		earliest.Add(10 * time.Minute),
		earliest.Add(20 * time.Minute),
		earliest.Add(30 * time.Minute),
		earliest.Add(40 * time.Minute),
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	l := esl.Default()
	l.Info("Test with seed", esl.Int64("seed", seed), esl.Time("earliest", earliest))

	if x := Earliest(entries...); !earliest.Equal(x) {
		t.Error(x)
	}

	entriesPtr := make([]*time.Time, 0)
	for _, entry := range entries {
		entriesPtr = append(entriesPtr, ClonePtr(&entry))
	}
	entriesPtr = append(entriesPtr, nil, nil)

	r.Shuffle(len(entriesPtr), func(i, j int) {
		entriesPtr[i], entriesPtr[j] = entriesPtr[j], entriesPtr[i]
	})

	if x := EarliestPtr(entriesPtr...); !earliest.Equal(*x) {
		t.Error(earliest, x)
	}

	if x := Earliest(); !x.IsZero() {
		t.Error(x)
	}
	if x := EarliestPtr(); x != nil {
		t.Error(x)
	}
	if x := EarliestPtr(nil); x != nil {
		t.Error(x)
	}
}

func TestLatest(t *testing.T) {
	latest := time.Now()
	entries := []time.Time{
		latest,
		latest.Add(-10 * time.Minute),
		latest.Add(-20 * time.Minute),
		latest.Add(-30 * time.Minute),
		latest.Add(-40 * time.Minute),
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	l := esl.Default()
	l.Info("Test with seed", esl.Int64("seed", seed), esl.Time("latest", latest))

	if x := Latest(entries...); !latest.Equal(x) {
		t.Error(x)
	}

	entriesPtr := make([]*time.Time, 0)
	for _, entry := range entries {
		entriesPtr = append(entriesPtr, ClonePtr(&entry))
	}
	entriesPtr = append(entriesPtr, nil, nil)

	r.Shuffle(len(entriesPtr), func(i, j int) {
		entriesPtr[i], entriesPtr[j] = entriesPtr[j], entriesPtr[i]
	})

	if x := LatestPtr(entriesPtr...); !latest.Equal(*x) {
		t.Error(latest, x)
	}

	if x := Latest(); !x.IsZero() {
		t.Error(x)
	}
	if x := LatestPtr(); x != nil {
		t.Error(x)
	}
	if x := LatestPtr(nil); x != nil {
		t.Error(x)
	}
}
