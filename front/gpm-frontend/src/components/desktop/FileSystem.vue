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
        /*text-overflow: clip;*/
        /*overflow:visible;*/
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
    #showEditorMadal /deep/ .ivu-modal-header{
        /*注意选择当前元素必须在最前面，过滤子组件元素并生效使用/deep/ */
        display: none;
    }
    #showEditorMadal /deep/ .ivu-modal-content{
        border-radius: 2px;
        height: 100%;
    }
    #showEditorMadal /deep/  .ivu-modal{
        height: 100%;
        top: 0px;
    }
    #showEditorMadal /deep/ .ivu-modal-body{
        padding: 0px;
    }
</style>
<template>
    <div class="fileSystem">
        <div style="margin-bottom: 20px">
            当前目录: <span v-html="curDirHtml"></span>
        </div>
        <hr/>
        <div class="fileContainer" v-if="selectedItem && selectedItem.children" >
            <div v-for="item in selectedItem.children" :key="item.title" class="fileItem" :title="item.title" :ref="'fileItem'+item.title" @mouseover="mouseOverRef('fileItem'+item.title,'rgb(229,243,255)')"
                 @mouseleave="mouseLeaveRef('fileItem'+item.title)">
                <a @dblclick="clickFile(item)" @click="selectClick(item,'fileItem'+item.title)">
                    <svg class="icon" aria-hidden="true">
                        <use :xlink:href="curFileIcon(item)"></use>
                    </svg>
                    <div class="fileText" :oriText="item.title" :contenteditable="selectNode!=null && !selectNode.isDir" @focus="focurCurFile"  @blur="renameCurrFile">
                        {{item.title}}
                    </div>
                </a>
            </div>
        </div>
        <Modal id="showEditorMadal" width="80%"
               v-model="showEditorMadal"
               title="编辑器"
               :footer-hide="true"
        >
            <div :style="{height: documentHeight+'px'}">
            <router-view></router-view>
            </div>
        </Modal>
        <Modal
                v-model="showTemplate"
                style="height:200px"
                title="模板选择"
                @on-ok="createFileFromTemplate"
                :z-index="10002">
            <Form :model="templateObject" :label-width="80">
                <FormItem label="模板组">
                    <Select v-model="templateObject.templateGroupId" style="width:200px" @on-change="changeTemplateGroup">
                        <Option v-for="groupItem in templateGroupData" :value="groupItem.value" :key="groupItem.value">{{
                            groupItem.label}}
                        </Option>
                    </Select>
                </FormItem>
                <FormItem label="模板">
                    <Select v-model="templateObject.templateId" style="width:200px" @on-change="changeTemplateGroup">
                        <Option v-for="item in templateData" :value="item.value" :key="item.value">{{
                            item.label}}
                        </Option>
                    </Select>
                </FormItem>
                <FormItem label="文件名称">
                    <Input v-model="templateObject.fileName"  style="width: 300px" />
                </FormItem>
            </Form>
        </Modal>
    </div>

