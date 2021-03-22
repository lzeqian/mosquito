<style scoped>

</style>
<template>

    <div style="height: 100%" @keydown="editorKeyDownSave">
        <!-- 顶部导航栏 -->
        <div class="headers">
            <el-menu mode="horizontal" @select="onMenu" background-color="#f8f8f8">
                <el-menu-item class="logo">

                    <img src="/img/log.png"/>

                </el-menu-item>
                <el-submenu index="file">
                    <template slot="title">文件</template>
                    <!--                    <el-menu-item index="new">新建文件</el-menu-item>-->
                    <!--                    <el-menu-item index="open">打开本地文件（新建）</el-menu-item>-->
                    <!--                    <el-menu-item index="replace">导入本地文件...</el-menu-item>-->

                    <el-menu-item index="save">保存</el-menu-item>
                    <el-menu-item index="saveNativePng">保存为png</el-menu-item>
                    <el-menu-item class="separator"></el-menu-item>
                    <el-menu-item index="downloadSrc">下载原始文件</el-menu-item>
                    <el-menu-item index="savePng">下载为PNG</el-menu-item>
                </el-submenu>
                <el-submenu index="edit">
                    <template slot="title">编辑</template>
                    <el-menu-item index="undo">撤消</el-menu-item>
                    <el-menu-item index="redo">重做</el-menu-item>
                    <el-menu-item class="separator"></el-menu-item>
                    <el-menu-item index="copy">复制</el-menu-item>
                    <el-menu-item index="cut">剪切</el-menu-item>
                    <el-menu-item index="parse">粘贴</el-menu-item>
                </el-submenu>
                <el-submenu index="help">
                    <template slot="title">帮助</template>

                </el-submenu>
            </el-menu>
            <el-menu mode="horizontal" class="full" background-color="#f8f8f8"></el-menu>
            <el-menu mode="horizontal" background-color="#f8f8f8">
                <el-menu-item>视图：{{scale}}%</el-menu-item>
                <el-submenu index="state" title="默认连线类型">
                    <template slot="title">
                        <i :class="`iconfont icon-${lineName}`"></i>
                    </template>
                    <el-menu-item
                            v-for="(item, index) in lineNames"
                            :key="index"
                            :index="`line-${item}`"
                            @click="onState('lineName', item)"
                    >
                        <i :class="`iconfont icon-${item}`"></i>
                    </el-menu-item>
                </el-submenu>
            </el-menu>
            <el-menu mode="horizontal" background-color="#f8f8f8">
                <el-submenu index="state" title="默认起点箭头">
                    <template slot="title">
                        <i :class="`iconfont icon-from-${fromArrowType}`"></i>
                    </template>
                    <el-menu-item
                            v-for="(item, index) in arrowTypes"
                            :key="index"
                            :index="`fromArrow-${item}`"
                            @click="onState('fromArrowType', item)"
                    >
                        <i :class="`iconfont icon-from-${item}`"></i>
                    </el-menu-item>
                </el-submenu>
            </el-menu>
            <el-menu mode="horizontal" background-color="#f8f8f8">
                <el-submenu index="state" title="默认终点箭头">
                    <template slot="title">
                        <i :class="`iconfont icon-to-${toArrowType}`"></i>
                    </template>
                    <el-menu-item
                            v-for="(item, index) in arrowTypes"
                            :key="index"
                            :index="`toArrow-${item}`"
                            @click="onState('toArrowType', item)"
                    >
                        <i :class="`iconfont icon-to-${item}`"></i>
                    </el-menu-item>
                </el-submenu>
            </el-menu>
            <el-menu mode="horizontal" background-color="#f8f8f8">
                <el-submenu index="user" v-if="user">
                    <template slot="title">
                        <el-avatar
                                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                                :size="24"
                        ></el-avatar>
                        {{user.username}}
                    </template>
                    <el-menu-item @click="onSignOut">退出</el-menu-item>
                </el-submenu>

            </el-menu>
            <el-menu mode="horizontal" background-color="#f8f8f8">
                <Button :icon="scalePropIcon" size="small" type="text" @click.stop="scaleProp" style="height: 100%"></Button>
            </el-menu>
        </div>

        <!-- body部分 -->
        <div class="body"  style="height:95%">
            <div class="page">
                <div class="tools">
                    <div v-for="(item, index) in tools" :key="index">
                        <div class="title">{{ item.group }}</div>
                        <div class="buttons">
                            <a
                                    v-for="(btn, i) in item.children"
                                    :key="i"
                                    :title="btn.name"
                                    :draggable="btn.data"
                                    @dragstart="onDrag($event, btn)"
                            >
                                <i :class="`iconfont ${btn.icon}`"></i>
                            </a>
                        </div>
                    </div>
                </div>
                <div id="topology-canvas" class="full" @contextmenu="onContextMenu($event)"></div>
                <div ref="props" class="props" :style="props.expand ? 'overflow: visible' : ''">
                    <CanvasProps :props.sync="props" @change="onUpdateProps"></CanvasProps>
                </div>
                <div class="context-menu" v-if="contextmenu.left" :style="this.contextmenu">
                    <CanvasContextMenu :canvas="canvas" :props.sync="props"></CanvasContextMenu>
                </div>
            </div>
        </div>
    </div>

