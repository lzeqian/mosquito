package models
type  Markdown struct{
	//必须的大写开头
	Value string `json:"value"`
	Html  string `json:"html"`
	DirPath string `json:"dirPath"`
	FileName string `json:"fileName"`
}