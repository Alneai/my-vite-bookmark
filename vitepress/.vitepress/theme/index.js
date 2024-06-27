import DefaultTheme from 'vitepress/theme'
import NavComponent from './components/NavComponent.vue'
import NoteComponent from './components/NoteComponent.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

export default {
  extends: DefaultTheme,
  enhanceApp({ app }) {
    app.use(ElementPlus)
    app.component('NavComponent', NavComponent)
    app.component('NoteComponent', NoteComponent)
  }
}