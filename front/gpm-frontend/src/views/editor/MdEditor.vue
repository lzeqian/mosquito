<style scoped>
    .layout {
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        position: relative;
        border-radius: 4px;
        overflow: hidden;
        height: 100%;
    }

    .layout-header-bar {
        background: #fff;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .1);
    }

    .layout-logo-left {
        width: 90%;
        height: 130px;
        background: #5b6270;
        border-radius: 3px;
        margin: 15px auto;
    }

    .menu-icon {
        transition: all .3s;
    }

    .rotate-icon {
        transform: rotate(-90deg);
    }

    .menu-item span {
        display: inline-block;
        overflow: hidden;
        width: 69px;
        text-overflow: ellipsis;
        white-space: nowrap;
        vertical-align: bottom;
        transition: width .2s ease .2s;
    }

    .menu-item i {
        transform: translateX(0px);
        transition: font-size .2s ease, transform .2s ease;
        vertical-align: middle;
        font-size: 16px;
    }

    .collapsed-menu span {
        width: 0px;
        transition: width .2s ease;
    }

    .collapsed-menu i {
        transform: translateX(5px);
        transition: font-size .2s ease .2s, transform .2s ease .2s;
        vertical-align: middle;
        font-size: 22px;
    }

    ivu-layout-sider {
        width: 300px
    }
</style>
<template>

    <div style="height: 100%;padding-left: 7px">
        <mavon-editor v-model="content" ref="md" :style="{height:'100%',maxHeight:'100%'}" @save="saveCode"  @imgAdd="handleEditorImgAdd"/>
        <button ref="diy" type="button" @click="downloadFile"
                class="op-icon fa fa-mavon-floppy-o"
                aria-hidden="true" title="下载"></button>
        <button ref="transDoc" type="button" @click="transDoc"
                class="op-icon fa fa-mavon-floppy-o"
                aria-hidden="true" title="转换doc"></button>
    </div>

</template>
<script>
    export default {
        data() {
            return {
                isCollapsed: false,
                data5: [],
                content: ""
            }
        },
        computed:{
        },
        watch: {
        },
        methods: {
            downloadFile(){
                let selectedNode=this.$store.getters.getSelectedNode
                let token=localStorage.getItem("token")
                if(this.$store.getters.getEditorMode=="share") {
                    let shareKey=this.$store.getters.getShareData["ShareKey"]
                    window.location = this.$globalConfig.goServer + "file/download?fileDir=" + selectedNode.dirPath + "&fileName=" + selectedNode.fileName + (token ? "&token=" + token : "") + "&shareKey=" + shareKey
                }else{
                    window.location = this.$globalConfig.goServer + "file/download?fileDir=" + selectedNode.dirPath + "&fileName=" + selectedNode.fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace
                }
            },
            transDoc(){
                let selectedNode=this.$store.getters.getSelectedNode
                let token=localStorage.getItem("token")
                if(this.$store.getters.getEditorMode=="share") {
                    let shareKey = this.$store.getters.getShareData["ShareKey"]
                    window.location = this.$globalConfig.goServer + "file/transDoc?fileDir=" + selectedNode.dirPath + "&fileName=" + selectedNode.fileName + (token ? "&token=" + token : "") + "&shareKey=" + shareKey
                }else {
                    window.location = this.$globalConfig.goServer + "file/transDoc?fileDir=" + selectedNode.dirPath + "&fileName=" + selectedNode.fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace
                }
            },
            handleEditorImgAdd(pos, $file){
                var _this=this;
                let selectedNode=this.$store.getters.getSelectedNode
                const param = new FormData();
                param.append('myfile', $file)
                param.append('projectName', selectedNode.fileName)
                this.$axios.post(this.$globalConfig.goServer + "file/uploadToServer", param).then(res => {
                    let imageData=res.data.data;
                    _this.$refs.md.$imglst2Url([[pos, imageData]])
                })
            },
            /**
             * 保存触发事件
             * @param value 原始markdown文件
             * @param render 解析的html文件
             */
            saveCode(saveVaue, render) {
                this.saveEditorContent({
                    value: saveVaue,
                })
            },
            initData(data){
                this.content = data
            }
        },
        mounted() {
            var md = this.$refs.md;
            var toolbar_left = md.$refs.toolbar_left;
            var diy = this.$refs.diy;
            toolbar_left.$el.append(diy);
            var transDoc = this.$refs.transDoc;
            toolbar_left.$el.append(transDoc);
        }
    }
</script>
