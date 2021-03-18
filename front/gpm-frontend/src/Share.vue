<template>
    <div id="app">
        <router-view></router-view>
    </div>
</template>

<style>
    html, body, #app {
        height: 100%;
    }

    .demo-spin-icon-load {
        animation: ani-demo-spin 1s linear infinite;
    }

    @keyframes ani-demo-spin {
        from {
            transform: rotate(0deg);
        }
        50% {
            transform: rotate(180deg);
        }
        to {
            transform: rotate(360deg);
        }
    }

    .demo-spin-col {
        height: 100px;
        position: relative;
        border: 1px solid #eee;
    }
</style>

<script>
    import {routePush} from "./utils/utils";

    export default {
        components: {},
        computed: {},
        data() {
            return {
                loading: true,
                contentTitle: '',
                shareKey: window.shareKey || '5b76930c'
            }
        },
        watch: {},
        methods: {
            initData() {
                let vueThis = this;
                this.$store.commit("setEditorMode", "share")
                if (localStorage.getItem("token")) {
                    this.$store.state.isLogin = true
                }
                vueThis.loadEditorContentByShareKey(vueThis.shareKey, (vueThis, shareData) => {
                    if (shareData.IsPublic != 0) {
                        vueThis.$globalConfig.goServer = vueThis.$globalConfig.externGoServer
                        debugger
                        vueThis.$globalConfig.documentServer = vueThis.$globalConfig.externDocumentServer
                    }
                    if (shareData.ID == 0) {
                        vueThis.$Message.error({
                            content: "分享key不存在，请尝试其他key"
                        });
                        vueThis.routePush({}, '/404', "空白预览")
                        return;
                    }
                    vueThis.$store.commit("setShareData", shareData)
                    let node = {
                        dirPath: shareData["FileDir"],
                        fileName: shareData["FileName"],
                        title: shareData["FileName"],
                    }
                    this.$store.commit("setSelecedNode", node)
                    let mapping = vueThis.$globalConfig.editorMapping
                    for (let key in mapping) {
                        let re;
                        eval("re=/^.+(" + key + ")$/")
                        if (re.test(node.title)) {
                            this.routePush(node, ...mapping[key])
                            if (node.title.endsWith(".pdf")) {
                                return;
                            }
                        }

                    }
                    vueThis.loadEditorContent((vueThis, data) => {
                        vueThis.$store.commit("setSelectedNodeCacheData", data)
                        let initInterval = setInterval(() => {
                            if (vueThis.$route && vueThis.$route.matched && vueThis.$route.matched.length > 0) {
                                let vueRouteComponents = vueThis.$route.matched[0].instances.default
                                if (vueRouteComponents) {
                                    if (vueRouteComponents.initData) {
                                        vueRouteComponents.initData(data)
                                        clearInterval(initInterval)
                                    }
                                }
                            }
                        }, 100)
                    })
                })

            },
            initGoServer() {
                let _this = this;
                _this.$axios.get(_this.$globalConfig.goServer + "share/getShareUrl").then((response) => {
                    if (response.data.code == 0) {
                        let serverJson = response.data.data;
                        _this.$globalConfig.goServer = serverJson["goServer"]
                        _this.$globalConfig.externGoServer = serverJson["externGoServer"]
                        _this.$globalConfig.documentServer=serverJson["documentServer"]
                        _this.$globalConfig.externDocumentServer=serverJson["externDocumentServer"]
                    }
                })
            }
        },
        mounted() {
            window.vueComponents = this;
            this.initData()
            this.initGoServer()
        }
    }
</script>
