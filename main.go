package validity

const (

	// --- Float rules

	// FloatValue -
	FloatValue = "value"
	// FloatStrictValue -
	FloatStrictValue = "value_strict"
	// FloatMax -
	FloatMax = "max"
	// FloatMin -
	FloatMin = "min"
	// FloatDigits -
	FloatDigits = "digits"

	// --- Integer rules

	// IntValue -
	IntValue = "value"
	// IntValueStrict -
	IntValueStrict = "value_strict"
	// IntDigitsBetween -
	IntDigitsBetween = "digits_between"
	// IntDigitsBetweenStrict -
	IntDigitsBetweenStrict = "digits_between_strict"
	// IntMax -
	IntMax = "max"
	// IntMin -
	IntMin = "min"
	// IntDigits -
	IntDigits = "digits"

	// --- Special rules

	// SpecialIBAN -
	SpecialIBAN = "iban"
	// SpecialCIF -
	SpecialCIF = "cif"
	// SpecialCnp -
	SpecialCnp = "cnp"
	// SpecialShortDate -
	SpecialShortDate = "shortDate"
	// SpecialLongDate -
	SpecialLongDate = "longDate"
	// SpecialEmail -
	SpecialEmail = "email"
)
