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

    <div style="height: 100%">
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
            routeQueryContent() {
                    return  this.$route.query.dirPath+
                    this.$route.query.fileName
            }
        },
        watch: {
            routeQueryContent(newVal, oldVal) {
                    this.initData();
            }
        },
        methods: {
            downloadFile(){
                window.location=this.$globalConfig.goServer+"file/download?fileDir=" + this.$route.query.dirPath + "&fileName=" + this.$route.query.fileName
            },
            transDoc(){
                window.location=this.$globalConfig.goServer+"file/transDoc?fileDir=" + this.$route.query.dirPath + "&fileName=" + this.$route.query.fileName
            },
            handleEditorImgAdd(pos, $file){
                var _this=this;
                const param = new FormData();
                param.append('myfile', $file)
                param.append('projectName', this.$route.query.fileName)
                this.$axios.post(this.$globalConfig.goServer + "/file/uploadToServer", param).then(res => {
                    debugger
                    let imageData=res.data.data;
                    _this.$refs.md.$imglst2Url([[pos, imageData]])
                })
            },
            /**
             * 保存触发事件
             * @param value 原始markdown文件
             * @param render 解析的html文件
             */
            saveCode(value, render) {
                this.$axios({
                    url: this.$globalConfig.goServer+"file/save",
                    method: 'post',
                    data: {
                        value: value,
                        html: render,
                        dirPath: this.$route.query.dirPath,
                        fileName: this.$route.query.fileName
                    },
                    header: {
                        'Content-Type': 'application/json'  //如果写成contentType会报错
                    }
                }).then((response) => {
                    this.$Message.info("保存成功")
                });
            },
            initData(){
                this.loadEditorContent((vueThis,data)=>{
                    vueThis.content = data
                })
            }
        },
        mounted() {
            var md = this.$refs.md;
            var toolbar_left = md.$refs.toolbar_left;
            var diy = this.$refs.diy;
            toolbar_left.$el.append(diy);
            var transDoc = this.$refs.transDoc;
            toolbar_left.$el.append(transDoc);
            this.initData()
        }
    }
</script>