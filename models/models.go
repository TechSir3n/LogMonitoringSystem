package models 

type LogEntry struct { 
	Level string `json:"level"`
	Message string `json:"message"`
}