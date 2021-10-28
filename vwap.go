package main

import "fmt"

type match struct {
	Size  float64
	Price float64
}

type matchAggregator struct {
	Matches                  []*match
	TotalSize                float64
	TotalVolumeWeightedPrice float64
}

const windowsSizeLimit = 200

var pairMatchAggregator = make(map[string]*matchAggregator, 0)

func addMatch(productID string, m *match) *matchAggregator {
	ma, ok := pairMatchAggregator[productID]
	if !ok {
		ma = new(matchAggregator)
		pairMatchAggregator[productID] = ma
	}

	ma.updateVwapBasedOnMatch(m)

	if len(ma.Matches) > windowsSizeLimit {
		ma.removeOldestMatch()
	}

	return nil
}

func (ma *matchAggregator) removeOldestMatch() {
	if len(ma.Matches) == 0 {
		return
	}
	oldestMatch := ma.Matches[0]
	ma.TotalSize -= oldestMatch.Size
	ma.TotalVolumeWeightedPrice -= oldestMatch.Size * oldestMatch.Price
	ma.Matches = ma.Matches[1:]
}

func (ma *matchAggregator) updateVwapBasedOnMatch(m *match) {
	ma.TotalSize += m.Size
	ma.TotalVolumeWeightedPrice += m.Size * m.Price
	ma.Matches = append(ma.Matches, m)
}

func (ma *matchAggregator) getVwap() float64 {
	return ma.TotalVolumeWeightedPrice / ma.TotalSize
}

func (ma *matchAggregator) getString() string {
	return fmt.Sprintf("VWAP: %f Window Size: %d", ma.getVwap(), len(ma.Matches))
}
