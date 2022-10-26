import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
import { loadJs } from '../assets/js/app.js' 
Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'index',
        component: () => import('../views/index.vue'),
        meta: { title: '首页'}
    },
    // {
    //     path: '/about',
    //     name: 'index',
    //     component: () => import('../views/about.vue'),
    //     meta: { title: '关于'}
    // }
]

const router = new VueRouter({
    mode: 'history',
    //base: process.env.BASE_URL,
    base: '/web/',  // 二级目录
    routes
})
router.beforeEach((to, from, next) => {
    let title = 'RanBlogs'
    if (to.meta.params){
        title = `${to.meta.title}:${to.params[to.meta.params] || ''} - ${title}`
    }else {
        title = `${to.meta.title} - ${title}`
    }
    document.title = title
    if (to.path !== from.path) {
        store.dispatch('setLoading', true);
    }
    next();
})
router.afterEach((to, from) => {
    if (to.path == '/' && to.hash == '') {
        loadJs();
      }
    // 最多延迟 关闭 loading
    setTimeout(() => {
        store.dispatch('setLoading', false);
    }, 1500)
})
export default router
