import Vue from 'vue'
import Vuex from 'vuex'
import desktop from './desktop.js'
import dtype from "./dirtype.js";
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isSpinShow: false,
    spinShowText:"编译中，请稍后。。。",
    selectedNode:null,
    selectedNodeCacheData:null,
    isLogin:false,
    editorMode:'back',//编辑器模式默认是后台模式:back,还有一种共享模式:share
    shareData:{
    },
    data: {
      scale: 1,
      lineName: 'curve',
      fromArrowType: '',
      toArrowType: 'triangleSolid',
      locked: 0
    },
    event: {
      name: '',
      data: null
    },
    error: {
      text: ''
    },

    count: 2,
    minder: {},
    editor: {},
    working: {
      editing: false,
      saving: false,
      draging: false
    },
    callbackQueue: [],
    config: {
      // 右侧面板最小宽度
      ctrlPanelMin: 250,

      // 右侧面板宽度
      ctrlPanelWidth: parseInt(window.localStorage.__dev_minder_ctrlPanelWidth) || 250,

      // 分割线宽度
      dividerWidth: 3,

      // 默认语言
      defaultLang: 'zh-cn',

      // 放大缩小比例
      zoom: [
        10,
        20,
        30,
        50,
        80,
        100,
        120,
        150,
        200
      ]
    }
  },
  getters:{
    count(state) {
      return state.count;
    },
    working(state) {
      return {
        saving: state.working.saving,
        draging: state.working.draging,
        editing: state.working.editing
      }
    },
    config(state) {
      return {
        ctrlPanelMin: state.config.ctrlPanelMin,
        ctrlPanelWidth: state.config.ctrlPanelWidth,
        dividerWidth: state.config.dividerWidth,
        defaultLang: state.config.defaultLang,
        zoom: state.config.zoom
      }
    },
    getMinder(state){
      return state.minder
    },

    getEditor(state){
      return state.editor
    },
    getSelectedNode(state){
      let seNo=state.selectedNode
      if(seNo==null){
        let lsSelectNode=localStorage.getItem("selectedNode");
        if(lsSelectNode!=null){
          seNo=JSON.parse(lsSelectNode)
        }
      }
      return seNo;
    },
    getSelectedNodeCacheData(state){
      return state.selectedNodeCacheData
    },
    getEditorMode(state){
      return state.editorMode
    },
    getShareData(state){
      return state.shareData;
    },
  },
  mutations: {
    login(state,userInfo){
      state.isLogin=true
      localStorage.setItem("token",userInfo.token)
      localStorage.setItem("userName",userInfo.userName)
    },
    showLoading(state){
      state.isSpinShow = true
    },
    hideLoading (state) {
      state.isSpinShow = false
    },
    setLoadingText(state,title){
      state.spinShowText=title;
    },
    setSelecedNode(state,selectedNode){
      state.selectedNode = selectedNode
      if(selectedNode==null){
        localStorage.setItem("selectedNode",null)
      }else {
        localStorage.setItem("selectedNode", JSON.stringify(selectedNode))
      }
    },
    setSelectedNodeCacheData(state,selectedNodeCacheData){
      state.selectedNodeCacheData = selectedNodeCacheData
    },
    setEditorMode(state,editorMode){
      state.editorMode=editorMode;
    },
    setShareData(state,shareData){
      state.shareData=shareData;
    },
    data(state, data) {
      state.data = data
    },
    emit(state, event) {
      state.event = event
    },
    error(state, error) {
      state.error = error
    },


    changeDrag(state, bool) {
      state.working.draging = bool;
    },

    setMinder(state, data) {
      state.minder = data
    },

    setEditor(state, data) {
      state.editor = data
    },

    changeSave(state, bool) {
      state.working.saving = bool;
    },

    changeCount(state) {
      state.count++;
    },

    increment(state) {
      state.count++
    },

    decrement(state) {
      state.count--
    },

    registerEvent(state, callback) {
      state.callbackQueue.push(callback);
    },

    setConfig(state) {
      var supported = Object.keys(state.config);
      var configObj = {};

      // 支持全配置
      if (typeof key === 'object') {
        configObj = key;
      } else {
        configObj[key] = value;
      }

      for (var i in configObj) {
        if (configObj.hasOwnProperty(i) && supported.indexOf(i) !== -1) {
          state.config[i] = configObj[i];
        } else {
          console.error('Unsupported config key: ', key, ', please choose in :', supported.join(', '));
          return false;
        }
      }
      return true;
    }


  },
  actions: {
    changeCount: ({
                    commit
                  }) => commit('changeCount'),

    increment: ({
                  commit
                }) => commit('increment'),

    decrement: ({
                  commit
                }) => commit('decrement'),

    incrementIfOdd({
                     commit,
                     state
                   }) {
      if ((state.count + 1) % 2 === 0) {
        commit('increment')
      }
    },

    incrementAsync({
                     commit
                   }) {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          commit('increment')
          resolve()
        }, 1000)
      })
    },

    setConfig: ({
                  commit
                }) => commit('setConfig'),

    registerEvent: ({
                      commit
                    }) => commit('registerEvent', callback),

    executeCallback({
                      commit,
                      state
                    }) {
      state.callbackQueue.forEach(function (ele) {
        ele.apply(this, arguments);
      })
    },

    isQuotaExceeded(e) {
      var quotaExceeded = false;
      if (e) {
        if (e.code) {
          switch (e.code) {
            case 22:
              quotaExceeded = true;
              break;
            case 1014:
              // Firefox
              if (e.name === 'NS_ERROR_DOM_QUOTA_REACHED') {
                quotaExceeded = true;
              }
              break;
          }
        } else if (e.number === -2147024882) {
          // Internet Explorer 8
          quotaExceeded = true;
        }
      }
      return quotaExceeded;
      index
    },

    getMemory({
                commit,
                state
              }, key) {
      var value = window.localStorage.getItem(key);
      var result = null || JSON.parse(value)
      console.log('action:' + result);
      return result;
    },

    setMemory({
                commit,
                state
              }, data) {
      try {
        window.localStorage.setItem(data.key, JSON.stringify(data.value));
        return true;
      } catch (e) {
        if (this.isQuotaExceeded(e)) {
          return false;
        }
      }
    },

    removeMemory(key) {
      var value = window.localStorage.getItem(key);
      window.localStorage.removeItem(key);
      return value;
    },

    clearMemory() {
      window.localStorage.clear();
    }
  },
  modules: {
    desktop,
    dtype
  }
})
