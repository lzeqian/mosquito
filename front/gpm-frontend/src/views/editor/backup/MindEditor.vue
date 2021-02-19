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

    </div>

</template>
<script>
    //https://inspiring-golick-3c01b9.netlify.app/
    import Vue from 'vue'
    import 'element-ui/lib/theme-chalk/index.css'
    require('@/assets/script/expose-editor.js')

    import kity from 'kity'
    import kityminder from 'kityminder-core'
    export default {
        data() {
            return {

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
        }
    }
</script>