<style scoped>
    .icon {
        width: 1em;
        height: 1em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }
</style>
<template>

    <div ref="element" style="height: 100%;overflow: auto">
        <Row style="margin-top: 50px;margin-left: 50px;margin-right: 50px">
            <Col span="7" style="margin-right:10px">
                <Card>
                    <p slot="title">
                        <Icon type="ios-settings"></Icon>
                        收藏列表
                    </p>
                    <a href="#" slot="extra" >
                        <input type="text"  v-model="favKeyword" style="width: 80px" v-show="showFavSearch" @keyup.enter="searchFavoriteByKeyword"/>
                        <Button icon="ios-search" size="small" type="text" @click="showFavSearch=!showFavSearch;"></Button>
                        <Button icon="ios-refresh-circle" size="small" type="text" @click="searchFavoriteInner('')"></Button>
                    </a>
                    <ul>
                        <li v-for="item in favList" :key="item.name">
                            <a @click="gotoFileLocation(item)" target="_blank">{{ item.FileName }}</a>
                            <Button icon="logo-linkedin" size="small" type="text" style="position:absolute;right:20px" @click="gotoFileLocation(item)"></Button>
                            <Button icon="md-remove-circle" size="small" type="text" style="position:absolute;right:40px;color:red" @click="cancelFav(item.ID)"></Button>

                        </li>
                    </ul>
                </Card>

            </Col>
            <Col span="7" style="margin-right:10px">
                <Card>
                    <p slot="title">
                        <Icon type="ios-share"></Icon>
                        分享列表
                    </p>
                    <a href="#" slot="extra" >
                        <input type="text"  v-model="shareKeyword" style="width: 80px" v-show="showShareSearch" @keyup.enter="searchShareFileByKeyword"/>
                        <Button icon="ios-search" size="small" type="text" @click="showShareSearch=!showShareSearch;"></Button>
                        <Button icon="ios-refresh-circle" size="small" type="text" @click="searchShareFileInner('')"></Button>
                    </a>
                    <ul>
                        <li v-for="item in shareList" :key="item.name">
                            <a @click="gotoUrl(item.ShareKey)" target="_blank">{{ item.FileName }}</a>
                            <Button icon="logo-linkedin" size="small" type="text" style="position:absolute;right:20px" @click="gotoFileLocation(item)"></Button>
                            <Button icon="md-remove-circle" size="small" type="text" style="position:absolute;right:40px;color:red" @click="cancelShare(item.ID)"></Button>

                        </li>
                    </ul>
                </Card>
            </Col>
            <Col span="7">
                <Card>
                    <p slot="title">
                        <Icon type="ios-settings"></Icon>
                        vuepress列表
                    </p>
                    <a href="#" slot="extra" >
                        <input type="text"  v-model="vpKeyword" style="width: 80px" v-show="showVpSearch" @keyup.enter="searchVuePressByKeyword"/>
                        <Button icon="ios-search" size="small" type="text" @click="showVpSearch=!showVpSearch;"></Button>
                        <Button icon="ios-refresh-circle" size="small" type="text" @click="searchVuePressInner('')"></Button>
                    </a>
                    <ul>
                        <li v-for="item in vpList" :key="item.name">
                            <svg class="icon" aria-hidden="true" style="color: green">
                                <use :xlink:href="item.Workspace == '0'?'#icon-duoren-renqun':'#icon-renshu'"></use>
                            </svg> <a @click="gotoVpUrl(item.AppPath)" target="_blank">{{ item.FileName }}</a>
                            <Button icon="logo-linkedin" size="small" type="text" style="position:absolute;right:20px" @click="gotoFileLocation(item)"></Button>
                            <Button icon="md-remove-circle" size="small" type="text" style="position:absolute;right:40px;color:red" @click="cancelVp(item)"></Button>

                        </li>
                    </ul>
                </Card>
            </Col>
        </Row>


    </div>

</template>
<script>
    export default {
        data() {
            return {
                shareKeyword:"",//共享文档
                showShareSearch:false,
                favKeyword:"",//收藏文档
                showFavSearch:false,
                vpKeyword:"", //vuepress列表
                showVpSearch:false,
                shareList: [
                    ],
                favList: [
                    ],
                vpList: [
                    ],
            }
        },
        watch:{

        },
        methods:{
            cancelShare(id){
                let _this=this;
                let shareData=_this.shareList.filter((item)=>{
                    return item.ID==id
                })
                if(shareData.length>0) {
                    _this.cancelShareFile(shareData[0].ShareKey, () => {
                        _this.searchShareFileInner()
                    })
                }
            },
            searchVuePressInner(keyword){
                let _this=this;
                _this.searchVuePress(keyword||"", (data) => {
                    _this.vpList = data
                })
            },
            searchVuePressByKeyword(){
                let _this=this;
                _this.searchVuePressInner(_this.vpKeyword)
            },
            searchShareFileInner(keyword){
                let _this=this;
                _this.searchShareFile(keyword||"",(data)=>{
                    _this.shareList=data
                })
            },
            searchShareFileByKeyword(){
                let _this=this;
                _this.searchShareFileInner(_this.shareKeyword)
            },
            searchFavoriteInner(keyword){
                let _this=this;
                _this.searchFavorite(keyword||"",(data)=>{
                    _this.favList=data
                })
            },
            searchFavoriteByKeyword(){
                let _this=this;
                _this.searchFavoriteInner(_this.favKeyword)
            },
            cancelVp(item){
                let _this=this;
                this.gotoFileLocation(item);
                this.cancelVpFile(this.$store.getters.getSelectedNode,()=>{
                    _this.searchVuePressInner();
                })

            },
            cancelFav(id){
                let _this=this;
                this.cancelFavorite(id,()=>{
                    _this.searchFavoriteInner();
                })
            },
            gotoUrl(shareKey) {
                this.getGoServer(shareKey,(shareServer)=>{
                    let jurl=window.location.protocol + shareServer + "docs/" + shareKey;
                    window.open(jurl)
                })

            },
            gotoVpUrl(appPath) {
                let jurl=window.location.protocol + this.$globalConfig.goServer  + appPath.substring(1);
                window.open(jurl)
            },
            gotoFileLocation(item){
                let _this=this;
                let workspace=(item.Workspace!=undefined?item.Workspace:1);
                let dirPath = item.FileDir
                let fileName = item.FileName
                if(workspace==1) {
                    dirPath = dirPath.substring(1)
                    dirPath = dirPath.substring(dirPath.indexOf("/"))
                }
                let selectNode={
                    dirPath:dirPath,
                    fileName:fileName,
                    title:fileName
                }
                let leftTree=this.$root.$children[0].$refs.home.$refs.leftTree;
                if(workspace!= this.$store.getters.currentWorkspace) {
                    this.$store.commit("updateWorkspace",workspace)
                    let title = this.$store.getters.currentWorkspace == "0" ? "公共文档库" : "个人文档库";
                    _this.$store.commit("setSelecedNode", selectNode)
                    leftTree.listRoot(title);

                }else{
                    _this.$store.commit("setSelecedNode", selectNode)
                    leftTree.initData()
                }


            }
        },
        mounted() {
            let _this=this;
            _this.searchVuePressInner()
            _this.searchShareFileInner();
            _this.searchFavoriteInner();
        }
    }
</script>
