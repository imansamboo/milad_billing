package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	_ "time"
	_ "fmt"
	_ "reflect"
	_ "github.com/go-sql-driver/mysql"
	xlsxreader "../excelreader"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:4005)/backend?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(){
	db := getConnection()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

type CDREntity struct {
	sip_number string
	account_number string
	invoice_number string
	ocean_region string
	date string
	time string
	originator_number string
	subscriber string
	destination_number string
	volume string
	unit string
	rate string
	total_charge string
	equipment_type string
	call_2_call_voice string
	call_identifier_id string
	originator_country string
	destination_country string
	provider string
	upstream_rate string
	downstream_rate string
	data_session_id string
	apn string
}

func BatchInsert(cdrInput []CDREntity) error {
	db := getConnection()
	// make n1, n2 from for in entity field
	entity_fields := []string{}
	columns_string := ""
	input_blanks := ""
	count := 0
	for _, value := range xlsxreader.Header2Entity {
		entity_fields = append(entity_fields, value)
		columns_string += columns_string + value
		input_blanks += input_blanks + "?"
		count += 1
		if count < len(entity_fields){
			columns_string += ", "
		}

	}
	sqlStr := "INSERT INTO custom_cdr ("+ columns_string +") VALUES "
	valueStrings := make([]string, 0, len(cdrInput))
	valueArgs := make([]interface{}, 0, len(cdrInput) * len(xlsxreader.Header2Entity))
	for _, cdr := range cdrInput {
		valueStrings = append(valueStrings, "(" + input_blanks + ")")
		for entity_field := range entity_fields{
			valueArgs = append(valueArgs, getField(&cdr, string(entity_field)))
		}
	}

	stmt := fmt.Sprintf(sqlStr + " %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func getField(v *CDREntity, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}