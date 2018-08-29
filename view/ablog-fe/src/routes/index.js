import Vue from 'vue';
import VueRouter from 'vue-router';

// 首页
import Index from '@/pages/Index.vue';

// 博客页面
import Blog from '@/pages/blog/Index.vue';
import Article from '@/pages/blog/article/Index.vue';
import Category from '@/pages/blog/category/Index.vue';
import Tag from '@/pages/blog/tag/Index.vue';

// 测试页面
import Test from '@/pages/Test.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    component: Index
  },
  {
    path: '/blog',
    redirect: '/blog/articles'
  },
  {
    path: '/blog/articles',
    component: Blog
  },
  {
    path: '/blog/articles/:filter/:value',
    component: Blog
  },
  {
    path: '/blog/articles/:filter/:value/:search',
    component: Blog
  },
  {
    path: '/blog/article/:id',
    component: Article
  },
  {
    path: '/blog/categories',
    component: Category
  },
  {
    path: '/blog/tags',
    component: Tag
  },
  {
    path: '/test',
    component: Test
  }
];

const router = new VueRouter({
  routes
});

export default router;
