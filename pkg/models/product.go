package models

import (
	"encoding/json"
	"errors"
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

func (p *Product) RawToJson(fieldName string) (err error) {
	// check if xxRaw exist
	rawName := fieldName + "Raw"
	_, exist := reflect.TypeOf(p).Elem().FieldByName(rawName)
	if !exist {
		err = fmt.Errorf("%w[%s]", ErrFieldNotFound, rawName)
	}

	// check if xxJson exist
	jsonName := fieldName + "Json"
	_, exist = reflect.TypeOf(p).Elem().FieldByName(jsonName)
	if !exist {
		err = errors.Join(err, fmt.Errorf("%w[%s]", ErrFieldNotFound, jsonName))
	}

	if err != nil {
		return err
	}

	// get raw/json valued fields
	rawField := reflect.ValueOf(p).Elem().FieldByName(rawName)
	jsonField := reflect.ValueOf(p).Elem().FieldByName(jsonName)

	// check if xxRaw is zerod value
	if rawField.IsZero() {
		return fmt.Errorf("%w[%s]", ErrZeroedValue, rawName)
	}

	// extract "underlying" value from rawField Value
	buf := rawField.Interface().([]uint8)

	err = json.Unmarshal(buf, jsonField.Addr().Interface())
	if err != nil {
		return fmt.Errorf("%w[%s->%s]: %w", ErrUnmarshal, rawName, jsonName, err)
	}

	return err
}
