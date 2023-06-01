package testhelper

import (
	"database/sql/driver"
	"reflect"

	"github.com/DATA-DOG/go-sqlmock"
)

const gormTag = "gorm"

func GenerateGormMockRowFromStructArray(datas interface{}) *sqlmock.Rows {
	return GenerateSQLMockFromStructArray(datas, gormTag)
}

const dbTag = "db"

func GenerateNativeMockRowFromStructArray(datas interface{}) *sqlmock.Rows {
	return GenerateSQLMockFromStructArray(datas, dbTag)
}

func GenerateSQLMockFromStructArray(datas interface{}, tag string) *sqlmock.Rows {
	value := reflect.ValueOf(datas)
	valueKind := value.Kind()
	if valueKind != reflect.Slice && valueKind != reflect.Array {
		return nil
	}

	if value.Len() < 1 {
		return nil
	}
	firstData := value.Index(0).Interface()

	headers := GetHeaders(firstData, tag)
	rows := sqlmock.NewRows(headers)
	for i := 0; i < value.Len(); i++ {
		valueField := value.Index(i).Interface()
		mockData := GenerateSQLMockDataFromStruct(valueField)
		rows.AddRow(mockData...)
	}

	return rows
}

func GetHeaders(data interface{}, tag string) []string {
	headers := make([]string, 0)
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = reflect.ValueOf(data).Elem().Type()
	}
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tagData := field.Tag.Get(tag)
		headers = append(headers, tagData)
	}
	return headers
}

func GenerateSQLMockDataFromStruct(data interface{}) []driver.Value {
	result := []driver.Value{}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return result
		}
		val = reflect.ValueOf(data).Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i).Interface()
		result = append(result, valueField)
	}
	return result
}
