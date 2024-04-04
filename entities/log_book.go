package entities

import "time"

type LogbookInsertParams struct {
	LogSysNo       int64     `gorm:"column:LOG_SYS_NO"`      // Numeric(18,0)
	RecordStatus   string    `gorm:"column:RECORD_STATUS"`   // Char(1)
	ApiName        string    `gorm:"column:API_NAME"`        // Varchar(255)
	ApiMethod      string    `gorm:"column:API_METHOD"`      // Varchar(15)
	ApiUrl         string    `gorm:"column:API_URL"`         // Varchar(255)
	ApiParameters  string    `gorm:"column:API_PARAMETERS"`  // Varchar(1000)
	ApiStart       time.Time `gorm:"column:API_START"`       // Datetime
	ApiFinish      time.Time `gorm:"column:API_FINISH"`      // Datetime
	ApiStatus      string    `gorm:"column:API_STATUS"`      // Varchar(50)
	RequestHeader  string    `gorm:"column:REQUEST_HEADER"`  // Varchar(MAX)
	RequestBody    string    `gorm:"column:REQUEST_BODY"`    // Varchar(MAX)
	ResponseHeader string    `gorm:"column:RESPONSE_HEADER"` // Varchar(MAX)
	ResponseBody   string    `gorm:"column:RESPONSE_BODY"`   // Varchar(MAX)
	RowsAffected   int       `gorm:"column:ROWS_AFFECTED"`   // Int
	Success        bool      `gorm:"column:SUCCESS"`         // Bit
	Message        string    `gorm:"column:MESSAGE"`         // Varchar(MAX)
	ProcessFrom    time.Time `gorm:"column:PROCESS_FROM"`    // Datetime
	ProcessBy      string    `gorm:"column:PROCESS_BY"`      // Varchar(25)
}