</template>
<script>

    export default {
        name: 'FileSystem',
        data() {
            return {
                dataArray: [],
                curDirHtml: '<a href=javascript:window.fileSystemVueThis.clickToPath("/","",true,true)>/ </a>',
                curDir: '/',
                selectedItem:null, //当前选中的目录，selectNode表示当前选择的节点可以是目录也可以是文件
                selectedRef:null,
                showEditorMadal:false,
                documentHeight:window.innerHeight,
                currentVisible:false,
                showShare: false,
                showTemplate: false,
                templateGroupData: [],
                templateData: [],
                templateObject: {
                    templateGroupId: "",
                    templateId: "",
                    fileName: "",
                    fileDir:""
                },
                shareObject: {
                    shareMode: 2,
                    shareKey: "",
                    joinKey: "",
                    assignUserMode: 0,
                    shareUrl: "",
                    joinUrl: ""
                },
            }
        },
        computed:{
            curFileSwitch(){
                return this.$store.getters.currentWorkspace
            },
            selectNode() {
                //当前选中的文件或者目录，需要打开
                return this.$store.getters.getSelectedNode
            },
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
                        this.curDirHtml = '<a href=javascript:window.fileSystemVueThis.clickToPath("/","",true)>/ </a>';
                    } else {
                        this.curDirHtml = this.curDirHtml + (parentDir == "/" ? "" : "/") + '<a href=javascript:window.fileSystemVueThis.clickToPath("' + parentDir + '","' + sa + '",' + false + ')>' + sa + '</a>'
                        parentDir = parentDir + (parentDir == "/" ? "" : "/") + sa;
                    }
                }
            }
        },
        methods: {
            focurCurFile(e){
                e.target.style.overflow='visible'
                e.target.style.textOverflow='clip'
            },
            renameCurrFile(e){
                let oriText=e.target.getAttribute("oriText");
                let newText=e.target.innerText;
                e.target.style.overflow='hidden'
                e.target.style.textOverflow='ellipsis'
                let selectNode=this.$store.getters.getSelectedNode;
                if(!selectNode.isDir && oriText!=newText){
                    let _this=this;
                    this.renameFile(newText,()=>{
                        _this.$set(selectNode, 'title', newText)
                        e.target.setAttribute("oriText",newText)
                    },(reason)=>{
                        e.target.setAttribute("oriText",oriText)
                        e.target.innerText=oriText;
                    })

                }
            },
            copyCurFile(){
                let _this = this;
                if (_this.selectNode==null) {
                    this.$Message.error("请选择一个文件");
                    return;
                }
                if (_this.selectNode.isDir) {
                    this.$Message.error("不允许复制目录");
                    return;
                }
                this.copyFile(function () {
                    _this.refreshCurView();
                })
            },
            deleteCurFile(){
                let _this=this;
                if (_this.selectNode==null) {
                    _this.$Message.error("请选择至少一个目标");
                    return;
                }
                if (_this.selectNode.isDir) {
                    this.deleteDir(_this.selectNode,()=>{
                        _this.refreshCurView();
                        this.$store.commit("setSelecedNode", null)
                    })
                    return;
                }else{
                    this.deleteFile(_this.selectNode,()=>{
                        _this.refreshCurView();
                        this.$store.commit("setSelecedNode", null)
                    })
                }
            },
            createCurDir(){
                let _this=this;
                let selectNode=_this.selectNode
                this._createDir(_this.selectNode,"请输入设置的文件夹名称",(fileDir,code)=>{
                    _this.refreshCurView();
                })
            },
            uploadFile(file){
                let _this=this;
                let index = this.curDir.lastIndexOf("/")
                let dirPath=this.curDir.substring(0,index)||"/"
                let fileName=this.curDir.substring(index+1)
                let root=false;
                if(!this.curDir || this.curDir=="/") root=true;
                this.uploadFileFun(file,dirPath,fileName,root,()=>{
                    _this.refreshCurView();
                })

            },
            createMdFileInCur() {
                this.createFileInCur("请输入文件名称：",".md")
            },
            createFlowFileInCur() {
                this.createFileInCur("请输入flow名称：",".flow");
            },
            createSnowFileInCur(){
                this.createFileInCur("请输入思维导图：",".mind");
            },
            createTextFileInCur() {
                this.createFileInCur("请输入文件名称：",null)
            },
            createWordFileInCur() {
                this.createFileInCur("请输入Word：", ".docx");
            },
            createExcelFileInCur() {
                this.createFileInCur("请输入Excel：", ".xlsx");
            },
            createPptFileInCur() {
                this.createFileInCur("请输入Ppt：", ".pptx");
            },
            canceldVpFileInCur(){
                this.cancelVpFile()
            },
            /*
            * 点击右键 从模板创建触发事件
            * */
            createFileFromTempalteInCur(){
                let _this = this;
                this.showTemplate = true;
                _this.loadTemplateGroup((templateGroup)=>{
                    _this.templateGroupData = templateGroup;
                    if (_this.templateGroupData.length > 0) {
                        _this.templateObject.templateGroupId = _this.templateGroupData[0].value;
                    }
                    _this.changeTemplateGroup();
                    let selectNode = _this.$store.getters.getSelectedNode
                    _this.templateObject.fileDir=selectNode.dirPath

                });
            },
            /*
            * 选择模板后点击确定后创建文件逻辑
            * */
            createFileFromTemplate() {
                let _this = this;
                let selectNode = this.$store.getters.getSelectedNode
                _this.templateObject.fileDir=selectNode.dirPath+"/"+selectNode.fileName
                this.createFileFromTemplateBack(_this.templateObject,()=>{
                    _this.refreshCurView(()=>{
                        _this.$nextTick(()=>{
                            let items=_this.dataArray.filter(item=>{
                                return item.title==_this.templateObject.fileName
                            })
                            _this.selectClick(items[0],'fileItem'+_this.templateObject.fileName)
                        })
                    });
                })
            },
            /*
            * 当模板组选择后，自动选择第一个模板默认
            * */
            changeTemplateGroup() {
                let _this = this;
                _this.loadTemplate(_this.templateObject.templateGroupId,(templdateData)=>{
                    _this.templateData = templdateData;
                    if (_this.templateData.length > 0) {
                        _this.templateObject.templateId = _this.templateData[0].value;
                        _this.templateObject.fileName=_this.templateData[0].templatePath.substring(_this.templateData[0].templatePath.lastIndexOf("/")+1)
                    }
                })
            },
            createVpFileInCur() {
                let _this = this;
                let selectNode = _this.selectNode
                this.createVpFile(selectNode,(code)=>{
                    _this.refreshCurView(()=>{
                        _this.$nextTick(()=>{
                            let items=_this.dataArray.filter(item=>{
                                return item.title==code
                            })
                            _this.clickFile(items[0])
                        })
                    });
                })
            },
            buildVpFileInCur(){
                let _this = this;
                //当前目录
                let selectNode = _this.selectNode;
                //当前选择的节点
                let selectedItem=_this.selectedItem;
                _this.$store.commit('showLoading')
                _this.buildVpFile(selectNode)
            },
            createFileInCur(title,suffix) {
                let _this = this;
                //手工选中某个节点
                this.createTextFun(_this.curDir||"/",suffix,title,(code)=>{
                    _this.refreshCurView(()=>{
                        _this.$nextTick(()=>{
                            let items=_this.dataArray.filter(item=>{
                                return item.title==code
                            })
                            _this.selectClick(items[0],'fileItem'+code)
                        })
                    });
                })
            },
            returnPreStep(func){
                let index = this.curDir.lastIndexOf("/")
                let dirPath=this.curDir.substring(0,index)||"/"
                let root=false;
                if(this.curDir=="/") root=true;
                if(!root){
                    if(dirPath=="/"){
                        this.clickToPath(dirPath,null,true,func)
                    }else {
                        let index = dirPath.lastIndexOf("/")
                        let preDirPath = dirPath.substring(0, index) || "/"
                        let preFileName = dirPath.substring(index + 1)
                        this.clickToPath(preDirPath, preFileName, false, func)
                    }
                }
            },
            refreshCurView(func){
                if(this.curDir=="/"){
                    this.clickToPath(this.curDir,"",true,func)
                }else {
                    let index = this.curDir.lastIndexOf("/")
                    let dirPath=this.curDir.substring(0,index)||"/"
                    let fileName=this.curDir.substring(index+1)
                    this.clickToPath(dirPath,fileName,false,func)
                }

            },
            clickToPath(dirPath, title, root,func){
                let fileSystemVueThis=this;
                if (root) {
                    fileSystemVueThis.curDir = "/"
                    fileSystemVueThis.initData(func);
                } else {
                    fileSystemVueThis.$axios.get(fileSystemVueThis.$globalConfig.goServer + "home/listSub?fileDir=" + dirPath + "&fileName=" + title + "&root=" + root).then((response) => {
                        fileSystemVueThis.dataArray = response.data.data //挂载子节点
                        fileSystemVueThis.selectedItem={
                            title: title,
                            fileName:title,
                            expand: true,
                            dirPath:dirPath,
                            contextmenu: true,
                            isDir: true,
                            root: false,
                            children:response.data.data
                        }
                        fileSystemVueThis.$store.commit("setSelecedNode", fileSystemVueThis.selectedItem)
                        func && func()
                    })
                    fileSystemVueThis.curDir = dirPath + (dirPath == "/" ? "" : "/") + title
                }
            },
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
            loadChildren(item,func){
                let _this=this;
                _this.$axios.get(_this.$globalConfig.goServer + "home/listSub?fileDir=" + item.dirPath + "&fileName=" + item.title + "&root=" + item.root).then((response) => {
                    item.children=response.data.data //挂载子节点
                    _this.$store.commit("setSelecedNode", item)
                    func && func()
                })
            },
            /**
             * 单击选择后需要调整选择的颜色和设置当前选择的item
             * @param item 选中item
             * @param refName 引用的元素
             */
            selectClick(item,refName){
                // this.loadChildren(item);
                this.$store.commit("setSelecedNode", item)//当前单击选中的节点。
                //this.selectedItem=item;
                //this.selectedRef=refName;
                for(let refTmp in this.$refs){
                    if(this.$refs[refTmp][0]) {
                        this.$refs[refTmp][0].style.border = "";
                        this.$refs[refTmp][0].style.backgroundColor = "";
                    }
                }
                this.$refs[refName][0].style.border="1px double rgb(153,209,255)";
                this.$refs[refName][0].style.backgroundColor="rgb(204,232,255)";
            },
            /**
             *  双击选择对应文件，并且进入文件夹内部或者打开文件
             * @param item
             * @param func
             */
            clickFile(item,func) {
                let _this=this;
                //如果是目录如果操作
                if(item.isDir) {
                    this.curDir = item.dirPath + (this.curDir == "/" ? "" : "/") + item.title
                    _this.loadChildren(item,function(){
                        _this.selectedItem=item
                        func && func();
                    })
                }else{
                    _this.$store.commit("setSelecedNode", item)
                    //打开文件
                    let mapping=this.$globalConfig.editorMapping
                    for(let key in mapping){
                        let re;
                        eval("re=/^.+("+key+")$/")
                        if (re.test(item.title) ) {
                            this.showEditorMadal=true
                            this.routePush(item,...mapping[key])
                            _this.loadEditorContent((vueThis, data) => {
                                let initInterval = setInterval(() => {
                                    if (vueThis.$route && vueThis.$route.matched && vueThis.$route.matched.length > 0) {
                                        let vueRouteComponents = vueThis.$route.matched[0].instances.default
                                        if (vueRouteComponents) {
                                            if (vueRouteComponents.initData) {
                                                vueRouteComponents.initData(data)
                                                clearInterval(initInterval)
                                            }
                                        }
                                    }
                                }, 100)
                            });
                            return;
                        }
                    }
                }
            },
            initData(func) {
                let vueThis = this;
                window.fileSystemVueThis = vueThis
                vueThis.curDir='/';
                vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?root=true&fileDir=/").then((response) => {
                    let curRoot = {
                        title: "",
                        fileName:"",
                        expand: true,
                        dirPath:'/',
                        contextmenu: true,
                        isDir: true,
                        root: true,
                        children:response.data.data
                    }
                    vueThis.dataArray = response.data.data //挂载子节点
                    vueThis.selectedItem=curRoot
                    vueThis.$store.commit("setSelecedNode", curRoot)
                    func && func();
                })
            },
        },
        components: {},
        created(){
            this.initData();
        },
        mounted() {

        }
    }
</script>
