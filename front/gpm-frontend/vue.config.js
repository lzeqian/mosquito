// vue.config.js
const HtmlWebpackPlugin = require("html-webpack-plugin");
module.exports = {
    devServer: {
        // open: process.platform === 'demo',
        // host: 'localhost',
        port: 8085,
        open: true, //配置自动启动浏览器
        proxy: {
            "/opAdmin": {
                target: "http://116.65.61.193:8081", //对应服务器的接口
                changeOrigin: true,
                pathRewrite: {
                    "^/opAdmin": "/opAdmin" //将以 /opAdmin 开头的接口重写http://116.65.61.193:8081/opAdmin ,调用时直接以 /opAdmin开头即表示调用http://116.65.61.193:8081/opAdmin
                    // "^/opAdmin": "/" //将以 /opAdmin 开头的接口重写http://116.65.61.193:8081/ ,调用时直接以 /opAdmin开头即表示调用http://116.65.61.193:8081/
                }
            },

            "/openParkApi": {
                target: "http://116.65.61.193:8087", //对应服务器的接口
                changeOrigin: true,
                pathRewrite: {
                    "^/openParkApi": "/openParkApi" //重写接口
                }
            },
        }
    },
    chainWebpack: config =>{

    },
    pages: {
        index: {
            entry: 'src/main.js',
            filename: 'index.html',
            title:"文档管理平台"
        },
        share:{
            entry: 'src/share.js',
            filename: 'share.html',
            title:"文档管理平台"
        }

    }
}
