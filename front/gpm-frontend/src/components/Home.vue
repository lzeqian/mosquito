<style scoped>

    .layout {
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        position: relative;
        border-radius: 4px;
        overflow: hidden;
        height: 97%;
    }

    .ivu-layout.ivu-layout-has-sider {
        height: 90%;
    }

    .ivu-layout {
        height: 100%;
    }

    .ivu-layout-header {
        height: 10%
    }

    .layout-header-bar {
        background: #fff;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .1);
    }

    .layout-logo-left {
        width: 90%;
        height: 130px;
        background: #5b6270;
        border-radius: 3px;
        margin: 15px auto;
    }

    .ivu-layout-sider {
        height: 99%;
        overflow-y: auto;
        overflow-x: hidden;
    }


</style>
<template>
    <div class="layout">
        <Layout>
            <Header style="display: flex;justify-content:  flex-start ;align-items: center;padding-left: 10px">
                    <span style="width: 80%;height: 100%;display: flex;justify-content:  flex-start ;align-items: center;padding-left: 10px">
                   <img src="img/log.png" style="padding-left: -100px">
                   <img src="img/logfont.png" style="padding-left: -100px">
                        </span>
                <span style="width: 20%;height: 100%;display: flex;justify-content:flex-end;align-items: center ">
                    <div style="display: flex;justify-content:flex-end;align-items: flex-end;height: 100%;padding-bottom: 10px ">
                     <Avatar src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
                         <Button type="text" @click="exitLogin" ghost>退出登录</Button>
                       </div>
                    </span>

            </Header>
            <Layout style="height: 85%">
                <Sider :style="{background: '#fff',maxWidth:'300px',flex:'0 0 300px'}"
                       :collapsedWidth="300">
                    <LeftTree></LeftTree>
                </Sider>
                <Layout :style="{padding: '0 0px 0px',zIndex: 10}" style="height: 99%">
                    <Breadcrumb :style="{margin: '0px 0'}" v-if="contentTitle!=''">
                        <BreadcrumbItem>Home</BreadcrumbItem>
                        <BreadcrumbItem>Components</BreadcrumbItem>
                        <BreadcrumbItem>{{contentTitle}}</BreadcrumbItem>
                    </Breadcrumb>
                    <Content :style="{padding: '5 0 0 0', background: '#fff'}" style="height: 100%">
                        <router-view></router-view>
                    </Content>
                </Layout>
            </Layout>
        </Layout>
    </div>

</template>
<script>
    import LeftTree from "./LeftTree";

    export default {
        data() {
            return {
                isCollapsed: false,
                contentTitle: ''
            }
        },
        methods: {
            exitLogin() {
                if (localStorage.getItem("token")) {
                    localStorage.removeItem("token")
                    this.$store.state.isLogin = false
                }
            }
        },
        components: {
            LeftTree
        },
        mounted() {
        }
    }
</script>