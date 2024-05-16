# APK签名检查工具

此工具旨在简化APK文件签名的查看过程，通过一个简单易用的Web界面，使非技术人员也能轻松查看APK文件的签名信息。

## 功能简介

本工具通过Web端接收上传的APK文件，后端运行`keytool -printcert -jarfile <apk>`命令后，将签名信息显示在Web页面上，简化了签名检查流程。

## 快速开始

要使用此工具，请按以下步骤操作：

1. **设置端口**：修改`main.go`文件中的`http.ListenAndServe(":8080", nil)`，将`:8080`替换为所需的端口号。
2. **编译程序**：在命令行中运行`go build main.go`。
3. **启动服务**：通过访问`IP:Port`（替换为实际的IP地址和你设置的端口）来启动Web服务。
4. **上传APK文件**：在Web界面中上传你的APK文件。
5. **检查签名**：点击`上传并检查`按钮，查看签名信息。

## 缘起

开发这个工具的初衷是为了帮助公司非技术部门的同事简化查看APK签名的过程。通过提供一个无需安装Java环境和使用命令行的解决方案，降低了使用门槛。

## 效果展示

以下是工具的部分截图：

### 基本效果

<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/效果.png" alt="效果图" width="600"/>

### 高亮示例

高亮特定签名信息可以提升阅读体验。以通过修改`templates/index.html`中的以下代码（`xxx`替换为需要高亮的结果）来实现高亮显示：

```html 
/xxx/g, '<span class="highlight">xxx</span>'
```
以下是开启高亮后的效果示例：

<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/高亮1.png" alt="效果图" width="600"/>

<img src="https://github.com/Ed1s0nZ/APK-SignCheck/blob/main/高亮2.png" alt="效果图" width="600"/>

## 改进建议

欢迎通过GitHub Issues或Pull Requests提供反馈和改进建议。


