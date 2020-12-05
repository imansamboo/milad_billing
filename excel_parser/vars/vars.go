package vars

import "reflect"

type CDREntity struct {
	SimNumber string
	AccountNumber string
	InvoiceNumber string
	OceanRegion string
	Date string
	Time string
	OriginatorNumber string
	Subscriber string
	DestinationNumber string
	Volume string
	Unit string
	Rate string
	TotalCharge string
	EquipmentType string
	Call_2_call_voice string
	CallIdentifierId string
	OriginatorCountry string
	DestinationCountry string
	Provider string
	UpstreamRate string
	DownstreamRate string
	DataSessionId string
	Apn string
}

var Address2Header map[string]string
var Header2Entity map[string]string
var ABC string
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
		"Sim Number": "SimNumber",
		"Account No": "AccountNumber",
		"Invoice No": "InvoiceNumber",
		"Ocean Region": "OceanRegion",
		"Date": "Date",
		"Time": "Time",
		"Originator No": "OriginatorNumber",
		"Subscriber": "Subscriber",
		"Destination No": "DestinationNumber",
		"Volume": "Volume",
		"Unit": "Unit",
		"Rate": "Rate",
		"Total Charge": "TotalCharge",
		"Equipment Type": "EquipmentType",
		"Call To Call Type": "Call_2_call_voice",
		"Call Identifier ID": "CallIdentifierId",
		"Originator Country": "OriginatorCountry",
		"Destination Country": "DestinationCountry",
		"Provider": "Provider",
		"Upstream Rate": "UpstreamRate",
		"Downstream Rate": "DownstreamRate",
		"Data Session ID": "DataSessionId",
		"APN": "Apn",
	}
	ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

}

func GetAddress2Header() map[string]string {
	return Address2Header
}
func GetHeader2Entity() map[string]string {
	return Header2Entity
}

func SetField(v *CDREntity, field string, value string) {
	r := reflect.ValueOf(v).Elem()
	r.FieldByName(field).SetString(value)
}

func GetField(v *CDREntity, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}



