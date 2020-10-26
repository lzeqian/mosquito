<style scoped>

</style>
<template>

    <div style="height: 600px" >
        <pdf
                v-for="i in numPages"
                :key="i"
                :page="i"
                :src="pdfUrl" style="width: 100%; height: auto;" @num-pages="pageCount=$event">
        </pdf>
    </div>

</template>
<script>
    import pdf from 'vue-pdf'
    export default {
        data() {
            return {
                pageCount:0,
                pdfUrl:'',
                src: '', // pdf文件地址
                numPages:0,

            }
        },
        computed: {
            routeQueryContent() {
                return (this.$route.query.dirPath+
                    this.$route.query.fileName)
            }
        },
        watch: {
            routeQueryContent(newVal, oldVal) {
                this.initData()
            }
        },
        methods: {
            async initData() {
                var dirPath = this.$route.query.dirPath
                var fileName = this.$route.query.fileName
                this.src=this.$globalConfig.goServer+"/file/download?fileDir="+dirPath+"&fileName="+fileName
                this.pdfUrl = pdf.createLoadingTask(this.src);
                this.pdfUrl.promise.then(pdf => {
                    this.numPages = pdf.numPages
                })
            }

        },
        components: {pdf},
        mounted() {
            this.initData()
        }
    }
</script>