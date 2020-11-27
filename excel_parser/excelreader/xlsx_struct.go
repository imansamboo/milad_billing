package excelreader


var Address2Header map[string]string
var Header2Entity map[string]string

func init()  {
	Address2Header = map[string]string {
		"A": "Sim Number",
		"B": "Account No",
		"C": "Invoice No",
		"D": "Ocean Region",
		"E": "Date",
		"F": "Time",
		"G": "Originator No",
		"H": "Subscriber",
		"I": "Destination No",
		"J": "Volume",
		"K": "Unit",
		"L": "Rate",
		"M": "Total Charge",
		"N": "Equipment Type",
		"O": "Call To Call Type",
		"P": "Call Identifier ID",
		"Q": "Originator Country",
		"R": "Destination Country",
		"S": "Provider",
		"T": "Upstream Rate",
		"U": "Downstream Rate",
		"V": "Data Session ID",
		"W": "APN",
	}

	Header2Entity = map[string]string{
		"Sim Number": "sip_number",
		"Account No": "account_number",
		"Invoice No": "invoice_number",
		"Ocean Region": "ocean_region",
		"Date": "date",
		"Time": "time",
		"Originator No": "originator_number",
		"Subscriber": "subscriber",
		"Destination No": "destination_number",
		"Volume": "volume",
		"Unit": "unit",
		"Rate": "rate",
		"Total Charge": "total_charge",
		"Equipment Type": "equipment_type",
		"Call To Call Type": "call_2_call_voice",
		"Call Identifier ID": "call_identifier_id",
		"Originator Country": "originator_country",
		"Destination Country": "destination_country",
		"Provider": "provider",
		"Upstream Rate": "upstream_rate",
		"Downstream Rate": "downstream_rate",
		"Data Session ID": "data_session_id",
		"APN": "apn",
	}
}



//sqlStr := "INSERT INTO test(n1, n2, n3) VALUES "
//vals := []interface{}{}
//
//for _, row := range data {
//sqlStr += "(?, ?, ?),"
//vals = append(vals, row["v1"], row["v2"], row["v3"])
//}
////trim the last ,
//sqlStr = sqlStr[0:len(sqlStr)-1]
////prepare the statement
//stmt, _ := db.Prepare(sqlStr)
//
////format all vals at once
//res, _ := stmt.Exec(vals...)
