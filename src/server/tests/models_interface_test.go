package tests

import (
	"testing"

	"server/models"

	"github.com/stretchr/testify/assert"
)

func TestModelsEqual(t *testing.T) {
	address1 := models.Address{}
	address2 := models.Address{}
	business1 := models.Business{}

	unequalFields1, equal1 := models.Equal(&address1, &address2)
	unequalFields2, equal2 := models.Equal(&business1, &address2)

	t.Logf("unequalFields1:  %v", unequalFields1)
	t.Logf("unequalFields2:  %v", unequalFields2)

	assert.Truef(t, equal1, "ERROR: Types do not match (EqualTypes result 1 -- Address vs. Address).")
	assert.Falsef(t, equal2, "ERROR: Equal function returned 'true' for mismatched object types (Business and Address).")
}

//  EqualTypes FUNCTION INITIALLY USED FOR DEBUGGING MODEL INTERFACE FUNCTION
// func TestEqualTypes(t *testing.T) {
// 	address1 := Address{}
// 	address2 := Address{}
// 	business1 := Business{}

// 	type1 := reflect.TypeOf(&address1)
// 	type2 := reflect.TypeOf(&address2)
// 	type3 := reflect.TypeOf(&business1)

// 	methodType1, methodType2, equalTypes1 := EqualTypes(&address1, &address2)
// 	methodType3, methodType4, equalTypes2 := EqualTypes(&business1, &address2)

// 	t.Logf("Type1 (Address1):  %s", type1)
// 	t.Logf("Type2 (Address2):  %s", type2)
// 	t.Logf("Type3 (Business1):  %s", type3)

// 	t.Logf("MethodType1 (Address1):  %s", methodType1)
// 	t.Logf("MethodType2 (Address2):  %s", methodType2)
// 	t.Logf("MethodType3 (Business1):  %s", methodType3)
// 	t.Logf("MethodType4 (Address2):  %s", methodType4)

// 	assert.Truef(t, equalTypes1, "ERROR: Types do not match (EqualTypes result 1 -- Address vs. Address).")
// 	assert.Falsef(t, equalTypes2, "ERROR: Types match (EqualTypes result 2 -- Business vs. Address).")
// 	assert.Equalf(t, type1, methodType1, "ERROR: type1 and methodType1 mismatch. '%s' does not match '%s'.", type1, methodType1)
// 	assert.Equalf(t, type2, methodType2, "ERROR: type2 and methodType2 mismatch. '%s' does not match '%s'.", type2, methodType2)
// 	assert.Equalf(t, type3, methodType3, "ERROR: type3 and methodType3 mismatch. '%s' does not match '%s'.", type3, methodType3)
// 	assert.Equalf(t, type2, methodType4, "ERROR: type2 and methodType4 mismatch. '%s' does not match '%s'.", type2, methodType4)
// }
