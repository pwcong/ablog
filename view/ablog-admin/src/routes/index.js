import Vue from 'vue';
import VueRouter from 'vue-router';

// 首页
import Index from '@/pages/Index.vue';

// 登录页
import Login from '@/pages/Login.vue';

// 注入路由权限
import injectPermission from './permission';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    component: Index
  },
  {
    path: '/login',
    component: Login
  }
];

const router = new VueRouter({
  routes
});

injectPermission(router);

export default router;
