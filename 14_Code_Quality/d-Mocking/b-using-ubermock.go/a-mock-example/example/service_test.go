package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestService_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Add(3, 5).Return(8)

	result := mockService.Add(3, 5)
	assert.Equal(t, 8, result)
}

func TestService_Multiply(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Multiply(4, 6).Return(24)

	result := mockService.Multiply(4, 6)
	assert.Equal(t, 24, result)
}

func TestService_Add_Zero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Add(0, 0).Return(0)

	result := mockService.Add(0, 0)
	assert.Equal(t, 0, result)
}

func TestService_Add_NegativeNumbers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Add(-3, -7).Return(-10)

	result := mockService.Add(-3, -7)
	assert.Equal(t, -10, result)
}

func TestService_Multiply_ByZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Multiply(10, 0).Return(0)

	result := mockService.Multiply(10, 0)
	assert.Equal(t, 0, result)
}

func TestService_Multiply_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Multiply(4, -3).Return(-12)

	result := mockService.Multiply(4, -3)
	assert.Equal(t, -12, result)
}

func TestService_Combination(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	mockService.EXPECT().Add(10, 20).Return(30)
	mockService.EXPECT().Multiply(30, 2).Return(60)

	sum := mockService.Add(10, 20)
	result := mockService.Multiply(sum, 2)
	assert.Equal(t, 60, result)
}
