package dto

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

func ConvertSlice[S any, D any](srcSlice []S) ([]D, error) {
	out := make([]D, 0, len(srcSlice))
	for _, item := range srcSlice {
		dto, err := ConvertToDTO[S, D](item)
		if err != nil {
			return nil, err
		}
		out = append(out, dto)
	}
	return out, nil
}

// ConvertToDTO takes a source value of type S and returns a newly‐built D.
// It will copy over any fields that share the same name, handling:
//   - identical types (AssignableTo)
//   - sql.NullString → string
//   - sql.NullTime   → *time.Time
//
// It ignores any destination fields that do not have a matching source field.
//
// Usage:
//
//	dto, err := ConvertToDTO[models.Blogposts, models.BlogpostDTO](srcValue)
func ConvertToDTO[S any, D any](src S) (D, error) {
	var dst D

	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() == reflect.Ptr {
		// If someone passed *S, we want to work with the underlying value
		srcVal = srcVal.Elem()
		srcType = srcType.Elem()
	}

	dstVal := reflect.ValueOf(&dst).Elem() // a Value we can .Set() into
	dstType := dstVal.Type()               // reflect.Type of D

	// Make sure src is a struct (or pointer to struct)
	if srcType.Kind() != reflect.Struct {
		return dst, fmt.Errorf("ConvertToDTO: src must be a struct or *struct, got %s", srcType.Kind())
	}
	if dstType.Kind() != reflect.Struct {
		return dst, fmt.Errorf("ConvertToDTO: dst must be a struct, got %s", dstType.Kind())
	}

	// Pre‐compute the reflect.Type for NullString / NullTime so we can compare
	nullStringType := reflect.TypeOf(sql.NullString{})
	nullTimeType := reflect.TypeOf(sql.NullTime{})
	ptrTimeType := reflect.TypeOf((*time.Time)(nil))

	// Iterate destination fields and try to set each from the source
	for i := 0; i < dstVal.NumField(); i++ {
		dstFieldVal := dstVal.Field(i)
		dstFieldType := dstType.Field(i)
		dstFieldKind := dstFieldVal.Type().Kind()

		// Look up a source field with the same name
		srcFieldVal := srcVal.FieldByName(dstFieldType.Name)
		if !srcFieldVal.IsValid() {
			// No matching field in src → skip
			continue
		}
		srcFieldType := srcFieldVal.Type()

		// 1) If they are directly assignable (same type or dst is interface{} that src implements)
		if srcFieldType.AssignableTo(dstFieldVal.Type()) {
			dstFieldVal.Set(srcFieldVal)
			continue
		}

		// 2) Handle sql.NullString → string
		if srcFieldType == nullStringType && dstFieldKind == reflect.String {
			ns := srcFieldVal.Interface().(sql.NullString)
			if ns.Valid {
				dstFieldVal.SetString(ns.String)
			} else {
				dstFieldVal.SetString("") // or whatever your “null” default is
			}
			continue
		}

		// 3) Handle sql.NullTime → *time.Time
		if srcFieldType == nullTimeType && dstFieldVal.Type() == ptrTimeType {
			nt := srcFieldVal.Interface().(sql.NullTime)
			if nt.Valid {
				// We need to create a *time.Time pointing at the inner Time
				t := nt.Time
				dstFieldVal.Set(reflect.ValueOf(&t))
			} else {
				// leave it as nil (zero value for *time.Time)
				dstFieldVal.Set(reflect.Zero(ptrTimeType))
			}
			continue
		}

		// 4) Handle sql.NullTime → time.Time (non‐pointer) if needed
		if srcFieldType == nullTimeType && dstFieldVal.Type().Kind() == reflect.Struct && dstFieldVal.Type().AssignableTo(reflect.TypeOf(time.Time{})) {
			nt := srcFieldVal.Interface().(sql.NullTime)
			if nt.Valid {
				dstFieldVal.Set(reflect.ValueOf(nt.Time))
			} else {
				dstFieldVal.Set(reflect.Zero(dstFieldVal.Type()))
			}
			continue
		}

		// 5) Handle *T → T if src is a pointer to the same underlying type
		if srcFieldType.Kind() == reflect.Ptr && srcFieldType.Elem().AssignableTo(dstFieldVal.Type()) {
			if !srcFieldVal.IsNil() {
				dstFieldVal.Set(srcFieldVal.Elem())
			} else {
				dstFieldVal.Set(reflect.Zero(dstFieldVal.Type()))
			}
			continue
		}

		// Fallback: not a supported conversion, just skip silently.
		// (You could also record/return an error if you want “strict mode.”)
	}

	return dst, nil
}
