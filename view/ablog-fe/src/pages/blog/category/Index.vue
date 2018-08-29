<template>
  <div class="container">

    <my-header :class="{
      hide: hideNavs
    }">
      <navbar :navs="navs"></navbar>
    </my-header>

    <ul class="categories">

      <li class="category" v-for="(category, index) in categories" :key="'category-item-' + index" @click="handleClickCategory(category)">
        <span class="count">{{category.articlesCount}}</span>
        <span class="name">{{category.name}}</span>
      </li>

    </ul>

  </div>
</template>

<script>
import { BLOG_NAVS } from '@/const';

import { CATEGORY_ACTION_GET_CATEGORIES } from '@/store/types';

export default {
  data() {
    return {
      navs: (BLOG_NAVS || []).map(nav => ({
        ...nav,
        active: nav.link === '/blog/categories'
      })),
      hideNavs: false
    };
  },

  methods: {
    loadCategories() {
      this.$store
        .dispatch(CATEGORY_ACTION_GET_CATEGORIES)
        .then(res => {})
        .catch(err => {});
    },
    handleClickCategory(category) {
      this.$router.push(`/blog/articles/category/${category.id}`);
    }
  },

  computed: {
    categories() {
      return this.$store.getters.categories;
    }
  },

  mounted() {
    this.loadCategories();
    this.$on('root-scroll', res => {
      if (res.scrollY > 88) {
        if (res.direction === 'up') {
          this.hideNavs = false;
        } else {
          this.hideNavs = true;
        }
      } else {
        this.hideNavs = false;
      }
    });
  }
};
</script>

<style lang="scss" scoped>
@media screen and(min-width: 980px) {
  .categories {
    max-width: 70vw;
  }
}

.container {
  background-color: #eee;
  min-height: 100vh;
  padding-top: 88px;
  box-sizing: border-box;

  .header {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    z-index: 9999;
    transition: all 0.3s;
    transition-delay: 0.1s;

    .header-main {
      .navbar {
        height: 100%;
        float: left;
      }
    }

    &.hide {
      top: -88px;
    }
  }

  .categories {
    margin: 0 auto;
    padding: 0;
    list-style: none;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    .category {
      cursor: pointer;
      display: flex;
      flex-flow: row nowrap;
      align-items: center;
      padding: 12px 16px;
      background-color: #fff;

      transition: all 0.3s;

      .count,
      .name {
        display: inline-block;
      }

      .count {
        width: 64px;
        text-align: center;
        color: $themeColor;
        font-size: 24px;
      }

      .name {
        color: #333;
        flex: 1;
        font-size: 14px;
      }

      &:not(:last-child) {
        border-bottom: 1px solid #f5f5f5;
      }

      &:hover {
        background-color: #f5f5f5;
      }

      &:active {
        background-color: #e5e5e5;
      }
    }
  }
}
</style>
