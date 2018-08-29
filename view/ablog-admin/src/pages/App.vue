<template>
  <div class="container">
    <transition enter-active-class="animated fadeIn" leave-active-class="animated fadeOut">
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
export default {};
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
  font-family: 'Microsoft YaHei', sans-serif;
  position: relative;
}

.container {
  position: relative;
}
</style>
