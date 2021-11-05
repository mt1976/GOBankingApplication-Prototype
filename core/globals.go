package core

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jimlawless/cfg"
)

var startTime = time.Now()
var SessionToken = ""
var UUID = "authorAdjust"
var SecurityViolation = ""
var DB *sql.DB

var ApplicationProperties map[string]string
var ApplicationPropertiesDB map[string]string
var ApplicationDB *sql.DB

var InstanceProperties map[string]string
var MasterPropertiesDB map[string]string
var MasterDB *sql.DB

var SienaProperties map[string]string
var SienaPropertiesDB map[string]string
var SienaDB *sql.DB
var SienaSystemDate DateItem

var SessionManager *scs.SessionManager

var IsChildInstance bool

const (
	DATEFORMATPICK           = "20060102T150405"
	DATEFORMATSIENA          = "2006-01-02"
	DATEFORMATYMD            = "20060102"
	DATEFORMATUSER           = "02/01/2006"
	DFNANO                   = "2006-01-02T15:04:05.999999"
	DFDD                     = "02"
	DFMM                     = "01"
	DFYY                     = "06"
	DFYYYY                   = "2006"
	DFmm                     = "04"
	DFss                     = "05"
	DFhh                     = "15"
	DATETIMEFORMATSQLSERVER  = "2006-01-02 15:04:05"
	SIENACPTYSEP             = "\u22EE"
	APPCONFIG                = "application.cfg"
	SQLCONFIG                = "sienaDB.cfg"
	SIENACONFIG              = "siena.cfg"
	DATASTORECONFIG          = "applicationDB.cfg"
	INSTANCECONFIG           = "instance.cfg"
	DATETIMEFORMATUSER       = DATEFORMATUSER + " 15:04:05"
	TIMEHMS                  = "15:04:05"
	ColorReset               = "\033[0m"
	ColorRed                 = "\033[31m"
	ColorGreen               = "\033[32m"
	ColorYellow              = "\033[33m"
	ColorBlue                = "\033[34m"
	ColorPurple              = "\033[35m"
	ColorCyan                = "\033[36m"
	ColorWhite               = "\033[37m"
	SessionRole              = "1891835972"
	SessionNavi              = "6782444386"
	SessionKnowAs            = "0627218437"
	SessionUserName          = "9214790474"
	SessionTokenID           = "1516099114"
	SessionUUID              = "0663644127"
	SessionSecurityViolation = "4097340829"
	SessionAppToken          = "1117429826"
	IDSep                    = "|"
	Tick                     = "\u2713"
	WarningLabel             = "\u2757"
	Bike                     = "\u1F3CD"
	Scheduled                = "Scheduled"
	Adhoc                    = "AdHoc"
	Monitor                  = "Monitor"
	Dispatcher               = "Dispatcher"
	Aquirer                  = "Aquirer"
	Character_MapTo          = "\u21D4"
)

type Connection struct {
	count int
	Pool  []ConnectionItem
}

type ConnectionItem struct {
	ID               string
	Name             string
	DealimportIn     string
	DealimportOut    string
	StaticIn         string
	StaticOut        string
	Database         *sql.DB
	ConnectionString sienaDBItem
}

