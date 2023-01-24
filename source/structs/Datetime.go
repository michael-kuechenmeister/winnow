
package structs;

import "winnow/console";
import "bytes";
import "strconv";
import "strings";
import "time";



var WEEKDAYS  = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"};
var MONTHS    = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"};
var MONTHDAYS = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
var TIMEZONES = []string{

	"BST", "BDT",                       // Bering
	"HST", "HDT", "HWT", "HPT",         // Hawaii
	"AHST", "AHDT", "AKST", "AKDT",     // Alaska
	"NST", "NDT", "NWT", "NPT",         // Nome
	"YST", "YDT", "YWT", "YPT", "YDDT", // Yukon
	"PST", "PDT", "PWT", "PPT", "PDDT", // Pacific
	"MST", "MDT", "MWT", "MPT", "MDDT", // Mountain
	"CST", "CDT", "CWT", "CPT", "CDDT", // Central America
	"EST", "EDT", "EWT", "EPT", "EDDT", // Eastern America
	"NST", "NDT", "NWT", "NPT", "NDDT", // Newfoundland
	"AST", "ADT", "APT", "AWT", "ADDT", // Atlantic

	"GMT", "BST", "IST", "BDST",        // Great Britain
	"WET", "WEST", "WEMT",              // Western Europe
	"CET", "CEST", "CEMT",              // Central Europe
	"MET", "MEST",                      // Middle Europe
	"EET", "EEST",                      // Eastern Europe

	"WAT", "WAST",                      // Western Africa
	"CAT", "CAST",                      // Central Africa
	"EAT",                              // Eastern Africa
	"SAST",                             // South Africa

	"MSK", "MSD",                       // Moscow
	"IST", "IDT", "IDDT",               // Israel
	"CST", "CDT",                       // China
	"PKT", "PKST",                      // Pakistan
	"IST",                              // India
	"HKT", "HKST", "HPT",               // Hong Kong
	"KST", "KDT",                       // Korea
	"JST", "JDT",                       // Japan

	"AWST", "AWDT",                     // Western Australia
	"ACST", "ACDT",                     // Central Australia
	"AEST", "AEDT",                     // Eastern Australia
	"WIB", "WIT", "WITA",               // Waktu Indonesia Barat/Timur/Tengah
	"PST", "PDT",                       // Philippines
	"GST", "GDT", "CHST",               // Guam / Chamorro
	"NZST", "NZDT",                     // New Zealand
	"SST",                              // Samoa

	"UTC",                              // Universal

};



func formatUint (value uint, length int) string {

	var result string;

	chunk := strconv.FormatUint(uint64(value), 10);

	if len(chunk) < length {

		var prefix strings.Builder;

		for p := 0; p < length - len(chunk); p++ {
			prefix.WriteString("0");
		}

		result = prefix.String() + chunk;

	} else {

		result = chunk;

	}

	return result;

}

func isDate (value string) bool {

	var result bool = false;

	if strings.Contains(value, "-") {

		var chunks = strings.Split(value, "-");

		if len(chunks) == 3 {
			result = true;
		}

	}

	return result;

}

func isDay (value string) bool {

	var result bool = false;

	num, err := strconv.ParseUint(value, 10, 64);

	if err == nil {

		if num >= 1 && num <= 31 {
			result = true;
		}

	}

	return result;

}

func isISO8601 (value string) bool {

	var result bool = false;

	// "2006-05-30T10:02:00"
	// "2006-05-30T10:02:00.000"
	// "2006-05-30T10:02:00.000Z"

	if strings.Contains(value, "T") && strings.HasSuffix(value, "Z") {

		var date = strings.Split(value, "T")[0];
		var time = strings.Split(value, "T")[1];

		if strings.HasSuffix(time, "Z") {
			time = time[0:len(time) - 1];
		}

		var check_date = strings.Split(date, "-");
		var check_time = strings.Split(time, ":");

		if len(check_date) == 3 && len(check_time) == 3 {
			result = true;
		}

	} else if strings.Contains(value, "T") {

		var date = strings.Split(value, "T")[0];
		var time = strings.Split(value, "T")[1];

		var check_date = strings.Split(date, "-");
		var check_time = strings.Split(time, ":");

		if len(check_date) == 3 && len(check_time) == 3 {
			result = true;
		}

	}

	return result;

}

