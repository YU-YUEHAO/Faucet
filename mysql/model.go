package mysql


type YGH struct {
	AddressYGH string `sql:"address" json:"address" form:"address"`
	transationNumberYGH   int64  `sql:"transationNumber" json:"transationNumber" form:"transationNumber"`
	transationtimeYGH    string `sql:"transationtime" json:"transationtime" form:"transationtime"`
}
