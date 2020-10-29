<style>
    .x-spreadsheet{
        height: 100%;
    }
    .x-spreadsheet-sheet{
        height: 90%;
    }
    .x-spreadsheet-toolbar{
        height: 5%;
    }
    .x-spreadsheet-bottombar{
        height: 5%;
    }
</style>
<template>

    <div style="height: 100%">
        <div id="x-spreadsheet" ref="xss" style="height: 100%" ></div>
    </div>

</template>
<script>
    import Spreadsheet from 'x-data-spreadsheet';
    import zhCN from 'x-data-spreadsheet/dist/locale/zh-cn';
    import XLSX from 'xlsx'
    import $ from 'jquery'
    //设置中文
    Spreadsheet.locale('zh-cn', zhCN);
    export default {
        data() {
            return {
            }
        },
        computed: {
            routeQueryContent() {
                return (this.$route.query.dirPath+
                    this.$route.query.fileName)
            }
        },
        watch: {
            routeQueryContent(newVal, oldVal) {
                this.initData()
            }
        },
        methods: {
            xtos(sdata) {
                var out = XLSX.utils.book_new();
                sdata.forEach(function(xws) {
                    var aoa = [[]];
                    var rowobj = xws.rows;
                    for(var ri = 0; ri < rowobj.len; ++ri) {
                        var row = rowobj[ri];
                        if(!row) continue;
                        aoa[ri] = [];
                        Object.keys(row.cells).forEach(function(k) {
                            var idx = +k;
                            if(isNaN(idx)) return;
                            aoa[ri][idx] = row.cells[k].text;
                        });
                    }
                    var ws = XLSX.utils.aoa_to_sheet(aoa);
                    XLSX.utils.book_append_sheet(out, ws, xws.name);
                });
                return out;
            },
            stox(wb) {
                var out = [];
                wb.SheetNames.forEach(function (name) {
                    var o = {name: name, rows: {}};
                    var ws = wb.Sheets[name];
                    var aoa = XLSX.utils.sheet_to_json(ws, {raw: false, header: 1});
                    aoa.forEach(function (r, i) {
                        var cells = {};
                        r.forEach(function (c, j) {
                            cells[j] = ({text: c});
                        });
                        o.rows[i] = {cells: cells};
                    })
                    out.push(o);
                });
                return out;
            },
            initData(){
                var vueThis=this;
                var dirPath=this.$route.query.dirPath
                var fileName=this.$route.query.fileName
                var req = new XMLHttpRequest();
                let token=localStorage.getItem("token")
                var url=vueThis.$globalConfig.goServer +"/file/download?fileDir=" + dirPath + "&fileName=" + fileName+(token?"&token="+token:"")
                req.open("GET", url, true);
                req.responseType = "arraybuffer";
                req.onload = function(e) {
                    var data = new Uint8Array(req.response);
                    var workbook = XLSX.read(data, {type:"array"});
                    vueThis.xs.loadData(vueThis.stox(workbook));

                }
                req.send();
            },
            export_xlsx() {
                /* build workbook from the grid data */
                var new_wb = this.xtos(this.xs.getData());
                /* generate download */
                // this is what you would normally use
                //XLSX.writeFile(new_wb, "SheetJS.xlsx");
                // codesandbox messes with the logic, so we need to do it manually
                var ab = XLSX.write(new_wb, { bookType: "xlsx", type: "array" });
                var blob = new Blob([ab]);
                var url = URL.createObjectURL(blob);
                var a = document.createElement("a");
                a.download = "SheetJS.xlsx";
                a.href = url;
                document.body.appendChild(a);
                a.click();
            }
        },
        mounted() {
            this.xs = new Spreadsheet('#x-spreadsheet', {
                showToolbar: true,
                showGrid: true,
                showContextmenu: true,
                style: {
                }
            })
                .loadData([]).change((cdata) => {
                    // console.log(cdata);
                    console.log('>>>', this.xs.getData());
                });

            this.xs.on('cell-selected', (cell, ri, ci) => {
                console.log('cell:', cell, ', ri:', ri, ', ci:', ci);
            }).on('cell-edited', (text, ri, ci) => {
                console.log('text:', text, ', ri: ', ri, ', ci:', ci);
            });
            this.$nextTick(function(){
                // $(".x-spreadsheet-toolbar-btns").append("<div class=\"x-spreadsheet-toolbar-btn\" data-tooltip=\"下载 (Ctrl+P)\"><div class=\"x-spreadsheet-icon\"><div class=\"x-spreadsheet-icon-img print\"></div></div></div>")
            })
            this.initData()
        }
    }
</script>