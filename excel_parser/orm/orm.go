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
	vars "../vars"
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

func BatchInsert(cdrInput []vars.CDREntity) error {
	db := getConnection()
	// make n1, n2 from for in entity field
	entity_fields := []string{}
	columns_string := ""
	input_blanks := ""
	count := 0
	for _, value := range vars.Header2Entity {
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
	valueArgs := make([]interface{}, 0, len(cdrInput) * len(vars.Header2Entity))
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

func getField(v *vars.CDREntity, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}