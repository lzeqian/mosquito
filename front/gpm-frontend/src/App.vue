<template>
  <div id="app">
    <Home></Home>
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
  export default {
    components:{
       Home
    },
    data() {
      return {
        loading: true,
        contentTitle:''
      }
    },
    mounted() {

    }
  }
</script>
