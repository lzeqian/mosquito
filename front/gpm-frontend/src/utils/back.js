function loadTemplateGroup(func){
    let _this = this;
    this.$axios.get(_this.$globalConfig.goServer + "template/groups").then((response) => {
        if (response.data.code == 0) {
            func&&func(response.data.data);
        }
    })
}
function loadTemplate(groupId,func){
    let _this = this;
    this.$axios.get(_this.$globalConfig.goServer + "/template/list?groupId="+groupId).then((response) => {
        if (response.data.code == 0) {
            func && func(response.data.data)
        }
    })
}
function searchVuePress(keyword,func){
    let _this = this;
    this.$axios.get(_this.$globalConfig.goServer + "/md/search?keyword="+keyword).then((response) => {
        if (response.data.code == 0) {
            func && func(response.data.data)
        }
    })
}
function searchShareFile(keyword,func){
    let _this = this;
    this.$axios.get(_this.$globalConfig.goServer + "share/search?keyword="+keyword).then((response) => {
        if (response.data.code == 0) {
            func && func(response.data.data)
        }
    })
}
function cancelShareFile(shareKey,func) {
    let _this = this;
    this.$axios.put(this.$globalConfig.goServer + "share/cancelShareFile?preShareKey=" + shareKey).then((resp) => {
        if (resp.data.code == 0) {
            _this.$Message.info('取消分享成功');
            func && func();
        }
    })
}
function collectFavorite(obj,func){
    let _this = this;
    this.$axios.post(_this.$globalConfig.goServer + "/fav/collectFile",obj).then((response) => {
        if (response.data.code == 0) {
            _this.$Message.info('收藏成功');
            func && func(response.data.data)
        }
    })
}
function searchFavorite(keyword,func){
    let _this = this;
    this.$axios.get(_this.$globalConfig.goServer + "fav/search?keyword="+keyword).then((response) => {
        if (response.data.code == 0) {
            func && func(response.data.data)
        }
    })
}
function cancelFavorite(id,func) {
    let _this = this;
    this.$axios.delete(this.$globalConfig.goServer + "/fav/cancelFavFile?id=" + id).then((resp) => {
        if (resp.data.code == 0) {
            _this.$Message.info('取消分享成功');
            func && func();
        }
    })
}
async function loadSubNode(node,curParent,func){
    let _this = this;
    await _this.$axios.get(_this.$globalConfig.goServer + "home/listSub?fileDir=" + node.dirPath + "&fileName=" + node.title + "&root=" + node.root).then((response) => {
        func && func(node,curParent,response.data.data)
    });
}
function initGoServer(){
    let _this=this;
    _this.$axios.get(_this.$globalConfig.goServer + "share/getShareUrl").then((response) => {
        if (response.data.code == 0) {
            let serverJson=response.data.data;
            _this.$globalConfig.goServer=serverJson["goServer"]
            _this.$globalConfig.externGoServer=serverJson["externGoServer"]
            _this.$globalConfig.documentServer=serverJson["documentServer"]
            _this.$globalConfig.externDocumentServer=serverJson["externDocumentServer"]
        }
    })
}
function getGoServer(shareKey,func){
    let _this=this;
    _this.$axios.get(_this.$globalConfig.goServer + "share/getShareUrl?preShareKey="+shareKey).then((response) => {
        if (response.data.code == 0) {
            let shareServer=response.data.data;
            func && func(shareServer)
        }
    })
}
export default{
    loadTemplateGroup,
    loadTemplate,
    searchVuePress,
    searchShareFile,
    cancelShareFile,
    loadSubNode,
    collectFavorite,
    searchFavorite,
    cancelFavorite,
    initGoServer,
    getGoServer
}
