package appsupport

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
)

//CONST_CONFIG_FILE is cheese
const (
	APPCONFIG       = "properties.cfg"
	SQLCONFIG       = "mssql.cfg"
	SIENACONFIG     = "siena.cfg"
	DATEFORMATPICK  = "20060102T150405"
	DATEFORMATSIENA = "2006-01-02"
	DATEFORMATYMD   = "20060102"
	DATEFORMATUSER  = "Monday, 02 Jan 2006"
	SIENACPTYSEP    = "\u22EE"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Comment
func GetURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	log.Println("Url Param '" + paramID + "' is: " + string(key))
	return key
}

//Comment
func Snooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	log.Println("Snoooze... Zzzzzz.... ", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}

//Comment
func ArrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

//RemoveContents has a comment
func RemoveContents(dir string) error {
	log.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return err
}

//Com
func GetTemplateID(tmpl string, userRole string) string {
	templateName := "html/" + tmpl + ".html"
	roleTemplate := "html" + userRole + "/" + tmpl + ".html"
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))
	if FileExists(roleTemplate) {
		//	templateName = roleTemplate
	}
	log.Println("Using Template :", templateName)
	return templateName
}

//Coment
func GetMenuID(tmpl string, userRole string) string {
	templateName := "config/menu/" + tmpl + ".json"
	roleTemplate := "config/menu" + userRole + "/" + tmpl + ".json"
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))
	if FileExists(roleTemplate) {
		templateName = roleTemplate
	}
	log.Println("Using Menu :", templateName)
	return templateName
}

//comment
func GetNavigationID(inUserRole string) string {
	templateName := "../assets/navigation.html"
	roleTemplate := "../assets" + inUserRole + "_navigation.html"
	log.Println("Testing", templateName, FileExists(templateName))
	log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	if FileExists(roleTemplate) {
		//templateName = roleTemplate
	}
	log.Println("NAVIGATION", templateName)
	return templateName
}

//comment
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//comment
func PickEpochToDateTimeString(pickEpoch string) string {
	//pickEpochLayout := "20060102T150405"
	t, err := time.Parse(DATEFORMATPICK, pickEpoch)
	if err != nil {
		fmt.Println(err)
	}
	tString := t.Format(time.RFC1123)
	return tString
}

//comment
func StrArrayToString(inArray []string) string {

	return StrArrayToStringWithSep(inArray, "\n")
}

//comment
func StrArrayToStringWithSep(inArray []string, inSep string) string {

	outString := ""
	noRows := len(inArray)
	for ii := 0; ii < noRows; ii++ {
		outString += inArray[ii] + inSep
	}
	return outString
}

//comment
func QmBundleAdd(inBundle []string, name string, value string) []string {
	return append(inBundle, name+"¡"+value)
}

//comment
func QmBundleToString(inBundle []string) string {
	return StrArrayToStringWithSep(inBundle, ";")
}

//ipRange - a structure that holds the start and end of a range of ip addresses
type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

//comment
func GetIPAdress(r *http.Request) string {
	var ipAddress string
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, ip := range strings.Split(r.Header.Get(h), ",") {
			// header can contain spaces too, strip those out.
			ip = strings.TrimSpace(ip)
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			} else {
				ipAddress = ip
				goto Done
			}
		}
	}
Done:
	return ipAddress
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	//handle err...

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

//comment
func SqlDateToHTMLDate(inDate string) (outDate string) {
	var rtnDate string
	if inDate != "" {
		rtnDate = inDate[0:10]
	}
	return rtnDate
}

//comment
func FormatCurrency(inAmount string, inCCY string) string {
	ac := accounting.Accounting{Symbol: inCCY, Precision: 2, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 ;\u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

//comment
func FormatCurrencyFull(inAmount string, inCCY string) string {
	//thisConnection, _ := sienaConnect()
	//_, ccyData, _ := getSienaCurrency(thisConnection, inCCY)
	prec, _ := strconv.Atoi("7")
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

//comment
func FormatCurrencyDps(inAmount string, inCCY string, inPrec string) string {
	prec, _ := strconv.Atoi(inPrec)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}
