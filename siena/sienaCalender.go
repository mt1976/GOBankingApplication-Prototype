package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var calenderSQL = "ID, 	Type, 	Date, 	Time, 	ShortName, 	Description"
var sqlCALID, sqlCALType, sqlCALDate, sqlCALTime, sqlCALShortName, sqlCALDescription sql.NullString

//calenderPage is cheese
type calenderListPage struct {
	UserMenu      []application.AppMenuItem
	UserRole      string
	UserNavi      string
	Title         string
	PageTitle     string
	CalenderCount int
	CalenderList  []calenderItem
}

//calenderPage is cheese
type calenderPage struct {
	UserMenu      []application.AppMenuItem
	UserRole      string
	UserNavi      string
	Title         string
	PageTitle     string
	CalenderCount int
	CalendarList  []calenderItem
}

//calenderItem is cheese
type calenderItem struct {
	ID          string
	Type        string
	Date        string
	DateInt     time.Time
	Time        string
	TimeInt     time.Time
	ShortName   string
	Description string
}

func listcalenderHandler(w http.ResponseWriter, r *http.Request) {
// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
// Code Continues Below

	wctProperties := application.GetProperties(application.APPCONFIG)
	tmpl := "listcalender"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []calenderItem
	noItems, returnList, _ := getcalenderList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pagecalenderList := calenderListPage{
		Title:         wctProperties["appname"],
		PageTitle:     "List Siena Calenders",
		CalenderCount: noItems,
		CalenderList:  returnList,
		UserRole:      globals.UserRole,
		UserNavi:      globals.UserNavi,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pagecalenderList)

}

// getcalenderList read all employees
func getcalenderList(db *sql.DB) (int, []calenderItem, error) {
	mssqlConfig := application.GetProperties(application.SQLCONFIG)

	var calenderList []calenderItem

	tsqlDeals := fmt.Sprintf("SELECT %s FROM %s.sienaCalenderDeals;", calenderSQL, mssqlConfig["schema"])
	countDeal, calenderDealList, _, _ := fetchcalenderData(db, tsqlDeals)
	calenderList = append(calenderList, calenderDealList...)

	tsqlDealLegs := fmt.Sprintf("SELECT %s FROM %s.sienaCalenderDealLegs;", calenderSQL, mssqlConfig["schema"])
	countDealLegs, calenderDealLegsList, _, _ := fetchcalenderData(db, tsqlDealLegs)
	calenderList = append(calenderList, calenderDealLegsList...)

	tsqlHolidays := fmt.Sprintf("SELECT %s FROM %s.sienaCalenderHoliday;", calenderSQL, mssqlConfig["schema"])
	countHolidays, calenderHolidaysList, _, _ := fetchcalenderData(db, tsqlHolidays)
	calenderList = append(calenderList, calenderHolidaysList...)

	count := countDeal + countDealLegs + countHolidays

	return count, calenderList, nil
}

// getcalenderList read all employees
func getcalender(db *sql.DB, id string) (int, calenderItem, error) {
	mssqlConfig := application.GetProperties(application.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.calender WHERE Code='%s';", calenderSQL, mssqlConfig["schema"], id)
	_, _, calender, _ := fetchcalenderData(db, tsql)
	return 1, calender, nil
}

// fetchcalenderData read all employees
func fetchcalenderData(db *sql.DB, tsql string) (int, []calenderItem, calenderItem, error) {

	var calender calenderItem
	var calenderList []calenderItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, calender, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCALID, &sqlCALType, &sqlCALDate, &sqlCALTime, &sqlCALShortName, &sqlCALDescription)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, calender, err
		}

		calender.ID = sqlCALID.String
		calender.Type = sqlCALType.String
		calender.Date = sqlCALDate.String
		calender.Time = sqlCALTime.String
		calender.ShortName = sqlCALShortName.String
		calender.Description = sqlCALDescription.String

		calenderList = append(calenderList, calender)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, calenderList, calender, nil
}
