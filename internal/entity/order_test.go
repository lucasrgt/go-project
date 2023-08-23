package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	// arrange
	order := Order{}

	// act
	result := order.Validate()

	// assert
	assert.Error(t, result, "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	// arrange
	order := Order{ID: "123"}

	// act
	result := order.Validate()

	// assert
	assert.Error(t, result, "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	// arrange
	order := Order{
		ID:    "123",
		Price: 10.0,
	}

	// act
	result := order.Validate()

	// assert
	assert.Error(t, result, "tax is blank")
}

func TestFinalPrice(t *testing.T) {
	// arrange
	order := Order{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	// act
	err := order.Validate()
	result := order.CalculateFinalPrice()

	// assert
	expected := 11.0

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