</template>
<script>
    import Element from 'element-ui'
    import 'element-ui/lib/theme-chalk/index.css'
    import '@/assets/css/font_1113798_0532l8oa6jqp.css'
    import '@/assets/css/font_1331132_5lvbai88wkb.css'
    import locale from 'element-ui/lib/locale/lang/zh-CN'
    import C2S from '@/assets/canvas2svg.js'

    import Vue from 'vue'

    Vue.use(Element, {locale})
    import * as FileSaver from 'file-saver';
    import {Topology, Options, registerNode, Node, Line} from '@topology/core';
    import {Tools, canvasRegister} from '@/services/canvas';
    import CanvasProps from '@/components/flow/CanvasProps';
    import CanvasContextMenu from '@/components/flow/CanvasContextMenu';
    import '@/assets/css/base.scss'

    let canvas;
    const canvasOptions = {
        rotateCursor: '@/assets/rotate.cur'
    };
    export default {
        data() {
            return {
                scalePropIcon:'md-arrow-dropright',
                tools: Tools,
                canvas:null,
                props: {
                    node: null,
                    line: null,
                    nodes: null,
                    multi: false,
                    expand: false,
                    locked: false
                },
                contextmenu: {
                    left: null,
                    top: null,
                    bottom: null
                },
                about: false,
                license: false,
                joinin: false,
                lineNames: ['curve', 'polyline', 'line'],
                arrowTypes: [
                    '',
                    'triangleSolid',
                    'triangle',
                    'diamondSolid',
                    'diamond',
                    'circleSolid',
                    'circle',
                    'line',
                    'lineUp',
                    'lineDown'
                ],
                user: null
            };
        },
        components: {
            CanvasProps,
            CanvasContextMenu
        },
        computed: {
            event() {
                return this.$store.state.event;
            },
            scale() {
                return Math.floor(this.$store.state.data.scale * 100)
            },
            lineName() {
                return this.$store.state.data.lineName
            },
            fromArrowType() {
                return this.$store.state.data.fromArrowType
            },
            toArrowType() {
                return this.$store.state.data.toArrowType
            },
            error() {
                return this.$store.state.error
            }
        },
        watch: {
            event(curVal) {
                if (this['handle_' + curVal.name]) {
                    this['handle_' + curVal.name](curVal.data);
                }
            },
            $route(val) {
                // this.open();
            },
            error(curVal) {
                this.$notify({
                    title: '错误',
                    type: 'error',
                    message: curVal.text
                })
            }
        },
        created() {
            canvasRegister();
            document.onclick = event => {
                this.contextmenu = {
                    left: null,
                    top: null,
                    bottom: null
                };
            };
        },
        mounted() {
            canvasOptions.on = this.onMessage;
            canvas = new Topology('topology-canvas', canvasOptions);
            this.canvas=canvas
            document.documentElement.style.fontSize="100px";
            // this.open();
        },
        methods: {
            editorKeyDownSave(e) {
                let _this=this;
                let currenKey = e.keyCode || e.which || e.charCode;
                if (currenKey == 83 && e.ctrlKey) {
                    e.preventDefault()
                    _this.saveEditorContent({
                        value: JSON.stringify(canvas.data),
                    })
                }
            },
            scaleProp(){
                if(this.scalePropIcon=="md-arrow-dropright"){
                    this.$refs.props.style.width=0;
                }else{
                    this.$refs.props.style.width="2.4rem";
                }
                this.scalePropIcon=(this.scalePropIcon=="md-arrow-dropright"?"md-arrow-dropleft":"md-arrow-dropright")
            },
            onMenu(key, keyPath) {
                if (!key || key.indexOf('/') === 0) {
                    return
                }

                switch (key) {
                    case 'new':
                        this.$router.push('/workspace')
                        break
                    case 'open':
                        this.$router.push('/workspace')
                        setTimeout(() => {
                            this.$store.commit('emit', {
                                name: key
                            })
                        }, 100)
                        break
                    case 'about':
                    case 'about2':
                        this.about = true
                        break
                    case 'license':
                        this.license = true
                        break
                    case 'joinin':
                        this.joinin = true
                        break
                    default:
                        this.$store.commit('emit', {
                            name: key
                        })
                        break
                }
            },
            onState(key, value) {
                this.$store.commit('emit', {
                    name: 'state',
                    data: {
                        key,
                        value
                    }
                })
            },
            onLogin() {
                if (process.client) {
                    location.href = `https://account.le5le.com?cb=${encodeURIComponent(
                        location.href
                    )}`
                }
            },
            onSignOut() {
                this.$cookies.remove('token')
                this.user = null
            },
            initData(data){
                canvas.open(data);
            },
            // async open() {
            //     this.loadEditorContent((vueThis,data)=>{
            //
            //     })
            // },
            onDrag(event, node) {
                event.dataTransfer.setData('Text', JSON.stringify(node.data));
            },
            onMessage(event, data) {
                // console.log('onMessage', event, data);
                // 右侧输入框编辑状态时点击编辑区域其他元素，onMessage执行后才执行onUpdateProps方法，通过setTimeout让onUpdateProps先执行
                setTimeout(() => {
                    switch (event) {
                        case 'node':
                        case 'addNode':
                            this.props = {
                                node: data,
                                line: null,
                                multi: false,
                                expand: this.props.expand,
                                nodes: null,
                                locked: data.locked
                            };
                            break;
                        case 'line':
                        case 'addLine':
                            this.props = {
                                node: null,
                                line: data,
                                multi: false,
                                nodes: null,
                                locked: data.locked
                            };
                            break;
                        case 'multi':
                            this.props = {
                                node: null,
                                line: null,
                                multi: true,
                                nodes: data.length > 1 ? data : null,
                                locked: this.getLocked({nodes: data})
                            };
                            break;
                        case 'space':
                            this.props = {
                                node: null,
                                line: null,
                                multi: false,
                                nodes: null,
                                locked: false
                            };
                            break;
                        case 'moveOut':
                            break;
                        case 'moveNodes':
                        case 'resizeNodes':
                            if (data.length > 1) {
                                this.props = {
                                    node: null,
                                    line: null,
                                    multi: true,
                                    nodes: data,
                                    locked: this.getLocked({nodes: data})
                                };
                            } else {
                                this.props = {
                                    node: data[0],
                                    line: null,
                                    multi: false,
                                    nodes: null,
                                    locked: false
                                };
                            }
                            break;
                        case 'resize':
                        case 'scale':
                        case 'locked':
                            if (canvas && canvas.data) {
                                this.$store.commit('data', {
                                    scale: canvas.data.scale || 1,
                                    lineName: canvas.data.lineName,
                                    fromArrowType: canvas.data.fromArrowType,
                                    toArrowType: canvas.data.toArrowType,
                                    fromArrowlockedType: canvas.data.locked
                                });
                            }
                            break;
                    }
                }, 0);
            },
            getLocked(data) {
                let locked = true;
                if (data.nodes && data.nodes.length) {
                    for (const item of data.nodes) {
                        if (!item.locked) {
                            locked = false;
                            break;
                        }
                    }
                }
                if (locked && data.lines) {
                    for (const item of data.lines) {
                        if (!item.locked) {
                            locked = false;
                            break;
                        }
                    }
                }
                return locked;
            },
            onUpdateProps(node) {
                // 如果是node属性改变，需要传入node，重新计算node相关属性值
                // 如果是line属性改变，无需传参
                canvas.updateProps(node);
            },
            handle_new(data) {
                canvas.open();
            },
            handle_open(data) {
                this.handle_replace(data);
            },
            handle_replace(data) {
                const input = document.createElement('input');
                input.type = 'file';
                input.onchange = event => {
                    const elem = event.srcElement || event.target;
                    if (elem.files && elem.files[0]) {
                        const name = elem.files[0].name.replace('.json', '');
                        const reader = new FileReader();
                        reader.onload = e => {
                            const text = e.target.result + '';
                            try {
                                const data = JSON.parse(text);
                                canvas.open(data);
                            } catch (e) {
                                return false;
                            }
                        };
                        reader.readAsText(elem.files[0]);
                    }
                };
                input.click();
            },
            handle_downloadSrc(data) {
                let _this=this;
                let selectedNode=_this.$store.getters.getSelectedNode
                let dirName=selectedNode.dirPath;
                let fileName=selectedNode.fileName;
                FileSaver.saveAs(
                    new Blob([JSON.stringify(canvas.data)], {
                        type: 'text/plain;charset=utf-8'
                    }),
                    fileName
                );
            },
            handle_save(data) {
                this.saveEditorContent({
                    value: JSON.stringify(canvas.data),
                })
            },
            handle_savePng(data) {
                let _this=this;
                let selectedNode=_this.$store.getters.getSelectedNode
                let fileName=selectedNode.fileName;
                canvas.saveAsImage(fileName.split(".flow")[0]+'.png');
            },
            handle_saveNativePng(data) {
                let _this=this;
                let selectedNode=_this.$store.getters.getSelectedNode
                let fileName=selectedNode.fileName;
                canvas.toImage(0, 'image/png',(blobData)=>{
                    let formData = new FormData();
                    fileName=fileName.substring(0,fileName.indexOf("."))+".png"
                    formData.append('myfile',blobData,fileName)
                    _this.$axios({
                        url: _this.$globalConfig.goServer +'file/upload?fileDir='+selectedNode.dirPath,
                        method: 'post',
                        processData:  {"Content-Type":"multipart/form-data",},
                        data: formData
                    }).then(()=>{
                        if(_this.$root.$children[0].$refs.home && _this.$root.$children[0].$refs.home.$refs.leftTree){
                            let leftTree=_this.$root.$children[0].$refs.home.$refs.leftTree;
                            let {index, parentNode} = leftTree.getParent(leftTree.$refs.tree.data[0], selectedNode)
                            leftTree.selectChange([parentNode])
                        }

                    })
                })
            },
            handle_undo(data) {
                canvas.undo();
            },
            handle_redo(data) {
                canvas.redo();
            },
            handle_copy(data) {
                canvas.copy();
            },
            handle_cut(data) {
                canvas.cut();
            },
            handle_parse(data) {
                canvas.parse();
            },
            handle_state(data) {
                canvas.data[data.key] = data.value;
                this.$store.commit('data', {
                    scale: canvas.data.scale || 1,
                    lineName: canvas.data.lineName,
                    fromArrowType: canvas.data.fromArrowType,
                    toArrowType: canvas.data.toArrowType,
                    fromArrowlockedType: canvas.data.locked
                });
            },
            onContextMenu(event) {
                event.preventDefault();
                event.stopPropagation();
                if (event.clientY + 360 < document.body.clientHeight) {
                    this.contextmenu = {
                        left: event.clientX + 'px',
                        top: event.clientY + 'px'
                    };
                } else {
                    this.contextmenu = {
                        left: event.clientX + 'px',
                        bottom: document.body.clientHeight - event.clientY + 'px'
                    };
                }
            }
        },
        destroyed() {
            canvas.destroy();
        }
    }
