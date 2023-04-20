package tests

import (
	"testing"
	"time"

	"server/models"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestModelsEqual

Tests the Equal method for the generic 'Model' interface type. Method confirms that objects with different types return as not equal and that identically defined objects return as being equal.
*/
func TestModelsEqual(t *testing.T) {
	// Initialize test objects (types that implement 'Model' interface)
	user1 := models.User{}
	user2 := models.User{}
	business1 := models.Business{}
	service1 := models.Service{
		BusinessID:    128,
		Name:          "Planks & Pilates",
		Description:   "I've heard Kylo Ren has an 8-pack. That Kylo Ren is shredded.",
		StartDateTime: time.Date(2023, 04, 20, 17, 30, 00, 00, time.Local),
		Length:        30,
		Capacity:      20,
		CancelFee:     0,
		Price:         2000,
	}
	service2 := models.Service{
		BusinessID:    128,
		Name:          "Planks & Pilates",
		Description:   "I've heard Kylo Ren has an 8-pack. That Kylo Ren is shredded.",
		StartDateTime: time.Date(2023, 04, 20, 17, 30, 00, 00, time.Local),
		Length:        30,
		Capacity:      20,
		CancelFee:     0,
		Price:         2000,
	}

	unequalFields1, equal1 := models.Equal(&user1, &user2)
	unequalFields2, equal2 := models.Equal(&business1, &user2)
	unequalFields3, equal3 := models.Equal(&service1, &service2)

	assert.Truef(t, equal1, "User objects should have matching types (EqualTypes result 1 -- User vs. User). Unequal fields:  %v", unequalFields1)
	assert.Falsef(t, equal2, "Business and User objects should have mismatched types. Unequal fields:  %v", unequalFields2)
	assert.Truef(t, equal3, "Identically defined Service objects should be equal. Unequal fields:  %v", unequalFields3)
}
