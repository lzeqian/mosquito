package routers

import (
	"github.com/astaxie/beego"
	"gpm/controllers"
)

func init() {
	beego.Router("/home/tree", &controllers.TreeController{})
	beego.Router("/home/listSub", &controllers.TreeController{}, "get:ListSubTree")
	beego.Router("/md/createVp", &controllers.MarkDownController{}, "post:CreateVuePress")
	beego.Router("/md/buildVp", &controllers.MarkDownController{}, "post:BuildVuePress")
	beego.Router("/file/download", &controllers.FileController{}, "get:DownloadFile")
	beego.Router("/file/upload", &controllers.FileController{}, "post:UploadFile")
	beego.Router("/file/save", &controllers.FileController{}, "post:SaveFile")
	beego.Router("/file/copy", &controllers.FileController{}, "post:CopyFile")
	beego.Router("/file/delete", &controllers.FileController{}, "delete:DeleteFile")
	beego.Router("/file/query", &controllers.FileController{}, "get:QueryFile")
	beego.Router("/file/create", &controllers.FileController{}, "post:CreateFile")
	beego.Router("/file/mkdir", &controllers.FileController{}, "post:CreateDir")
	beego.Router("/file/rmdir", &controllers.FileController{}, "delete:DeleteDir")
	beego.Router("/file/rename", &controllers.FileController{}, "post:RenameFile")
	beego.Router("/file/uploadToBase64Img", &controllers.FileController{}, "post:UploadToBase64Img")
	beego.Router("/file/transDoc", &controllers.FileController{}, "get:TranslateDoc")
	beego.Router("/file/transPdf", &controllers.FileController{}, "get:TranslatePdf")
	beego.Router("/file/uploadToServer", &controllers.FileController{}, "post:UploadToServer")
	beego.Router("/file/viewerFromServer", &controllers.FileController{}, "get:ViewerFromServer")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
}
