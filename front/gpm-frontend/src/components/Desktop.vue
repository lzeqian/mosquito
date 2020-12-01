<style scoped>
    .desktop {
        position: relative;
        height: 100%;
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        align-content: flex-start;
        background: url("../../public/img/windowbg.jpg") no-repeat;
        background-size: 100% 100%;
    }

    .icon {
        width: 5em;
        height: 5em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }

    .single {
        position: absolute;
        left: 50px;
        top: 50px;
        padding-left: 10px;
        padding-right: 10px;
    }
    .person {
        position: absolute;
        left: 50px;
        top: 180px;
        padding-left: 10px;
        padding-right: 10px;
    }
    .textDiv{
        padding-left: 5px;
        color: white;
        user-select:none;
    }
    .contextmenu{
        position: absolute;
        left: 50px;
        top: 50px;
        background-color: rgb(242,242,242);
        user-select:none;
        padding-right: 40px;
        padding-left: 10px;
    }
    .fileSystemContextmenu{
        position: absolute;
        z-index: 500;
        left: 50px;
        top: 50px;
        background-color: rgb(242,242,242);
        user-select:none;
        padding-right: 40px;
        padding-left: 10px;
    }
    .contextmenu div{
        margin-bottom: 5px;
    }
    /deep/ .ivu-modal-header{
        display: none;
    }
    /deep/ .ivu-modal-content{
        border-radius: 2px;
        height: 90%;
    }
    /deep/ .ivu-modal{
        height: 90%;
    }
</style>
<template>

    <div class="desktop">
        <div class="single"  @mouseover="mouseOver('.single','rgb(60,95,130)')"
             @mouseleave="mouseLeave('.single')">
            <a @dblclick="dirClick" >
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-wenjianjia"></use>
                </svg>
                <div class="textDiv">
                公共文档
                </div>
            </a>
        </div>
        <div class="person"  @mouseover="mouseOver('.person','rgb(60,95,130)')"
             @mouseleave="mouseLeave('.person')">
            <a @dblclick="dirPersonClick" >
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-wenjianjia"></use>
                </svg>
                <div class="textDiv">
                    个人文档
                </div>
            </a>
        </div>
        <!--桌面右键菜单-->
        <div class="contextmenu" v-if="contextmenuVisible">
            <div class="switch"  @mouseover="mouseOver('.switch','rgb(217,217,217)')" @click="switchToDirectory"
                 @mouseleave="mouseLeave('.switch')">
                切换目录树方式(S)
            </div>
            <hr/>
            <div class="refreshCurren"  @mouseover="mouseOver('.refreshCurren','rgb(217,217,217)')"
                 @mouseleave="mouseLeave('.refreshCurren')">
                刷新
            </div>
        </div>

        <Modal width="70%"
                v-model="showMadal"
                title="文件系统"
                :footer-hide="true"
                >
           <FileSystem></FileSystem>
            <!--文件系统右键菜单-->
            <div class="fileSystemContextmenu" v-if="fileSystemContextmenuVisible">
                <div class="fileSystemUploadFile" @mouseover="mouseOver('.fileSystemUploadFile','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemUploadFile')">
                    <Upload :style="{width:'100%'}" :show-upload-list="false" action=""
                            ref="upload"

                    >上传文件(UF)
                    </Upload>
                </div>
                <div class="fileSystemRename" @mouseover="mouseOver('.fileSystemRename','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemRename')">
                    重命名(RN)
                </div>
                <div class="fileSystemDelete" @mouseover="mouseOver('.fileSystemDelete','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemDelete')">
                    删除文件(DE)
                </div>
                <hr/>

                <div class="fileSystemCreateMd" @mouseover="mouseOver('.fileSystemCreateMd','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemCreateMd')">
                    <font color="red"> 新建md(MD)</font>
                </div>
                <div class="fileSystemCreateFlow" @mouseover="mouseOver('.fileSystemCreateFlow','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemCreateFlow')">
                    <font color="red"> 新建flow(FW)</font>
                </div>
                <div class="fileSystemCreateSnow" @mouseover="mouseOver('.fileSystemCreateSnow','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemCreateSnow')">
                    <font color="red">  新建思维导图(NN)</font>
                </div>
                <div class="fileSystemCreateFile" @mouseover="mouseOver('.fileSystemCreateFile','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemCreateFile')">
                    <font color="red">  新建文件(NF)</font>
                </div>
                <hr/>
                <div class="fileSystemCreateVp" @mouseover="mouseOver('.fileSystemCreateVp','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemCreateVp')">
                    <font color="green">新建vuepress(VP)</font>
                </div>
                <div class="fileSystemBuildVp" @mouseover="mouseOver('.fileSystemBuildVp','rgb(217,217,217)')"
                     @mouseleave="mouseLeave('.fileSystemBuildVp')">
                    <font color="green"> 构建vuepress(BP)</font>
                </div>
            </div>
        </Modal>
    </div>


