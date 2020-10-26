package models

type Node struct {
	//必须的大写开头
	Title string `json:"title"`
	Expand  bool `json:"expand"`
	DirPath string `json:"dirPath"`
	IsDir bool `json:"isDir"`
	Contextmenu bool `json:"contextmenu"`
	Children []Node `json:"children"`//key重命名,最外面是反引号
}