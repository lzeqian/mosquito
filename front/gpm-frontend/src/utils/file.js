function downloadFile(selectNode) {
    if (selectNode.isDir) {
        this.$Message.error("不允许直接下载目录，请选择文件");
        return;
    }
    let token=localStorage.getItem("token")
    window.location = this.$globalConfig.goServer + "file/download?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title+(token?"&token="+token:"")+"&Workspace="+this.$store.getters.currentWorkspace
}

function uploadFile(file,func) {
    const param = new FormData();
    let fileDir=this.selectNode.dirPath + "/" +this.selectNode.title;
    if(this.selectNode.root){
        fileDir=this.selectNode.dirPath
    }
    param.append('myfile', file)
    param.append('fileDir', fileDir)
    this.$axios.post(this.$globalConfig.goServer + "/file/upload", param).then(res => {
        func && func()
    })
}
function uploadFileFun(file,dirPath,fileName,root,func) {
    const param = new FormData();
    let fileDir=dirPath + "/" +fileName;
    if(root){
        fileDir=dirPath
    }
    param.append('myfile', file)
    param.append('fileDir', fileDir)
    this.$axios.post(this.$globalConfig.goServer + "/file/upload", param).then(res => {
        func && func()
    })
}
/**
 * 删除文件
 */
function deleteFile(selectNode,func) {
    if(confirm("确认是否删除")) {
        this.$axios.delete(this.$globalConfig.goServer + "file/delete?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title).then((response) => {
            func && func()
        })
    }
}
/**
 * 删除文件
 */
function deleteDir(selectNode,func) {
    if(confirm("确认是否删除目录")) {
        this.$axios.delete(this.$globalConfig.goServer + "file/rmdir?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.fileName).then((response) => {
            func && func()
        })
    }
}
function buildVpFile(selectNode) {
    let _this=this;
    if (selectNode.isDir) {
        this.$axios.post(this.$globalConfig.goServer + "md/buildVp",{fileDir:selectNode.dirPath,fileName:selectNode.title}).then((response) => {
            _this.$store.commit('hideLoading')
        }).catch(()=>{
            _this.$store.commit('hideLoading')
        });
    }
}
function cancelVpFile(inputSelectNode,func){
    let _this=this;
    let selectNode=inputSelectNode||this.$store.getters.getSelectedNode
    this.$axios.post(this.$globalConfig.goServer + "md/cancelVp",{fileDir:selectNode.dirPath,fileName:selectNode.title}).then((response) => {
        if(response.data.code==0) {
            _this.$Message.info("取消映射成功")
            func && func()
        }
    }).catch(()=>{

    });
}
function createVpFile(selectNode,func){
    if (selectNode.isDir) {
        let code = prompt("请输入vuepress名称：");
        if (code != null && code.trim() != "") {
            let fileDir=selectNode.dirPath + "/" +selectNode.title;
            if(selectNode.root){
                fileDir=selectNode.dirPath
            }
            this.$axios.post(this.$globalConfig.goServer + "md/createVp",{fileDir:fileDir,fileName:code}).then((response) => {
                func && func(code)
            });
        }
    }
}

