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
            <div id="map"></div>
    </div>

</template>
<script>
    //https://inspiring-golick-3c01b9.netlify.app/
    import MindElixir, { E } from 'mind-elixir'
    import hotkeys from 'hotkeys-js';
    import html2canvas from "html2canvas"
    export default {
        data() {
            return {
                mind:null
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
                    let topicData=MindElixir.new('new topic')
                    if(data!=null && data!=""){
                        topicData=JSON.parse(data)
                    }
                    vueThis.mirrorCode = data
                    let mind = new MindElixir({
                        el: '#map',
                        direction: MindElixir.LEFT,
                        //data: MindElixir.new('new topic'), // 也可以把 getDataAll 得到的数据初始化
                        data: topicData, // 也可以把 getDataAll 得到的数据初始化
                        draggable: true, // 启用拖动 default true
                        contextMenu: true, // 启用右键菜单 default true
                        toolBar: true, // 启用工具栏 default true
                        nodeMenu: true, // 启用节点菜单 default true
                        keypress: true, // 启用快捷键 default true
                    })
                    mind.init()
                    window.mind=mind
                })
            },
        },
        mounted() {
            this.initData()
            if (!window.regCtrlSHotKey) {
                window.vueThis=this
                window.regCtrlSHotKey=true
                hotkeys('ctrl+d', function (event, handler) {
                    // var svgXml=window.mind.svg2nd.innerHTML
                    // var image = new Image();
                    // image.src = 'data:image/svg+xml;base64,' + window.btoa(unescape(encodeURIComponent(svgXml))); //给图片对象写入base64编码的svg流
                    // var canvas = document.createElement('canvas');  //准备空画布
                    // canvas.width =1800;
                    // canvas.height = 800;
                    // var context = canvas.getContext('2d');  //取得画布的2d绘图上下文
                    // context.drawImage(image, 0, 0);
                    html2canvas(document.querySelector("#map")).then(canvas => {
                        var a = document.createElement('a');
                        a.href = canvas.toDataURL('image/png');  //将画布内的信息导出为png图片数据
                        a.download = "a.png";  //设定下载名称
                        a.click();
                    });


                })
                hotkeys('ctrl+s', function (event, handler) {
                    if (handler.key == "ctrl+s") {
                        event.preventDefault()
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