package utils

import (
	"encoding/csv"
	"encoding/xml"
	"logs-monitoring/models"
	"os"
	"strings"
)

func GetFormatTypeLog(str string) string {
	content, err := os.ReadFile("/var/log/nginx/error.log.1")
	if err != nil {
		logger.Error("Error read file function -> [GetFormatTypeLog]", err.Error())
		return ""
	}

	if str == "json" { // this is the default format data in the log file, just need to return
		return string(content)

	} else if str == "xml" {
		var format models.LogFormatXML
		err = xml.Unmarshal(content, &format)
		if err != nil {
			logger.Error("Error decoding XML:", err.Error())
			return ""
		}
		
		output, err := xml.MarshalIndent(format, "", " ")
		if err != nil {
			logger.Error("Error encoding XML: ", err.Error())
			return ""
		}
		return string(output)

	} else if str == "csv" {
		reader := csv.NewReader(strings.NewReader(string(content)))
		records, err := reader.ReadAll()
		if err != nil {
			logger.Error("Error read file [csv]", err.Error())
		}

		var sb strings.Builder
		for _, record := range records {
			sb.WriteString(strings.Join(record, ","))
			sb.WriteString("\n")
		}
		return sb.String()
	}

	return ""
}
