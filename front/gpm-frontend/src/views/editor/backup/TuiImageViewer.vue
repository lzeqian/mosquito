<style >
    .imageEditorApp {
        width: 1000px;
        height: 800px;
    }
</style>
<template>

   <div ref="element" id="tui-image-editor" style="height: 100%;overflow: auto;margin-left: 5px;padding-bottom: 100px">
        <tui-image-editor ref="editor"  :include-ui="useDefaultUI" :options="options" style="padding-top: 100px"></tui-image-editor>
    </div>

</template>
<script>
    //https://github.com/nhn/toast-ui.vue-image-editor/blob/master/docs/Basic-Tutorial.md
    import hotkeys from 'hotkeys-js';
    import 'tui-image-editor/dist/svg/icon-a.svg';
    import 'tui-image-editor/dist/svg/icon-b.svg';
    import 'tui-image-editor/dist/svg/icon-c.svg';
    import 'tui-image-editor/dist/svg/icon-d.svg';
    // Load Style Code
    import 'tui-image-editor/dist/tui-image-editor.css';
    import {ImageEditor} from '@toast-ui/vue-image-editor';
    export default {
        components: {
            'tui-image-editor': ImageEditor
        },
        data() {
            return {
                imgSrc: "",
                lc: null,
                useDefaultUI: true,
                options: { // for tui-image-editor component's "options" prop
                    cssMaxWidth: 700,
                    cssMaxHeight: 800,
                    includeUI: {
                        theme: {
                            'common.bi.image': '',
                            'common.bisize.width': '0px',
                            'common.bisize.height': '0px',
                            'common.backgroundImage': 'none',
                            'common.backgroundColor': '#DDDDDD',
                            'common.border': '0px',
                            'header.backgroundImage': 'none',
                            'header.backgroundColor': 'transparent',
                            'header.border': '0px',
                            'loadButton.backgroundColor': '#fff',
                            'loadButton.border': '1px solid #ddd',
                            'loadButton.color': '#222',
                            'loadButton.fontFamily': 'NotoSans, sans-serif',
                            'loadButton.fontSize': '12px',
                            'downloadButton.backgroundColor': '#fdba3b',
                            'downloadButton.border': '1px solid #fdba3b',
                            'downloadButton.color': '#fff',
                            'downloadButton.fontFamily': 'NotoSans, sans-serif',
                            'downloadButton.fontSize': '12px',
                            'menu.normalIcon.color': '#8a8a8a',
                            'menu.activeIcon.color': '#555555',
                            'menu.disabledIcon.color': '#434343',
                            'menu.hoverIcon.color': '#e9e9e9',
                            'submenu.normalIcon.color': '#8a8a8a',
                            'submenu.activeIcon.color': '#e9e9e9',
                            'menu.iconSize.width': '24px',
                            'menu.iconSize.height': '24px',
                            'submenu.iconSize.width': '32px',
                            'submenu.iconSize.height': '32px',
                            'submenu.backgroundColor': '#1e1e1e',
                            'submenu.partition.color': '#858585',
                            'submenu.normalLabel.color': '#858585',
                            'submenu.normalLabel.fontWeight': 'lighter',
                            'submenu.activeLabel.color': '#fff',
                            'submenu.activeLabel.fontWeight': 'lighter',
                            'checkbox.border': '1px solid #ccc',
                            'checkbox.backgroundColor': '#fff',
                            'range.pointer.color': '#fff',
                            'range.bar.color': '#666',
                            'range.subbar.color': '#d1d1d1',
                            'range.disabledPointer.color': '#414141',
                            'range.disabledBar.color': '#282828',
                            'range.disabledSubbar.color': '#414141',
                            'range.value.color': '#fff',
                            'range.value.fontWeight': 'lighter',
                            'range.value.fontSize': '11px',
                            'range.value.border': '1px solid #353535',
                            'range.value.backgroundColor': '#151515',
                            'range.title.color': '#fff',
                            'range.title.fontWeight': 'lighter',
                            'colorpicker.button.border': '1px solid #1e1e1e',
                            'colorpicker.title.color': '#fff'
                        }
                    },

                }
            }
        },
        computed: {
            routeQueryContent() {
                return this.$route.query.dirPath +
                    this.$route.query.fileName
            }
        },
        watch: {
            routeQueryContent(newVal, oldVal) {
                this.initData()
            }
        },
        methods: {
            initData() {
                let _this = this;
                let dirName = this.$route.query.dirPath;
                let fileName = this.$route.query.fileName;
                let token = localStorage.getItem("token")
                var imageUrl=this.$globalConfig.goServer + '/file/download?fileDir=' + dirName + '&fileName=' + fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace;
                _this.$refs.editor.invoke('loadImageFromURL', imageUrl, 'My sample image')
                // let img = new Image()
                // img.src = this.$globalConfig.goServer + '/file/download?fileDir=' + dirName + '&fileName=' + fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace;

            }
        },
        created() {


        },
        mounted() {
            let _this = this;
            _this.initData()
        }
    }
</script>
