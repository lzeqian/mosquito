export default {
    state: {
        //目录类型 tree表示目录树，desktop表示window桌面模式
        dirType: 'tree',
        //0表示公共文档库，1表示个人文档库
        workspace:'0'
    },
    mutations: {
        updateDirTree(state, dirType) {
            state.dirType = dirType
        },
        updateWorkspace(state, workSpace) {
            state.workspace = workSpace
        },

    },
    actions: {

    },
    modules: {}
}
