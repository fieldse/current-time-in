package citylookup

var Example_LosAngeles = CityRow{
	City:      "Los Angeles",
	CityAscii: "Los Angeles",
	Country:   "United States of America",
	Iso2:      "US",
	Iso3:      "USA",
	Province:  "California",
	StateAnsi: "CA",
	Timezone:  "America/Los_Angeles",
}
var Example_Washington = CityRow{
	City:      "Washington, D.C.",
	CityAscii: "Washington, D.C.",
	Country:   "United States of America",
	Iso2:      "US",
	Iso3:      "USA",
	Province:  "District of Columbia",
	Timezone:  "America/New_York",
}

var Example_Christiansted = CityRow{
	City:      "Christiansted",
	CityAscii: "Christiansted",
	Country:   "United States Virgin Islands",
	Iso2:      "VI",
	Iso3:      "VIR",
	Province:  "Virgin Islands",
	StateAnsi: "VI",
	Timezone:  "America/St_Thomas",
}
var Example_BacLieu = CityRow{
	City:      "Bac Lieu",
	CityAscii: "Bac Lieu",
	Country:   "Vietnam",
	Iso2:      "VN",
	Iso3:      "VNM",
	Province:  "Bạc Liêu",
	Timezone:  "Asia/Ho_Chi_Minh",
}
var Example_SaoPaulo = CityRow{
	City:      "Sao Paulo",
	CityAscii: "Sao Paulo",
	Country:   "Brazil",
	Iso2:      "BR",
	Iso3:      "BRA",
	Province:  "São Paulo",
	Timezone:  "America/Sao_Paulo",
}
var Example_FrayBentos = CityRow{
	City:      "Fray Bentos",
	CityAscii: "Fray Bentos",
	Country:   "Uruguay",
	Iso2:      "UY",
	Iso3:      "URY",
	Province:  "Río Negro",
	Timezone:  "America/Montevideo",
}

var Example_QalEhYe = CityRow{
	City:      "Qal eh-ye Now",
	CityAscii: "Qal eh-ye",
	Country:   "Afghanistan",
	Iso2:      "AF",
	Iso3:      "AFG",
	Province:  "Badghis",
	Timezone:  "Asia/Kabul",
}

var ExampleCities = []CityRow{
	Example_LosAngeles,
	Example_Washington,
	Example_Christiansted,
	Example_BacLieu,
	Example_SaoPaulo,
	Example_FrayBentos,
	Example_QalEhYe,
}
