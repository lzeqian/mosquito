<style scoped>

</style>
<template>

    <div ref="element" style="height: 100%">
        <div  id="div1" style="height: 100%"></div>
    </div>

</template>
<script>
    import E from "wangeditor";
    const { $ } = E
    const { BtnMenu, DropListMenu, PanelMenu, DropList, Panel, Tooltip } = E.menuConstructors
    var vueThis=null;
    class AlertMenu extends BtnMenu {
        constructor(editor) {
            const $elem = E.$(
                `<div class="w-e-menu">
                <button>save</button>
            </div>`
            )
            super($elem, editor)
        }
        // 菜单点击事件
        clickHandler() {
            // 做任何你想做的事情
            // 可参考【常用 API】文档，来操作编辑器
            let html=this.editor.txt.html()
            vueThis.$axios({
                url: vueThis.$globalConfig.goServer+"file/save",
                method: 'post',
                data: {
                    value: html,
                    html: '',
                    dirPath: vueThis.$route.query.dirPath,
                    fileName: vueThis.$route.query.fileName
                },
                header: {
                    'Content-Type': 'application/json'  //如果写成contentType会报错
                }
            }).then((response) => {
                vueThis.$Message.info("保存成功")
            });
        }
        // 菜单是否被激活（如果不需要，这个函数可以空着）
        // 1. 激活是什么？光标放在一段加粗、下划线的文本时，菜单栏里的 B 和 U 被激活，如下图
        // 2. 什么时候执行这个函数？每次编辑器区域的选区变化（如鼠标操作、键盘操作等），都会触发各个菜单的 tryChangeActive 函数，重新计算菜单的激活状态
        tryChangeActive() {
            // 激活菜单
            // 1. 菜单 DOM 节点会增加一个 .w-e-active 的 css class
            // 2. this.this.isActive === true
            this.active()

            // // 取消激活菜单
            // // 1. 菜单 DOM 节点会删掉 .w-e-active
            // // 2. this.this.isActive === false
            // this.unActive()
        }
    }
    export default {
        data() {
            return {
                content: "",
                editor:null
            }
        },
        computed:{
            routeQueryContent() {
                return this.$route.query.dirPath+
                    this.$route.query.fileName
            }
        },
        watch:{
            content(n,o){
                this.editor.txt.html(n)
            },
            routeQueryContent(newVal, oldVal) {
                this.initData()
                this.editor.config.uploadImgParams = {
                    projectName:this.$route.query.fileName
                }
            }
        },
        components: {

        },
        methods: {
            initData() {
                this.loadEditorContent((vueThis,data)=>{
                    vueThis.content = data
                })
            },
        },
        created(){


        },
        mounted() {
            vueThis=this;
            setTimeout(function(){
                const editor = new E("#div1");
                vueThis.editor=editor
                vueThis.editor.config.height = vueThis.$refs.element.offsetHeight-50;

                const menuKey = 'alertMenuKey' // 菜单 key ，各个菜单不能重复
                editor.menus.extend('alertMenuKey', AlertMenu)
                editor.config.uploadImgServer =vueThis.$globalConfig.goServer + '/file/uploadToServer'
                editor.config.uploadImgParams = {
                    projectName:vueThis.$route.query.fileName
                }

                editor.config.uploadFileName = 'myfile'
                editor.config.menus = editor.config.menus.concat(menuKey)
                editor.create();
                vueThis.initData()
            },1000)
        }
    }
</script>