<template lang="html">
    <div class="mind-editor"></div>
</template>

<script>
    import {
        mapActions,
        mapMutations,
        mapGetters
    } from 'vuex'
    import hotkeys from 'hotkeys-js';
    export default {
        mounted() {
            var Editor = require('../../../assets/script/editor');
            var el = this.$el;
            var editor = window.editor = new Editor(el);
            this.setEditor(editor);
            this.initData()
            // if (window.localStorage.mindText) {
            //   editor.minder.importJson(JSON.parse(window.localStorage.mindText));
            // }

            editor.minder.on('contentchange', function () {
                //window.localStorage.mindText = JSON.stringify(editor.minder.exportJson());
            });

            window.minder = window.km = editor.minder;
            this.setMinder(editor.minder);
            this.executeCallback();
        },

        computed: {
            ...mapGetters([
                'minder',
            ]),
        },
        watch: {
            $route: {
                handler: function (val, oldVal) {
                    let _this = this;
                    this.$nextTick(function () {  //页面加载完成后执行
                        _this.initData()
                    })
                },
                // 深度观察监听
                deep: true
            }
        },
        methods: {
            ...mapActions([
                'executeCallback'
            ]),
            ...mapMutations([
                'setMinder',
                'setEditor'
            ]),
            initData() {
                this.loadEditorContent((vueThis, data) => {
                    window.vueThis.loadEditorContent((vueThis, data) => {
                        if (data == null || data == "") {
                            let data = {"root": {"data": {"id": "c6j2vns4ms00", "created": 1603335159200, "text": "标准主题"}}};
                            editor.minder.importJson(data);
                            return;
                        }
                        editor.minder.importJson(JSON.parse(data))
                    })
                })
            },
        },
    }
</script>

<style lang="scss">
    @import "../../../assets/style/editor.scss";
</style>
