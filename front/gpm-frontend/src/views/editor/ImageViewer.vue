<style scoped>

</style>
<template>

    <div ref="element" style="height: 100%;overflow: auto">
        <img :src="imgSrc"/>
    </div>

</template>
<script>
    export default {
        data() {
            return {
                imgSrc: ""
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
        components: {

        },
        methods: {
            initData() {
                let dirName=this.$route.query.dirPath;
                let fileName=this.$route.query.fileName;
                this.imgSrc=this.$globalConfig.goServer+'/file/download?fileDir='+dirName+'&fileName='+fileName;
            },
        },
        created(){


        },
        mounted() {
            this.initData()
        }
    }
</script>