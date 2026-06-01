package storage

import (
	"errors"
	"sync"
)

type MarketMatrix struct {
	mu     sync.RWMutex
	Prices [][]float64
}

func NewMarketMatrix() *MarketMatrix {
	const totalCombinations = 4455100
	
	return &MarketMatrix{
		Prices: make([][]float64, totalCombinations),
	}
}

func (mm *MarketMatrix) InsertPrice(index int, price float64) error {
	if index < 0 || index >= len(mm.Prices) {
		return errors.New("matrix index out of bounds") 
	}

	mm.mu.Lock()
	defer mm.mu.Unlock()

	mm.Prices[index] = append(mm.Prices[index], price)
	return nil
}

func (mm *MarketMatrix) GetPrices(index int) ([]float64, error) {
	if index < 0 || index >= len(mm.Prices) {
		return nil, errors.New("matrix index out of bounds")
	}

	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if len(mm.Prices[index]) == 0 {
		return nil, nil
	}
	
	pricesCopy := make([]float64, len(mm.Prices[index]))
	copy(pricesCopy, mm.Prices[index])
	return pricesCopy, nil
}