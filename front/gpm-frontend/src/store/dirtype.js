export default {
    state: {
        //目录类型 tree表示目录树，desktop表示window桌面模式
        dirType: '',
        //0表示公共文档库，1表示个人文档库
        workspace:''
    },
    mutations: {
        updateDirTree(state, dirType) {
            state.dirType = dirType
            localStorage.setItem("dirType",dirType)
        },
        updateWorkspace(state, workSpace) {
            state.workspace = workSpace
            localStorage.setItem("workspace",workSpace)
        },

    },
    getters:{
        currentDirType(state){
            let dirType=state.dirType
            dirType=state.dirType||localStorage.getItem("dirType")||"tree"
            state.dirType=dirType
            return state.dirType
        },
        currentWorkspace(state){
            let workspace=state.workspace
            workspace=state.workspace||localStorage.getItem("workspace")||"0"
            state.workspace=workspace
            return state.workspace
        }
    },
    actions: {

    },
    modules: {}
}
