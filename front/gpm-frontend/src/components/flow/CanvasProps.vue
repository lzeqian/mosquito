<template>
    <div>
        <!-- 选中为空 -->
        <div v-if="!props.node && !props.line && !props.multi">
            <div class="title">欢迎使用饺子编辑器！</div>

        </div>
        <!-- 选中节点 -->
        <div v-if="props.node">
            <div class="title">位置和大小</div>
            <div class="items">
                <div class="flex grid">
                    <div>X（px）</div>
                    <div class="ml5">Y（px）</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input-number
                                v-model="props.node.rect.x"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                    <div class="ml5">
                        <el-input-number
                                v-model="props.node.rect.y"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div>宽（px）</div>
                    <div class="ml5">高（px）</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input-number
                                v-model="props.node.rect.width"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                    <div class="ml5">
                        <el-input-number
                                v-model="props.node.rect.height"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                </div>
            </div>
            <div class="items" v-if="props.node.is3D">
                <div class="flex grid">
                    <div>Z（px）</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input-number
                                v-model="props.node.z"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="百分比%对应的小数值">圆角（0 - 1）</div>
                    <div class="ml5">旋转（°）</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input-number
                                v-model="props.node.borderRadius"
                                controls-position="right"
                                @change="onChange"
                                :min="0"
                                :max="1"
                                :step="0.1"
                        ></el-input-number>
                    </div>
                    <div class="ml5">
                        <el-input-number
                                v-model="props.node.rotate"
                                controls-position="right"
                                @change="onChange"
                        ></el-input-number>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-left">内边距-左</div>
                    <div title="padding-right" class="ml5">内边距-右</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.paddingLeft"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                    <div class="ml5">
                        <el-input
                                size="small"
                                v-model="props.node.paddingRight"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top">内边距-上</div>
                    <div title="padding-bottom" class="ml5">内边距-下</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.paddingTop"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                    <div class="ml5">
                        <el-input
                                size="small"
                                v-model="props.node.paddingBottom"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div class="items gray" style="line-height: 1.5">
                内边距：输入数字表示像素；输入%表示百分比
            </div>
            <div class="title"></div>
            <div class="items">
                <div class="flex grid">
                    <div class="custom-data">自定义数据 <i :class="props.expand ? 'el-icon-zoom-out' : 'el-icon-zoom-in'"
                                                      @click="changeExpand" size='small'>{{props.expand ? '缩小' :
                        '放大'}}</i></div>
                </div>
                <div class="flex grid">
                    <div :class="props.expand ? 'expand-data' : ''">
                        <el-input
                                type="textarea"
                                v-model="nodeData"
                                :rows="props.expand ? 15 : 3"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>

            <div class="title">样式</div>
            <div class="items">
                <div class="flex grid">
                    <div>线条样式</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-select v-model="lineStyle" placeholder="请选择" @change="onSelectChange">
                            <el-option
                                    v-for="item in options"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top">线条颜色</div>
                    <div title="padding-bottom" class="ml5">线条宽度(px)</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-color-picker v-model="props.node.strokeStyle" @change="onChange"></el-color-picker>
                    </div>
                    <div class="ml5">
                        <el-input
                                size="small"
                                v-model="props.node.lineWidth"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div>背景</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-select v-model="props.node.bkType" placeholder="请选择"
                                   @change="onChange(props.node.bkType,'bkType')">
                            <el-option
                                    v-for="item in bgOptions"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top" v-if="props.node.bkType=='0' || props.node.bkType==null">背景颜色</div>
                    <div title="padding-bottom" v-if="props.node.bkType!=null && props.node.bkType!='0'" class="ml5">
                        开始颜色
                    </div>
                    <div title="padding-bottom" v-if="props.node.bkType!=null && props.node.bkType!='0'" class="ml5">
                        结束颜色
                    </div>
                </div>
                <div class="flex grid">
                    <div v-if="props.node.bkType=='0' || props.node.bkType==null">
                        <el-color-picker v-model="props.node.fillStyle" @change="onChange"></el-color-picker>
                    </div>
                    <div class="ml5" v-if="props.node.bkType!=null && props.node.bkType!='0'">
                        <el-color-picker v-model="props.node.gradientFromColor" @change="onChange"></el-color-picker>
                    </div>
                    <div class="ml5" v-if="props.node.bkType!=null && props.node.bkType!='0'">
                        <el-color-picker v-model="props.node.gradientToColor" @change="onChange"></el-color-picker>
                    </div>

                </div>
            </div>
            <div class="items" v-if="props.node.bkType!=null && props.node.bkType=='1'">
                <div class="flex grid">
                    <div>渐变角度</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.gradientAngle"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div class="items" v-if="props.node.bkType!=null && props.node.bkType=='2'">
                <div class="flex grid">
                    <div>渐变半径</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.gradientRadius"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>


            <div class="title">文字</div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top">字体</div>
                    <div title="padding-bottom" class="ml5">大小</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.font.fontFamily"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                    <div class="ml5">
                        <el-input
                                size="small"
                                v-model="props.node.font.fontSize"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top">颜色</div>
                    <div title="padding-bottom" class="ml5">背景</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-color-picker v-model="props.node.font.color" @change="onChange"></el-color-picker>
                    </div>
                    <div class="ml5">
                        <el-color-picker v-model="props.node.font.background" @change="onChange"></el-color-picker>
                    </div>
                </div>
            </div>
            <div class="items">
                <div class="flex grid">
                    <div title="padding-top">水平对齐</div>
                    <div title="padding-bottom" class="ml5">垂直对齐</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-select v-model="props.node.font.textAlign" placeholder="请选择"
                                   @change="onChange">
                            <el-option
                                    v-for="item in textAlignOptions"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                    <div class="ml5">
                        <el-select v-model="props.node.font.textBaseline" placeholder="请选择"
                                   @change="onChange">
                            <el-option
                                    v-for="item in baseAlignOptions"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>
            </div>

            <div class="items" v-if="props.node.name!=null">
                <div class="flex grid">
                    <div>图片地址</div>
                </div>
                <div class="flex grid">
                    <div>
                        <el-input
                                size="small"
                                v-model="props.node.image"
                                controls-position="right"
                                @change="onChange"
                        ></el-input>
                    </div><div>
                        <el-upload
                                :action="$globalConfig.goServer + '/file/uploadToServer'"
                                :before-upload="beforeUploadForm" :show-upload-list="false"
                                :http-request="imageChange"
                               >
                            <el-button size="small" type="primary">点击上传</el-button>
                            <div slot="tip" class="el-upload__tip">只能上传图片文件</div>
                        </el-upload>
                </div>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                nodeId: null,
                nodeIsJson: false,
                nodeData: '',
                lineStyle: null,
                bkType: "0",
                options: [{
                    value: '[1,0]',
                    label: '_________________'
                }, {
                    value: '[5,10]',
                    label: '- - - - - - - - - - - - - -'
                }],
                bgOptions: [{
                    value: 0,
                    label: '纯色背景'
                }, {
                    value: 1,
                    label: '线性渐变'
                }, {
                    value: 2,
                    label: '径向渐变'
                }],
                textAlignOptions: [{
                    value: 'left',
                    label: '左对齐'
                }, {
                    value: 'center',
                    label: '居中对齐'
                }, {
                    value: 'right',
                    label: '右对齐'
                }],

                baseAlignOptions: [{
                    value: 'top',
                    label: '顶部对齐'
                }, {
                    value: 'middle',
                    label: '居中对齐'
                }, {
                    value: 'bottom',
                    label: '底部对齐'
                }],
            }
        },
        props: {
            props: {
                type: Object,
                require: true,
            }
        },
        computed:{
        },
        updated() {
            if (!this.props.node || this.nodeId === this.props.node.id) {
                return;
            }
            this.props.expand = false;
            this.nodeId = this.props.node.id;
            let originData = this.props.node.data;
            this.nodeIsJson = this.isJson(originData);
            this.nodeData = this.nodeIsJson ?
                JSON.stringify(originData, null, 4) :
                this.nodeData = originData;
        },
        methods: {
            // 开始上传前验证
            beforeUploadForm (file) {
                // 验证文件类型
                var testmsg = file.name.substring(file.name.lastIndexOf('.') + 1)
                const extension = testmsg === 'jpg' || testmsg === 'png' || testmsg === 'gif'
                if (!extension) {
                    this.$message({
                        message: '上传文件只能是jpg/png/gif格式!',
                        duration: 1000,
                        showClose: true,
                        type: 'warning'
                    })
                }
                return extension
            },

            // 提交图片
            imageChange(param,type){
                const formData = new FormData()
                formData.append('myfile', param.file)
                formData.append('projectName', this.$route.query.fileName)
                var _this=this;
                _this.$axios.post(this.$globalConfig.goServer + '/file/uploadToServer', formData).then(res => {
                    if (res.data.errno==0) {
                        _this.props.node.image=res.data.data[0]
                    }
                });
                // console.log(this.imgList);
            },
            onSelectChange(value) {
                this.props.node.lineDash = eval(value)
                this.$emit('change', this.props.node);
            },
            onChange(value) {
                if (this.props.node) {
                    this.props.node.data = this.nodeIsJson ? JSON.parse(this.nodeData) : this.nodeData;
                }
                if (arguments.length > 1) {
                    //如果使用了线性渐变没有设置开始和结束颜色会卡住页面
                    if (arguments[1] == "bkType" && value > 0) {
                        if (this.props.node.gradientFromColor == null) {
                            this.props.node.gradientFromColor = '#99FF33'
                        }
                        if (this.props.node.gradientToColor == null) {
                            this.props.node.gradientToColor = '#FF6633'
                        }
                    }
                }
                this.$emit('change', this.props.node);
            },
            changeExpand() {
                this.props.expand = !this.props.expand;
            },
            isJson(obj) {
                return typeof (obj) == "object" && Object.prototype.toString.call(obj).toLowerCase() == "[object object]" && !obj.length;
            }
        }
    }
</script>

<style lang="scss">
    .star {
        display: block;
        color: #f60 !important;
        text-decoration: underline;
        margin-bottom: 0.1rem;
    }

    .title {
        color: #0d1a26;
        font-weight: 600;
        padding: 0.05rem 0.15rem;
        border-bottom: 1px solid #ccc;
    }

    .group {
        margin: 0.1rem 0 0.2rem 0.3rem;
        padding: 0;

        a,
        li {
            line-height: 2;
        }

        li {
            list-style: initial;
        }
    }

    .bottom {
        position: absolute;
        bottom: 0.1rem;
    }

    .items {
        padding: 0.05rem 0.15rem;

        .el-input-number {
            width: 1.02rem;
            line-height: 0.32rem;

            .el-input__inner {
                padding-left: 0;
                padding-right: 40px;
                height: 0.34rem;
                line-height: 0.34rem;
            }

            .el-input-number__increase {
                line-height: 0.16rem;
            }
        }

        .custom-data i {
            cursor: pointer;
            float: right;
            color: #409eff;
            height: 2em;
            line-height: 2em;
        }

        .expand-data {
            position: absolute;
            right: 0.15rem;
            width: 5rem;
        }
    }

    .formItem {
        margin-bottom: 0.1rem;
    }
</style>