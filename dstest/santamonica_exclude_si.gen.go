// Automatically generated file. DO NOT EDIT.

package dstest

import (
	"time"

	darksky "github.com/twpayne/go-darksky"
)

func init() {
	request := Request{
		Latitude:  34.0219,
		Longitude: -118.4814,
		Time: darksky.Time{
			Time: time.Time{},
		},
		Exclude: "alerts,currently,daily,flags,minutely",
		Extend:  darksky.Extend(""),
		Lang:    darksky.Lang("en"),
		Units:   darksky.Units("si"),
	}
	// https://api.darksky.net/forecast/${DARKSKY_KEY}/34.021900,-118.481400?exclude=alerts%2Ccurrently%2Cdaily%2Cflags%2Cminutely&units=si
	forecastStr := `{"latitude":34.0219,"longitude":-118.4814,"timezone":"America/Los_Angeles","hourly":{"summary":"Partly cloudy starting tonight.","icon":"partly-cloudy-night","data":[{"time":1558713600,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":15.88,"apparentTemperature":15.88,"dewPoint":9.87,"humidity":0.68,"pressure":1015.41,"windSpeed":1.69,"windGust":1.93,"windBearing":235,"cloudCover":0.02,"uvIndex":3,"visibility":15.18,"ozone":375.95},{"time":1558717200,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":17.18,"apparentTemperature":17.18,"dewPoint":10.17,"humidity":0.63,"pressure":1015.16,"windSpeed":2.91,"windGust":2.91,"windBearing":201,"cloudCover":0.02,"uvIndex":5,"visibility":15.02,"ozone":376.55},{"time":1558720800,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":18.29,"apparentTemperature":18.29,"dewPoint":10.41,"humidity":0.6,"pressure":1014.85,"windSpeed":3.33,"windGust":3.33,"windBearing":228,"cloudCover":0.02,"uvIndex":7,"visibility":15.85,"ozone":376.31},{"time":1558724400,"summary":"Clear","icon":"clear-day","precipIntensity":0.0279,"precipProbability":0.01,"precipType":"rain","temperature":19.06,"apparentTemperature":19.06,"dewPoint":10.2,"humidity":0.56,"pressure":1014.74,"windSpeed":3.66,"windGust":4,"windBearing":243,"cloudCover":0.03,"uvIndex":8,"visibility":16.09,"ozone":374.95},{"time":1558728000,"summary":"Clear","icon":"clear-day","precipIntensity":0.0178,"precipProbability":0.01,"precipType":"rain","temperature":19.27,"apparentTemperature":19.27,"dewPoint":10.15,"humidity":0.56,"pressure":1014.54,"windSpeed":4.11,"windGust":4.99,"windBearing":242,"cloudCover":0.04,"uvIndex":8,"visibility":16.09,"ozone":372.86},{"time":1558731600,"summary":"Clear","icon":"clear-day","precipIntensity":0.0152,"precipProbability":0.01,"precipType":"rain","temperature":19.87,"apparentTemperature":19.87,"dewPoint":10.23,"humidity":0.54,"pressure":1013.95,"windSpeed":4.64,"windGust":6.03,"windBearing":240,"cloudCover":0.06,"uvIndex":8,"visibility":16.09,"ozone":370.48},{"time":1558735200,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":20.3,"apparentTemperature":20.3,"dewPoint":10.53,"humidity":0.53,"pressure":1013.46,"windSpeed":4.95,"windGust":6.58,"windBearing":246,"cloudCover":0.07,"uvIndex":6,"visibility":16.09,"ozone":367.82},{"time":1558738800,"summary":"Clear","icon":"clear-day","precipIntensity":0.0152,"precipProbability":0.01,"precipType":"rain","temperature":20.53,"apparentTemperature":20.53,"dewPoint":10.48,"humidity":0.53,"pressure":1012.98,"windSpeed":5.02,"windGust":6.97,"windBearing":252,"cloudCover":0.07,"uvIndex":4,"visibility":16.09,"ozone":364.96},{"time":1558742400,"summary":"Clear","icon":"clear-day","precipIntensity":0.0406,"precipProbability":0.01,"precipType":"rain","temperature":20.15,"apparentTemperature":20.15,"dewPoint":10.56,"humidity":0.54,"pressure":1012.62,"windSpeed":4.91,"windGust":6.99,"windBearing":250,"cloudCover":0.07,"uvIndex":2,"visibility":16.09,"ozone":362.64},{"time":1558746000,"summary":"Clear","icon":"clear-day","precipIntensity":0.0076,"precipProbability":0.01,"precipType":"rain","temperature":19.33,"apparentTemperature":19.33,"dewPoint":10.68,"humidity":0.57,"pressure":1012.21,"windSpeed":4.51,"windGust":6.83,"windBearing":226,"cloudCover":0.07,"uvIndex":1,"visibility":16.09,"ozone":361.49},{"time":1558749600,"summary":"Clear","icon":"clear-day","precipIntensity":0,"precipProbability":0,"temperature":18.06,"apparentTemperature":18.06,"dewPoint":10.74,"humidity":0.62,"pressure":1012.09,"windSpeed":3.9,"windGust":6.37,"windBearing":269,"cloudCover":0.08,"uvIndex":0,"visibility":16.09,"ozone":360.86},{"time":1558753200,"summary":"Clear","icon":"clear-night","precipIntensity":0.0254,"precipProbability":0.01,"precipType":"rain","temperature":16.67,"apparentTemperature":16.67,"dewPoint":10.83,"humidity":0.68,"pressure":1012.15,"windSpeed":3.27,"windGust":5.5,"windBearing":249,"cloudCover":0.09,"uvIndex":0,"visibility":16.09,"ozone":360.05},{"time":1558756800,"summary":"Clear","icon":"clear-night","precipIntensity":0.0432,"precipProbability":0.01,"precipType":"rain","temperature":15.45,"apparentTemperature":15.45,"dewPoint":10.98,"humidity":0.75,"pressure":1012.52,"windSpeed":2.51,"windGust":4.35,"windBearing":244,"cloudCover":0.09,"uvIndex":0,"visibility":16.09,"ozone":358.43},{"time":1558760400,"summary":"Clear","icon":"clear-night","precipIntensity":0.033,"precipProbability":0.01,"precipType":"rain","temperature":14.72,"apparentTemperature":14.72,"dewPoint":11.15,"humidity":0.79,"pressure":1013.01,"windSpeed":1.86,"windGust":2.99,"windBearing":178,"cloudCover":0.09,"uvIndex":0,"visibility":16.09,"ozone":356.71},{"time":1558764000,"summary":"Clear","icon":"clear-night","precipIntensity":0.0102,"precipProbability":0.01,"precipType":"rain","temperature":14.08,"apparentTemperature":14.08,"dewPoint":11.27,"humidity":0.83,"pressure":1013.24,"windSpeed":1.61,"windGust":2.48,"windBearing":200,"cloudCover":0.09,"uvIndex":0,"visibility":16.09,"ozone":355.69},{"time":1558767600,"summary":"Clear","icon":"clear-night","precipIntensity":0.0178,"precipProbability":0.01,"precipType":"rain","temperature":13.43,"apparentTemperature":13.43,"dewPoint":11.08,"humidity":0.86,"pressure":1013.04,"windSpeed":1.18,"windGust":1.74,"windBearing":187,"cloudCover":0.19,"uvIndex":0,"visibility":16.09,"ozone":356.09},{"time":1558771200,"summary":"Clear","icon":"clear-night","precipIntensity":0.0279,"precipProbability":0.01,"precipType":"rain","temperature":12.75,"apparentTemperature":12.75,"dewPoint":11.02,"humidity":0.89,"pressure":1012.58,"windSpeed":1.21,"windGust":1.76,"windBearing":155,"cloudCover":0.21,"uvIndex":0,"visibility":16.09,"ozone":357.34},{"time":1558774800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.033,"precipProbability":0.02,"precipType":"rain","temperature":12.09,"apparentTemperature":12.09,"dewPoint":10.96,"humidity":0.93,"pressure":1012.24,"windSpeed":1.32,"windGust":1.88,"windBearing":135,"cloudCover":0.26,"uvIndex":0,"visibility":16.09,"ozone":358.41},{"time":1558778400,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0279,"precipProbability":0.02,"precipType":"rain","temperature":11.56,"apparentTemperature":11.56,"dewPoint":10.9,"humidity":0.96,"pressure":1012.11,"windSpeed":1.44,"windGust":2.13,"windBearing":120,"cloudCover":0.31,"uvIndex":0,"visibility":16.09,"ozone":358.64},{"time":1558782000,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0203,"precipProbability":0.02,"precipType":"rain","temperature":11.29,"apparentTemperature":11.29,"dewPoint":10.83,"humidity":0.97,"pressure":1012.14,"windSpeed":1.53,"windGust":2.49,"windBearing":108,"cloudCover":0.36,"uvIndex":0,"visibility":16.09,"ozone":358.79},{"time":1558785600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0127,"precipProbability":0.02,"precipType":"rain","temperature":11.19,"apparentTemperature":11.19,"dewPoint":10.81,"humidity":0.97,"pressure":1012.19,"windSpeed":1.71,"windGust":2.79,"windBearing":98,"cloudCover":0.39,"uvIndex":0,"visibility":16.09,"ozone":359.61},{"time":1558789200,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0076,"precipProbability":0.01,"precipType":"rain","temperature":11.29,"apparentTemperature":11.29,"dewPoint":10.85,"humidity":0.97,"pressure":1012.3,"windSpeed":1.75,"windGust":2.93,"windBearing":100,"cloudCover":0.43,"uvIndex":0,"visibility":16.09,"ozone":361.91},{"time":1558792800,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0051,"precipProbability":0.01,"precipType":"rain","temperature":11.79,"apparentTemperature":11.79,"dewPoint":10.95,"humidity":0.95,"pressure":1012.53,"windSpeed":1.78,"windGust":3,"windBearing":112,"cloudCover":0.53,"uvIndex":0,"visibility":16.09,"ozone":364.98},{"time":1558796400,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0076,"precipProbability":0.02,"precipType":"rain","temperature":12.84,"apparentTemperature":12.84,"dewPoint":11.01,"humidity":0.89,"pressure":1012.81,"windSpeed":1.92,"windGust":3.06,"windBearing":127,"cloudCover":0.65,"uvIndex":1,"visibility":16.09,"ozone":367.59},{"time":1558800000,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":14.17,"apparentTemperature":14.17,"dewPoint":11.03,"humidity":0.81,"pressure":1013.03,"windSpeed":2.07,"windGust":3.12,"windBearing":136,"cloudCover":0.61,"uvIndex":2,"visibility":16.09,"ozone":369.32},{"time":1558803600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":15.73,"apparentTemperature":15.73,"dewPoint":11.03,"humidity":0.74,"pressure":1013.22,"windSpeed":2.24,"windGust":3.18,"windBearing":150,"cloudCover":0.54,"uvIndex":4,"visibility":16.09,"ozone":370.76},{"time":1558807200,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":16.94,"apparentTemperature":16.94,"dewPoint":11.03,"humidity":0.68,"pressure":1013.33,"windSpeed":2.48,"windGust":3.29,"windBearing":165,"cloudCover":0.5,"uvIndex":5,"visibility":16.09,"ozone":372.31},{"time":1558810800,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":18.01,"apparentTemperature":18.01,"dewPoint":11.03,"humidity":0.64,"pressure":1013.13,"windSpeed":2.84,"windGust":3.46,"windBearing":176,"cloudCover":0.39,"uvIndex":6,"visibility":16.09,"ozone":374.25},{"time":1558814400,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":18.68,"apparentTemperature":18.68,"dewPoint":11.03,"humidity":0.61,"pressure":1012.8,"windSpeed":3.26,"windGust":3.68,"windBearing":187,"cloudCover":0.36,"uvIndex":7,"visibility":16.09,"ozone":376.3},{"time":1558818000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":19.12,"apparentTemperature":19.12,"dewPoint":11.06,"humidity":0.6,"pressure":1012.39,"windSpeed":3.46,"windGust":3.83,"windBearing":195,"cloudCover":0.32,"uvIndex":6,"visibility":16.09,"ozone":378.04},{"time":1558821600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0076,"precipProbability":0.01,"precipType":"rain","temperature":19.16,"apparentTemperature":19.16,"dewPoint":11.16,"humidity":0.6,"pressure":1012.02,"windSpeed":3.59,"windGust":3.88,"windBearing":204,"cloudCover":0.3,"uvIndex":5,"visibility":16.09,"ozone":379.29},{"time":1558825200,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0152,"precipProbability":0.01,"precipType":"rain","temperature":18.67,"apparentTemperature":18.67,"dewPoint":11.31,"humidity":0.62,"pressure":1011.57,"windSpeed":3.59,"windGust":3.86,"windBearing":212,"cloudCover":0.29,"uvIndex":4,"visibility":16.09,"ozone":380.15},{"time":1558828800,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0229,"precipProbability":0.01,"precipType":"rain","temperature":18.03,"apparentTemperature":18.03,"dewPoint":11.41,"humidity":0.65,"pressure":1011.26,"windSpeed":3.5,"windGust":3.69,"windBearing":216,"cloudCover":0.28,"uvIndex":2,"visibility":16.09,"ozone":380.91},{"time":1558832400,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0254,"precipProbability":0.01,"precipType":"rain","temperature":17.06,"apparentTemperature":17.06,"dewPoint":11.42,"humidity":0.69,"pressure":1011.03,"windSpeed":3.31,"windGust":3.38,"windBearing":218,"cloudCover":0.29,"uvIndex":1,"visibility":16.09,"ozone":382.35},{"time":1558836000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0254,"precipProbability":0.01,"precipType":"rain","temperature":15.79,"apparentTemperature":15.79,"dewPoint":11.39,"humidity":0.75,"pressure":1011.02,"windSpeed":2.88,"windGust":2.95,"windBearing":215,"cloudCover":0.3,"uvIndex":0,"visibility":16.09,"ozone":383.61},{"time":1558839600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0229,"precipProbability":0.02,"precipType":"rain","temperature":14.62,"apparentTemperature":14.62,"dewPoint":11.37,"humidity":0.81,"pressure":1011.22,"windSpeed":2.39,"windGust":2.49,"windBearing":210,"cloudCover":0.36,"uvIndex":0,"visibility":16.09,"ozone":383.71},{"time":1558843200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0127,"precipProbability":0.02,"precipType":"rain","temperature":13.76,"apparentTemperature":13.76,"dewPoint":11.38,"humidity":0.86,"pressure":1011.55,"windSpeed":1.98,"windGust":2.43,"windBearing":195,"cloudCover":0.39,"uvIndex":0,"visibility":16.09,"ozone":381.78},{"time":1558846800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0076,"precipProbability":0.02,"precipType":"rain","temperature":13.35,"apparentTemperature":13.35,"dewPoint":11.42,"humidity":0.88,"pressure":1012.02,"windSpeed":1.81,"windGust":2.57,"windBearing":184,"cloudCover":0.55,"uvIndex":0,"visibility":16.09,"ozone":378.62},{"time":1558850400,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0102,"precipProbability":0.02,"precipType":"rain","temperature":13.08,"apparentTemperature":13.08,"dewPoint":11.42,"humidity":0.9,"pressure":1012.24,"windSpeed":1.69,"windGust":2.86,"windBearing":173,"cloudCover":0.65,"uvIndex":0,"visibility":16.09,"ozone":374.73},{"time":1558854000,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.033,"precipProbability":0.03,"precipType":"rain","temperature":12.89,"apparentTemperature":12.89,"dewPoint":11.41,"humidity":0.91,"pressure":1012.16,"windSpeed":1.69,"windGust":3.36,"windBearing":162,"cloudCover":0.77,"uvIndex":0,"visibility":16.09,"ozone":370.09},{"time":1558857600,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.1092,"precipProbability":0.05,"precipType":"rain","temperature":12.78,"apparentTemperature":12.78,"dewPoint":11.38,"humidity":0.91,"pressure":1011.8,"windSpeed":1.8,"windGust":3.99,"windBearing":129,"cloudCover":0.91,"uvIndex":0,"visibility":16.09,"ozone":364.76},{"time":1558861200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.1803,"precipProbability":0.06,"precipType":"rain","temperature":12.58,"apparentTemperature":12.58,"dewPoint":11.3,"humidity":0.92,"pressure":1011.41,"windSpeed":2.01,"windGust":4.51,"windBearing":115,"cloudCover":0.96,"uvIndex":0,"visibility":16.09,"ozone":360.19},{"time":1558864800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.1473,"precipProbability":0.06,"precipType":"rain","temperature":12.4,"apparentTemperature":12.4,"dewPoint":11.13,"humidity":0.92,"pressure":1011.12,"windSpeed":2.19,"windGust":4.82,"windBearing":117,"cloudCover":0.97,"uvIndex":0,"visibility":16.09,"ozone":356.79},{"time":1558868400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0838,"precipProbability":0.05,"precipType":"rain","temperature":12.31,"apparentTemperature":12.31,"dewPoint":10.91,"humidity":0.91,"pressure":1010.83,"windSpeed":2.27,"windGust":5.02,"windBearing":114,"cloudCover":0.96,"uvIndex":0,"visibility":16.09,"ozone":354.03},{"time":1558872000,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0584,"precipProbability":0.05,"precipType":"rain","temperature":12.34,"apparentTemperature":12.34,"dewPoint":10.71,"humidity":0.9,"pressure":1010.7,"windSpeed":2.54,"windGust":5.12,"windBearing":120,"cloudCover":0.96,"uvIndex":0,"visibility":16.09,"ozone":352.17},{"time":1558875600,"summary":"Overcast","icon":"cloudy","precipIntensity":0.1575,"precipProbability":0.1,"precipType":"rain","temperature":12.59,"apparentTemperature":12.59,"dewPoint":10.53,"humidity":0.87,"pressure":1010.91,"windSpeed":2.69,"windGust":5.1,"windBearing":162,"cloudCover":0.97,"uvIndex":0,"visibility":16.09,"ozone":350.88},{"time":1558879200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.1549,"precipProbability":0.14,"precipType":"rain","temperature":12.64,"apparentTemperature":12.64,"dewPoint":10.37,"humidity":0.86,"pressure":1011.32,"windSpeed":2.74,"windGust":5,"windBearing":189,"cloudCover":0.98,"uvIndex":0,"visibility":16.09,"ozone":350.38},{"time":1558882800,"summary":"Overcast","icon":"cloudy","precipIntensity":0.1422,"precipProbability":0.17,"precipType":"rain","temperature":12.83,"apparentTemperature":12.83,"dewPoint":10.26,"humidity":0.84,"pressure":1011.8,"windSpeed":2.55,"windGust":4.95,"windBearing":212,"cloudCover":0.99,"uvIndex":1,"visibility":16.09,"ozone":351.12},{"time":1558886400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.3099,"precipProbability":0.27,"precipType":"rain","temperature":13.42,"apparentTemperature":13.42,"dewPoint":10.22,"humidity":0.81,"pressure":1012.05,"windSpeed":2.59,"windGust":5,"windBearing":191,"cloudCover":0.99,"uvIndex":2,"visibility":16.09,"ozone":353.43}]},"offset":-7}`
	defaultForecasts[request] = forecastStr
}
