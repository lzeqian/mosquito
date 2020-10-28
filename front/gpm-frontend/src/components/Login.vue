<style scoped>
    .wrapper{
        width: 100%;
        height: 100%;
    }
    .content{
        width: 100%;
        height: 100%;
    }
    .bg1{
        width: 100%;
        height: 100%;
    }

    .loginBox{
        position: absolute;
        right: 200px;
        top: 200px;
        width: 300px;
        background-color: white;
        border-radius: 8px;
        padding-left: 20px;
        padding-right: 20px;
    }

</style>
<template>

    <div class="content" >
        <div class="wrapper">
            <img class="bg1" src="img/bg1.jpg"/>
            <div class="loginBox">
                <div class="logo">
                    <img src="img/log.png"/>
                </div>
                <div class="formDiv">
                    <Form ref="loginForm" :model="form" :rules="rules" @keydown.enter.native="loginForm">
                        <FormItem prop="userName">
                            <Input v-model="form.userName" placeholder="请输入用户名">
                 <span slot="prepend">
           <Icon :size="16" type="ios-person"></Icon>
          </span></Input>
                        </FormItem>
                        <FormItem prop="password">
                            <Input type="password" v-model="form.password" placeholder="请输入密码">
                 <span slot="prepend">
           <Icon :size="14" type="md-lock"></Icon>
        </span>
                            </Input>
                        </FormItem>
                        <FormItem label="记住密码">
                            <i-switch v-model="form.remember" size="large">
                                <span slot="open">开</span>
                                <span slot="close">关</span>
                            </i-switch>
                        </FormItem>
                        <FormItem>
                            <Button @click="loginForm" type="primary" long>登录</Button>
                        </FormItem>
                    </Form>

                </div>
            </div>
        </div>
    </div>


</template>
<script>
    import LeftTree from "./LeftTree";

    export default {
        data() {
            return {
                form: {
                    userName: null,
                    password: null,
                    remember: null
                }
            }
        },
        props: {
             userNameRules: {
                   type: Array,
                       default: () => {
                         return [{ required: true, message: '账号不能为空', trigger: 'blur' }]
                       }
                 },
             passwordRules: {
                   type: Array,
                       default: () => {
                         return [{ required: true, message: '密码不能为空', trigger: 'blur' }]
                       }
                 },
        },
        computed:{
            rules () {
               return {
                  userName: this.userNameRules,
                  password: this.passwordRules
               }
            }
        },
        components: {},
        methods: {
            loginForm() {
                this.$refs.loginForm.validate(valid => {
                    if (valid) {
                        this.$axios.post(this.$globalConfig.goServer + "login", {
                            ...this.form
                        }).then((resp) => {
                                this.$store.commit("login",resp.data.data)
                        });
                    }
                });
            }

        },
        mounted() {
        }
    }
</script>