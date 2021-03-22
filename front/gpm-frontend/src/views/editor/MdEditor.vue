<style scoped>
    .vditor-toc:before{
        font-size: 0.15rem;
    }
</style>
<template>

    <div id="vditor" style="height: 100%;padding-left: 0px" @keydown="editorKeyDownSave">

    </div>

</template>
<script>
    import Vditor from 'vditor'

    import "vditor/src/assets/scss/index.scss";
    let vueThis=null;
    // import "~vditor/src/assets/scss/index"
    export default {
        data() {
            return {
                isCollapsed: false,
                data5: [],
                content: "",
                contentEditor:null
            }
        },
        computed:{
        },
        watch: {
        },
        methods: {
            editorKeyDownSave(e) {
                let _this=this;
                let currenKey = e.keyCode || e.which || e.charCode;
                if (currenKey == 83 && e.ctrlKey) {
                    e.preventDefault()
                    _this.saveEditorContent({
                        value: _this.contentEditor.getValue(),
                    })
                }
            },
            initData(data){
                let _this=this;
                if(!this.contentEditor) {
                    this.contentEditor = new Vditor('vditor', {
                        width:'100%',
                        height:'100%',
                        paddingLeft: "0",
                        typewriterMode: true,
                        outline: {
                            enable: true,
                            position: 'left'
                        },
                        toolbarConfig: {
                            pin: false,
                        },
                        cache: {
                            enable: false,
                        },
                        toolbar: [
                            "emoji",
                            "headings",
                            "bold",
                            "italic",
                            "strike",
                            "link",
                            "|",
                            "list",
                            "ordered-list",
                            "check",
                            "outdent",
                            "indent",
                            "|",
                            "quote",
                            "line",
                            "code",
                            "inline-code",
                            "insert-before",
                            "insert-after",
                            "|",
                            "upload",
                            "record",
                            "table",
                            "|",
                            "undo",
                            "redo",
                            "|",
                            "fullscreen",
                            "edit-mode",
                            {
                                hotkey: "Ctrl-S",
                                name: "save",
                                tipPosition: "s",
                                tip: "保存",
                                className: "right",
                                icon: `<img style="height: 16px" src='https://img.58cdn.com.cn/escstatic/docs/imgUpload/idocs/save.svg'/>`,
                                click() {
                                    vueThis.saveEditorContent({
                                        value: vueThis.contentEditor.getValue(),
                                    })
                                }
                            }, {
                                name: "upload",
                                tip: "上传word转换为md"
                            },
                            {
                                name: "more",
                                toolbar: [
                                    "both",
                                    "code-theme",
                                    "content-theme",
                                    "export",
                                    "outline",
                                    "preview",
                                    "devtools",
                                    "info",
                                    "help",
                                ],
                            },
                            ],
                        upload: {
                            accept: 'image/*,application/vnd.openxmlformats-officedocument.wordprocessingml.document',
                            multiple: false,
                            url: this.$globalConfig.goServer + "file/uploadToServer",
                            linkToImgUrl: 'file/uploadToServer',
                            filename (name) {
                                return name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').
                                replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').
                                replace('/\\s/g', '')
                            },
                            handler(files) {
                                let selectedNode=vueThis.$store.getters.getSelectedNode
                                let fileName=files[0].name
                                const param = new FormData();
                                param.append('myfile', files[0])
                                param.append('projectName', selectedNode.fileName)
                                let requestUrl=vueThis.$globalConfig.goServer + "file/uploadToServer";
                                if (fileName.endsWith("doc") || fileName.endsWith("docx")){
                                    requestUrl=vueThis.$globalConfig.goServer + "file/translateToMarkdown";
                                }
                                vueThis.$axios.post(requestUrl, param).then(res => {
                                    if (fileName.endsWith("doc") || fileName.endsWith("docx")){
                                        let mdData = res.data.data;
                                        vueThis.contentEditor.setValue(mdData)
                                    }else {
                                        let imageData = res.data.data;
                                        let name = files[0] && files[0].name;
                                        let succFileText = "";
                                        if (vueThis.contentEditor && vueThis.contentEditor.currentMode === "wysiwyg") {
                                            succFileText += `\n <img alt=${name} src="${imageData}">`;
                                        } else {
                                            succFileText += `  \n![${name}](${imageData})`;
                                        }
                                        document.execCommand("insertHTML", false, succFileText);
                                    }

                                })
                            },
                        },
                        after: () => {
                            this.contentEditor.setValue(data)
                        },
                    })
                }else{
                    this.contentEditor.setValue(data)
                }
            }
        },
        mounted() {
            vueThis=this;
            document.documentElement.style.fontSize="28px";
        }
    }
</script>
