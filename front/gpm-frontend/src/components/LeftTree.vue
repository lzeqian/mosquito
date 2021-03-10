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
                <DropdownItem @click.native="handleContextCollect" v-if="selectNode!=null">收藏
                </DropdownItem>
                <DropdownItem @click.native="handleContextDownload" v-if="selectNode!=null && !selectNode.isDir">下载
                </DropdownItem>
                <DropdownItem @click.native="handleContextShare"
                              v-if="selectNode!=null && !selectNode.isDir && $store.getters.currentWorkspace==1">分享
                </DropdownItem>
                <DropdownItem @click.native="handleContextSendEmail"
                              v-if="selectNode!=null && !selectNode.isDir">发送邮件
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
                <DropdownItem @click.native="handleContextMenuCreateFileFromTemplate" style="color: #ed4014"
                              v-if="selectNode!=null && selectNode.isDir">从模板新建
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
                <DropdownItem @click.native="handleContextMenuCannelVp" style="color: green"
                              v-if="selectNode!=null && checkIfVb(selectNode)">取消映射
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
                @on-ok="sharePersonFile"
                :z-index="10002">
            谁可以查看/编辑文档<br/>
            <RadioGroup v-model="shareObject.shareMode">
                <Radio :label="0">仅仅我自己</Radio>
                <br/>
                <Radio :label="3">仅我分享的好友</Radio>
                <br v-if="shareObject.shareMode==3"/>
                <RadioGroup v-model="shareObject.assignUserMode" v-if="shareObject.shareMode==3"
                            style="margin-left:50px">
                    <Radio :label="0">可查看</Radio>
                    <br/>
                    <Radio :label="1">可编辑</Radio>
                    <br/>
                    分享加入url: <a :href="shareObject.shareUrl">{{shareObject.joinUrl}}</a>
                </RadioGroup>
                <br/>
                <Radio :label="1">所有人可查看</Radio>
                <br/>
                <Radio :label="2">所有人可编辑</Radio>
                <br/>
            </RadioGroup>
            <br/><br/>
            <div v-if="preShareKey==null || preShareKey==''">文档url: <a :href="shareObject.shareUrl">{{shareObject.shareUrl}}</a>
            </div>
            <div v-if="preShareKey!=null && preShareKey!=''" style="color:red">
                已分享:<a style="color:red" :href="getPreShareUrl()">{{getPreShareUrl()}}</a> ,<a style="color:gray"
                                                                                               @click="cancelShare">取消分享</a>
            </div>
        </Modal>
        <Modal
                v-model="showTemplate"
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
                    <Select v-model="templateObject.templateId" style="width:200px">
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
        <Modal
                v-model="showSendEmail"
                title="模板选择"
                @on-ok="sendEmail"
                :z-index="10002">
            <Form :model="emailObject" :label-width="80">
                <FormItem label="主题">
                    <Input v-model="emailObject.subject" style="width:80%">
                    </Input>
                </FormItem>
                <FormItem label="收件人">
                    <Input v-model="emailObject.receiver" style="width:80%">
                    </Input>
                </FormItem>
            </Form>
        </Modal>
        <div style="position: absolute;top:0px;right:5px">
            <a @click="gotoDesktop">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-dasuolvetuliebiao"></use>
                </svg>
            </a>&nbsp;
            <a @click="gotoPerson">
                <svg class="icon" aria-hidden="true">
                    <use :xlink:href="$store.getters.currentWorkspace == '0'?'#icon-duoren-renqun':'#icon-renshu'" style="color: blue"></use>
                </svg>
            </a>
        </div>
    </div>

</template>

