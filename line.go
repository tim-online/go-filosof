package filosof

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Line struct {
	RecordCode      RecordCode
	InvoiceDate     Date
	ValueDate       Date
	DebitAccount    int
	CreditAccount   int
	CostCenter      int
	CostObject      int
	InvoiceNumber   string
	ReferenceNumber string
	GrossAmount     Amount // accounting currency
	FCCode          Currency
	FCGrossAmount   Amount // foreign currency
	VATAmount       Amount // accounting currency
	FCVATAmount     Amount // foreign currency
	PostingText     string
	ARTitle         string
	ARName1         string
	ARName2         string
	ARName3         string
	ARStreetName    string
	ARZipCode       string
	ARCity          string
	ARCountry       string
	ARLanguageCode  CountryCode
	ARContactPerson string
	ARTelephone     string
	ARFax           string
	AREmailAddress  Email
	ARURL           string
	ARNetDays       int // dats to due date
}

func (l *Line) Validate() []error {
	var errs []error

	errs = append(errs, l.RecordCode.Validate()...)

	if l.ValueDate.IsZero() {
		errs = append(errs, errors.New("ValueDate is required"))
	}

	if l.DebitAccount == 0 {
		errs = append(errs, errors.New("DebitAccount is required"))
	}

	if l.CreditAccount == 0 {
		errs = append(errs, errors.New("CreditAccount is required"))
	}

	if l.PostingText == "" {
		errs = append(errs, errors.New("PostingText is required"))
	}

	return errs
}

func (l *Line) Headers() []string {
	return []string{
		"RecordCode",
		"InvoiceDate",
		"ValueDate",
		"DebitAccount",
		"CreditAccount",
		"CostCenter",
		"CostObject",
		"InvoiceNumber",
		"ReferenceNumber",
		"GrossAmount",
		"FCCode",
		"FCGrossAmount",
		"VATAmount",
		"FCVATAmount",
		"PostingText",
		"ARTitle",
		"ARName1",
		"ARName2",
		"ARName3",
		"ARStreetName",
		"ARZipCode",
		"ARCity",
		"ARCountry",
		"ARLanguageCode",
		"ARContactPerson",
		"ARTelephone",
		"ARFax",
		"AREmailAddress",
		"ARURL",
		"ARNetDays",
	}
}

func (l *Line) ToStrings() []string {
	m := l.ToMap()
	return []string{
		m["RecordCode"],
		m["InvoiceDate"],
		m["ValueDate"],
		m["DebitAccount"],
		m["CreditAccount"],
		m["CostCenter"],
		m["CostObject"],
		m["InvoiceNumber"],
		m["ReferenceNumber"],
		m["GrossAmount"],
		m["FCCode"],
		m["FCGrossAmount"],
		m["VATAmount"],
		m["FCVATAmount"],
		m["PostingText"],
		m["ARTitle"],
		m["ARName1"],
		m["ARName2"],
		m["ARName3"],
		m["ARStreetName"],
		m["ARZipCode"],
		m["ARCity"],
		m["ARCountry"],
		m["ARLanguageCode"],
		m["ARContactPerson"],
		m["ARTelephone"],
		m["ARFax"],
		m["AREmailAddress"],
		m["ARURL"],
		m["ARNetDays"],
	}
}

func (l *Line) ToMap() map[string]string {
	m := map[string]string{
		"RecordCode":      string(l.RecordCode),
		"InvoiceDate":     l.InvoiceDate.String(),
		"ValueDate":       l.ValueDate.String(),
		"DebitAccount":    strconv.Itoa(l.DebitAccount),
		"CreditAccount":   strconv.Itoa(l.CreditAccount),
		"CostCenter":      strconv.Itoa(l.CostCenter),
		"CostObject":      strconv.Itoa(l.CostObject),
		"InvoiceNumber":   l.InvoiceNumber,
		"ReferenceNumber": l.ReferenceNumber,
		"GrossAmount":     l.GrossAmount.String(),
		"FCCode":          string(l.FCCode),
		"FCGrossAmount":   l.FCGrossAmount.String(),
		"VATAmount":       l.VATAmount.String(),
		"FCVATAmount":     l.FCVATAmount.String(),
		"PostingText":     l.PostingText,
		"ARTitle":         l.ARTitle,
		"ARName1":         l.ARName1,
		"ARName2":         l.ARName2,
		"ARName3":         l.ARName3,
		"ARStreetName":    l.ARStreetName,
		"ARZipCode":       l.ARZipCode,
		"ARCity":          l.ARCity,
		"ARCountry":       string(l.ARCountry),
		"ARLanguageCode":  string(l.ARLanguageCode),
		"ARContactPerson": l.ARContactPerson,
		"ARTelephone":     l.ARTelephone,
		"ARFax":           l.ARFax,
		"AREmailAddress":  string(l.AREmailAddress),
		"ARURL":           l.ARURL,
		"ARNetDays":       strconv.Itoa(l.ARNetDays),
	}

	if m["ARNetDays"] == "0" {
		m["ARNetDays"] = ""
	}

	if m["CostCenter"] == "0" {
		m["CostCenter"] = ""
	}

	if m["CostObject"] == "0" {
		m["CostObject"] = ""
	}

	if m["FCGrossAmount"] == "0.00" {
		m["FCGrossAmount"] = ""
	}

	if m["FCVATAmount"] == "0.00" {
		m["FCVATAmount"] = ""
	}

	return m
}

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	// YYYYMMDD format
	return json.Marshal(d.Time.Format("20060102"))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try parsing filosof format
	d.Time, err = time.Parse("20060102", value)
	if err == nil {
		return err
	}

	// try parsing time in RFC format
	d.Time, err = time.Parse(time.RFC3339, value)
	return err
}

func (d *Date) String() string {
	return d.Format("20060102")
}

type RecordCode rune

func (rc *RecordCode) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*rc = RecordCode(s[0])
	return nil
}

func (rc RecordCode) Validate() []error {
	var errs []error

	if string(rc) == "" {
		errs = append(errs, errors.New("RecordCode is required"))
	}

	return errs
}

type Amount float64

func (a Amount) String() string {
	return strconv.FormatFloat(float64(a), 'f', 2, 64)
}

type Currency string

type CountryCode string

type Email string
