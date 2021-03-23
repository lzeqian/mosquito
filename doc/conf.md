

# gpm文件管理项目

## .envfile配置详解

| 配置| 默认值| 描述|
| --- | --- | --- |
| httpport | 80 | 服务器启动端口号 |
| runmode | prod | 服务器启动环境，生产为prod,开发为dev |
| frontGoServer| 无| 前端指定的后端服务器地址（建议域名），一般定义为一个域名指向gpm服务器对应的ip地址，这里地址是前端调用后台的地址，目前前端和后端位于同一服务器，所以该http://${frontGoServer}/console可直接访问前端首页。 |
| externFrontGoServer| 无| 该参数（建议域名）和${frontGoServer}，域名直接指向gpm服务器对应的地址，该地址只有在分享选择外网时使用 |
| frontDocumentServer| 无| 该参数（建议域名）必须指向onlyoffice对应的访问地址，否则office无法编辑和预览 |
| externFrontDocumentServer| 无| 该参数（建议域名）必须指向onlyoffice对应的访问地址，必须是公网域名，用于外网共享时onlyoffice调用 |
| pandocpath| 无| 安装的pandoc的地址，docker容器中默认是，/usr/bin/pandoc，如果源代码编译安装，可通过该参数自定义指定pandoc的位置 |
| libreofficepath| 无| libreoffice安装的地址，docker容器中默认是，/usr/bin/libreoffice，系统暂时不依赖libreoffice，参数可选 |
| libreofficeTmpPath| 无| libreoffice操作临时文件生成位置,默认/tmp即可 |
| uploadDir| 无| html和markdown编辑器上传图片对应目录，默认/upload即可 |
| uploadAccessAdress| 无| 图片访问地址，可使用域名指定到服务器对应ip地址，配置该域名 |
| rootpath| 无| 公共共享目录挂载目录，如果是/ 将列表所有shareNme，也可以指定/shareNme |
| sambahost| 无| 公共共享目录的主机名 |
| sambaport| 无| 公共共享目录的主机端口，默认445 |
| sambauser| 无| 公共享目录的主机用户名 |
| sambapassword| 无| 共共享目录的主机密码 |
| personRootpath| 无| 个人共享目录挂载目录，必须是/shareNme/，不能使用/ |
| personSambahost| 无| 个人共共享目录的主机名 |
| personSambaport| 无| 个人共享目录的主机端口，默认445 |
| personSambauser| 无| 个人共享目录的主机用户名|
| personSambapassword| 无| 个人共享目录的主机密码|
| emailSmtpHost| 无| 用于发送邮件的主机smtp信息，如qq：smtp.exmail.qq.com:25|
| emailSmtpSender| 无| 用于发送邮件的邮箱用户名|
| personSambapassword| 无| 用于发送邮件的邮箱密码|

完整配置：
```
httpport=80
runmode=prod
frontGoServer=//docs.jieztech-internal.com/
externFrontGoServer=//docs.kldidi.com/
frontDocumentServer=http://10.10.0.100
externFrontDocumentServer=http://xxx.xxx.com
pandocpath=/usr/bin/pandoc
libreofficepath=/usr/bin/libreoffice
libreofficeTmpPath=/tmp
uploadDir=/upload
uploadAccessAdress=http://xxx.xxxx.com
rootpath=${rootpath||/}
sambahost=192.168.1.xxx
sambaport=445
sambauser=share
sambapassword=xxxxxxx
personRootpath=/person/
personSambahost=192.168.1.xxx
personSambaport=445
personSambauser=personshare
personSambapassword=xxxx
emailSmtpHost=smtp.exmail.qq.com:25
emailSmtpPassword=app@xxx.com
emailSmtpPassword=xxxx
```
## conf/app.conf详解
该参数为beego标准格式的参数，参数与.envfile一一对应。

## 权限配置文件详解
- 系统启动后会自动在宿主机/etc/mosquito生成文件rbac.yml,该文件提供了权限管理配置。
- 开发模式下默认的配置文件位于files/default_rbac.yml,启动应用后自动写入到rbac/rbac.yml中。
### 角色目录权限控制
```javascript
roles:
  admin:
    - path: '/.*'
      act:
        - read
        - listDir
        - write
        - "*"
  dev-tmp-test:
    - path: '/.*'
      act:
        - listDir
        - 'read'
    - path: '/个人文档'
      act:
        - 'write'
        - createDir
        - createFile
  product-market:
    - path: '05\.产品基础资料库'
      act:
        - 'read'
        - 'write'
```
以上配置内容含义:
 - 角色:admin对公共目录/所有子目录都拥有读写和列表权限和所有权限。
 - 角色：dev和tmp和test对公共目录/所有子目录只有列表和读权限,对/个人文档目录有写，创建目录，创建文件窗前。

所有可操作权限包括：

| 权限| 描述|
| --- | --- |
| read | 读取权限 |      
| write | 写入权限 |          
| createDir | 创建目录权限 |     
| createFile | 创建文件权限 |     
| deleteDir | 删除目录权限 |     
| deleteFile | 删除文件权限 |  
| listDir | 列表目录权限 |  
| * | 所有权限 |  
### 用户管理
```javascript
userGroups:
- groupCode: admin
  groupName: 管理员
- groupCode: apollodev
  groupName: Apollo开发部
- groupCode: operation
  groupName: 运维组
- groupCode: test
  groupName: 测试客服部
users:
  - ifActivate: 1
    password: e10adc3949ba59abbe56e057f20f883e
    position: 管理员
    role: admin
    userFullName: 管理员
    userGroup: admin
    userLog: /images/userlog.jpg
    email: admin@jieztech.com
    userName: admin
```
以上配置内容含义:
 - 创建admin管理员的分组
 - 创建用户admin，密码为：123456（注意配置是md5码）
