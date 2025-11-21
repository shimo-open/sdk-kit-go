package common

type Lang string

const (
	Lang_ZH_CN Lang = "zh-CN"
	Lang_EN    Lang = "en"
	Lang_JP    Lang = "jp"
)

func (l Lang) String() string {
	return string(l)
}
