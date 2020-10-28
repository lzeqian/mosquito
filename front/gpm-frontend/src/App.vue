<template>
  <div id="app">
    <Login v-if="!$store.state.isLogin"></Login>
    <Home  v-if="$store.state.isLogin"></Home>
    <Spin fix :style="{zIndex:1000}" v-show="$store.state.isSpinShow">
      <Icon type="ios-loading" size=18 class="demo-spin-icon-load"></Icon>
      <div style="color: red">编译中，请稍后。。。</div>
    </Spin>
  </div>
</template>

<style>
  html,body,#app{
    height: 100%;
  }
  .demo-spin-icon-load{
    animation: ani-demo-spin 1s linear infinite;
  }
  @keyframes ani-demo-spin {
    from { transform: rotate(0deg);}
    50%  { transform: rotate(180deg);}
    to   { transform: rotate(360deg);}
  }
  .demo-spin-col{
    height: 100px;
    position: relative;
    border: 1px solid #eee;
  }
</style>

<script>
  import Home from "./components/Home";
  import Login from "./components/Login";
  export default {
    components:{
      Login,
      Home
    },
    computed:{
      routeQueryContent() {
        return this.$route.query.dirPath+
                this.$route.query.fileName
      }
    },
    data() {
      return {
        loading: true,
        contentTitle:''
      }
    },
    watch:{
      routeQueryContent(newVal, oldVal) {
        this.initData()
      },
      $route: {
        handler:function(val, oldVal){
          let _this=this;
          this.$nextTick(function(){  //页面加载完成后执行
            _this.initData()
          })
        },
        // 深度观察监听
        deep: true
      }
    },
    methods:{
      initData() {
        let vueThis = this;
        if (localStorage.getItem("token")) {
          this.$store.state.isLogin = true
        }
        if (vueThis.$route.name != "blankViewer"){
            if (vueThis.$route && vueThis.$route.matched && vueThis.$route.matched.length > 0) {
              let vueRouteComponents = vueThis.$route.matched[0].instances.default
              if (vueRouteComponents) {
                vueThis.loadEditorContent((vueThis, data) => {
                  if (vueRouteComponents.initData) {
                    vueRouteComponents.initData(data)
                  }
                })
              }
            }
        }
      }

    },
    mounted() {
      window.vueComponents=this;
      this.initData()
    }
  }
</script>
