export function fileIcon(title){
        if(!title) return 'icon-wenjian'
        if (title.endsWith(".doc") || title.endsWith(".docx")) {
            return "icon-doc"
        }
        if (title.endsWith(".xls") || title.endsWith(".xlsx")) {
            return "icon-xls"
        }
        if (title.endsWith(".ppt") || title.endsWith(".pptx")) {
            return "icon-ppt"
        }
        if (title.endsWith(".json")) {
            return "icon-json"
        }
        if (title.endsWith(".js")) {
            return "icon-js-square"
        }
        if (title.endsWith(".pdf")) {
            return "icon-pdf"
        }
        if (/.*\.(png|PNG|jpg|JPG|JPEG|jpeg|gif|GIF)/.test(title)) {
            return "icon-picture"
        }
        if (/.*\.(zip|7z|rar)/.test(title)) {
            return "icon-zip"
        }
        if (title.endsWith(".md")) {
            return "icon-file-markdown"
        }
        if (title.endsWith(".html")) {
            return "icon-HTML"
        }
        if (title.endsWith(".xml")) {
            return "icon-xml"
        }
        if (title.endsWith(".flow")) {
            return "icon-liuchengtu"
        }
        if (title.endsWith(".mind")) {
            return "icon-xmind"
        }
        return 'icon-wenjian'
}

export function routePush(node,routerAddress,title){
    this.$router.push({
        path: routerAddress,
        // query: {dirPath: node.dirPath, fileName: node.title}
        params: {dirPath: node.dirPath, fileName: node.title}
    });
}
export function randomUuid(len) {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0,
            v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    }).substr(0,len);
}
