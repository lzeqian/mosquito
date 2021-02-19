<style scoped>

</style>
<template>

    <div style="height: 100%">
        <div>
            <div id="wordView" v-html="vHtml" />
        </div>
    </div>

</template>
<script>
    import mammoth from "mammoth";
    export default {
        data() {
            return {
                vHtml: "",
                wordURL:''//文件地址
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
                let vm=this;
                let token=localStorage.getItem("token")
                vm.wordURL =  this.$globalConfig.goServer+"/file/download?fileDir="+dirPath+"&fileName="+fileName+(token?"&token="+token:"")+"&Workspace="+this.$store.getters.currentWorkspace;
                const xhr = new XMLHttpRequest();
                xhr.open("get", this.wordURL, true);
                xhr.responseType = "arraybuffer";
                xhr.onload = function () {
                    debugger
                    if (xhr.status == 200) {
                        mammoth
                            .convertToHtml({ arrayBuffer: new Uint8Array(xhr.response) })
                            .then(function (resultObject) {
                                vm.$nextTick(() => {
                                    // document.querySelector("#wordView").innerHTML =
                                    //   resultObject.value;
                                    vm.vHtml=resultObject.value;
                                });
                            });
                    }
                };
                xhr.send();
            }

        },

        mounted() {
            this.initData()
        }
    }
</script>
