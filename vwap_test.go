package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMatchFromStream(t *testing.T) {
	matches := make([]*match, 121)
	pairMatchAggregator[productBTCUSD] = &matchAggregator{
		Matches:                  matches,
		TotalSize:                150.7,
		TotalVolumeWeightedPrice: 42953.5,
	}

	addMatch(productBTCUSD, &match{
		Size:  3.3,
		Price: 792.5,
	})

	assert.Equal(t, 154.0, pairMatchAggregator[productBTCUSD].TotalSize)
	assert.Equal(t, 42953.5+(3.3*792.5), pairMatchAggregator[productBTCUSD].TotalVolumeWeightedPrice)
	assert.Equal(t, 122, len(pairMatchAggregator[productBTCUSD].Matches))
}

func TestAddMatchFromStream_exceedWindow(t *testing.T) {
	matches := make([]*match, 200)
	matches[0] = &match{
		Size:  2.7,
		Price: 534.7,
	}

	pairMatchAggregator[productBTCUSD] = &matchAggregator{
		Matches:                  matches,
		TotalSize:                150.7,
		TotalVolumeWeightedPrice: 42953.5,
	}

	addMatch(productBTCUSD, &match{
		Size:  3.3,
		Price: 792.5,
	})

	assert.Equal(t, 150.7+3.3-matches[0].Size, pairMatchAggregator[productBTCUSD].TotalSize)
	assert.Equal(t, 42953.5+(3.3*792.5)-(matches[0].Size*matches[0].Price), pairMatchAggregator[productBTCUSD].TotalVolumeWeightedPrice)
	assert.Equal(t, 200, len(pairMatchAggregator[productBTCUSD].Matches))
}
