package models

type LogFormatJSON struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}


type LogFormatXML struct {
	Level   string `xml:"level"`
	Message string `xml:"message"`
}


type LogFormatCSV struct {
	Level   string `csv:"level"`
	Message string `csv:"message"`
}


type LogLevel struct {
	Level string `json:"level"`
}


type LogType struct {
	Type string `json:"type"`
}

