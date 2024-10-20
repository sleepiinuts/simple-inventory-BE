package models

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// xxxRaw: must be declared of type []uint for method RawToJson to work properly
type Product struct {
	Id             int     `db:"ID" json:"id"`
	Sku            string  `db:"SKU" json:"sku"`
	Name           string  `db:"NAME" json:"name"`
	ImageUrl       string  `db:"IMAGE_URL" json:"image_url"`
	DepartmentRaw  []uint8 `db:"DEP" json:"dep"`
	DepartmentJson Deps
	Price          float32 `db:"PRICE" json:"price"`
}

type Deps struct {
	Department []string `json:"department"`
}

var (
	ErrFieldNotFound = fmt.Errorf("the required field is missing")
	ErrZeroedValue   = fmt.Errorf("the required field has zeroed value")
	ErrUnmarshal     = fmt.Errorf("marshalling of raw to json has error")
)

func (p *Product) RawToJson(fieldName string) error {
	// check if xxRaw exist
	rawName := fieldName + "Raw"
	_, exist := reflect.TypeOf(*p).FieldByName(rawName)
	if !exist {
		return fmt.Errorf("%w[%s]", ErrFieldNotFound, rawName)
	}

	// check if xxJson exist
	jsonName := fieldName + "Json"
	_, exist = reflect.TypeOf(*p).FieldByName(jsonName)
	if !exist {
		return fmt.Errorf("%w[%s]", ErrFieldNotFound, jsonName)
	}

	// get raw/json valued fields
	rawField := reflect.ValueOf(*p).FieldByName(rawName)

	// check if xxRaw is zerod value
	if rawField.IsZero() {
		return fmt.Errorf("%w[%s]", ErrZeroedValue, rawName)
	}

	// extract "underlying" value from rawField Value
	buf := rawField.Interface().([]uint8)

	var err error
	switch fieldName {
	case "Department":
		err = json.Unmarshal(buf, &p.DepartmentJson)
	default:
		err = fmt.Errorf("%w[Missing unmarshal case for field %s],", ErrUnmarshal, fieldName)
	}

	if err != nil {
		return fmt.Errorf("%w[%s->%s]: %w", ErrUnmarshal, rawName, jsonName, err)
	}

	return nil
}
