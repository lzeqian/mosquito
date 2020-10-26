<style scoped>
    .outer {
        position: relative;
        margin: 0px;
    }
    #map {
        height: 99%;
        width: 100%;
        overflow: auto;
    }
</style>
<template>

    <div ref="element" class="outer" style="height: 100%;overflow: auto">

        <js-mind :values="mind" :options="options" ref="jsMind" height="100%"></js-mind>
    </div>

</template>
<script>
    //https://inspiring-golick-3c01b9.netlify.app/
    import Vue from 'vue'
    import jm from 'vue-jsmind'

    Vue.use(jm)
    import hotkeys from 'hotkeys-js';
    import html2canvas from "html2canvas"
    export default {
        data() {
            return {
                mind:{
                    /* 元数据，定义思维导图的名称、作者、版本等信息 */
                    "meta":{
                        "name":"jsMind-demo-tree",
                        "author":"hizzgdev@163.com",
                        "version":"0.2"
                    },
                    /* 数据格式声明 */
                    "format":"node_tree",
                    /* 数据内容 */
                    "data":{"id":"root","topic":"jsMind"}
                },
                options:{
                    editable:true,                // [可选] 是否启用编辑
                    theme:'orange',
                    draggable:true,
                    shortcut: {
                        enable: true, // 是否启用快捷键
                        handles: {
                            SavePNG(){
                                this.jm.screenshot.shootDownload();
                            }
                        }, // 命名的快捷键事件处理器
                        mapping: {
                            // 快捷键映射
                            addchild: 45, // <Insert>
                            addbrother: 13, // <Enter>
                            editnode: 113, // <F2>
                            delnode: 46, // <Delete>
                            toggle: 32, // <Space>
                            left: 37, // <Left>
                            up: 38, // <Up>
                            right: 39, // <Right>
                            down: 40 ,// <Down>
                            SavePNG:40
                        }
                    }
                }
            }
        },
        computed:{
            routeQueryContent() {
                return this.$route.query.dirPath+
                    this.$route.query.fileName
            }
        },
        watch:{
            routeQueryContent(newVal, oldVal) {
                this.initData()
            }
        },
        methods:{
            initData() {
                this.loadEditorContent((vueThis,data)=>{

                })
            },
        },
        mounted() {
            this.initData()
            if (!window.regCtrlSHotKey) {
                window.vueThis=this
                window.regCtrlSHotKey=true
                hotkeys('ctrl+d', function (event, handler) {



                })
                hotkeys('ctrl+s', function (event, handler) {
                    if (handler.key == "ctrl+s") {
                        let dataString = window.mind.getAllDataString()
                        window.vueThis.$axios({
                            url: window.vueThis.$globalConfig.goServer + "file/save",
                            method: 'post',
                            data: {
                                value: dataString,
                                dirPath: window.vueThis.$route.query.dirPath,
                                fileName: window.vueThis.$route.query.fileName
                            },
                            header: {
                                'Content-Type': 'application/json'  //如果写成contentType会报错
                            }
                        }).then((response) => {
                            window.vueThis.$Message.info("保存成功")
                        });
                    }
                    return false
                });
            }
        }
    }
</script>