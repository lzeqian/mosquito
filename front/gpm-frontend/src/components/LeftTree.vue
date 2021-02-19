<template>
    <div id="area">
        <Tree :data="data5" ref="tree" :render="renderContent" @on-select-change="selectChange" v-if="!isCollapsed"
              class="demo-tree-render" @on-contextmenu="handleContextMenu">
            <template slot="contextMenu">
                <DropdownItem @click.native="handleContextMenuEdit" v-if="selectNode!=null && !selectNode.isDir">重命名
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCopy" v-if="selectNode!=null && !selectNode.isDir">复制
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuDelete" v-if="selectNode!=null && !selectNode.isDir">删除
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuUpload" v-if="selectNode!=null && selectNode.isDir">
                    <Upload :style="{width:'100%'}" :show-upload-list="false" action=""
                            ref="upload"
                            :before-upload="handleUpload"
                    >上传
                    </Upload>
                </DropdownItem>
                <DropdownItem @click.native="handleContextDownload" v-if="selectNode!=null && !selectNode.isDir">下载
                </DropdownItem>
                <DropdownItem @click.native="handleContextShare" v-if="selectNode!=null && !selectNode.isDir && $store.getters.currentWorkspace==1">分享
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateMd" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建md
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateFlow" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建flow
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateSnow" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建思维导图
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateDir" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建目录
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuDeleteDir" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">删除目录
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateFile" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建文件
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateWord" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建word
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateExcel" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建excel
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreatePpt" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建ppt
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateVp" style="color: green"
                              v-if="selectNode!=null && selectNode.isDir">新建vuepress
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuBuildVp" style="color: green"
                              v-if="selectNode!=null && checkIfVb(selectNode)">构建vuepress
                </DropdownItem>
            </template>
        </Tree>
        <Spin fix v-show="isSpinShow">
            <Icon type="load-c" size="30" class="demo-spin-icon-load"></Icon>
            <div>Loading...</div>
        </Spin>
        <Modal
                v-model="showShare"
                title="分享"
                @on-ok="ok"
                :z-index="10002"
                @on-cancel="cancel">
            谁可以查看/编辑文档<br/>
            <RadioGroup v-model="shareObject.shareMode">
                <Radio :label="0">仅仅我自己</Radio><br/>
                <Radio :label="1">仅我分享的好友</Radio>
                <br  v-if="shareObject.shareMode==1"/>
                <RadioGroup v-model="shareObject.assignUserMode" v-if="shareObject.shareMode==1" style="margin-left:50px">
                    <Radio :label="0">可查看</Radio><br/>
                    <Radio :label="1">可编辑</Radio><br/>
                    分享加入url: <a :href="shareObject.shareUrl">{{shareObject.joinUrl}}</a>
                </RadioGroup><br/>
                <Radio :label="2">所有人可查看</Radio><br/>
                <Radio :label="3">所有人可编辑</Radio><br/>
            </RadioGroup><br/><br/>
            文档url: <a :href="shareObject.shareUrl">{{shareObject.shareUrl}}</a>
        </Modal>
        <div style="position: absolute;top:0px;right:5px">
            <a @click="gotoDesktop">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-dasuolvetuliebiao"></use>
                </svg>
            </a>&nbsp;
            <a @click="gotoPerson">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-renshu" style="color: blue"></use>
                </svg>
            </a>
        </div>
    </div>

</template>

