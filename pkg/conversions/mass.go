package conversions

const poundToGram = 453.59237
const ounceToGram = 28.3495231
const unitPound = "lbs"
const unitOunce = "oz"
const unitGram = "g"

var Masses = []Converter{
	PoundGram,
	OunceGram,
	GramPound,
	GramOunce,
}

func PoundGram(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitPound,
		},
		Output: Amount{
			Value: in.Value * poundToGram,
			Unit:  unitGram,
		},
	}
}

func OunceGram(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitOunce,
		},
		Output: Amount{
			Value: in.Value * ounceToGram,
			Unit:  unitGram,
		},
	}
}

func GramPound(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitGram,
		},
		Output: Amount{
			Value: in.Value / poundToGram,
			Unit:  unitPound,
		},
	}
}

func GramOunce(in Amount) Conversion {
	return Conversion{
		Input: Amount{
			Value: in.Value,
			Unit:  unitGram,
		},
		Output: Amount{
			Value: in.Value / ounceToGram,
			Unit:  unitOunce,
		},
	}
}
