module.exports = {
    title: '测试',
    description: '测试',
    base: '${base}',
    head: [
        // ...
        ['link', { rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.7.1/katex.min.css' }],
        ['link', { rel: "stylesheet", href: "https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.10.0/github-markdown.min.css" }]
    ],
    plugins: [
        'vuepress-plugin-mermaidjs',
        '@maginapp/vuepress-plugin-flowchart',
        {
            openMarker: '```mermaid',
            closeMarker: '```',
            scondMarker: 'flowchat',
            ignoreSecondLine: false
        }
    ],
    markdown: {
        anchor:{ permalink: false },
        toc:{ includeLevel: [1, 2] },
        extendMarkdown: md => {
            md.use(require('markdown-it-katex'))
                .use(require('markdown-it-footnote'))
                .use(require('markdown-it-ins'))
                .use(require('markdown-it-mark'))
                .use(require('markdown-it-sub'))
                .use(require('markdown-it-sup'))
                .use(require('markdown-it-abbr'))
        }
    },
    themeConfig: {
        sidebar: "auto",
        displayAllHeaders: true,
        sidebarDepth: 6,
        nav: [
            { text: 'Home', link: '/' },
            { text: 'baidu', link: 'https://www.baidu.com' }
        ]
    }
}