function createTextFile(selectNode,title,suffix,func) {
    let _this=this;
    if (selectNode.isDir) {
        let code = prompt(title);
        if (code != null && code.trim() != "") {
            let suffixRe=_this.$globalConfig.supportFile
            if(!suffix && !suffixRe.test(code)){
                _this.$Message.error("该文件目不支持创建,只支持:"+suffixRe)
                return;
            }
            if (suffix && !code.endsWith(suffix)) {
                code = code + suffix;
            }
            let fileDir=selectNode.dirPath + "/" +this.selectNode.title;
            if(selectNode.root){
                fileDir=selectNode.dirPath
            }
            this.$axios.post(this.$globalConfig.goServer + "file/create",{fileDir:fileDir,fileName:code}).then((response) => {
                func && func(code)
            })
        }
    } else {
        _this.$Message.error("请选选择展开子目录");
    }
}
function createTextFun(dirPath,suffix,title,func) {
    let _this=this;
    let code = prompt(title);
    if (code != null && code.trim() != "") {
        let suffixRe=_this.$globalConfig.supportFile
        if(!suffix && !suffixRe.test(code)){
            _this.$Message.error("该文件目不支持创建,只支持:"+suffixRe)
            return;
        }
        if (suffix && !code.endsWith(suffix)) {
            code = code + suffix;
        }
        this.$axios.post(this.$globalConfig.goServer + "file/create",{fileDir:dirPath,fileName:code}).then((response) => {
            func && func(code)
        })
    }

}
function createDir(selectNode,title,func) {
    let _this=this;
    if (selectNode.isDir) {
        _this._createDir(selectNode,title,func);
    } else {
        _this.$Message.error("请选择一个目录");
    }
}
function _createDir(selectNode,title,func) {
    let code = prompt(title);
    if (code != null && code.trim() != "") {
        let fileDir=null;
        if(!selectNode){
            fileDir="/"
        }else{
            fileDir=selectNode.dirPath + "/" +this.selectNode.fileName;
        }
        this.$axios.post(this.$globalConfig.goServer + "/file/mkdir",{fileDir:fileDir,fileName:code}).then((response) => {
            func && func(fileDir,code)
        })
    }
}
function editFile(func){
    let selectNode=this.$store.getters.getSelectedNode
    let code = prompt("请输入名称：",selectNode.title);
    let _this=this;
    if (code != null && code.trim() != "") {
        this.$axios.post(this.$globalConfig.goServer + "file/rename",{
            fileDir:selectNode.dirPath,
            fileName:selectNode.title,
            newFileName:code
        }).then((response) => {
            func && func(code)
        });
    }
}
function renameFile(newName,func,errFunc){
    let selectNode=this.$store.getters.getSelectedNode
    let _this=this;
    if (newName != null && newName.trim() != "") {
        this.$axios.post(this.$globalConfig.goServer + "file/rename",{
            fileDir:selectNode.dirPath,
            fileName:selectNode.title,
            newFileName:newName
        }).then((response) => {
            func && func()
        }).catch((reason)=>{
            errFunc && errFunc(reason)
        });
    }
}
function loadEditorContent(func) {
    let vueThis = this;
    if(vueThis.$store.getters.getSelectedNode) {
        let fileDir = vueThis.$store.getters.getSelectedNode.dirPath;
        let fileName = vueThis.$store.getters.getSelectedNode.fileName;
        this.$axios.get(this.$globalConfig.goServer + "file/query?fileDir=" + fileDir + "&fileName=" + fileName).then((response) => {
            vueThis.content = response.data
            func(vueThis, response.data.data)
        })
    }
}
function loadEditorContentByShareKey(key,func) {
    let vueThis = this;
    this.$axios.get(this.$globalConfig.goServer + "share/getShareFile?shareKey=" + key).then((response) => {
        let fileDir = response.data.fileDir;
        let fileName = response.data.fileName;
        func(vueThis, response.data.data)
    })
}
function saveEditorContent (data, func) {
    let vueThis = this;
    let fileDir=vueThis.$store.getters.getSelectedNode.dirPath;
    let fileName=vueThis.$store.getters.getSelectedNode.fileName;
    if(vueThis.$store.getters.getSelectedNode.root){
        fileDir="/"
    }
    vueThis.$axios({
        url: vueThis.$globalConfig.goServer + "file/save",
        method: 'post',
        data: {
            ...data,
            fileDir: fileDir,
            fileName: fileName
        },
        header: {
            'Content-Type': 'application/json'  //如果写成contentType会报错
        }
    }).then((response) => {
        vueThis.$Message.info("保存成功")
        if (func) {
            func(response)
        }
    }).catch((err) => {
        // vueThis.$Message.info("保存失败" + err)
    });
}
function copyFile (func) {
    let _this=this;
    let selectNode=_this.$store.getters.getSelectedNode
    if (!selectNode.isDir) {
        let fileName=selectNode.fileName;
        let fileNamePre=fileName.substring(0,fileName.lastIndexOf("."))
        let fileExt=fileName.substr(fileName.lastIndexOf(".")+1)
        let code = prompt("请输入名称：",fileNamePre+"_bak."+fileExt);
        if (code != null && code.trim() != "") {
            let fileDir = selectNode.dirPath;
            this.$axios.post(this.$globalConfig.goServer + "file/copy", {
                fileDir: fileDir,
                fileName: fileName,
                newFileName:code
            }).then((response) => {
                func && func(code)
            })
        }
    } else {
        _this.$Message.error("请选择文件");
    }
}
function createFileFromTemplateBack (templateObject,func) {
    let _this = this;
    this.$axios.post(_this.$globalConfig.goServer + "/template/gen",templateObject).then((response) => {
        if (response.data.code == 0) {
            func && func()
        }
    })
}
export default{
    downloadFile,
    uploadFile,
    uploadFileFun,
    deleteFile,
    buildVpFile,
    cancelVpFile,
    createVpFile,
    createTextFile,
    createTextFun,
    editFile,
    renameFile,
    loadEditorContent,
    loadEditorContentByShareKey,
    saveEditorContent,
    copyFile,
    createDir,
    _createDir,
    deleteDir,
    createFileFromTemplateBack
}
