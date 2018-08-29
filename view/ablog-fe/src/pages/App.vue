<template>
  <div class="wrapper">
    <progress-bar />
    <toast />
    <transition enter-active-class="animated fadeIn" leave-active-class="animated fadeOut">
      <router-view ref="container"></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  mounted() {
    // 监听根窗口滚动

    let scrollY = 0;
    window.addEventListener('scroll', () => {
      this.$refs.container.$emit('root-scroll', {
        scrollY: window.scrollY,
        direction: window.scrollY > scrollY ? 'down' : 'up'
      });
      scrollY = window.scrollY;
    });
  }
};
</script>

<style lang="scss">
.animated {
  animation-duration: 0.3s;
  position: absolute;
  width: 100%;
  left: 0;
  top: 0;
}

html,
body {
  font-family: 'Montserrat', 'Helvetica Neue', 'Hiragino Sans GB', 'LiHei Pro',
    Arial, sans-serif;
  position: relative;
}

.wrapper {
  position: relative;
}
</style>
