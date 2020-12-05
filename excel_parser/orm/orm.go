package orm

import (
	tools "../tools"
	vars "../vars"
	"database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "reflect"
	"strings"
	_ "time"
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
	input_blanks := "("
	count := 0
	for _, value := range vars.GetHeader2Entity() {
		entity_fields = append(entity_fields, value)
		columns_string += tools.UnderScore(value)
		input_blanks +=  "?"
		count += 1
		if count < len(vars.GetHeader2Entity()){
			columns_string += ", "
			input_blanks +=  ","
		}

	}
	input_blanks +=  ")"
	sqlStr := "INSERT INTO cdr ("+ columns_string +") VALUES "
	valueStrings := make([]string, 0, len(cdrInput))
	valueArgs := make([]interface{}, 0, len(cdrInput) * len(vars.Header2Entity))
	for _, cdr := range cdrInput {
		valueStrings = append(valueStrings, input_blanks)
		for _, entity_field := range entity_fields{
			valueArgs = append(valueArgs, vars.GetField(&cdr, entity_field))
		}
	}
	stmt := fmt.Sprintf(sqlStr + " %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}


