package payloads

type BinningDetailResponses struct {
	BinDocNo  string `gorm:"column:BIN_DOC_NO"`
	BinLineNo string `gorm:"column:BIN_LINE_NO"`
	PoLineNo  string `gorm:"column:PO_LINE_NO"`
	ItemCode  string `gorm:"column:ITEM_CODE"`
	LocCode   string `gorm:"column:LOC_CODE"`
	CaseNo    string `gorm:"column:CASE_NO"`
	GrpoQty   int    `gorm:"column:GRPO_QTY"`
}
type BinningHeaderResponses struct {
	CompanyCode string
	PoDocNo     string
	WHSGroup    string
	WHSCode     string
	Item        []BinningDetailResponses
}
type ErrorRespones struct {
	LogSysNo int    `json:"log_sys_no"`
	Message  string `json:"message"`
	Success  bool   `json:"success"`
}
