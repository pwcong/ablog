<template>
  <div class="container">

    <my-header :class="{
      hide: hideNavs
    }">
      <navbar :navs="navs"></navbar>
    </my-header>

    <div class="tags">

      <div class="tag" v-for="(tag, index) in tags" @click="handleClickTag(tag)" :key="'tag-item-' + index" :style="{
          'background-color': '#' + getColor(index),
          'padding-left': ((8 + tag.articlesCount * 2) > 48 ? 48 : (8 + tag.articlesCount * 2)) + 'px',
          'padding-right': ((8 + tag.articlesCount * 2) > 48 ? 48 : (8 + tag.articlesCount * 2)) + 'px'
        }">
        {{tag.name}}
      </div>

    </div>

  </div>
</template>

<script>
import { BLOG_NAVS, COMFORTABLE_COLORS } from '@/const';

import { TAG_ACTION_GET_TAGS } from '@/store/types';

export default {
  data() {
    return {
      navs: (BLOG_NAVS || []).map(nav => ({
        ...nav,
        active: nav.link === '/blog/tags'
      })),
      hideNavs: false
    };
  },

  methods: {
    loadTags() {
      this.$store
        .dispatch(TAG_ACTION_GET_TAGS)
        .then(res => {})
        .catch(err => {});
    },
    getColor(index) {
      return COMFORTABLE_COLORS[index % COMFORTABLE_COLORS.length];
    },
    handleClickTag(tag) {
      this.$router.push(`/blog/articles/tag/${tag.id}`);
    }
  },

  computed: {
    tags() {
      return this.$store.getters.tags;
    }
  },

  mounted() {
    this.loadTags();

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
  .tags {
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

  .tags {
    margin: 0 auto;
    display: flex;
    flex-flow: row wrap;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    background-color: white;
    padding: 16px;
    box-sizing: border-box;
    .tag {
      display: flex;
      font-size: 12px;
      height: 36px;
      flex-flow: row nowrap;
      align-items: center;
      justify-content: center;
      margin: 8px;
      padding: 0 8px;
      border-radius: 18px;
      color: white;
      position: relative;
      transition: all 0.3s;
      cursor: pointer;
      top: 0;

      &:hover {
        top: -2px;
        box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.4);
      }
    }
  }
}
</style>
