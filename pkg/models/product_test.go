package models

import (
	"encoding/json"
	"errors"
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
			name:    "happycase: field Department",
			field:   "Department",
			rawJson: `{"department":["Mens","Shoes","Running Shoes"]}`,
			p: Product{
				Id:             1,
				Sku:            "NIKE-BLCK-42-M-1",
				Name:           "stan smith sneaker",
				ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
				DepartmentRaw:  []byte(depJson), // []byte is an alias of []uint8
				DepartmentJson: Deps{Department: []string{}},
			},
			err: nil,
		},
		{
			name:    "negativecase: field - Departments",
			field:   "Departments",
			rawJson: `{"department":["Mens","Shoes","Running Shoes"]}`,
			p: Product{
				Id:             1,
				Sku:            "NIKE-BLCK-42-M-1",
				Name:           "stan smith sneaker",
				ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
				DepartmentRaw:  []byte(depJson), // []byte is an alias of []uint8
				DepartmentJson: Deps{Department: []string{}},
			},
			err: ErrFieldNotFound,
		},
		{
			name:    "negativecase: field - Department with zeroed value",
			field:   "Department",
			rawJson: `{"department":["Mens","Shoes","Running Shoes"]}`,
			p: Product{
				Id:             1,
				Sku:            "NIKE-BLCK-42-M-1",
				Name:           "stan smith sneaker",
				ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
				DepartmentRaw:  []byte(nil), // []byte is an alias of []uint8
				DepartmentJson: Deps{Department: []string{}},
			},
			err: ErrZeroedValue,
		},
		{
			name:    "negativecase: field - Department with wrong json raw value",
			field:   "Department",
			rawJson: `{"departmentz":["Mens","Shoes","Running Shoes"]}`,
			p: Product{
				Id:             1,
				Sku:            "NIKE-BLCK-42-M-1",
				Name:           "stan smith sneaker",
				ImageUrl:       "Stan_Smith_Lux_Shoes_Black_IH2450_01_standard.avif",
				DepartmentRaw:  []byte(`["Mens","Shoes","Running Shoes"]`), // []byte is an alias of []uint8
				DepartmentJson: Deps{Department: []string{}},
			},
			err: ErrUnmarshal,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.p.RawToJson(c.field)

			// not expect error
			if c.err == nil && err != nil {
				t.Errorf("not expecte error, but got %v", err)
			}

			// get wrong error
			if c.err != nil && !errors.Is(err, c.err) {
				t.Errorf("mismatch error: expected (%v), but got (%v)\n", c.err, err)
			}

			// get incorrect unmarshall result
			var b []byte
			if !(errors.Is(c.err, ErrFieldNotFound) || errors.Is(c.err, ErrZeroedValue) || errors.Is(c.err, ErrUnmarshal)) {
				b, err = json.Marshal(reflect.ValueOf(c.p).FieldByName(c.field + "Json").Interface())
				if err != nil {
					t.Errorf("not expect error, but got %v", err)
				}

				if string(b) != c.rawJson {
					t.Errorf("[Marshal mismatch] expected %s\n,but got %s\n", c.rawJson, b)
				}
			}
		})
	}
}
