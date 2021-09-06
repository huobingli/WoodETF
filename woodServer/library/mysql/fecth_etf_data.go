package mysql

// ARKF_ETF;
// ARKG_ETF;
// ARKK_ETF;
// ARKQ_ETF;
// ARKW_ETF;
// ARKX_ETF;

type ARK_ETF struct {
	Ark_Date         string `db:"ark_date"`         // 日期
	Ark_Stock_Name   string `db:"ark_stock_name"`   // 代码
	Ark_Shares       string `db:"ark_shares"`       // 持仓数量
	Ark_Market_Value string `db:"ark_market_value"` // 持仓市值
	Ark_Weight       string `db:"ark_weight"`       // 占比ETF
}

type ARK_STOCK struct {
}