<script>
    import {randomUuid} from "../utils/utils";

    export default {
        name: 'LeftTree',
        props: {
            isCollapsed: {
                type: Boolean,
                default: false
            },
        },
        data() {
            return {
                showSendEmail:false,
                emailObject:{
                    subject:"",
                    receiver:"",
                    fileName: "",
                    fileDir:""
                },
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
                preShareKey: null,
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
            handleContextSendEmail(){
                this.showSendEmail=true;
            },
            sendEmail(){
                let _this = this;
                let selectNode = this.$store.getters.getSelectedNode
                _this.emailObject.fileDir=selectNode.dirPath
                _this.emailObject.fileName=selectNode.fileName
                this.sendEmailToBack(_this.emailObject,()=>{
                    _this.$Message.info('发送成功');
                });
            },
            getShareUrl() {
                this.shareObject.shareKey = randomUuid(8);
                return window.location.protocol + this.$globalConfig.goServer + "docs/" + this.shareObject.shareKey;
            },
            getJoinUrl() {
                this.shareObject.joinKey = randomUuid(8);
                return window.location.protocol + this.$globalConfig.goServer + "docJoin/" + this.shareObject.joinKey;
            },
            getPreShareUrl() {
                return window.location.protocol + this.$globalConfig.goServer + "docs/" + this.preShareKey;
            },
            handleContextShare() {
                let _this = this;
                this.showShare = true;
                this.shareObject.shareUrl = this.getShareUrl();
                this.shareObject.joinUrl = this.getJoinUrl();
                let selectNode = this.$store.getters.getSelectedNode
                this.$axios.get(this.$globalConfig.goServer + "share/isShareFile?fileDir=" + selectNode.dirPath + "&fileName=" + selectNode.fileName).then((response) => {
                    if (response.data.code == 0) {
                        _this.preShareKey = response.data.data.ShareKey;
                    }
                })
            },
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
            /**
             * 点击模板创建文件菜单弹出
             * */
            handleContextMenuCreateFileFromTemplate() {
                let _this = this;
                this.showTemplate = true;
                _this.loadTemplateGroup((templateGroup)=>{
                    _this.templateGroupData = templateGroup;
                    if (_this.templateGroupData.length > 0) {
                        _this.templateObject.templateGroupId = _this.templateGroupData[0].value;
                    }
                    _this.changeTemplateGroup();
                    let selectNode = this.$store.getters.getSelectedNode
                    _this.templateObject.fileDir=selectNode.dirPath

                });
            },
            /**
             * 模板创建文件点
             * */
            createFileFromTemplate() {
                let _this = this;
                let selectNode = this.$store.getters.getSelectedNode
                _this.templateObject.fileDir=selectNode.dirPath+"/"+selectNode.fileName
                this.createFileFromTemplateBack(_this.templateObject,()=>{
                    _this.$Message.info('创建成功');
                    _this.selectChange([selectNode],()=>{
                        _this.$set(selectNode, 'selected', false)
                        for(let c of selectNode.children){
                            console.log(c.fileName)
                            if(c.fileName==_this.templateObject.fileName){
                                _this.$store.commit("setSelecedNode", c)
                                _this.selectChange([ _this.$store.getters.getSelectedNode])
                                break;
                            }
                        }
                    })
                });
            },
            /**
             * 分享文件
             * */
            sharePersonFile() {
                let _this = this;
                let selectNode = this.$store.getters.getSelectedNode
                let requestUrl = (this.preShareKey != null && this.preShareKey != "" ? "updateShareFile" : "shareFile");
                this.$axios.post(this.$globalConfig.goServer + "share/" + requestUrl, {
                    fileDir: selectNode.dirPath,
                    fileName: selectNode.fileName,
                    shareUserName: localStorage.getItem("userName"),
                    ..._this.shareObject
                }).then((resp) => {
                    if (resp.data.code == 0) {
                        this.$Message.info('分享成功');
                    }
                })
            },
            cancelShare() {
                let _this = this;
                _this.cancelShareFile(this.preShareKey,()=>{
                    _this.preShareKey = null;
                })
            },
            gotoDesktop() {
                this.$store.commit("updateDirTree", "desktop")
                this.routePush({}, '/default', "空白预览")
            },
            gotoPerson() {
                let workspace = this.$store.getters.currentWorkspace == "1" ? "0" : "1";
                let title = this.$store.getters.currentWorkspace == "1" ? "公共文档库" : "个人文档库";
                this.$store.commit("updateWorkspace", workspace)
                this.routePush({}, '/default', "空白预览")
                this.listRoot(title)
            },
            handleUpload(file) {
                let _this = this;
                this.uploadFile(file, () => {
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
                let selectNode = this.$store.getters.getSelectedNode
                let _this = this;
                this.editFile((code) => {
                    _this.$set(selectNode, 'title', code)
                    _this.$set(selectNode, 'fileName', code)
                })
            },
            handleContextMenuUpload() {
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
            handleContextCollect(){
                let selectNode = this.$store.getters.getSelectedNode
                this.collectFavorite({
                    fileDir:selectNode.dirPath,
                    fileName:selectNode.fileName
                })
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
                this.deleteFile(selectNode, () => {
                    parentNode.children.splice(index, 1)
                    _this.$set(parentNode, 'selected', true)
                    _this.$store.commit("setSelecedNode", parentNode)
                    _this.routePush({}, '/default', "空白预览")
                })
            },
            /**
             * 删除文件
             */
            handleContextMenuDeleteDir() {
                let _this = this;
                let selectNode = this.selectNode;
                let {index, parentNode} = _this.getParent(_this.$refs.tree.data[0], selectNode)
                this.deleteDir(selectNode, () => {
                    parentNode.children.splice(index, 1)
                    _this.$set(parentNode, 'selected', true)
                    _this.$store.commit("setSelecedNode", parentNode)
                    _this.routePush({}, '/default', "空白预览")
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
            handleContextMenuCannelVp(){
                this.cancelVpFile()
            },
            handleContextMenuCopy() {
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
                this.createVpFile(selectNode, (code) => {
                    selectNode.children.push({
                        title: code,
                        fileName:code,
                        dirPath: selectNode.root?"/":fileDir,
                        expand: true,
                        contextmenu: true,
                        isDir: true,
                        selected: true,
                        children: []
                    })
                    _this.$set(selectNode, 'selected', false)
                    _this.$store.commit("setSelecedNode", selectNode.children[selectNode.children.length - 1])
                    _this.selectChange([ _this.$store.getters.getSelectedNode])
                })
            },
            handleContextMenuCreateDir() {
                let _this=this;
                let selectNode = this.selectNode
                this.createDir(this.selectNode, "请输入设置的文件夹名称", (fileDir, code) => {
                    if (!selectNode.children) {
                        selectNode.children = []
                    }
                    selectNode.children.push({
                        title: code,
                        fileName: code,
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
                this.handleContextMenuCreateText("请输入文件名称：", null, null);
            },
            handleContextMenuCreateText(title, suffix) {
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
                this.createTextFile(selectNode, title, suffix, (code) => {
                    if (!selectNode.children) {
                        selectNode.children = []
                    }
                    selectNode.children.push({
                        title: code,
                        fileName: code,
                        dirPath: fileDir,
                        expand: false,
                        contextmenu: true,
                        isDir: false,
                        selected: true,
                        children: []
                    })
                    _this.$set(selectNode, 'selected', false)
                    let newSelected = selectNode.children[selectNode.children.length - 1]
                    _this.$store.commit("setSelecedNode", newSelected)
                    _this.selectChange([newSelected])
                })
            },
            /**
             * 创建markdown文件
             */
            handleContextMenuCreateMd() {
                this.handleContextMenuCreateText("请输入markdown名称：", ".md");
            },
            /**
             * 创建markdown文件
             */
            handleContextMenuCreateFlow() {
                this.handleContextMenuCreateText("请输入flow名称：", ".flow");
            },
            handleContextMenuCreateSnow() {
                this.handleContextMenuCreateText("请输入思维导图：", ".mind");
            },
            handleContextMenuCreateWord() {
                this.handleContextMenuCreateText("请输入Word：", ".docx");
            },
            handleContextMenuCreateExcel() {
                this.handleContextMenuCreateText("请输入Excel：", ".xlsx");
            },
            handleContextMenuCreatePpt() {
                this.handleContextMenuCreateText("请输入Ppt：", ".pptx");
            },
            /**
             * 检查是否是vuepress项目
             * @param data 当前节点数据
             * @returns {boolean}
             */
            checkIfVb(data) {
                if (data && data.isDir) {
                    if (data.expand && data.children && data.children.length > 0) {
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
            selectChange(selectedList,func) {
                if (selectedList.length == 0) {
                    this.routePush({}, '/default', "空白预览")
                    return;
                }
                const node = selectedList[selectedList.length - 1]
                if (node) {
                    //取消选择
                    for(let sn of this.$refs.tree.getSelectedNodes()){
                        this.$set(sn, 'selected', false)
                    }
                    this.$store.commit("setSelecedNode", node)
                    var vueThis = this;
                    this.$set(node, 'selected', true)
                    if (node.isDir) {
                        vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?fileDir=" + node.dirPath + "&fileName=" + node.fileName + "&root=" + node.root).then((response) => {
                            node.children = response.data.data //挂载子节点
                            node.expand = true    //展开子节点
                            func && typeof(func)=="function" && func();
                        })
                    } else {
                        let mapping = this.$globalConfig.editorMapping
                        for (let key in mapping) {
                            let re;
                            eval("re=/^.+(" + key + ")$/")
                            if (re.test(node.title)) {
                                this.routePush(node, ...mapping[key])
                                return;
                            }
                        }
                    }
                    //没有push直接跳转到白板页面
                    this.routePush({}, '/default', "空白预览")
                }
            },
            preventDefault() {
                this.$(document).on({
                    dragleave: function (e) {      //拖离
                        e.preventDefault();
                    },
                    drop: function (e) {           //拖后放
                        e.preventDefault();
                    },
                    dragenter: function (e) {      //拖进
                        e.preventDefault();
                    },
                    dragover: function (e) {       //拖来拖去
                        e.preventDefault();
                    }
                });
            },
            async uploadEntry(parentDir, entry) {
                let name = entry.name;
                let _this = this;
                if (entry.isFile) {
                    entry.file(async function (file) {
                        const param = new FormData();
                        param.append('myfile', file)
                        param.append('fileDir', parentDir)
                        await _this.$axios.post(_this.$globalConfig.goServer + "/file/upload", param).then(res => {

                        })
                    })
                } else {
                    //服务器创建目录
                    await this.$axios.post(this.$globalConfig.goServer + "file/mkdir", {
                        fileDir: parentDir,
                        fileName: name
                    }).then((response) => {
                    });
                    let dirReader = entry.createReader()
                    dirReader.readEntries(async function (entries) {
                        for (let centry of entries) {
                            _this.uploadEntry(parentDir + "/" + name, centry)
                        }
                    })
                }
            },
            dropUpload(e) {
                e.preventDefault(); //取消默认浏览器拖拽效果
                let selectNode = this.$store.getters.getSelectedNode
                var fileDir = selectNode.dirPath + "/" + selectNode.title;
                let _this = this;
                if (!selectNode) {
                    this.$message.error("请选择上传的目录")
                    return;
                }
                let fileList = e.dataTransfer.files; //获取文件对象
                //检测是否是拖拽文件到页面的操作
                if (fileList.length == 0) {
                    return false;
                }
                for (var i = 0; i < fileList.length; i++) {
                    let item = e.dataTransfer.items[i]
                    let entry = item.webkitGetAsEntry()
                    this.uploadEntry(fileDir, entry)
                    _this.selectChange([selectNode])
                }

            },
            initWorkspace(workspace,func){

            },
            async initData() {
                let _this = this;
                let selectedNode = _this.$store.getters.getSelectedNode
                if (selectedNode != null) {
                    let dirPath = selectedNode.dirPath
                    let fileName = selectedNode.fileName
                    let root = this.$refs.tree.data[0]
                    let curParent = root;
                    if (dirPath && fileName) {
                        let newDirPath = dirPath.replace(/[/\\|/|\\]+/g, "\\")
                        if (newDirPath == "\\") {
                            for (let [index, tnode] of new Map(curParent.children.map((tnode, i) => [i, tnode]))) {
                                if (fileName == tnode.title) {
                                    _this.selectChange([tnode])
                                    break;
                                }
                            }
                            return;
                        }
                        // newDirPath=\abc\uuu  fileName=aa.md
                        let dirPathSplit = newDirPath.split("\\");
                        //展开所有目录，注意只是展开到\abc\uuu
                        for (let j = 0; j < dirPathSplit.length; j++) {
                            let curDirPath = dirPathSplit[j]
                            if (curDirPath != null && curDirPath != "") {
                                for (let [index, node] of new Map(curParent.children.map((node, i) => [i, node]))) {
                                    if (curDirPath == node.title) {
                                        if (node.isDir) {
                                            await _this.loadSubNode(node,curParent,(replateNode,rcurParent,childrenData)=>{
                                                replateNode.children = childrenData //挂载子节点
                                                replateNode.expand = true    //展开子节点
                                                _this.$set(replateNode, 'expand', true)
                                                let callback=()=>{
                                                    let funParent=rcurParent;
                                                    let funFileName=fileName;
                                                    //选择目录下的文件
                                                    for (let [index, tnode] of new Map(funParent.children.map((tnode, i) => [i, tnode]))) {
                                                        if (fileName == tnode.title) {
                                                            _this.selectChange([tnode])
                                                            break;
                                                        }
                                                    }
                                                }
                                                _this.$nextTick(callback)
                                            })
                                        }
                                        curParent = node
                                        break;
                                    }
                                }
                            }
                        }
                        //选中对应curParent相应的文件或者目录
                        for (let [index, tnode] of new Map(curParent.children.map((tnode, i) => [i, tnode]))) {
                            if (fileName == tnode.title) {
                                _this.selectChange([tnode])
                                break;
                            }
                        }


                    }
                }
            },
            listRoot(title,func) {
                let vueThis = this;
                vueThis.$axios.get(this.$globalConfig.goServer + "home/listSub?root=true&fileDir=/").then((response) => {
                    var resultData = response.data
                    vueThis.data5 = [
                        {
                            title: title,
                            fileName: "",
                            expand: true,
                            dirPath: '/',
                            contextmenu: vueThis.$store.getters.currentWorkspace == 0 ? false : true,
                            isDir: true,
                            root: true,
                            children: resultData.data
                        }];
                    vueThis.$nextTick(() => {
                        vueThis.initData()
                    })
                    func && func()

                })
            }
        }
        ,
        mounted() {
            let title = this.$store.getters.currentWorkspace == "0" ? "公共文档库" : "个人文档库";
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
</style>
