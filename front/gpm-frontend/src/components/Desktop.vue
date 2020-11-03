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
            <a @dblclick="dirClick">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-wenjianjia"></use>
                </svg>
                <div class="textDiv">
                共享目录
                </div>
            </a>
        </div>
        <div class="contextmenu" v-if="currentVisible">
            <div class="switch"  @mouseover="mouseOver('.switch','rgb(217,217,217)')"
                  @mouseleave="mouseLeave('.switch')" @click="switchToDirectory">
                切换目录树方式(S)
            </div>
            <hr/>
            <div class="uploadFile"  @mouseover="mouseOver('.uploadFile','rgb(217,217,217)')"
                 @mouseleave="mouseLeave('.uploadFile')">
                上传文件(U)
            </div>
        </div>
        <Modal width="70%"
                v-model="showMadal"
                title="文件系统"
                :footer-hide="true"
                >
           <FileSystem></FileSystem>
        </Modal>
    </div>


</template>
<script>
    import FileSystem from "./desktop/FileSystem";
    export default {
        data() {
            return {
                currentVisible:false,
                showMadal:false
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
            dirClick(){
                this.showMadal=true
            },
            dragDiv(className){
                let helperdialogwrapper =$(className);
                let x = 0;
                let y = 0;
                let l = 0;
                let t = 0;
                let isDown = false;
                //鼠标按下事件
                $(className).bind("mousedown",function(e) {
                    //获取x坐标和y坐标
                    x = e.clientX;
                    y = e.clientY;

                    //获取左部和顶部的偏移量
                    l = helperdialogwrapper.offset().left;
                    t = helperdialogwrapper.offset().top;
                    //开关打开
                    isDown = true;
                    //设置样式
                });
                //鼠标移动
                window.onmousemove = function(e) {
                    if (isDown == false) {
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
                    $(className).offset({top:nt,left:nl});
                }
                //鼠标抬起事件
                $(className).bind("mouseup",function() {
                        //开关关闭
                        isDown = false;
                    }
                );
            },
            contextMenu(){
                var _this=this
                $(document).ready(function() {
                    // 鼠标右键事件
                    $(document).contextmenu(function(e) {
                        // 获取窗口尺寸
                        let winWidth = $(document).width();
                        let winHeight = $(document).height();
                        // 鼠标点击位置坐标
                        let mouseX = e.pageX;
                        let mouseY = e.pageY;
                        // ul标签的宽高
                        let menuWidth = $(".contextmenu").width();
                        let menuHeight = $(".contextmenu").height();
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
                        _this.currentVisible=true;
                        _this.$nextTick(()=>{
                            // ul菜单出现
                            $(".contextmenu").css({
                                "left": menuLeft,
                                "top": menuTop
                            }).show();
                        })
                        // 阻止浏览器默认的右键菜单事件
                        return false;
                    });
                    // 点击之后，右键菜单隐藏
                    $(document).click(function() {
                        $(".contextmenu").hide();
                        _this.currentVisible=false;
                    });
                });
            }
        },
        mounted() {
            this.dragDiv(".single")
            this.contextMenu()
        }
    }
</script>
