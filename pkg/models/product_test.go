package models

import (
	"encoding/json"
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

	cases := []struct {
		name    string
		field   string
		rawJson string
		p       Product
		err     error
	}{
		{
			name:    "happycase: field - Department",
			field:   "Department",
			rawJson: `{"department":["Mens","Shoes","Running Shoes"]}`,
			p: Product{
				Id:             1,
				Sku:            "NIKE-BLCK-42-M-1",
				Name:           "stan smith sneaker",
				ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
				DepartmentRaw:  []uint8([]byte(depJson)),
				DepartmentJson: Deps{Department: []string{}},
			},
			err: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.p.RawToJson(c.field)
			if c.err == nil && err != nil {
				t.Errorf("not expecte error, but got %v", err)
			}

			if reflect.ValueOf(c.p.DepartmentJson).IsZero() {
				t.Errorf("expect value, but got nil")
			}

			b, err := json.Marshal(reflect.ValueOf(c.p).FieldByName(c.field + "Json").Interface())
			if err != nil {
				t.Errorf("not expecte error, but got %v", err)
			}

			if string(b) != c.rawJson {
				t.Errorf("Marshal mismatch] expected %s\n,but got %s\n", c.rawJson, b)
			}

		})
	}
}
