package models

type Markdown struct {
	//必须的大写开头
	Value       string `json:"value"`
	Html        string `json:"html"`
	FileDir     string `json:"fileDir"`
	FileName    string `json:"fileName"`
	NewFileName string `json:"newFileName"`
}
