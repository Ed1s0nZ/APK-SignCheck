## 功能
这个工具的原理很简单，通过web端上传apk文件，后端执行`keytool -printcert -jarfile <apk>`命令，将返回的签名结果展示到web端；

## 为什么做这个工具？
公司其他部门非技术同事有需要查看签名的需求，让非技术同事去安装Java环境，再通过终端执行命令去查看有一定门槛，于是写了这个小工具。

## 效果
<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/效果.png" alt="效果图" width="600"/>

## 其他
1、 可通过templates/index.html 中第73行`/xxx/g, '<span class="highlight">xxx</span>');`中的xxx替换为需要高亮的结果，可以高亮展示：

<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/高亮1.png" alt="效果图" width="600"/>
<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/高亮2.png" alt="效果图" width="600"/>