type sienaDBItem struct {
	ID         string
	Server     string
	Port       string
	User       string
	Password   string
	Database   string
	Schema     string
	Active     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//SienaBusinessDateItem is cheese
type DateItem struct {
	Today     string
	Internal  time.Time
	Siena     string
	YYYYMMDD  string
	PICKEpoch string
}

func Initialise() {
	LOG_msg("Initialisation", "Vroom")

	SessionToken = ""
	UUID = "authorAdjust"
	SecurityViolation = ""

	PreInitialise()

	//SienaSystemDate DateItem
	ApplicationProperties = getProperties(APPCONFIG)
	SienaProperties = getProperties(SIENACONFIG)
	SienaPropertiesDB = getProperties(SQLCONFIG)
	ApplicationPropertiesDB = getProperties(DATASTORECONFIG)
	InstanceProperties = getProperties(INSTANCECONFIG)
	//MasterPropertiesDB = getProperties(DATASTORECONFIG)
	//MasterPropertiesDB["database"] = "master"

	IsChildInstance = false
	if len(ApplicationPropertiesDB["instance"]) != 0 {
		IsChildInstance = true
	}
	//

	LOG_msg("Initialisation", "Connecting to application databases...")

	//MasterDB, _ = GlobalsDatabaseConnect(MasterPropertiesDB)

	ApplicationDB, _ = GlobalsDatabaseConnect(ApplicationPropertiesDB)

	SienaDB, _ = GlobalsDatabaseConnect(SienaPropertiesDB)

	LOG_success("Connections established")

	//

	//spew.Dump(ApplicationDB)
	// TODO: get a list of additional DB's to connect to (from the SRS.sienadbStore)
	// TODO: load them into the var sourceAccess []*sql.DB slice

	SessionManager = scs.New()
	life, err := time.ParseDuration(ApplicationProperties["sessionlife"])
	if err != nil {
		log.Fatal("No Session Life Found", err, life)
	}
	SessionManager.Lifetime = life
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Name = ApplicationProperties["releaseid"]
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Persist = true
	//SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Cookie.Secure = false

	LOG_success("Vroooom Vrooooom VROOOOM! " + Bike)
}

// Load a Properties File
func getProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	// For docker - if can't find properties file (create one from the template properties file)
	propertiesFileName := "config/" + inPropertiesFile
	if fileExists(propertiesFileName) {
		// Do nothign this is ok
	} else {

		ok := copyDataFile(inPropertiesFile, "config/", "config/fileSystem/config")
		if !ok {
			log.Println("Issue in Copy Function")
		}
	}

	err := cfg.Load(propertiesFileName, wctProperties)
	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
}

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PreInitialise() {

}

func writeDataFile(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return false, nil
}

func deleteDataFile(fileName string, path string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Delete        :", filePath)

	// delete file

	if fileExists(filePath) {
		var err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	LOG_info("File Deleted", fileName+" - "+path)
	return 1
}

func copyDataFile(fileName string, fromPath string, toPath string) bool {

	log.Panicln("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := ReadDataFile(fileName, fromPath)
	if err != nil {
		log.Panicln(err.Error())
	}

	ok, err2 := writeDataFile(fileName, toPath, content)
	if err2 != nil {
		log.Panicln(err2.Error())
	}

	if !ok {
		log.Panicln("Unable to Copy " + fileName + " from " + fromPath + " to " + toPath)
	}

	return true
}

// getFundsCheckList read all employees
func GetDataList(basePath string) (int, []string, error) {

	var listing []string
	//	log.Println(basePath, kind, direction, requestPath)
	pwd, _ := os.Getwd()
	//log.Println(pwd + requestPath + "/")
	files, err := ioutil.ReadDir(pwd + basePath)
	if err != nil {
		log.Fatal(err)
	}
	//spew.Dump(files)

	for _, k := range files {
		//fmt.Println("key:", k)
		listing = append(listing, k.Name())
	}

	//count, simFundsCheckList, _, _ := fetchFundsCheckData("")
	return len(files), listing, nil
}

func ReplaceWildcard(orig string, replaceThis string, withThis string) string {
	wrkThis := "{{" + replaceThis + "}}"
	//log.Printf("Replace %s with %q", wrkThis, withThis)
	return strings.ReplaceAll(orig, wrkThis, withThis)
}

func Uptime() time.Duration {
	return time.Since(startTime)
}

func Log_uptime() {
	log.Println("App Uptime    :", Uptime().String())
}

/*
* leftPad and rightPad just repoeat the padStr the indicated
* number of times
*
 */
func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
func rightPad(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen)
}

/* the Pad2Len functions are generally assumed to be padded with short sequences of strings
* in many cases with a single character sequence
*
* so we assume we can build the string out as if the char seq is 1 char and then
* just substr the string if it is longer than needed
*
* this means we are wasting some cpu and memory work
* but this always get us to want we want it to be
*
* in short not optimized to for massive string work
*
* If the overallLen is shorter than the original string length
* the string will be shortened to this length (substr)
*
 */
func rightPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
func leftPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}