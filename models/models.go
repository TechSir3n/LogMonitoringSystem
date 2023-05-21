package models 

type LogEntry struct { 
	Level string `json:"level"`
	Message string `json:"message"`
}

type LogLevel struct  { 
	Level string `json:"level"`
}
