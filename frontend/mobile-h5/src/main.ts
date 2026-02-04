import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import i18n from './i18n'

// Vant组件库
import {
  Button,
  Cell,
  CellGroup,
  NavBar,
  Tabbar,
  TabbarItem,
  List,
  PullRefresh,
  Card,
  Tag,
  Icon,
  Field,
  Form,
  Toast,
  Dialog,
  ActionSheet,
  Popup,
  Search,
  Tabs,
  Tab,
  Swipe,
  SwipeItem,
  Lazyload,
  Skeleton,
  Empty,
  Loading,
  Badge,
  Picker
} from 'vant'

// Vant样式
import 'vant/lib/index.css'

// 主题样式（必须在Vant样式之后导入以覆盖变量）
import './styles/theme.css'

// 全局样式
import './styles/main.scss'

const app = createApp(App)

// 注册Vant组件
const vantComponents = [
  Button,
  Cell,
  CellGroup,
  NavBar,
  Tabbar,
  TabbarItem,
  List,
  PullRefresh,
  Card,
  Tag,
  Icon,
  Field,
  Form,
  Toast,
  Dialog,
  ActionSheet,
  Popup,
  Search,
  Tabs,
  Tab,
  Swipe,
  SwipeItem,
  Lazyload,
  Skeleton,
  Empty,
  Loading,
  Badge,
  Picker
]

vantComponents.forEach(component => {
  app.use(component)
})

app.use(createPinia())
app.use(router)
app.use(i18n)

app.mount('#app')
