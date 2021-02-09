<style>
    .literally {
        width: 100%;
        height: 100%;
        position: relative;
    }
</style>
<template>

    <div ref="element" id="my-drawing" style="height: 100%;overflow: auto;margin-left: 5px">
    </div>

</template>
<script>
    import hotkeys from 'hotkeys-js';

    export default {
        data() {
            return {
                imgSrc: "",
                lc: null
            }
        },
        computed: {
            routeQueryContent() {
                return this.$route.query.dirPath +
                    this.$route.query.fileName
            }
        },
        watch: {
            routeQueryContent(newVal, oldVal) {
                this.initData()
            }
        },
        components: {},
        methods: {
            drawDashLine(ctx,[x1, y1], [x2, y2], step = 5) {
                const x = x2 - x1,
                    y = y2 - y1,
                    count = Math.floor(Math.sqrt(x * x + y * y) / step),
                    xv = x / count,
                    yv = y / count;
                ctx.beginPath();
                for (let i = 0; i < count; i++) {
                    if (i % 2 === 0) {
                        ctx.moveTo(x1, y1);
                    } else {
                        ctx.lineTo(x1, y1);
                    }
                    x1 += xv;
                    y1 += yv;
                }
                ctx.lineTo(x2, y2);
            },
            drawDashRect(ctx, left, top, width, height, step = 5) {
                this.drawDashLine(ctx,[left, top], [left + width, top], step);
                ctx.stroke();
                this.drawDashLine(ctx,
                    [left + width, top],
                    [left + width, top + height],
                    step
                );
                ctx.stroke();
                this.drawDashLine(ctx,
                    [left + width, top + height],
                    [left, top + height],
                    step
                );
                ctx.stroke();
                this.drawDashLine(ctx,[left, top + height], [left, top], step);
                ctx.stroke();
            },
            initData() {
                let _this = this;
                let dirName = this.$route.query.dirPath;
                let fileName = this.$route.query.fileName;
                let token = localStorage.getItem("token")
                let img = new Image()
                img.src = this.$globalConfig.goServer + '/file/download?fileDir=' + dirName + '&fileName=' + fileName + (token ? "&token=" + token : "") + "&Workspace=" + this.$store.getters.currentWorkspace;
                //定义tools
                let SelectTool = function (lc) {  // take lc as constructor arg
                    let self = this;
                    return {
                        usesSimpleAPI: false,  // DO NOT FORGET THIS!!!
                        name: 'MyTool',
                        iconName: 'select',
                        strokeWidth: lc.opts.defaultStrokeWidth,
                        optionsStyle: 'stroke-width',
                        didBecomeActive: function (lc) {
                            self.unsubscribeFuncs = [
                                lc.on('lc-pointerdown', (pt) => {
                                    self.x1 = pt.x
                                    self.y1 = pt.y;
                                }),
                                lc.on('lc-pointerdrag', (pt) => {
                                }),
                                lc.on('lc-pointerup', (pt) => {
                                    self.x2 = pt.x
                                    self.y2 = pt.y
                                    var context = lc.canvas.getContext('2d');
                                    _this.drawDashRect(context, self.x1, self.y1, self.x2 - self.x1, self.y2 - self.y1, 4)
                                }),
                                lc.on('lc-pointermove', (pt) => {
                                })
                            ];
                        },
                        willBecomeInactive: function (lc) {
                            self.unsubscribeFuncs.map(function (f) {
                                f()
                            });
                        }
                    }
                };

                this.lc = LC.init(document.getElementById("my-drawing"), {
                    imageURLPrefix: '/literallycanvas/lc-images',
                    toolbarPosition: 'bottom',
                    defaultStrokeWidth: 2,
                    keyboardShortcuts: true,
                    strokeWidths: [1, 2, 3, 5, 30],
                    watermarkScale: 1,
                    tools: LC.defaultTools.concat([SelectTool])
                });
                this.lc.saveShape(LC.createShape('Image', {x: 0, y: 0, width: 100, height: 100, image: img}));
            },
            injectCustomJs(jsSrc, func) {
                let _this = this;
                let head = document.getElementsByTagName('head')[0];
                let script = document.createElement('script');
                script.type = 'text/javascript';
                script.onload = function () {
                    func && func()
                }
                script.src = jsSrc;
                head.appendChild(script);
            },
            injectCustomCss(cssPath) {
                var temp = document.createElement('link');
                temp.setAttribute('rel', 'stylesheet');
                temp.setAttribute('href', cssPath);
                document.head.appendChild(temp);
            }
        },
        created() {


        },
        mounted() {
            let _this = this;
            _this.injectCustomCss("/literallycanvas/literallycanvas.css")
            _this.injectCustomJs("/literallycanvas/react-0.14.3.js", () => {
                _this.injectCustomJs("/literallycanvas/literallycanvas.js", () => {
                    _this.initData()
                    //ctrl+x 选择 ctrl+y 反选
                    if (!window.regImageViewHotKey) {
                        window.regImageViewHotKey = true
                        hotkeys('ctrl+d', function (event, handler) {
                            event.preventDefault()
                            alert("delete")
                        })

                        _this.lc.canvas.addEventListener("mousedown", (e) => {

                        }, false)
                    }
                })
            })


        }
    }
</script>
