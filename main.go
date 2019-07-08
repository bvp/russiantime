package russiantime

import (
	"regexp"
	"strconv"
	"time"
)

func ParseDateString(s string) (dt time.Time) {
	re := regexp.MustCompile(`(?P<dom>\d{1,2}) (?P<month>[а-я]+) (?P<year>\d{4}) г\.( (?P<hours>[0-2]\d):(?P<mins>[0-5]\d))?`)
	n1 := re.SubexpNames()
	res := re.FindStringSubmatch(s)

	md := map[string]string{}
	for i, n := range res {
		md[n1[i]] = n
	}

	year, _ := strconv.ParseInt(md["year"], 10, 32)
	month := stringInCase(md["month"])
	dom, _ := strconv.ParseInt(md["dom"], 10, 32)
	hours, _ := strconv.ParseInt(md["hours"], 10, 32)
	mins, _ := strconv.ParseInt(md["mins"], 10, 32)

	dt = time.Date(int(year), month, int(dom), int(hours), int(mins), 0, 0, time.Local)

	return
}

func stringInCase(s string) (month time.Month) {
	switch s {
	case "января":
		month = time.January
	case "февраля":
		month = time.February
	case "марта":
		month = time.March
	case "апреля":
		month = time.April
	case "мая":
		month = time.May
	case "июня":
		month = time.June
	case "июля":
		month = time.July
	case "августа":
		month = time.August
	case "сентября":
		month = time.September
	case "октября":
		month = time.October
	case "ноября":
		month = time.November
	case "декабря":
		month = time.December
	}
	return
}
