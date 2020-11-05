define(function (require, exports, module) {
  var png = require("../protocol/png");
  var svg = require("../protocol/svg");
  var json = require("../protocol/json");
  var plain = require("../protocol/plain");
  var md = require("../protocol/markdown");
  var mm = require("../protocol/freemind");

  function ExportRuntime() {
    var minder = this.minder;
    var hotbox = this.hotbox;
    var exps = [
      {label: '.json', key: 'j', cmd: exportJson},
      {label: '.png', key: 'p', cmd: exportImage},
      {label: '.svg', key: 's', cmd: exportSVG},
      {label: '.txt', key: 't', cmd: exportTextTree},
      {label: '.md', key: 'm', cmd: exportMarkdown},
      {label: '.mm', key: 'f', cmd: exportFreeMind}
    ];

    var main = hotbox.state('main');
    main.button({
      position: 'top',
      label: '导出',
      key: 'E',
      enable: canExp,
      next: 'exp'
    });
    var main1 = hotbox.state('main');
    main1.button({
      position: 'top',
      label: '保存',
      key: 'S',
      enable: canExp,
      action: function() {
        var saveData=JSON.stringify(minder.exportJson())
        window.vueThis.$axios({
          url: window.vueThis.$globalConfig.goServer+"file/save",
          method: 'post',
          data: {
            value: saveData,
            fileDir: window.vueThis.$route.query.dirPath,
            fileName: window.vueThis.$route.query.fileName
          },
          header: {
            'Content-Type': 'application/json'  //如果写成contentType会报错
          }
        }).then((response) => {
          window.vueThis.$Message.info("保存成功")
        });
      }
    });

    var exp = hotbox.state('exp');
    exps.forEach(item => {
      exp.button({
        position: 'ring',
        label: item.label,
        key: null,
        action: item.cmd
      });
    });

    exp.button({
      position: 'center',
      label: '取消',
      key: 'esc',
      next: 'back'
    });

    function canExp() {
      return true;
    }

    function exportJson(){
      json.exportJson(minder);
    }

    function exportImage (){
      //png.exportPNGImage(minder);
      minder.exportData("png").then((imgUrl)=>{
        // 这里是获取到的图片base64编码,这里只是个例子哈，要自行编码图片替换这里才能测试看到效果
       // 如果浏览器支持msSaveOrOpenBlob方法（也就是使用IE浏览器的时候），那么调用该方法去下载图片
        if (window.navigator.msSaveOrOpenBlob) {
          var bstr = atob(imgUrl.split(',')[1])
          var n = bstr.length
          var u8arr = new Uint8Array(n)
          while (n--) {
            u8arr[n] = bstr.charCodeAt(n)
          }
          var blob = new Blob([u8arr])
          window.navigator.msSaveOrOpenBlob(blob, 'chart-download' + '.' + 'png')
        } else {
          // 这里就按照chrome等新版浏览器来处理
          const a = document.createElement('a')
          a.href = imgUrl
          a.setAttribute('download', 'chart-download')
          a.click()
        }
      })

    }

    function exportSVG (){
      svg.exportSVG(minder);
    }

    function exportTextTree (){
      plain.exportTextTree(minder);
    }

    function exportMarkdown (){
      md.exportMarkdown(minder);
    }

    function exportFreeMind (){
      mm.exportFreeMind(minder);
    }
  }

  return module.exports = ExportRuntime;
});