func isMeridiem (value string) bool {

	if value == "AM" || value == "PM" {
		return true;
	}

	return false;

}

func isMonth (value string) bool {

	var result bool = false;

	for m := 0; m < len(MONTHS); m++ {

		if MONTHS[m] == value {
			result = true;
			break;
		}

	}

	return result;

}

func isTime (value string) bool {

	var result bool = false;

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":");

		if len(chunks) == 3 {
			result = true;
		}

	}

	return result;

}

func isTimezone (value string) bool {

	var result bool = false;

	for t := 0; t < len(TIMEZONES); t++ {

		if TIMEZONES[t] == value {
			result = true;
			break;
		}

	}

	return result;

}

func isWeekday (value string) bool {

	var result bool = false;

	for w := 0; w < len(WEEKDAYS); w++ {

		if WEEKDAYS[w] == value {
			result = true;
			break;
		}

	}

	return result;

}

func isLeapYear (value int) bool {

	if value % 4 != 0 {
		return false;
	} else if value % 100 != 0 {
		return true;
	} else if value % 400 != 0 {
		return false;
	} else {
		return true;
	}

}

func isYear (value string) bool {

	var result bool = false;

	num, err := strconv.ParseUint(value, 10, 64);

	if err == nil {

		if num >= 2000 {
			result = true;
		}

	}

	return result;

}

func parseDate (datetime *Datetime, value string) {

	if strings.Contains(value, "-") {

		var tmp = strings.Split(value, "-");

		if len(tmp) == 3 {
			parseYear(datetime, tmp[0]);
			parseMonth(datetime, tmp[1]);
			parseDay(datetime, tmp[2]);
		}

	}


}

func parseDay (datetime *Datetime, value string) {

	num, err := strconv.ParseUint(value, 10, 64);

	if err == nil {

		if num >= 1 && num <= 31 {
			datetime.Day = uint(num);
		}

	}

}

func parseISO8601 (datetime *Datetime, value string) {

	if strings.HasSuffix(value, "Z") {
		value = value[0:len(value) - 1];
	}

	if strings.Contains(value, "T") {

		var date = strings.Split(value, "T")[0];
		var time = strings.Split(value, "T")[1];
		var tmp  = strings.Split(date, "-");

		if len(tmp) == 3 {
			parseYear(datetime, tmp[0]);
			parseMonth(datetime, tmp[1]);
			parseDay(datetime, tmp[2]);
		}

		// Strip out milliseconds
		if strings.Contains(time, ".") {
			time = strings.Split(time, ".")[0];
		}

		if strings.Contains(time, ":") {
			parseTime(datetime, time);
		}

	}

}

func parseMonth (datetime *Datetime, value string) {

	for m := 0; m < len(MONTHS); m++ {

		if MONTHS[m] == value {
			datetime.Month = uint(m + 1);
			break;
		}

	}

	if datetime.Month == 0 {

		num, err := strconv.ParseUint(value, 10, 64);

		if err == nil {

			if num >= 1 && num <= 12 {
				datetime.Month = uint(num);
			}

		}

	}

}

func parseTime (datetime *Datetime, value string) {

	if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":");

		if len(chunks) == 3 {

			num1, err1 := strconv.ParseUint(chunks[0], 10, 64);
			num2, err2 := strconv.ParseUint(chunks[1], 10, 64);
			num3, err3 := strconv.ParseUint(chunks[2], 10, 64);

			if err1 == nil {

				if num1 >= 0 && num1 <= 24 {
					datetime.Hour = uint(num1);
				}

			}

			if err2 == nil {

				if num2 >= 0 && num2 <= 60 {
					datetime.Minute = uint(num2);
				}

			}

			if err3 == nil {

				if num3 >= 0 && num3 <= 60 {
					datetime.Second = uint(num3);
				}

			}

		}

	}

}

func parseYear (datetime *Datetime, value string) {

	num, err := strconv.ParseUint(value, 10, 64);

	if err == nil {

		if num >= 1970 {
			datetime.Year = uint(num);
		}

	}

}



