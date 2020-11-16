package main

import (
    "fmt"
    "log"
    "strconv"
    "net/http"
    "strings"
    "io/ioutil"
)

const TIME_SERVER = "http://worldtimeapi.org/api/timezone/America/Los_Angeles"

func main() {
    response, err := http.Get(TIME_SERVER)
    logError(err)
    contents, err := ioutil.ReadAll(response.Body)
    logError(err)
    
    var stringResponse string = string(contents)

    defer response.Body.Close()
    
    var timeInfo map[string]string = JSONDecode(stringResponse)

    // Convert to number if you want to use arithmetically, indexing directly results in string
    unixTime, err := strconv.Atoi(timeInfo["unixtime"])
    logError(err)

    fmt.Println(unixTime)
}

func JSONDecode(body string) map[string]string {
    // Initialize map
    JSONInfo := make(map[string]string)

    for strings.Index(body, ",") != -1 {
        var nextComma int = strings.Index(body, ",")
        var nextData string = body[1:nextComma]
        body = body[nextComma + 1:]
        nextData = strings.ReplaceAll(nextData, "\"", " ")

        var separatorPos int = strings.Index(nextData, ":")
        var index string = nextData[0:separatorPos - 1]
        var value string = nextData[separatorPos + 1:]
        index, value = strings.TrimSpace(index), strings.TrimSpace(value)

        JSONInfo[index] = value
    }

    return JSONInfo
}

func JSONEncode(JSONInfo map[string]string) string {
    var JSONString string = "{"
    
    for index, value := range(JSONInfo) {
        JSONString += "\"" + index +  "\":\"" + value + "\","
    }
    JSONString = JSONString[0:len(JSONString) - 1]
    JSONString += "}"

    return JSONString
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
