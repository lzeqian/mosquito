<style scoped>
    .fileSystem {
        padding-top: 31px;
        user-select: none;
    }

    .fileContainer {
        display: flex;
        flex-direction: row;
        user-select: none;
        flex-wrap: wrap;
        justify-content: flex-start;
        align-items: flex-start;
    }

    .fileItem {
        margin-top: 20px;
        text-align: center;
        width: 80px;
    }

    .icon {
        width: 2em;
        height: 2em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }

    .fileText {
        width: 100%;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        user-select: none;
        color: black
    }
    .contextmenu1{
        position: absolute;
        left: 50px;
        top: 50px;
        background-color: rgb(242,242,242);
        user-select:none;
        padding-right: 40px;
        padding-left: 10px;
    }
    /deep/ .ivu-modal-header{
        display: none;
    }
    /deep/ .ivu-modal-content{
        border-radius: 2px;
        height: 100%;
    }
    /deep/ .ivu-modal{
        height: 100%;
        top: 0px;
    }
    /deep/ .ivu-modal-body{
        padding: 0px;
    }
</style>
<template>
    <div class="fileSystem">
        <div style="margin-bottom: 20px">
            当前目录: <span v-html="curDirHtml"></span>
        </div>
        <hr/>
        <div class="fileContainer">
            <div v-for="(item,i) in dataArray" :key="i" class="fileItem" :title="item.title" :ref="'fileItem'+i" @mouseover="mouseOverRef('fileItem'+i,'rgb(229,243,255)')"
                 @mouseleave="mouseLeaveRef('fileItem'+i)">
                <a @dblclick="clickFile(item)" @click="selectClick(item,'fileItem'+i)">
                    <svg class="icon" aria-hidden="true">
                        <use :xlink:href="curFileIcon(item)"></use>
                    </svg>
                    <div class="fileText">
                        {{item.title}}
                    </div>
                </a>
            </div>
        </div>
        <Modal width="80%"
               v-model="showEditorMadal"
               title="编辑器"
               :footer-hide="true"
        >
            <div :style="{height: documentHeight+'px'}">
            <router-view></router-view>
            </div>
        </Modal>
    </div>

</template>
<script>

    window.clickToPath = (dirPath, title, root) => {
        if (root) {
            fileSystemVueThis.curDir = "/"
            fileSystemVueThis.initData();
        } else {
            fileSystemVueThis.$axios.get(fileSystemVueThis.$globalConfig.goServer + "home/listSub?fileDir=" + dirPath + "&fileName=" + title + "&root=" + root).then((response) => {
                fileSystemVueThis.dataArray = response.data.data //挂载子节点
            })
            fileSystemVueThis.curDir = dirPath + (dirPath == "/" ? "" : "/") + title
        }

    }
    export default {
        name: 'FileSystem',
        data() {
            return {
                dataArray: [],
                curDirHtml: '<a href=javascript:window.clickToPath("/","",true,true)>/ </a>',
                curDir: '/',
                selectedItem:null,
                selectedRef:null,
                showEditorMadal:false,
                documentHeight:window.innerHeight,
                currentVisible:false,
            }
        },
        computed:{
            curFileSwitch(){
                return this.$store.state.dtype.workspace
            }
        },
        watch: {
            curFileSwitch(newv, oldv){
                this.initData();
            },
            curDir(newv, oldv) {
                let splitArray = newv.split("/")
                let parentDir = null;
                for (let sa of splitArray) {
                    if (sa == "") {
                        parentDir = "/"
                        this.curDirHtml = '<a href=javascript:window.clickToPath("/","",true,true)>/ </a>';
                    } else {
                        this.curDirHtml = this.curDirHtml + (parentDir == "/" ? "" : "/") + '<a href=javascript:window.clickToPath("' + parentDir + '","' + sa + '",' + false + ',true)>' + sa + '</a>'
                        parentDir = parentDir + (parentDir == "/" ? "" : "/") + sa;
                    }
                }
            }
        },
        methods: {
            mouseOverRef(refName,bgColor){
                // this.$(className).css('background-color', 'rgb(60,95,130)');
                this.$refs[refName][0].style.backgroundColor=bgColor;
            },
            mouseLeaveRef(refName){
                if(this.selectedRef==refName)
                    this.$refs[refName][0].style.backgroundColor='rgb(204,232,255)';
                else
                    this.$refs[refName][0].style.backgroundColor='';
            },
            mouseOver(className,bgColor){
                // this.$(className).css('background-color', 'rgb(60,95,130)');
                this.$(className).css('background-color', bgColor);
            },
            mouseLeave(className){
                this.$(className).css('background-color', '');
            },
            curFileIcon(data) {
                if (data.isDir) {
                    return '#icon-wenjianjia'
                } else {
                    return '#'+this.fileIcon(data.title)
                }
            },
            selectClick(item,refName){
                this.$store.commit("setSelecedNode", item)
                this.selectedItem=item;
                this.selectedRef=refName;
                for(let refTmp in this.$refs){
                    if(this.$refs[refTmp][0]) {
                        this.$refs[refTmp][0].style.border = "";
                        this.$refs[refTmp][0].style.backgroundColor = "";
                    }
                }
                this.$refs[refName][0].style.border="1px double rgb(153,209,255)";
                this.$refs[refName][0].style.backgroundColor="rgb(204,232,255)";
            },
            clickFile(item) {
                this.$store.commit("setSelecedNode", item)
                if(item.isDir) {
                    this.curDir = item.dirPath + (this.curDir == "/" ? "" : "/") + item.title
                    let vueThis = this;
                    vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?fileDir=" + item.dirPath + "&fileName=" + item.title + "&root=" + item.root).then((response) => {
                        vueThis.dataArray = response.data.data //挂载子节点
                    })
                }else{
                    //打开文件
                    let mapping=this.$globalConfig.editorMapping
                    for(let key in mapping){
                        let re;
                        eval("re=/^.+("+key+")$/")
                        if (re.test(item.title) ) {
                            this.showEditorMadal=true
                            this.routePush(item,...mapping[key])
                            return;
                        }
                    }
                }
            },
            initData() {
                let vueThis = this;
                window.fileSystemVueThis = vueThis
                vueThis.curDir='/';
                vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?root=true&fileDir=/").then((response) => {
                    vueThis.dataArray = response.data.data
                })
            },
        },
        components: {},
        mounted() {
            this.initData();
        }
    }
</script>