type Datetime struct {

	Year   uint `json:"year"`;
	Month  uint `json:"month"`;
	Day    uint `json:"day"`;
	Hour   uint `json:"hour"`;
	Minute uint `json:"minute"`;
	Second uint `json:"second"`;

}

func IsDatetime (datetime Datetime) bool {

	if datetime.Year != 0 && datetime.Month != 0 && datetime.Day != 0 {

		if datetime.Hour != 0 || datetime.Minute != 0 || datetime.Second != 0 {
			return true;
		}

	}

	return false;

}

func NewDatetime (str string) Datetime {

	var datetime Datetime;

	var chunks []string = strings.Split(strings.TrimSpace(str), " ");
	var isZulu bool     = false;


	if isISO8601(strings.TrimSpace(str)) {

		if strings.HasSuffix(str, "Z") {
			isZulu = true;
		}

		parseISO8601(&datetime, chunks[0]);

	} else if len(chunks) == 2 {

		if strings.HasSuffix(chunks[1], "Z") {
			isZulu = true;
		}

		if isDate(chunks[0]) && isTime(chunks[1]) {
			parseDate(&datetime, chunks[0]);
			parseTime(&datetime, chunks[1]);
		}

	} else if len(chunks) == 5 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isYear(chunks[4]) {

			parseMonth(&datetime, chunks[1]);
			parseDay(&datetime, chunks[2]);
			parseTime(&datetime, chunks[3]);
			parseYear(&datetime, chunks[4]);

		}

	} else if len(chunks) == 6 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isMeridiem(chunks[4]) && isYear(chunks[5]) {

			parseMonth(&datetime, chunks[1]);
			parseDay(&datetime, chunks[2]);
			parseTime(&datetime, chunks[3]);
			parseYear(&datetime, chunks[5]);

			if chunks[4] == "PM" {
				datetime.Hour = datetime.Hour + 12;
			}

		} else if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isTimezone(chunks[4]) && isYear(chunks[5]) {

			parseMonth(&datetime, chunks[1]);
			parseDay(&datetime, chunks[2]);
			parseTime(&datetime, chunks[3]);
			parseYear(&datetime, chunks[5]);

		}

	} else if len(chunks) == 7 {

		if isWeekday(chunks[0]) && isMonth(chunks[1]) && isDay(chunks[2]) && isTime(chunks[3]) && isMeridiem(chunks[4]) && isTimezone(chunks[5]) && isYear(chunks[6]) {

			parseMonth(&datetime, chunks[1]);
			parseDay(&datetime, chunks[2]);
			parseTime(&datetime, chunks[3]);
			parseYear(&datetime, chunks[6]);

			if chunks[4] == "PM" {
				datetime.Hour = datetime.Hour + 12;
			}

		}

	}


	if isZulu == false {
		datetime.ToZulu();
	}

	return datetime;

}



func (datetime *Datetime) IsEarlierThan (other Datetime) bool {

	var result bool = false;

	if datetime.Year == other.Year {

		if datetime.Month == other.Month {

			if datetime.Day == other.Day {

				if datetime.Hour == other.Hour {

					if datetime.Minute == other.Minute {

						if datetime.Second == other.Second {
							result = false;
						} else if datetime.Second > other.Second {
							result = false;
						} else if datetime.Second < other.Second {
							result = true;
						}

					} else if datetime.Minute > other.Minute {
						result = false;
					} else if datetime.Minute < other.Minute {
						result = true;
					}

				} else if datetime.Hour > other.Hour {
					result = false;
				} else if datetime.Hour < other.Hour {
					result = true;
				}

			} else if datetime.Day > other.Day {
				result = false;
			} else if datetime.Day < other.Day {
				result = true;
			}

		} else if datetime.Month > other.Month {
			result = false;
		} else if datetime.Month < other.Month {
			result = true;
		}

	} else if datetime.Year > other.Year {
		result = false;
	} else if datetime.Year < other.Year {
		result = true;
	}

	return result;

}

