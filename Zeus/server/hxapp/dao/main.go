package dao



type HX_Draw struct {
	HX_ID       string
	HX_NAME string
}

type HX_Container struct {
	HX_ID		string
	HX_PROPERTY	string
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var Sql_file="C:\\Users\\huobingli\\go\\src\\Zeus\\export_sav.db"