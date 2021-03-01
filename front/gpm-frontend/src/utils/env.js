export const GlobalConfig={
    goServer:window.goServer||"//192.168.20.13:8089/",
    supportFile:/^.+(\.txt|\.md|\.flow|\.html|\.js|\.css|\.json|\.yaml|\.yml|\.xml|\.java|\.py|\.vue|\.mind|\.sh|\.cmd)$/,
    documentServer:"http://10.10.0.100",
    editorMapping:{
        "\\.md|\\.markdown":['/mdeditor',"markdown编辑器"],
        "\\.fl|\\.flow":['/floweditor',"流程编辑器"],
        "\\.xls|\\.xlsx":['/officeeditor',"excel编辑器"],
        "\\.html|\\.txt":['/htmleditor',"html编辑器"],
        "\\.json|\\.js|\\.py|\\.css|\\.java|\\.vue|\\.yml|\\.yaml|\\.xml|\\.cmd|\\.sh":['/codeeditor',"code编辑器"],
        "\\.pdf":['/pdfeditor',"pdf预览器"],
        "\\.doc|\\.docx|\\.ppt|\\.pptx":['/officeeditor',"office编辑器"],
        "\\.png|\\.PNG|\\.JPG|\\.jpg|\\.JPEG|\\.jpeg|\\.gif|\\.GIF":['/imageviewer',"pdf预览器"],
        "\\.mind":['/mindeditor',"pdf预览器"],
    }
}