</template>
<script>
    import FileSystem from "./desktop/FileSystem";
    export default {
        data() {
            return {
                contextmenuVisible:false,
                fileSystemContextmenuVisible:false,
                showMadal:false,
                registerContextMenu:false
            }
        },
        props: {},
        computed: {},
        components: {FileSystem},
        methods: {
            switchToDirectory(){
                this.$store.commit("updateDirTree","tree")
                this.routePush({},'/blank',"空白预览")
            },
            mouseOver(className,bgColor){
                // this.$(className).css('background-color', 'rgb(60,95,130)');
                this.$(className).css('background-color', bgColor);
            },
            mouseLeave(className){
                this.$(className).css('background-color', '');
            },
            selectClick(){

            },
            dirClick(){
                this.$store.commit("updateWorkspace","0")
                this.showMadal=true
                if(!this.registerContextMenu) {
                    this.contextMenu(".ivu-modal-content", ".fileSystemContextmenu")
                    this.registerContextMenu=true
                }
            },
            dirPersonClick(){
                this.$store.commit("updateWorkspace","1")
                this.showMadal=true
                if(!this.registerContextMenu) {
                    this.contextMenu(".ivu-modal-content", ".fileSystemContextmenu")
                    this.registerContextMenu=true
                }
            },
            dragDiv(className){
                let helperdialogwrapper =$(className);
                let x = 0;
                let y = 0;
                let l = 0;
                let t = 0;
                window.isDown = false;
                window.curClassName=null;
                //鼠标按下事件
                $(className).bind("mousedown",function(e) {
                    //获取x坐标和y坐标
                    x = e.clientX;
                    y = e.clientY;

                    //获取左部和顶部的偏移量
                    l = helperdialogwrapper.offset().left;
                    t = helperdialogwrapper.offset().top;
                    //开关打开
                    window.isDown = true;
                    window.curClassName=className
                    //设置样式
                });
                //鼠标移动
                window.onmousemove = function(e) {
                    if (window.isDown == false) {
                        return;
                    }
                    //获取x和y
                    let nx = e.clientX;
                    let ny = e.clientY;
                    //计算移动后的左偏移量和顶部的偏移量
                    let nl = parseInt(l)+(parseInt(nx) -parseInt(x));
                    let nt = parseInt(t)+(parseInt(ny) -parseInt(y));
                    let sss=parseInt(nx) -parseInt(x);
                    let lll=parseInt(ny) -parseInt(y);
                    //这里设置offset而不是css，因为获取时是根据offset获取的偏移量
                    $(window.curClassName).offset({top:nt,left:nl});
                }
                //鼠标抬起事件
                $(className).bind("mouseup",function() {
                        //开关关闭
                    window.isDown = false;
                    }
                );
            },
            getOffset(parentDivClassName,target){
                let x=0,y=0
                if(target!=null) {
                    let curClassName = ("." + target.getAttribute("class"))
                    if (curClassName != parentDivClassName) {
                        let offsetObject = this.getOffset(parentDivClassName, ((target.tagName=="use" || target.tagName=="svg")?target.parentNode:target.offsetParent))
                        x += (target.offsetLeft ? target.offsetLeft : 0) + offsetObject.x
                        y += (target.offsetTop ? target.offsetTop : 0) + offsetObject.y
                    }
                }
                return {
                    x:x,
                    y:y
                }
            },
            contextMenu(parentDivClassName,contextMenuDivClassName){
                var _this=this
                $(parentDivClassName).ready(function() {
                    // 鼠标右键事件
                    $(parentDivClassName).contextmenu(function(e) {
                        // 获取窗口尺寸
                        let winWidth = $(parentDivClassName).width();
                        let winHeight = $(parentDivClassName).height();
                        // 鼠标点击位置坐标
                        let offsetObject=_this.getOffset(parentDivClassName,e.target)
                        let mouseX = e.offsetX+offsetObject.x;
                        let mouseY = e.offsetY+offsetObject.y;
                        // ul标签的宽高
                        let menuWidth = $(contextMenuDivClassName).width();
                        let menuHeight = $(contextMenuDivClassName).height();
                        // 最小边缘margin(具体窗口边缘最小的距离)
                        let minEdgeMargin = 10;
                        // 以下判断用于检测ul标签出现的地方是否超出窗口范围
                        // 第一种情况：右下角超出窗口
                        let menuLeft,menuTop=0
                        if(mouseX + menuWidth + minEdgeMargin >= winWidth &&
                            mouseY + menuHeight + minEdgeMargin >= winHeight) {
                            menuLeft = mouseX - menuWidth - minEdgeMargin + "px";
                            menuTop = mouseY - menuHeight - minEdgeMargin + "px";
                        }
                        // 第二种情况：右边超出窗口
                        else if(mouseX + menuWidth + minEdgeMargin >= winWidth) {
                            menuLeft = mouseX - menuWidth - minEdgeMargin + "px";
                            menuTop = mouseY + minEdgeMargin + "px";
                        }
                        // 第三种情况：下边超出窗口
                        else if(mouseY + menuHeight + minEdgeMargin >= winHeight) {
                            menuLeft = mouseX + minEdgeMargin + "px";
                            menuTop = mouseY - menuHeight - minEdgeMargin + "px";
                        }
                        // 其他情况：未超出窗口
                        else {
                            menuLeft = mouseX + minEdgeMargin + "px";
                            menuTop = mouseY + minEdgeMargin + "px";
                        };
                       _this[contextMenuDivClassName.substring(1)+"Visible"]=true;
                        _this.$nextTick(()=>{
                            // ul菜单出现
                            $(contextMenuDivClassName).css({
                                "left": menuLeft,
                                "top": menuTop
                            }).show();
                        })
                        // 阻止浏览器默认的右键菜单事件
                        return false;
                    });
                    // 点击之后，右键菜单隐藏
                    $(parentDivClassName).click(function() {
                        $(contextMenuDivClassName).hide();
                        _this[contextMenuDivClassName.substring(1)+"Visible"]=false;
                    });
                });
            }
        },
        mounted() {
            this.dragDiv(".single")
            this.dragDiv(".person")
            this.contextMenu(".desktop",".contextmenu")
        }
    }
</script>
