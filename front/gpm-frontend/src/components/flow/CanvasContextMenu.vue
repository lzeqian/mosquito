<template>
    <div class="menus">
        <div>
            <a :class="{disabled:!props.node && !props.nodes}" @click="onTop()">置顶</a>
        </div>
        <div>
            <a :class="{disabled:!props.node && !props.nodes}" @click="onBottom()">置底</a>
        </div>
        <div class="line"></div>
        <div v-if="props.nodes">
            <a @click="onCombine()">组合</a>
            <div>
                <a @click="onCombine(true)">包含</a>
            </div>
        </div>
        <div v-if="props.node && props.node.name === 'combine'">
            <a @click="onUncombine()">取消组合/包含</a>
        </div>
        <div>
            <a
                    :class="{disabled:!props.node && !props.nodes}"
                    @click="onLock()"
            >{{ props.locked ? '解锁' : '锁定' }}</a>
        </div>
        <div class="line"></div>
        <div>
            <a :class="{disabled:!props.node && !props.nodes && !props.line}" @click="onDel()">删除</a>
        </div>
        <div class="line"></div>
        <div>
            <a @click="canvas.undo()" class="flex">
                <span class="full">撤消</span>
                <span class="ml50">Ctrl + Z</span>
            </a>
        </div>
        <div>
            <a @click="canvas.redo()">
                恢复
                <span class="ml50">Ctrl + Shift+ Z</span>
            </a>
        </div>
        <div class="line"></div>
        <div>
            <a @click="canvas.cut()" class="flex">
                <span class="full">剪切</span>
                <span class="ml50">Ctrl + X</span>
            </a>
        </div>
        <div>
            <a @click="canvas.copy()" class="flex">
                <span class="full">复制</span>
                <span class="ml50">Ctrl + C</span>
            </a>
        </div>
        <div>
            <a @click="canvas.paste()" class="flex">
                <span class="full">粘贴</span>
                <span class="ml50">Ctrl + V</span>
            </a>
        </div>
        <div class="line"></div>
        <div>
            <a :class="{disabled:!props.node || !props.node.image}" @click="onCopyImage()" class="flex">
                <span class="full">复制节点图片地址</span>
            </a>
        </div>
    </div>
</template>

<script >
    export default {
        data() {
            return {}
        },
        props: {
            canvas: {
                type: Object,
                require: true
            },
            props: {
                type: Object,
                require: true
            }
        },
        methods: {
            onTop() {
                if (this.props.node) {
                    this.canvas.top(this.props.node)
                }
                if (this.props.nodes) {
                    for (const item of this.props.nodes) {
                        this.canvas.top(item)
                    }
                }
                this.canvas.render()
            },
            onBottom() {
                if (this.props.node) {
                    this.canvas.bottom(this.props.node)
                }
                if (this.props.nodes) {
                    for (const item of this.props.nodes) {
                        this.canvas.bottom(item)
                    }
                }
                this.canvas.render()
            },
            onCombine(stand) {
                if (!this.props.nodes) {
                    return
                }
                this.canvas.combine(this.props.nodes, stand)
                this.canvas.render()
            },
            onUncombine() {
                if (!this.props.node) {
                    return
                }
                this.canvas.uncombine(this.props.node)
                this.canvas.render()
            },
            onLock() {
                this.props.locked = !this.props.locked
                if (this.props.node) {
                    this.props.node.locked = this.props.locked
                }
                if (this.props.nodes) {
                    for (const item of this.props.nodes) {
                        item.locked = this.props.locked
                    }
                }
                if (this.props.lines) {
                    for (const item of this.props.lines) {
                        item.locked = this.props.locked
                    }
                }
                this.canvas.render(true)
            },
            onDel() {
                this.canvas.delete()
            }
        }
    }
</script>

<style lang="scss">
    .menus {
        color: #000;
        background-color: #fff;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.3);
        min-width: 1.8rem;
        text-align: left;
        padding: 0.08rem 0;
        & > div {
            line-height: 2.2;
            a {
                color: #314659;
                display: block;
                padding: 0 0.2rem;
                text-decoration: none;
                cursor: pointer;
                &:hover {
                    color: #1890ff;
                }
                &.flex {
                    display: flex;
                }
                &.disabled {
                    color: #ccc;
                    cursor: default;
                }
            }
        }
        .line {
            background-color: transparent !important;
            padding: 0;
            margin: 0.05rem 0;
            border-top: 1px solid #eee;
        }
    }
</style>
