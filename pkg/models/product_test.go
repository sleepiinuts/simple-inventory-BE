package models

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func TestRawToJson(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", string(debug.Stack()))
		}
	}()

	depJson := `{
"department": [
	"Mens",
	"Shoes",
	"Running Shoes"
]}`
	// 	depJson := `[
	// 	"Mens",
	// 	"Shoes",
	// 	"Running Shoes"
	// ]`

	p := Product{
		Id:             1,
		Sku:            "NIKE-BLCK-42-M-1",
		Name:           "stan smith sneaker",
		ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
		DepartmentRaw:  []uint8([]byte(depJson)),
		DepartmentJson: Deps{Department: []string{}},
	}
	err := p.RawToJson("Department")
	if err != nil {
		t.Errorf("not expecte error, but got %v", err)
	}

	if reflect.ValueOf(p.DepartmentJson).IsZero() {
		t.Errorf("expect value, but got nil")
	}

	t.Logf("JSON value: %v", p.DepartmentJson)

}
