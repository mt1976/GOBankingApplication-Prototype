package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func getURLparam(r *http.Request, paramID string) string {
	fmt.Println(paramID)
	fmt.Println(r.URL)
	key := r.FormValue(paramID)
	log.Println("Url Param '" + paramID + "' is: " + string(key))
	return key
}

func doSnooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	fmt.Println("Snoooze... Zzzzzz.... ", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}

func arrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

//RemoveContents has a comment
func RemoveContents(dir string) error {
	fmt.Println("TRASH", dir)
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

func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := getProperties(CONST_CONFIG_FILE)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	fmt.Println("WCT : Serving :", inUTL)

	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := RemoveContents(wctProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}

	//fmt.Println("recPath", wctProperties["receivepath"])

	err2 := RemoveContents(wctProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])

	err3 := RemoveContents(wctProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}

	homePageHandler(w, r)

}

func getTemplateID(tmpl string) string {

	templateName := "html/" + tmpl + ".html"
	fmt.Println("HTMLPAGE", templateName)
	return templateName
}
