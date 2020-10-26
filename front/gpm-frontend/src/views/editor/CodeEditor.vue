<style>

    .CodeMirror {
        border: 1px solid #eee;
        height: 99%;
        min-height: 99%;
    }

    .CodeMirror-scroll {
        height: 100%;
        min-height: 100%;
        overflow-y: hidden;
        overflow-x: auto;
    }
    .vue-codemirror{
        height: 100%;
    }
</style>
<template>

    <div style="height: 100%">
        <codemirror
                ref="chartOption"
                v-model="mirrorCode"
                :value="mirrorCode"
                :options="cmOptions"

                @keyHandled="keyHandled"
        >
        </codemirror>
    </div>

</template>
<script>

    import { codemirror } from 'vue-codemirror'
    import 'codemirror/lib/codemirror.css'
    require("codemirror/mode/python/python.js")
    require("codemirror/mode/javascript/javascript.js")
    require("codemirror/mode/yaml/yaml.js")
    require("codemirror/mode/xml/xml.js")
    require('codemirror/addon/fold/foldcode.js')
    require('codemirror/addon/fold/foldgutter.js')
    require('codemirror/addon/fold/brace-fold.js')
    require('codemirror/addon/fold/xml-fold.js')
    require('codemirror/addon/fold/indent-fold.js')
    require('codemirror/addon/fold/markdown-fold.js')
    require('codemirror/addon/fold/comment-fold.js')
    import 'codemirror/theme/idea.css';
    import "codemirror/theme/ambiance.css"
    import "codemirror/lib/codemirror.css"
    import "codemirror/addon/hint/show-hint.css"
    export default {
        data () {
            return {
                mirrorCode:'',
                cmOptions: {
                    tabSize: 2, // Tab键空格数
                    mode: "", //模式
                    theme: 'idea', // 主题
                    lineNumbers: true, //是否显示行号
                    line: true,
                    extraKeys:{
                        "Ctrl-S":function () {

                        }
                    }
                },
            }
        },
        components:{
            codemirror
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
        methods: {
            keyHandled(){
                if(event.code=="KeyS"){
                    var vueThis = this;
                    this.$axios({
                        url: this.$globalConfig.goServer+"file/save",
                        method: 'post',
                        data: {
                            value: vueThis.mirrorCode,
                            dirPath: this.$route.query.dirPath,
                            fileName: this.$route.query.fileName
                        },
                        header: {
                            'Content-Type': 'application/json'  //如果写成contentType会报错
                        }
                    }).then((response) => {
                        this.$Message.info("保存成功")
                    });
                }
            },
            initData() {
                this.loadEditorContent((vueThis,data)=>{
                    vueThis.mirrorCode = data
                })
            },
        },
        created(){
            let fileExtArray=this.$route.query.fileName.split(".");
            let fileExt=fileExtArray[fileExtArray.length-1]
            if(fileExt=="js"||fileExt=="vue"){
                this.cmOptions.mode="javascript";
            }
            if(fileExt=="yaml" || fileExt=="yml" ){
                this.cmOptions.mode="yaml";
            }
            if(fileExt=="yaml" || fileExt=="yml" ){
                this.cmOptions.mode="yaml";
            }
            if(fileExt=="java"){
                this.cmOptions.mode="text/x-java";
            }
            if(fileExt=="xml"){
                this.cmOptions.mode="xml";
            }
            if(fileExt=="py"){
                this.cmOptions.mode="python";
            }
        },
        mounted() {
            this.initData()
        }
    }
</script>
