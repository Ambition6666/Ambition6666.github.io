package question

type Question struct {
	Qid      int    `gorm:"autoIncrement"`
	Templet  string `gorm:"index"`
	Language string `gorm:"index"`
	Data     string `json:"data"`
	Answer   string `json:"answer"`
}
type MidData struct {
	L string
	T string
}
type MidDailyData struct {
	L string
	T []string
}
type MidAnData struct {
	Q int
	L string
	T string
	A []string
}
type Que struct {
	Templet  string `json:"templet"`
	Language string `json:"language"`
	Data     string `json:"data"`
	Answer   string `json:"answer"`
}
