<template>
  <header class="header">
    <div class="header-wrapper">

      <router-link class="header-logo" to="/">
        <img :src="logo" alt="LOGO">
        <span>{{appName}}</span>
      </router-link>

      <div class="header-main">
        <slot>
        </slot>
      </div>

    </div>
  </header>

</template>

<script>
import logo from '@/assets/imgs/logo.png';

import { APP_NAME } from '@/config';

export default {
  props: {},
  data() {
    return {
      logo,
      appName: APP_NAME
    };
  }
};
</script>

<style lang="scss">
.header {
  background-color: white;
  .header-wrapper {
    max-width: 980px;
    margin: 0 auto;
    padding: 0 16px;

    align-items: center;
    flex-flow: row nowrap;
    box-sizing: border-box;
    height: 60px;
    display: flex;

    .header-logo {
      display: flex;
      flex-flow: row nowrap;
      align-items: center;
      margin-right: 32px;
      text-decoration: none;
      color: #333;
      img {
        width: 36px;
        height: 36px;
        margin-right: 16px;
      }

      span {
        font-size: 24px;
      }
    }

    .header-main {
      flex: 1;
      height: 100%;
    }
  }
}
</style>
