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

    .rotate-icon{
        transform: rotate(-90deg);
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
                <Split v-model="split1">
                <Sider slot="left" ref="leftSider" style="width: 100%" :style="{background: '#fff',minWidth:'0px',maxWidth:'auto',flex:'0 0 auto'}"
                       :collapsedWidth="0" v-model="isCollapsed" hide-trigger collapsible  :collapsed-width="0">
                    <LeftTree :isCollapsed="isCollapsed"></LeftTree>
                </Sider>
                <Layout slot="right" :style="{padding: '0 0px 0px',zIndex: 10}" style="height: 99%">
<!--                    <Header :style="{padding: 0}" class="layout-header-bar">-->
<!--                        <Icon @click.native="collapsedSider" :class="rotateIcon" :style="{margin: '0 20px'}" type="md-menu" size="24"></Icon>-->
<!--                       <Icon style="position: absolute;top:0;left:10px;z-index: 1000" :style="{margin: '0 2px'}" type="md-menu" size="24"></Icon>-->
<!--                    </Header>-->
<!--                    <Breadcrumb v-if="contentTitle!=''">-->
<!--                        <BreadcrumbItem>Home</BreadcrumbItem>-->
<!--                        <BreadcrumbItem>Components</BreadcrumbItem>-->
<!--                        <BreadcrumbItem>{{contentTitle}}</BreadcrumbItem>-->
<!--                    </Breadcrumb>-->
                    <Content :style="{padding: '5 0 0 0', background: '#fff'}" style="height: 100%">
                        <router-view></router-view>
                    </Content>
                </Layout>
                </Split>
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
                contentTitle: 'hello',
                split1:0.15
            }
        },
        computed: {
            rotateIcon () {
                return [
                    'menu-icon',
                    this.isCollapsed ? 'rotate-icon' : ''
                ];
            },
            menuitemClasses () {
                return [
                    'menu-item',
                    this.isCollapsed ? 'collapsed-menu' : ''
                ]
            }
        },
        methods: {
            exitLogin() {
                if (localStorage.getItem("token")) {
                    localStorage.removeItem("token")
                    this.$store.state.isLogin = false
                    this.$router.push("/")
                }
            },
            collapsedSider () {
                debugger
                this.$refs.leftSider.toggleCollapse();
            }
        },
        components: {
            LeftTree
        },
        mounted() {
        }
    }
</script>