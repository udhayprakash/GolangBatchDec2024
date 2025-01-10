package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPerformAddition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalc := NewMockCalculator(ctrl)
	service := NewService(mockCalc)

	// Expect the Add method to be called with 2 and 3, and return 5
	mockCalc.EXPECT().Add(2, 3).Return(5)

	result := service.PerformAddition(2, 3)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
