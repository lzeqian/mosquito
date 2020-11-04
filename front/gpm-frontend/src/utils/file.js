function downloadFile(selectNode) {
    if (selectNode.isDir) {
        this.$Message.error("不允许直接下载目录，请选择文件");
        return;
    }
    let token=localStorage.getItem("token")
    window.location = this.$globalConfig.goServer + "file/download?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title+(token?"&token="+token:"")
}

function uploadFile(file,func) {
    const param = new FormData();
    param.append('myfile', file)
    param.append('fileDir', this.selectNode.dirPath + "/" + this.selectNode.title)
    this.$axios.post(this.$globalConfig.goServer + "/file/upload", param).then(res => {
        func && func()
    })
}
/**
 * 删除文件
 */
function deleteFile(selectNode,func) {
    this.$axios.delete(this.$globalConfig.goServer + "file/delete?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title).then((response) => {
        func && func()
    })
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
function createVpFile(selectNode,func){
    if (selectNode.isDir) {
        let code = prompt("请输入vuepress名称：");
        if (code != null && code.trim() != "") {
            let fileDir = selectNode.dirPath + "/" + selectNode.title;
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
            var fileDir = selectNode.dirPath + "/" + selectNode.title;
            this.$axios.post(this.$globalConfig.goServer + "file/create",{fileDir:fileDir,fileName:code}).then((response) => {
                func && func(code)
            })
        }
    } else {
        _this.$Message.error("请选选择展开子目录");
    }
}
function editFile(func){
    let selectNode=this.$store.state.selectedNode
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
function loadEditorContent(func) {
    let dirName = this.$route.query.dirPath;
    let fileName = this.$route.query.fileName;
    let vueThis = this;
    this.$axios.get(this.$globalConfig.goServer + "file/query?fileDir=" + dirName + "&fileName=" + fileName).then((response) => {
        vueThis.content = response.data
        func(vueThis, response.data.data)
    })
}
function saveEditorContent (data, func) {
    let vueThis = this;
    vueThis.$axios({
        url: vueThis.$globalConfig.goServer + "file/save",
        method: 'post',
        data: {
            ...data,
            fileDir: vueThis.$route.query.dirPath,
            fileName: vueThis.$route.query.fileName
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
export default{
    downloadFile,
    uploadFile,
    deleteFile,
    buildVpFile,
    createVpFile,
    createTextFile,
    editFile,
    loadEditorContent,
    saveEditorContent
}
