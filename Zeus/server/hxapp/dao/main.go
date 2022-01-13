package dao



type HX_Draw struct {
	HX_ID       string
	HX_NAME string
}

type HX_Container struct {
	HX_ID		string
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}