<template>
    <div style="height: 99%;" @keydown="editorKeyDownSave">
        <header-menu></header-menu>

        <main-editor ref="mainEditor"></main-editor>
        <navigator></navigator>
    </div>

</template>

<script>
    import Vue from "vue"
    import locale from 'element-ui/lib/locale/lang/zh-CN'
    import Element from 'element-ui'
    import 'element-ui/lib/theme-chalk/index.css'
    Vue.use(Element, {locale})
    import $ from 'jquery'
    window.$ = $
    window.jQuery = $
    require('../../../node_modules/codemirror/lib/codemirror.js')
    require('../../../node_modules/codemirror/lib/codemirror.js')
    require('../../../node_modules/codemirror/mode/xml/xml.js')
    require('../../../node_modules/codemirror/mode/javascript/javascript.js')
    require('../../../node_modules/codemirror/mode/css/css.js')
    require('../../../node_modules/codemirror/mode/htmlmixed/htmlmixed.js')
    require('../../../node_modules/codemirror/mode/markdown/markdown.js')
    require('../../../node_modules/codemirror/addon/mode/overlay.js')
    require('../../../node_modules/codemirror/mode/gfm/gfm.js')
    require('../../../node_modules/marked/lib/marked.js')
    require('../../../node_modules/kity/dist/kity.js')
    require('../../../node_modules/hotbox/hotbox.js')
    require('../../../node_modules/kityminder-core/dist/kityminder.core.js')
    require('../../assets/script/expose-editor.js')
    import headerMenu from '@/components/mind/header'
    import mainEditor from '@/components/mind/main/mainEditor'
    import navigator from '@/components/mind/main/navigator'

    export default {
        name: 'editor',
        components: {
            headerMenu,
            mainEditor,
            navigator
        },
        created(){
            window.vueThis=this
        },
        methods:{
            initData(data){
                this.$refs.mainEditor.childInitData(data)
            },
            editorKeyDownSave(e) {
                let _this=this;
                let currenKey = e.keyCode || e.which || e.charCode;
                if (currenKey == 83 && e.ctrlKey) {
                    e.preventDefault()
                    let saveData=JSON.stringify(_this.$store.getters.getMinder.exportJson())
                    _this.saveEditorContent({
                        value: saveData,
                    })
                }
            }
        },
        mounted() {

        }
    }

</script>
