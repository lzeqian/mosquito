<template>
    <div id="area">
        <Tree :data="data5" ref="tree" :render="renderContent" @on-select-change="selectChange"
              class="demo-tree-render" @on-contextmenu="handleContextMenu">
            <template slot="contextMenu">
                <DropdownItem @click.native="handleContextMenuEdit" v-if="selectNode!=null && !selectNode.isDir">重命名
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
                <DropdownItem @click.native="handleContextMenuCreateMd" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建md
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateFlow" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建flow
                </DropdownItem>
                <DropdownItem @click.native="handleContextMenuCreateFile" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">新建文件
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
    </div>
</template>

<script>
    export default {
        name: 'LeftTree',
        data() {
            return {
                isSpinShow: false,
                isCollapsed: false,
                isImgShow: false,
                data5: [],
            }
        },
        computed: {
            selectNode() {
                return this.$store.state.selectedNode
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
            handleUpload(file) {
                const param = new FormData();
                param.append('myfile', file)
                param.append('fileDir', this.selectNode.dirPath + "/" + this.selectNode.title)
                this.$axios.post(this.$globalConfig.goServer + "/file/upload", param).then(res => {
                    this.selectChange([this.selectNode])
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
                let selectNode=this.$store.state.selectedNode
                let code = prompt("请输入名称：",selectNode.title);
                var _this=this;
                if (code != null && code.trim() != "") {
                    this.$axios.post(this.$globalConfig.goServer + "file/rename?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title+ "&newFileName=" + code).then((response) => {
                        _this.$set(selectNode, 'title', code)
                    });
                }
            },
            handleContextMenuUpload(){
            },
            /**
             * 下载文件方法事件
             * */
            handleContextDownload() {
                var selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选择一个文件");
                    return;
                }
                var selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    this.$Message.error("不允许直接下载目录，请选择文件");
                    return;
                }
                window.location = this.$globalConfig.goServer + "file/download?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title
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
                let vueThis = this;
                let selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                    return;
                }
                var selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    this.$Message.error("不允许直接删除目录，请选择文件");
                    return;
                }
                let {index, parentNode} = this.getParent(this.$refs.tree.data[0], selectNode)
                this.$axios.delete(this.$globalConfig.goServer + "file/delete?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title).then((response) => {
                    parentNode.children.splice(index, 1)
                    this.$set(parentNode, 'selected', true)
                    // vueThis.selectNode=parentNode;
                    vueThis.$store.commit("setSelecedNode", parentNode)
                    vueThis.routePush({},'/blank',"空白预览")
                })
            },
            /**
             * 编译vuepress项目
             */
            handleContextMenuBuildVp() {
                var vueThis = this;
                vueThis.$store.commit('showLoading')
                var selectNodes = this.$refs.tree.getSelectedNodes()
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                }
                var selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    this.$axios.post(this.$globalConfig.goServer + "md/buildVp?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.title).then((response) => {
                        window.open(this.$globalConfig.goServer + response.data.data)
                        vueThis.$store.commit('hideLoading')
                    }).catch(()=>{
                        vueThis.$store.commit('hideLoading')
                    });
                }
                console.log("执行完成")
            },
            /**
             * 创建vuepress项目
             */
            handleContextMenuCreateVp() {
                var selectNodes = this.$refs.tree.getSelectedNodes()
                var vueThis = this;
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                }
                var selectNode = selectNodes[0];
                if (selectNode.isDir) {
                    var code = prompt("请输入vuepress名称：");
                    if (code != null && code.trim() != "") {
                        var fileDir = selectNode.dirPath + "/" + selectNode.title;
                        this.$axios.post(this.$globalConfig.goServer + "md/createVp?fileDir=" + fileDir + "&fileName=" + code).then((response) => {
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
                                path: this.redirect || '/mdeditor',
                                query: {dirPath: fileDir, fileName: code, content: ""}
                            });
                            vueThis.$set(selectNode, 'selected', false)
                            //vueThis.selectNode=selectNode.children[selectNode.children.length-1]
                            vueThis.$store.commit("setSelecedNode", selectNode.children[selectNode.children.length - 1])
                            vueThis.selectChange([vueThis.selectNode])
                        });
                    }
                }
            },
            handleContextMenuCreateFile() {

                this.handleContextMenuCreateText("请输入文件名称：",null,null);
            },
            handleContextMenuCreateText(title,suffix) {
                //手工选中某个节点
                let selectNodes = this.$refs.tree.getSelectedNodes()
                let vueThis = this;
                if (selectNodes.length == 0) {
                    this.$Message.error("请选选择展开子目录");
                }
                let selectNode = selectNodes[0];
                // vueThis.selectNode=selectNode;
                vueThis.$store.commit("setSelecedNode", selectNode)
                if (selectNode.isDir) {
                    if (selectNode.expand == false) {
                        vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?fileDir=" + vueThis.selectNode.dirPath + "&fileName=" + vueThis.selectNode.title).then((response) => {
                            vueThis.selectNode.children = response.data.data //挂载子节点
                            vueThis.selectNode.expand = true    //展开子节点
                        })
                    }
                    selectNode.expand = true;
                    var code = prompt(title);
                    if (code != null && code.trim() != "") {
                        let suffixRe=this.$globalConfig.supportFile
                        if(!suffix && !suffixRe.test(code)){
                            vueThis.$Message.error("该文件目不支持创建,只支持:"+suffixRe)
                            return;
                        }
                        if (suffix && !code.endsWith(suffix)) {
                            code = code + suffix;
                        }
                        var fileDir = selectNode.dirPath + "/" + selectNode.title;
                        this.$axios.post(this.$globalConfig.goServer + "file/create?fileDir=" + fileDir + "&fileName=" + code).then((response) => {
                            debugger
                            selectNode.children.push({
                                title: code,
                                dirPath: fileDir,
                                expand: false,
                                contextmenu: true,
                                isDir: false,
                                selected: true,
                                children: []
                            })
                            vueThis.$set(selectNode, 'selected', false)
                            let newSelected = selectNode.children[selectNode.children.length - 1]
                            vueThis.$store.commit("setSelecedNode",newSelected )
                            vueThis.selectChange([newSelected])
                        })
                    }
                } else {
                    this.$Message.error("请选选择展开子目录");
                }
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
                                            if (data.title.endsWith(".doc") || data.title.endsWith(".docx")) {
                                                return "icon-doc"
                                            }
                                            if (data.title.endsWith(".xls") || data.title.endsWith(".xlsx")) {
                                                return "icon-xls"
                                            }
                                            if (data.title.endsWith(".ppt") || data.title.endsWith(".pptx")) {
                                                return "icon-ppt"
                                            }
                                            if (data.title.endsWith(".json")) {
                                                return "icon-json"
                                            }
                                            if (data.title.endsWith(".js")) {
                                                return "icon-js-square"
                                            }
                                            if (data.title.endsWith(".pdf")) {
                                                return "icon-pdf"
                                            }
                                            if (/.*\.(png|PNG|jpg|JPG|JPEG|jpeg|gif|GIF)/.test(data.title)) {
                                                return "icon-picture"
                                            }
                                            if (/.*\.(zip|7z|rar)/.test(data.title)) {
                                                return "icon-zip"
                                            }
                                            if (data.title.endsWith(".md")) {
                                                return "icon-file-markdown"
                                            }
                                            if (data.title.endsWith(".html")) {
                                                return "icon-HTML"
                                            }
                                            if (data.title.endsWith(".xml")) {
                                                return "icon-xml"
                                            }
                                            return 'icon-wenjian'
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
            routePush(node,routerAddress,title){
                this.contentTitle = title
                this.$router.push({
                    path: this.redirect || routerAddress,
                    query: {dirPath: node.dirPath, fileName: node.title}
                });
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
                        vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?fileDir=" + node.dirPath + "&fileName=" + node.title+ "&root=" + node.root).then((response) => {
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
                    await this.$axios.post(this.$globalConfig.goServer + "file/mkdir?fileDir=" + parentDir + "&fileName=" + name).then((response) => {
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
                let selectNode=this.$store.state.selectedNode
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
            }
        }
        ,
        props: {}
        ,
        mounted() {
            var vueThis = this;
            vueThis.$axios.get(this.$globalConfig.goServer + "home/tree").then((response) => {
                var resultData = response.data
                vueThis.data5 = [
                    {
                        title: "目录结构树",
                        expand: true,
                        dirPath:'/',
                        contextmenu: false,
                        isDir: true,
                        root: true,
                        children: resultData.data
                    }];
                vueThis.$nextTick(()=>{
                    vueThis.initData()
                })

            })
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
