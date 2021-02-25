<style scoped>

</style>
<template>

    <div style="height: 100%">
        <div style="height: 100%">
            <div id="wordView" style="height: 100%"/>
        </div>
    </div>

</template>
<script>

    var docEditor=null;
    import mammoth from "mammoth";

    export default {
        data() {
            return {
                vHtml: "",
                wordURL: '',//文件下载地址
                callbackURL:''//office检查状态和保存接口
            }
        },
        computed: {
        },
        watch: {
        },
        methods: {
            async initData() {
                if(docEditor){
                    docEditor.destroyEditor();
                }
                let selectedNode=this.$store.getters.getSelectedNode
                var dirPath = selectedNode.dirPath
                var fileName = selectedNode.fileName
                let vm = this;
                let token = localStorage.getItem("token")
                let fileExtension = fileName.substring(fileName.lastIndexOf('.') + 1);
                let key=vm.randomUuid(20).replaceAll("-","");
                if(this.$store.getters.getEditorMode=="share"){
                    let shareKey=this.$store.getters.getShareData["ShareKey"]
                    key=shareKey;
                    vm.wordURL = window.location.protocol + this.$globalConfig.goServer + "/file/download?fileDir=" + dirPath + "&fileName=" + fileName + (token ? "&token=" + token : "") + "&shareKey=" + shareKey;
                    vm.callbackURL = window.location.protocol+this.$globalConfig.goServer+ "file/uploadOfficeFile?fileDir=" + dirPath+ "&fileName=" + fileName + (token ? "&token=" + token : "") + "&shareKey=" + shareKey;
                }else {
                    vm.wordURL = window.location.protocol + this.$globalConfig.goServer + "/file/download?fileDir=" + dirPath + "&fileName=" + fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace;
                    vm.callbackURL = window.location.protocol+this.$globalConfig.goServer+ "file/uploadOfficeFile?fileDir=" + dirPath+ "&fileName=" + fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace;
                }
                let documentType = "";
                if (/(xls|xlsx)/.test(fileExtension)) {
                    documentType = "spreadsheet"
                }
                if (/(doc|docx)/.test(fileExtension)) {
                    documentType = "word"
                }
                if (/(ppt|pptx)/.test(fileExtension)) {
                    documentType = "presentation"
                }
                //公共文档库必须使用同一个秘钥，可以协作编辑。
                let configJson={
                    "document": {
                        "fileType": fileExtension,
                        "key": key,
                        "title": fileName,
                        "url":  vm.wordURL,
                        "permissions": {
                            "comment": true,
                            "download": true,
                            "edit": true,
                            "fillForms": true,
                            "modifyFilter": true,
                            "modifyContentControl": true,
                            "review": true
                        }
                    },
                    "documentType": documentType,
                    "editorConfig": {
                        "actionLink": null,
                        "mode": "edit", //view表示预览
                        "lang": "zh",
                        "user": {
                            "id": localStorage.getItem("userName"),
                            "name": localStorage.getItem("userName")
                        },
                        "callbackUrl": vm.callbackURL,
                        "customization": {
                            "about": true,
                            "chat": true,
                            "comments": true,
                            "feedback": true,
                            "forcesave": true,
                        },
                        "fileChoiceUrl": "",
                        "plugins": {"pluginsData":[]}
                    },
                    events: {
                        "onAppReady": function(){


                        },
                    }
                }
                docEditor =new DocsAPI.DocEditor("wordView",configJson);

            }

        },
        mounted() {
            var _this=this;
            var head= document.getElementsByTagName('head')[0];
            var script= document.createElement('script');
            script.type= 'text/javascript';
            script.onload= function(){
                _this.initData()
            }
            script.src= _this.$globalConfig.documentServer+"/web-apps/apps/api/documents/api.js";
            head.appendChild(script);
        }
    }
</script>
