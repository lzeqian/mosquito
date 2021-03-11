<template>
    <div id="app">
        <Login v-if="!$store.state.isLogin"></Login>
        <Desktop ref="desktop" v-if="$store.state.isLogin && $store.getters.currentDirType=='desktop'"></Desktop>
        <Home ref="home" v-if="$store.state.isLogin && $store.getters.currentDirType=='tree'"></Home>
        <Spin fix :style="{zIndex:10000}" v-show="$store.state.isSpinShow">
            <Icon type="ios-loading" size=18 class="demo-spin-icon-load"></Icon>
            <div style="color: red" v-html="$store.state.spinShowText"></div>
        </Spin>
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
    .icon {
        width: 1em;
        height: 1em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }
</style>

<script>
    import Home from "./components/Home";
    import Login from "./components/Login";
    import Desktop from "./components/Desktop";

    export default {
        components: {
            Login,
            Desktop,
            Home
        },
        computed: {
        },
        data() {
            return {
                loading: true,
                contentTitle: ''
            }
        },
        watch: {
            "$store.getters.getSelectedNode": {
                handler: function (val, oldVal) {
                    if(this.$store.getters.getSelectedNode && !this.$store.getters.getSelectedNode.isDir)
                        this.initData()
                }
            },
            '$route'(to,from) {
                if (this.$route.name == "blankViewer" || this.$route.name == "defaultViewer" ) {
                    this.$store.commit("setSelecedNode", null)
                }
            }
        },
        methods: {
            initData() {
                let vueThis = this;
                if (localStorage.getItem("token")) {
                    this.$store.state.isLogin = true
                }
                if (vueThis.$route.name != "blankViewer" && vueThis.$route.name != "defaultViewer" && vueThis.$route.name!=null) {
                    vueThis.loadEditorContent((vueThis, data) => {
                        vueThis.$store.commit("setSelectedNodeCacheData", data)
                        let initInterval=setInterval(()=>{
                          if (vueThis.$route && vueThis.$route.matched && vueThis.$route.matched.length > 0) {
                            let vueRouteComponents = vueThis.$route.matched[0].instances.default
                            if (vueRouteComponents) {
                              if (vueRouteComponents.initData) {
                                vueRouteComponents.initData(data)
                                clearInterval(initInterval)
                              }
                            }
                          }
                        },100)
                    })
                }
            }

        },
        mounted() {
            window.vueComponents = this;
            this.initData()
        }
    }
</script>
