export function loadEditorContent(func) {
    let dirName = this.$route.query.dirPath;
    let fileName = this.$route.query.fileName;
    let vueThis = this;
    this.$axios.get(this.$globalConfig.goServer + "file/query?fileDir=" + dirName + "&fileName=" + fileName).then((response) => {
        vueThis.content = response.data.data
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
