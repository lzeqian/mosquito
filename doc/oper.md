# mosquito
## 普通安装

### 系统依赖项目
#### vuepress
项目依赖vuepress构建markdown项目。

安装:
```
npm install -g vuepress
```
安装完成后确认是否可直接使用命令(环境变量PATH中配置)。
详细vuepress教程参考[官网](https://www.vuepress.cn/guide/getting-started.html)
#### pandoc
项目依赖pandoc将markdown转换为doc文件。

安装（直接下载到指定目录）:
```
https://github.com/jgm/pandoc/releases/tag/2.11 找到对应系统版本
```
解压后修改conf/app.conf
```
pandocpath=你解压的pandoc目录/pandoc
```
#### libreoffice
项目依赖libreoffice实现doc|xls|ppt转换为pdf。
##### window安装
https://zh-cn.libreoffice.org/download/libreoffice/ 找到对应版本的下载安装
桌面生成的快捷方式右键找到对应目标：
```
"C:\Program Files\LibreOffice\program\soffice.exe"
```
#### linux安装
``` 
yum -y install libreoffice
```
解压完成修改conf/app.conf
```
libreofficepath=C:/Program Files/LibreOffice/program/soffice.exe
```
### 系统部署
#### 克隆源代码
```$xslt
  git clone https://github.com/lzeqian/gpm.git
```
#### 前端部署

进入front/gpm-frontend，执行命令

```
npm i && npm run build
```

1. 将dist目录的index.html和share.html拷贝项目根目录的views目录。

2. 将dist目录的其他文件拷贝到static目录。

#### 后台部署

##### 编译打包

在部署系统(window|linux)中golang语言环境，[详情参考](https://blog.csdn.net/liaomin416100569/article/details/106082235)

执行命令编译:
 ```golang
   go build -x
 ```
window根目录下shengcheng gpm.exe，linux下生成gpm。
> 注意生成可执行文件和conf目录是必须同时存在的，否则配置无法生效
> 具体配置请参考conf/app.conf,端口配置：httpport = 8080

注意编译成功后，需要留下以下目录和文件可正常运行：

1. 可执行文件gpm|gpm.exe
2. conf目录：核心配置文件app.conf在该目录。
3. files目录：模板文件，以及默认的系统文件。
4. static目录：静态的css和js，图片等。
5. views目录：首页和共享页面的模板文件

#### 运行测试
执行gpm运行
```$xslt
$ ./gpm
2020/10/26 14:53:59.981 [I] [asm_amd64.s:1373]  http server Running on http://:8080
```
浏览器测试接口
```$xslt
http://localhost:8080/console
```



## docker安装

拷贝源代码目录下的.envfile和run.sh到任意工作目录，修改.envfile对应的参数

```
httpport=80
runmode=prod
frontGoServer=//docs.jieztech-internal.com/
frontDocumentServer=http://10.10.0.100
pandocpath=/usr/bin/pandoc
libreofficepath=/usr/bin/libreoffice
uploadDir=/upload
uploadAccessAdress=http://docs.jieztech-internal.com
sambahost=192.168.1.250
sambaport=445
sambauser=share
sambapassword=s[2]AV%E
personRootpath=/person/
personSambahost=192.168.1.250
personSambaport=445
personSambauser=share
personSambapassword=s[2]AV%E
```

run.sh脚本内容：

```
mkdir -p /etc/mosquito
setenforce 0
docker run -d -p 80:80  -v /etc/mosquito:/application/rbac --env-file ./.envfile --name mosquito liaomin789/mosquito:1.0.0
```

>  注意docs.jieztech-internal.com域名指定到要部署的docker服务器的ip地址

执行脚本

```
chmod +x ./run.sh && ./run.sh
```

