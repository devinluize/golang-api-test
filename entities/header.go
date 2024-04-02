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
type BinningStockHeader struct {
	CompanyCode string               `gorm:"column:COMPANY_CODE"`
	PoDocNo     string               `gorm:"column:PO_DOC_NO"`
	WHSGroup    string               `gorm:"column:WHS_GROUP"`
	WHSCode     string               `gorm:"column:WHS_CODE"`
	Item        []BinningStockDetail `gorm:"-"`
}

type BinningHeader struct {
	CompanyCode string               `gorm:"column:COMPANY_CODE"`
	PoDocNo     string               `gorm:"column:PO_DOC_NO"`
	WHSGroup    string               `gorm:"column:WHS_GROUP"`
	WHSCode     string               `gorm:"column:WHS_CODE"`
	Item        []BinningStockDetail `gorm:"-"`
}
