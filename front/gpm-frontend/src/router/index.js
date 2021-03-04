import Vue from 'vue'
import VueRouter from 'vue-router'
import MdEditor from '../views/editor/MdEditor.vue'
import XlsEditor from '../views/editor/backup/XlsEditor.vue'
import CodeEditor from '../views/editor/CodeEditor.vue'
import PdfEditor from '../views/editor/PdfEditor.vue'
import HtmlEditor from '../views/editor/HtmlEditor'
import WordEditor from '../views/editor/backup/WordEditor'
import ImageViewer from '../views/editor/ImageViewer'
import FlowEditor from '../views/editor/FlowEditor'
import BlankViewer from '../views/editor/BlankViewer'
import MindEditor from "../components/mind/editor";
import OfficeEditor from "../views/editor/OfficeEditor";
Vue.use(VueRouter)

  const routes = [
    {
      path: '/',
      name: 'default',
      redirect: '/blank'
    },
    {
      path: '/blank',
      name: 'blankViewer',
      component: BlankViewer
    },
    {
      path: '/mdeditor',
      name: 'MdEditor',
      component: MdEditor
    },
    {
      path: '/exceleditor',
      name: 'ExcelEditor',
      component: XlsEditor
    },
    {
      path: '/codeeditor',
      name: 'CodeEditor',
      component: CodeEditor
    }
    ,
    {
      path: '/pdfeditor',
      name: 'PdfEditor',
      component: PdfEditor
    },
    {
      path: '/htmleditor',
      name: 'HtmlEditor',
      component: HtmlEditor
    },
    {
      path: '/wordeditor',
      name: 'WordEditor',
      component: WordEditor
    },
    {
      path: '/imageviewer',
      name: 'ImageViewer',
      component: ImageViewer
    },
    {
      path: '/floweditor',
      name: 'FlowEditor',
      component: FlowEditor
    },
    {
      path: '/mindeditor',
      name: 'MindEditor',
      component: MindEditor
    },
    {
      path: '/officeeditor',
      name: 'OfficeEditor',
      component: OfficeEditor
    }
]
const originalPush = VueRouter.prototype.push
/**
 * 解决同一个路由多次被调用出现
 * @param location
 * @returns {Promise<Route>}
 */
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const router = new VueRouter({
  routes
})

export default router
