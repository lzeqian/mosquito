export function loadEditorContent(func) {
    let dirName = this.$route.query.dirPath;
    let fileName = this.$route.query.fileName;
    let vueThis = this;
    this.$axios.get(this.$globalConfig.goServer + "file/query?fileDir=" + dirName + "&fileName=" + fileName).then((response) => {
        vueThis.content = response.data
        func(vueThis, response.data.data)
    })
}
export function saveEditorContent (data, func) {
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

export function fileIcon(title){
        if (title.endsWith(".doc") || title.endsWith(".docx")) {
            return "icon-doc"
        }
        if (title.endsWith(".xls") || title.endsWith(".xlsx")) {
            return "icon-xls"
        }
        if (title.endsWith(".ppt") || title.endsWith(".pptx")) {
            return "icon-ppt"
        }
        if (title.endsWith(".json")) {
            return "icon-json"
        }
        if (title.endsWith(".js")) {
            return "icon-js-square"
        }
        if (title.endsWith(".pdf")) {
            return "icon-pdf"
        }
        if (/.*\.(png|PNG|jpg|JPG|JPEG|jpeg|gif|GIF)/.test(title)) {
            return "icon-picture"
        }
        if (/.*\.(zip|7z|rar)/.test(title)) {
            return "icon-zip"
        }
        if (title.endsWith(".md")) {
            return "icon-file-markdown"
        }
        if (title.endsWith(".html")) {
            return "icon-HTML"
        }
        if (title.endsWith(".xml")) {
            return "icon-xml"
        }
        return 'icon-wenjian'
}

export function routePush(node,routerAddress,title){
    this.$router.push({
        path: routerAddress,
        query: {dirPath: node.dirPath, fileName: node.title}
    });
}
