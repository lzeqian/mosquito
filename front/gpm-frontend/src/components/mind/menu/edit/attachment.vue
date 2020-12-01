<template lang="">
    <div class="attachment-group" :disabled="commandDisabled">
        <div class="link ">
        <el-button class="tab-icons insert" @click="showOverlay"></el-button>
        <el-dropdown trigger="click" :disabled="commandDisabled">
        <span class="el-dropdown-link">
        链接
        <i class="el-icon-caret-bottom el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown" class="link-dropdown-list">
        <el-dropdown-item  @click.native="handleCommand(3)">插入链接</el-dropdown-item>
        <el-dropdown-item  @click.native="handleCommand(4)">移除已有链接</el-dropdown-item>
        </el-dropdown-menu>
        </el-dropdown>
        </div>
        <div class="img">
        <el-button class="tab-icons insert" @click="showOverlay"></el-button>
        <el-dropdown trigger="click">
        <span class="el-dropdown-link">
        图片
        <i class="el-icon-caret-bottom el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown" class="img-dropdown-list">
        <el-dropdown-item @click.native="handleCommand(1)">
        <el-upload :disabled="commandDisabled"
    :action="$globalConfig.goServer + '/file/uploadToServer'"
    :before-upload="beforeUploadForm" :show-file-list="false"
    :http-request="imageChange"
        >
        插入图片
        </el-upload>
        </el-dropdown-item>
        <el-dropdown-item @click.native="handleCommand(2)">移除已有图片</el-dropdown-item>
        </el-dropdown-menu>
        </el-dropdown>
        </div>
        <div class="remark">
        <el-button class="tab-icons insert" @click="showOverlay"> </el-button>
        <el-dropdown trigger="click">
        <span class="el-dropdown-link">
        批注
        <i class="el-icon-caret-bottom el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown" class="remark-dropdown-list">
        <el-dropdown-item @click.native="handleCommand(5)">插入批注</el-dropdown-item>
        <el-dropdown-item @click.native="handleCommand(6)">移除已有批注</el-dropdown-item>
        <el-dropdown-item @click.native="handleCommand(7)">查看批注</el-dropdown-item>
        </el-dropdown-menu>
        </el-dropdown>
        </div>
        </div>
</template>
<script>
    import {
        mapGetters,
        mapActions,
        mapMutations
    } from 'vuex'

    export default {
        name: 'attachment',
        mounted() {
        },
        data() {
            return {
                options1: [{
                    value: '选项1',
                    label: '插入链接'
                }, {
                    value: '选项2',
                    label: '移除已有链接'
                }],
                options2: [{
                    value: '选项1',
                    label: '插入图片'
                }, {
                    value: '选项2',
                    label: '移除已有图片'
                }],
                options3: [{
                    value: '选项1',
                    label: '插入备注'
                }, {
                    value: '选项2',
                    label: '移除已有备注'
                }],
                value1: '',
                value2: '',
                value3: ''
            }
        },
        computed: {
            ...mapGetters({
                count: 'count',
                'minder': 'getMinder',
            }),
            commandDisabled() {
                var minder = this.minder
                minder.on && minder.on('interactchange', function () {
                    this.commandValue = minder.queryCommandValue('priority');
                });
                return minder.queryCommandState && minder.queryCommandState('priority') === -1;
            },
        },
        methods: {
            ...mapActions([
                'changeCount',
                'increment',
                'clearMemory',
                'setMemory',
            ]),
            handleCommand(type) {
                if(this.commandDisabled) return;
                if (type == 2) {
                    //清除图片
                    this.minder.execCommand('Image', '', '');
                }
                if (type == 3) {
                    //新增超链接
                    var hylink = prompt("请输入超链接路径：");
                    if (hylink != null && hylink.trim() != "") {
                        this.minder.execCommand('HyperLink', hylink, '');
                    }
                }
                if (type == 4) {
                    //删除超链接
                    this.minder.execCommand('HyperLink', '', '');
                }
                if (type == 5) {
                    //新增标注
                    var remark = prompt("请输入超链接路径：");
                    if (remark != null && remark.trim() != "") {
                        this.minder.execCommand('Note', remark);
                    }
                }
                if (type == 6) {
                    //删除标注
                    this.minder.execCommand('Note');
                }
                if (type == 7) {
                    //删除标注
                    if(this.minder.queryCommandState('Note')==0 && this.minder.queryCommandValue('Note'))
                        this.$Message.info("当前批注值："+this.minder.queryCommandValue('Note'));
                }
            },
            test(key, value) {
                this.clearMemory(key, value);
            },
            showOverlay() {
                this.$msgbox({
                    title: '输入',
                    message: '暂时未实现，敬请期待！',
                    showCancelButton: true,
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    beforeClose: (action, instance, done) => {
                        if (action === 'confirm') {
                            instance.confirmButtonLoading = true;
                            instance.confirmButtonText = '执行中...';
                            setTimeout(() => {
                                done();
                                setTimeout(() => {
                                    instance.confirmButtonLoading = false;
                                }, 300);
                            }, 3000);
                        } else {
                            done();
                        }
                    }
                }).then(action => {
                });
            },
            // 开始上传前验证
            beforeUploadForm(file) {
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
            imageChange(param, type) {
                debugger
                const formData = new FormData()
                formData.append('myfile', param.file)
                formData.append('projectName', this.$route.query.fileName)
                var _this = this;
                _this.$axios.post(this.$globalConfig.goServer + '/file/uploadToServer', formData).then(res => {
                    debugger
                    if (res.data.errno == 0) {
                        this.minder.execCommand('Image', res.data.data[0], "");
                    }
                });
                // console.log(this.imgList);
            },
        }
    }
</script>