</script>

<style lang="scss">
    .page {
        display: flex;
        width: 100%;
        height: 100%;

        .tools {
            flex-shrink: 0;
            width: 1.75rem;
            background-color: #f8f8f8;
            border-right: 1px solid #d9d9d9;
            overflow-y: auto;

            .title {
                color: #0d1a26;
                font-weight: 600;
                font-size: 0.12rem;
                line-height: 1;
                padding: 0.05rem 0.1rem;
                margin-top: 0.08rem;
                border-bottom: 1px solid #ddd;

                &:first-child {
                    border-top: none;
                }
            }

            .buttons {
                padding: 0.1rem 0;

                a {
                    display: inline-block;
                    color: #314659;
                    line-height: 1;
                    width: 0.4rem;
                    height: 0.4rem;
                    text-align: center;
                    text-decoration: none !important;
                    cursor: pointer;

                    .iconfont {
                        font-size: 0.24rem;
                    }

                    &:hover {
                        color: #1890ff;
                    }
                }
            }
        }

        .full {
            flex: 1;
            width: initial;
            position: relative;
            overflow: auto;
            background: #fff;
        }

        .props {
            flex-shrink: 0;
            width: 2.4rem;
            padding: 0.1rem 0;
            background-color: #f8f8f8;
            border-left: 1px solid #d9d9d9;
            overflow-y: auto;
            position: relative;
        }

        .context-menu {
            position: fixed;
            z-index: 10;
        }
    }

    html {
        font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
        Roboto, 'Helvetica Neue', Arial, sans-serif;
        font-size: 100px;
        word-spacing: 1px;
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%;
        -moz-osx-font-smoothing: grayscale;
        -webkit-font-smoothing: antialiased;
        box-sizing: border-box;
    }

    /*body {*/
    /*    font-size: 0.12rem;*/
    /*    line-height: 2;*/
    /*}*/

    *,
    *:before,
    *:after {
        box-sizing: border-box;
        margin: 0;
    }

    .headers {
        display: flex;
        background-color: #f8f8f8;
        font-size: 0.13rem;
        height: 0.4rem;

        .full {
            flex: 1;
        }

        .logo {
            img {
                height: 0.22rem;
                position: relative;
                top: -0.04rem;
            }
        }
    }

    .el-menu-item,
    .el-submenu__title {
        color: #314659 !important;
        font-size: 0.13rem;
        padding: 0 0.1rem;
        height: 0.39rem !important;
        line-height: 0.39rem !important;

        &.separator {
            border-top: 1px solid #e8e8e8;
            margin: 0 0.1rem;
            height: 0.01rem !important;
            line-height: 0.01rem !important;
        }

        a {
            text-decoration: none;
            color: #314659;
        }

        .iconfont {
            color: #314659;
            font-size: 0.26rem;
        }
    }

    .body {
        height: calc(100vh - 0.4rem);
    }
    .icon-mind{
        content: "\e629";
    }
</style>

