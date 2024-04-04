package entities

type BinningStockDetail struct {
	BinDocNo  string `gorm:"column:BIN_DOC_NO"`
	BinLineNo string `gorm:"column:BIN_LINE_NO"`
	PoLineNo  string `gorm:"column:PO_LINE_NO"`
	ItemCode  string `gorm:"column:ITEM_CODE"`
	LocCode   string `gorm:"column:LOC_CODE"`
	CaseNo    string `gorm:"column:CASE_NO"`
	GrpoQty   int    `gorm:"column:GRPO_QTY"`
	//HeaderID  uint   `gorm:"column:HeaderID"`
}

/*type LogbookInsertParams struct {
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
}*/

//type BinningStockHeader struct {
//	CompanyCode string               `gorm:"column:COMPANY_CODE"`
//	PoDocNo     string               `gorm:"column:PO_DOC_NO"`
//	WHSGroup    string               `gorm:"column:WHS_GROUP"`
//	WHSCode     string               `gorm:"column:WHS_CODE"`
//	Item        []BinningStockDetail `gorm:"-"`
//}

type BinningStockHeader struct {
	CompanyCode string               `gorm:"column:COMPANY_CODE"`
	PoDocNo     string               `gorm:"column:PO_DOC_NO"`
	WHSGroup    string               `gorm:"column:WHS_GROUP"`
	WHSCode     string               `gorm:"column:WHS_CODE"`
	Item        []BinningStockDetail `gorm:"-"`
}

func (BinningStockHeader) TableName() string {
	return "atitempo0" // Replace "atitempo0" with the actual table name
}
