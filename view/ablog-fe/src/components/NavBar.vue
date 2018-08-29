<template>
  <nav class="navbar">
    <router-link :class="{
        active: nav.active
      }" v-for="(nav, index) in navs" :to="nav.link" :key="'nav-item-' + index">
      {{nav.name}}
    </router-link>
  </nav>
</template>

<script>
export default {
  props: {
    navs: {
      type: Array,
      default() {
        return [];
      }
    }
  }
};
</script>

<style lang="scss">
.navbar {
  display: flex;
  flex-flow: row nowrap;
  a {
    height: 100%;
    padding: 0 24px;
    display: flex;
    align-items: center;
    text-decoration: none;
    color: #666;
    box-sizing: border-box;
    border-bottom: 2px solid transparent;
    transition: all 0.3s;
    &:hover {
      color: $themeColor;
      border-color: $themeColor;
    }

    &.active {
      color: $themeColor;
      border-color: $themeColor;
    }
  }
}
</style>
