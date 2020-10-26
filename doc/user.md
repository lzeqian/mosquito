# gpm文件管理项目

## 系统依赖项目
### vuepress
项目依赖vuepress构建markdown项目

安装
```
npm install -g vuepress
```
安装完成后确认是否可直接使用命令(环境变量PATH中配置)

### pandoc
项目依赖pandoc将markdown转换为doc文件

安装（直接下载到指定目录）
```
https://github.com/jgm/pandoc/releases/tag/2.11 找到对应系统版本
```
解压完成修改conf/app.conf
```
pandocpath=你解压的pandoc目录
```
### libreoffice
项目依赖libreoffice实现doc|xls|ppt转换为pdf
#### window安装
https://zh-cn.libreoffice.org/download/libreoffice/ 找到对应版本的下载安装
桌面生成的快捷方式右键找到对应目标：
```
"C:\Program Files\LibreOffice\program\soffice.exe"
```
#### linux安装
```$xslt
 yum -y install libreoffice
```

解压完成修改conf/app.conf
```
libreofficepath=C:/Program Files/LibreOffice/program/soffice.exe
```