<script>
    import {randomUuid} from "../utils/utils";

    export default {
        name: 'LeftTree',
        props:{
            isCollapsed: {
                type: Boolean,
                default: false
            },
        },
        data() {
            return {
                showShare:false,
                shareObject:{
                    shareMode:2,
                    shareKey:"",
                    joinKey:"",
                    assignUserMode:1,
                    shareUrl:"",
                    joinUrl:""
                },

                isSpinShow: false,
                isImgShow: false,
                data5: [],
            }
        },
        computed: {
            selectNode() {
                return this.$store.getters.getSelectedNode
            },
            rotateIcon() {
                return [
                    'menu-icon',
                    this.isCollapsed ? 'rotate-icon' : ''
                ];
            },
            menuitemClasses() {
                return [
                    'menu-item',
                    this.isCollapsed ? 'collapsed-menu' : ''
                ]
            }
        },
        methods: {
            getShareUrl(){
                this.shareObject.shareKey=randomUuid(8);
                return window.location.protocol+this.$globalConfig.goServer+"docs/"+this.shareObject.shareKey;
            },
            getJoinUrl(){
                this.shareObject.joinKey=randomUuid(8);
                return window.location.protocol+this.$globalConfig.goServer+"docJoin/"+this.shareObject.joinKey;
            },
            handleContextShare(){
                this.showShare=true;
                this.shareObject.shareUrl=this.getShareUrl();
                this.shareObject.joinUrl=this.getJoinUrl();
            },
            ok () {
                let _this=this;
                let selectNode=this.$store.getters.getSelectedNode
                this.$axios.post(this.$globalConfig.goServer+"share/shareFile",{
                    fileDir:selectNode.dirPath,
                    fileName:selectNode.fileName,
                    shareUserName:localStorage.getItem("userName"),
                    ..._this.shareObject
                }).then((resp)=>{
                    if(resp.data.code==0){
                        this.$Message.info('分享成功');
                    }
                })
            },
            cancel () {
                this.$Message.info('Clicked cancel');
            },
            gotoDesktop(){
                this.$store.commit("updateDirTree","desktop")
                this.routePush({},'/blank',"空白预览")
            },
            gotoPerson(){
                let workspace=this.$store.getters.currentWorkspace=="1"?"0":"1";
                let title=this.$store.getters.currentWorkspace=="1"?"公共文档库":"个人文档库";
                this.$store.commit("updateWorkspace",workspace)
                this.routePush({},'/blank',"空白预览")
                this.listRoot(title)
            },
            handleUpload(file) {
                let _this=this;
                this.uploadFile(file,()=>{
                    _this.selectChange([_this.selectNode])
                })
            },
            /**
             * 右键菜单被选中时自动选择当前节点
             * */
            handleContextMenu(data) {
                this.contextData = data;
                var curSelectNodes = this.$refs.tree.getSelectedNodes()
                if (curSelectNodes.length > 0) {
                    var curSelectNode = curSelectNodes[0];
                    this.$set(curSelectNode, 'selected', false)
                }

                this.$set(data, 'selected', true)
                // this.selectNode=data;
                this.$store.commit("setSelecedNode", data)
                if (!data.expand) {
                    this.selectChange([data])
                }
            },
            handleContextMenuEdit() {
                let selectNode=this.$store.getters.getSelectedNode
                let _this=this;
                this.editFile((code)=>{
                    _this.$set(selectNode, 'title', code)
                })
            },
            handleContextMenuUpload(){
            },
            /**
             * 下载文件方法事件
             * */
            handleContextDownload() {
                let selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选择一个文件");
                    return;
                }
                let selectNode = selectNodes[0];
                this.downloadFile(selectNode)
            },
            /**
             * 从跟节点遍历获取当前节点父节点
             */
            getParent(root, filterNode) {
                for (let [index, node] of new Map(root.children.map((node, i) => [i, node]))) {
                    if (filterNode.title == node.title && filterNode.dirPath == node.dirPath) {
                        return {index: index, parentNode: root};
                    } else {
                        if (node.isDir) {
                            var returnNode = this.getParent(node, filterNode)
                            if (returnNode != null)
                                return returnNode
                        }
                    }
                }
                return null;
            },
            /**
             * 删除文件
             */
            handleContextMenuDelete() {
                let _this = this;
                let selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                    return;
                }
                let selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    this.$Message.error("不允许直接删除目录，请选择文件");
                    return;
                }
                let {index, parentNode} = _this.getParent(_this.$refs.tree.data[0], selectNode)
                this.deleteFile(selectNode,()=>{
                    parentNode.children.splice(index, 1)
                    _this.$set(parentNode, 'selected', true)
                    _this.$store.commit("setSelecedNode", parentNode)
                    _this.routePush({},'/blank',"空白预览")
                })
            },
            /**
             * 删除文件
             */
            handleContextMenuDeleteDir() {
                let _this = this;
                let selectNode = this.selectNode;
                let {index, parentNode} = _this.getParent(_this.$refs.tree.data[0], selectNode)
                this.deleteDir(selectNode,()=>{
                    parentNode.children.splice(index, 1)
                    _this.$set(parentNode, 'selected', true)
                    _this.$store.commit("setSelecedNode", parentNode)
                    _this.routePush({},'/blank',"空白预览")
                })
            },
            /**
             * 编译vuepress项目
             */
            handleContextMenuBuildVp() {
                let _this = this;
                _this.$store.commit('showLoading')
                let selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                }
                let selectNode = selectNodes[0];
                _this.buildVpFile(selectNode)

            },
            handleContextMenuCopy(){
                debugger
                let _this = this;
                let selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选择一个文件");
                    return;
                }
                let selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    this.$Message.error("不允许复制目录");
                    return;
                }
                let {index, parentNode} = _this.getParent(_this.$refs.tree.data[0], selectNode)
                this.copyFile(function () {
                    _this.selectChange([parentNode])
                })
            },
            /**
             * 创建vuepress项目
             */
            handleContextMenuCreateVp() {
                let selectNodes = this.$refs.tree.getSelectedNodes()
                let _this = this;
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                }
                let selectNode = selectNodes[0];
                let fileDir = selectNode.dirPath + "/" + selectNode.title;
                this.createVpFile(selectNode,(code)=>{
                    selectNode.children.push({
                        title: code,
                        dirPath: fileDir,
                        expand: true,
                        contextmenu: true,
                        isDir: true,
                        selected: true,
                        children: []
                    })
                    this.$router.push({
                        path: _this.redirect || '/mdeditor',
                        query: {dirPath: fileDir, fileName: code, content: ""}
                    });
                    _this.$set(selectNode, 'selected', false)
                    _this.$store.commit("setSelecedNode", selectNode.children[selectNode.children.length - 1])
                    _this.selectChange([vueThis.selectNode])
                })
            },
            handleContextMenuCreateDir(){
                let selectNode=this.selectNode
                this.createDir(this.selectNode,"请输入设置的文件夹名称",(fileDir,code)=>{
                    if(!selectNode.children){
                        selectNode.children=[]
                    }
                    selectNode.children.push({
                        title: code,
                        fileName:code,
                        dirPath: fileDir,
                        expand: false,
                        contextmenu: true,
                        isDir: true,
                        selected: false,
                        children: []
                    })
                })
            },
            handleContextMenuCreateFile() {
                this.handleContextMenuCreateText("请输入文件名称：",null,null);
            },
            handleContextMenuCreateText(title,suffix) {
                //手工选中某个节点
                let selectNodes = this.$refs.tree.getSelectedNodes()
                let _this = this;
                if (selectNodes.length == 0) {
                    this.$Message.error("请先选择目录");
                }
                let selectNode = selectNodes[0];
                _this.$store.commit("setSelecedNode", selectNode)
                if (selectNode.isDir) {
                    if (selectNode.expand == false) {
                        vueThis.$axios.get(_this.$globalConfig.goServer + "home/listSub?fileDir=" + _this.selectNode.dirPath + "&fileName=" + _this.selectNode.title).then((response) => {
                            _this.selectNode.children = response.data.data //挂载子节点
                            _this.selectNode.expand = true    //展开子节点
                        })
                    }
                    selectNode.expand = true;
                }
                let fileDir = selectNode.dirPath + "/" + selectNode.fileName
                this.createTextFile(selectNode,title,suffix,(code)=>{
                    if(!selectNode.children){
                        selectNode.children=[]
                    }
                    selectNode.children.push({
                        title: code,
                        fileName:code,
                        dirPath: fileDir,
                        expand: false,
                        contextmenu: true,
                        isDir: false,
                        selected: true,
                        children: []
                    })
                    _this.$set(selectNode, 'selected', false)
                    let newSelected = selectNode.children[selectNode.children.length - 1]
                    _this.$store.commit("setSelecedNode",newSelected )
                    _this.selectChange([newSelected])
                })
            },
            /**
             * 创建markdown文件
             */
            handleContextMenuCreateMd() {
                this.handleContextMenuCreateText("请输入markdown名称：",".md");
            },
            /**
             * 创建markdown文件
             */
            handleContextMenuCreateFlow() {
                this.handleContextMenuCreateText("请输入flow名称：",".flow");
            },
            handleContextMenuCreateSnow(){
                this.handleContextMenuCreateText("请输入思维导图：",".mind");
            },
            handleContextMenuCreateWord(){
                this.handleContextMenuCreateText("请输入Word：",".docx");
            },
            handleContextMenuCreateExcel(){
                this.handleContextMenuCreateText("请输入Excel：",".xlsx");
            },
            handleContextMenuCreatePpt(){
                this.handleContextMenuCreateText("请输入Ppt：",".pptx");
            },
            /**
             * 检查是否是vuepress项目
             * @param data 当前节点数据
             * @returns {boolean}
             */
            checkIfVb(data) {
                if (data && data.isDir) {
                    if (data.expand && data.children && data.children.length>0) {
                        for (let c of data.children) {
                            if (c.title == ".vuepress") {
                                return true
                            }
                        }
                    }
                }
                return false
            },
            /**
             * render树结构的图标和内容
             * @param h
             * @param root
             * @param node
             * @param data
             * @returns {*}
             */
            renderContent(h, {root, node, data}) {
                var vueThis = this;
                return h('span', {
                    style: {
                        display: 'inline-block',
                        width: '100%'
                    }
                }, [
                    h('span', [
                        // h('Icon', {
                        //     props: {
                        //         type: function(){
                        //             if(data.root){
                        //                 return 'logo-windows'
                        //             }else if(data.isDir){
                        //                 if(vueThis.checkIfVb(data)){
                        //                     return "ios-bug-outline"
                        //                 }
                        //                 return 'iconfont icon-xls'
                        //             }else{
                        //                 return 'ios-paper-outline'
                        //             }
                        //         }()
                        //     },
                        //     style: {
                        //         fontFamily:'iconfont',
                        //         marginRight: '8px',
                        //         color:function(){
                        //             if(data.isDir) {
                        //                 if (data.expand) {
                        //                     for (let c of data.children) {
                        //                         if (c.title == ".vuepress") {
                        //                             return "red"
                        //                         }
                        //                     }
                        //                 }
                        //                 return ""
                        //             }
                        //
                        //         }()
                        //
                        //     }
                        // }),
                        h('svg', {
                            attrs: {
                                class: "icon",
                            },
                            style: {
                                width: function () {
                                    if (data.root) {
                                        return 25
                                    }
                                    return 20
                                }(),
                                height: function () {
                                    if (data.root) {
                                        return 25
                                    }
                                    return 20
                                }(),

                            }
                        }, [
                            h("use", {
                                style: {
                                    fill: function () {
                                        if (data.title.endsWith(".md")) {
                                            return "gray"
                                        }
                                        return ""
                                    }(),
                                },
                                attrs: {
                                    "xlink:href": "#" + function () {
                                        if (data.root) {
                                            return 'icon-wenjianjia'
                                        } else if (data.isDir) {
                                            if (vueThis.checkIfVb(data)) {
                                                return "icon-Vue"
                                            }
                                            return 'icon-wenjianjia'
                                        } else {
                                            return vueThis.fileIcon(data.title)
                                        }
                                    }(),
                                }
                            })
                        ]),
                        h('span', {
                            style: {
                                marginLeft: '18px',
                                position: 'relative',
                                zIndex: 10
                            }, attrs: {
                                title: data.title
                            },
                        }, data.title)
                    ]),

                ]);
            },
            /**
             * 树节点被选中时触发的编辑器打开和父节点展开事件
             * @param selectedList
             */
            selectChange(selectedList) {
                if(selectedList.length==0){
                    this.routePush({},'/blank',"空白预览")
                }
                const node = selectedList[selectedList.length - 1]
                if (node) {
                    this.$store.commit("setSelecedNode", node)
                    var vueThis = this;
                    this.$set(node, 'selected', true)
                    if (node.isDir) {
                        vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?fileDir=" + node.dirPath + "&fileName=" + node.fileName+ "&root=" + node.root).then((response) => {
                            node.children = response.data.data //挂载子节点
                            node.expand = true    //展开子节点
                        })
                    } else {
                        let mapping=this.$globalConfig.editorMapping
                        for(let key in mapping){
                            let re;
                            eval("re=/^.+("+key+")$/")
                            if (re.test(node.title) ) {
                                this.routePush(node,...mapping[key])
                                return;
                            }
                        }
                    }
                    //没有push直接跳转到白板页面
                    this.routePush({},'/blank',"空白预览")
                }
            },
            preventDefault(){
                this.$(document).on({
                    dragleave:function(e){      //拖离
                        e.preventDefault();
                    },
                    drop:function(e){           //拖后放
                        e.preventDefault();
                    },
                    dragenter:function(e){      //拖进
                        e.preventDefault();
                    },
                    dragover:function(e){       //拖来拖去
                        e.preventDefault();
                    }
                });
            },
            async uploadEntry(parentDir,entry){
                let name=entry.name;
                let _this=this;
                if(entry.isFile){
                    entry.file(async function (file) {
                        const param = new FormData();
                        param.append('myfile', file)
                        param.append('fileDir', parentDir)
                        await _this.$axios.post(_this.$globalConfig.goServer + "/file/upload", param).then(res => {

                        })
                    })
                }else{
                    //服务器创建目录
                    await this.$axios.post(this.$globalConfig.goServer + "file/mkdir" ,{fileDir:parentDir,fileName:name}).then((response) => {
                    });
                    let dirReader=entry.createReader()
                    dirReader.readEntries(async function (entries) {
                        for(let centry of entries){
                            _this.uploadEntry(parentDir+"/"+name,centry)
                        }
                    })
                }
            },
            dropUpload(e){
                e.preventDefault(); //取消默认浏览器拖拽效果
                let selectNode=this.$store.getters.getSelectedNode
                var fileDir = selectNode.dirPath + "/" + selectNode.title;
                let _this=this;
                if(!selectNode){
                    this.$message.error("请选择上传的目录")
                    return;
                }
                let fileList = e.dataTransfer.files; //获取文件对象
                //检测是否是拖拽文件到页面的操作
                if (fileList.length == 0) {
                    return false;
                }
                for(var i=0;i<fileList.length;i++){
                    let item=e.dataTransfer.items[i]
                    let entry=item.webkitGetAsEntry()
                    this.uploadEntry(fileDir,entry)
                    _this.selectChange([selectNode])
                }

            },
            initData(){
                let _this = this;
                let dirPath=this.$route.query.dirPath
                let fileName=this.$route.query.fileName
                let root=this.$refs.tree.data[0]
                let curParent=root;
                if(dirPath && fileName){
                    let newDirPath=dirPath.replace(/[/\\|/|\\]+/,"\\")
                    let dirPathSplit=newDirPath.split("\\");
                    //选择目录
                    for(let j=0;j<dirPathSplit.length;j++){
                        let curDirPath=dirPathSplit[j]
                        if(curDirPath!=null && curDirPath!=""){
                            for (let [index, node] of new Map(curParent.children.map((node, i) => [i, node]))) {
                                if (curDirPath == node.title) {
                                    if (node.isDir) {
                                        _this.$axios.get(_this.$globalConfig.goServer + "home/listSub?fileDir=" + node.dirPath + "&fileName=" + node.title+ "&root=" + node.root).then((response) => {
                                            node.children = response.data.data //挂载子节点
                                            node.expand = true    //展开子节点
                                            _this.$nextTick(()=>{
                                                //选择目录下的文件
                                                for (let [index, tnode] of new Map(curParent.children.map((tnode, i) => [i, tnode]))) {
                                                    if (fileName == tnode.title) {
                                                        _this.selectChange([tnode])
                                                        break;
                                                    }
                                                }
                                            })
                                        })
                                    }
                                    curParent=node
                                    break;
                                }
                            }
                        }
                    }

                }
            },
            listRoot(title){
                let vueThis = this;
                vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?root=true&fileDir=/").then((response) => {
                    var resultData = response.data
                    vueThis.data5 = [
                        {
                            title: title,
                            fileName:"",
                            expand: true,
                            dirPath:'/',
                            contextmenu: vueThis.$store.getters.currentWorkspace==0?false:true,
                            isDir: true,
                            root: true,
                            children: resultData.data
                        }];
                    vueThis.$nextTick(()=>{
                        vueThis.initData()
                    })

                })
            }
        }
        ,
        mounted() {
            let title=this.$store.getters.currentWorkspace=="0"?"公共文档库":"个人文档库";
            this.listRoot(title);
            this.preventDefault()
            document.getElementById('area').addEventListener("drop",
                this.dropUpload,
                false);

        }
    }
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
    .menu-item span {
        display: inline-block;
        overflow: hidden;
        width: 69px;
        text-overflow: ellipsis;
        white-space: nowrap;
        vertical-align: bottom;
        transition: width .2s ease .2s;
    }

    .menu-item i {
        transform: translateX(0px);
        transition: font-size .2s ease, transform .2s ease;
        vertical-align: middle;
        font-size: 16px;
    }

    .collapsed-menu span {
        width: 0px;
        transition: width .2s ease;
    }

    .collapsed-menu i {
        transform: translateX(5px);
        transition: font-size .2s ease .2s, transform .2s ease .2s;
        vertical-align: middle;
        font-size: 22px;
    }

    .icon {
        width: 1em;
        height: 1em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }
</style>