func (datetime *Datetime) IsLaterThan (other Datetime) bool {

	var result bool = false;

	if datetime.Year == other.Year {

		if datetime.Month == other.Month {

			if datetime.Day == other.Day {

				if datetime.Hour == other.Hour {

					if datetime.Minute == other.Minute {

						if datetime.Second == other.Second {
							result = false;
						} else if datetime.Second > other.Second {
							result = true;
						} else if datetime.Second < other.Second {
							result = false;
						}

					} else if datetime.Minute > other.Minute {
						result = true;
					} else if datetime.Minute < other.Minute {
						result = false;
					}

				} else if datetime.Hour > other.Hour {
					result = true;
				} else if datetime.Hour < other.Hour {
					result = false;
				}

			} else if datetime.Day > other.Day {
				result = true;
			} else if datetime.Day < other.Day {
				result = false;
			}

		} else if datetime.Month > other.Month {
			result = true;
		} else if datetime.Month < other.Month {
			result = false;
		}

	} else if datetime.Year > other.Year {
		result = true;
	} else if datetime.Year < other.Year {
		result = false;
	}

	return result;

}

func (datetime *Datetime) String () string {

	var buffer bytes.Buffer;

	buffer.WriteString(formatUint(datetime.Year, 4));
	buffer.WriteString("-");
	buffer.WriteString(formatUint(datetime.Month, 2));
	buffer.WriteString("-");
	buffer.WriteString(formatUint(datetime.Day, 2));
	buffer.WriteString("T");
	buffer.WriteString(formatUint(datetime.Hour, 2));
	buffer.WriteString(":");
	buffer.WriteString(formatUint(datetime.Minute, 2));
	buffer.WriteString(":");
	buffer.WriteString(formatUint(datetime.Second, 2));
	buffer.WriteString("Z");

	return buffer.String();

}

func (datetime *Datetime) ToZulu () {

	var offset = strings.Split(time.Now().String(), " ")[2];

	if len(offset) == 5 && datetime.Year > 0 && datetime.Month > 0 && datetime.Day > 0 {

		if strings.HasPrefix(offset, "+") || strings.HasPrefix(offset, "-") {

			operator             := string(offset[0]);
			offset_hours,   err1 := strconv.ParseUint(offset[1:3], 10, 64);
			offset_minutes, err2 := strconv.ParseUint(offset[3:],  10, 64);

			if err1 == nil && err2 == nil {

				if operator == "+" {

					var year      = int(datetime.Year);
					var month     = int(datetime.Month);
					var day       = int(datetime.Day);
					var hour      = int(datetime.Hour   - uint(offset_hours));
					var minute    = int(datetime.Minute - uint(offset_minutes));

					if minute < 0 {
						hour   = hour   - 1;
						minute = minute + 60;
					}

					if hour < 0 {
						day  = day  - 1;
						hour = hour + 24;
					}

					if day < 0 {

						month = month - 1;

						if isLeapYear(year) && month == 2 {
							day = day + MONTHDAYS[month - 1] + 1;
						} else {
							day = day + MONTHDAYS[month - 1];
						}

					}

					if month < 0 {
						year  = year  - 1;
						month = month + 12;
					}

					datetime.Year   = uint(year);
					datetime.Month  = uint(month);
					datetime.Day    = uint(day);
					datetime.Hour   = uint(hour);
					datetime.Minute = uint(minute);

				} else if operator == "-" {

					var year      = int(datetime.Year);
					var month     = int(datetime.Month);
					var day       = int(datetime.Day);
					var hour      = int(datetime.Hour   + uint(offset_hours));
					var minute    = int(datetime.Minute + uint(offset_minutes));

					if minute > 60 {
						hour   = hour   + 1;
						minute = minute - 60;
					}

					if hour > 24 {
						day  = day  + 1;
						hour = hour - 24;
					}

					var monthdays = MONTHDAYS[month - 1];

					if isLeapYear(year) && month == 2 {
						monthdays = monthdays + 1;
					}

					if day > monthdays {
						month = month + 1;
						day   = day   - monthdays;
					}

					if month > 12 {
						year  = year  + 1;
						month = month - 12;
					}

					datetime.Year   = uint(year);
					datetime.Month  = uint(month);
					datetime.Day    = uint(day);
					datetime.Hour   = uint(hour);
					datetime.Minute = uint(minute);

				}

			}

		}

	}

}

func (datetime *Datetime) Log () {

	var buffer = datetime.String();
	var chunks = strings.Split(buffer, "\n");

	for c := 0; c < len(chunks); c++ {
		console.Log(chunks[c]);
	}

}

