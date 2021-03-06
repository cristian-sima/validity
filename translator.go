package validity

import (
	"fmt"
	"strings"
)

// Translater ... defines the methods for a Language translato. The translator
// is the object which transoforms the error codes to human messages
// It may transform just a particular rule or the entire map
type Translater interface {
	Translate(results *Results)
}

type floatT struct {
	floatNumber string
	value       string
	valueStrict string
	min         string
	max         string
	digits      string
}

type intT struct {
	intNumber           string
	value               string
	valueStrict         string
	digits              string
	digitsBetween       string
	digitsBetweenStrict string
	min                 string
	max                 string
}

type stringT struct {
	regexp        string
	between       string
	betweenStrict string
	minLen        string
	maxLen        string
	len           string
}

type specialT struct {
	iban      string
	cif       string
	cnp       string
	shortDate string
	longDate  string
	email     string
}

// Translator is the basic type of a translator
// It must be inherited
type Translator struct {
	floatT     floatT
	intT       intT
	stringT    stringT
	specialT   specialT
	isRequired string
	itMustBe   string
	and        string
}

//
func (translator Translator) getMessageBetween(old string) string {
	newString := strings.Replace(old, ",", " "+translator.and+" ", -1)
	return newString
}

// Translate translates the messages
func (translator Translator) Translate(results *Results) {
	for _, fieldErrors := range results.Errors {
		fieldErrors.Messages = []string{}
		for _, fullMethod := range fieldErrors.Keys {
			parts := strings.SplitN(fullMethod, ":", 2)
			method := parts[0]
			options := ""
			if len(parts) == 2 {
				options = parts[1]
			}
			humanMessage := translator.itMustBe + " " + translator.translateRule(method, options)
			fieldErrors.Messages = append(fieldErrors.Messages, humanMessage)
		}
	}
}

// translateRule translates a method into a english human message
func (translator Translator) translateRule(method string, options string) string {

	generalMessage := "There is no translation rule for [" + method + ":" + options + "]"

	getFloatMessage := func(rule string) string {
		switch rule {
		case FloatValue:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.floatT.value, betweenMessage)
		case FloatStrictValue:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.floatT.valueStrict, betweenMessage)
		case FloatMin:
			return fmt.Sprintf(translator.floatT.min, options)
		case FloatMax:
			return fmt.Sprintf(translator.floatT.max, options)
		case FloatDigits:
			return fmt.Sprintf(translator.floatT.digits, options)
		}
		return generalMessage
	}
	getIntMessage := func(rule string) string {
		switch rule {
		case IntValue:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.intT.value, betweenMessage)
		case IntValueStrict:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.intT.valueStrict, betweenMessage)
		case IntDigitsBetween:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.intT.digitsBetween, betweenMessage)
		case IntDigitsBetweenStrict:
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.intT.digitsBetweenStrict, betweenMessage)
		case IntMax:
			return fmt.Sprintf(translator.intT.max, options)
		case IntMin:
			return fmt.Sprintf(translator.intT.min, options)
		case IntDigits:
			return fmt.Sprintf(translator.intT.digits, options)
		}
		return generalMessage
	}
	getStringMessage := func(rule string) string {
		switch rule {
		case "regexp":
			return fmt.Sprintf(translator.stringT.regexp, options)
		case "between":
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.stringT.between, betweenMessage)
		case "between_strict":
			betweenMessage := translator.getMessageBetween(options)
			return fmt.Sprintf(translator.stringT.betweenStrict, betweenMessage)
		case "min_len":
			return fmt.Sprintf(translator.stringT.minLen, options)
		case "max_len":
			return fmt.Sprintf(translator.stringT.maxLen, options)
		case "len":
			return fmt.Sprintf(translator.stringT.len, options)
		}
		return generalMessage
	}
	getSpecialMessage := func(rule string) string {
		switch rule {
		case SpecialIBAN:
			return translator.specialT.iban
		case SpecialCIF:
			return translator.specialT.cif
		case SpecialCnp:
			return translator.specialT.cnp
		case SpecialShortDate:
			return translator.specialT.shortDate
		case SpecialLongDate:
			return translator.specialT.longDate
		case SpecialEmail:
			return translator.specialT.email
		}
		return generalMessage
	}

	parts := strings.SplitN(method, "#", 2)

	if len(parts) == 1 {
		switch method {
		case "REQUIRED":
			return translator.isRequired
		case "FLOAT":
			return translator.floatT.floatNumber
		case "INT":
			return translator.intT.intNumber
		}
	}
	rule := parts[1]
	switch parts[0] {
	case "FLOAT":
		return getFloatMessage(rule)
	case "INT":
		return getIntMessage(rule)
	case "STRING":
		return getStringMessage(rule)
	case "SPECIAL":
		return getSpecialMessage(rule)
	}
	return generalMessage
}
