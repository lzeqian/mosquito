import Vue from 'vue'
import Share from './Share.vue'
import router from './router'
import store from './store'
import ViewUI from 'view-design';
import axios from 'axios'
import 'view-design/dist/styles/iview.css';
import {GlobalConfig} from './utils/env.js'
//markdown在线编辑器
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import 'codemirror/lib/codemirror.css'
import '@/assets/myfonts/iconfont.js'
import '@/assets/myfonts/iconfont.js'
import $ from 'jquery'

Vue.use(mavonEditor)
Vue.use(router);
Vue.config.productionTip = false
Vue.prototype.$axios = axios
Vue.prototype.$ = $
import {fileIcon,routePush,randomUuid} from  './utils/utils'
Vue.prototype.fileIcon=fileIcon
Vue.prototype.routePush=routePush
Vue.prototype.randomUuid=randomUuid
//文件操作相关函数
import fileFunction from  './utils/file'
Object.assign(Vue.prototype,fileFunction)
Vue.prototype.$globalConfig = GlobalConfig
Vue.prototype.$axios.interceptors.request.use(
    config => {
        if (config.method == 'get') {
            config.params = {
                _t: Date.parse(new Date()) / 1000,
                ...config.params
            }
        }
        if(localStorage.getItem("token")) {
            config.headers['Authorization'] = localStorage.getItem("token");
        }
        if(window.vueComponents.$store.getters.getShareData && window.vueComponents.$store.getters.getShareData.ShareKey) {
            config.headers['Share-Key'] = window.vueComponents.$store.getters.getShareData.ShareKey
        }
        return config
    }, function (error) {
        return Promise.reject(error)
    }
)
Vue.prototype.$axios.interceptors.response.use(
    function (response) {
        const status = response.status
        if (status === 200) {
            if ((response.data.code!=undefined && response.data.code != 0) ||
                (response.data.errno!=undefined && response.data.errno != 0)) {
                //token验证失败后需要清空token，让登录失败
                if(response.data.code==3 || response.data.code==4){
                    window.vueComponents.$store.state.isLogin=false;
                    localStorage.removeItem("token")
                }
                Vue.prototype.$Message.error({
                    content: response.data.data,
                    duration: 2,
                    closable: true
                });
                return Promise.reject(response.data.data);
            }
        }
        return response;
    }
)

router.beforeEach((to, from, next) => {
    if (!to.meta.title) {
        document.title = "共享文件"
    }
    next()

})
router.afterEach((to, from) => {
    if (!to.meta.title) {
        document.title = "共享文件"
    }
})
Vue.use(ViewUI);
new Vue({
    router,
    store,
    render: h => h(Share)
}).$mount('#app')
