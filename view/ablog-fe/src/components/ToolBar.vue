<template>
  <div class="toolbar">
    <div class="logo">
      <router-link to="/" class="app-name">
        {{appName}}
      </router-link>
    </div>
    <div class="link">
      <router-link :class="{
          active: nav.active
        }" v-for="(nav, index) in navs" :to="nav.link" :key="'nav-item-' + index">
        <icon v-if="nav.icon" :name="nav.icon" />
        <span class="name">{{nav.name}}</span>
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    navs: {
      type: Array,
      default() {
        return [];
      }
    },
    appName: {
      type: String
    }
  }
};
</script>

<style lang="scss">
.toolbar {
  .logo {
    font-size: 32px;
    text-align: center;

    .app-name {
      color: #333;
      text-decoration: none;
    }
  }

  .link {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;

    a {
      padding: 12px;
      text-decoration: none;
      color: #666;
      user-select: none;
      .name {
        transition: all 0.3s;
      }

      &.router-link-active {
        color: $themeColor;
      }
      &:hover {
        color: $themeColor;
      }
    }
  }
}
</style>
