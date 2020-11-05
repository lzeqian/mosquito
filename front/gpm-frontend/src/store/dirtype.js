export default {
    state: {
        dirType: 'tree',
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
