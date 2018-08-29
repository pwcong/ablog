<template>
  <div class="progress-bar" :style="style"></div>
</template>

<script>
export default {
  computed: {
    style() {
      const progress = this.progress;
      const options = this.progress.options;
      return {
        width: progress.percent + '%',
        display: options.show ? 'block' : 'none'
      };
    },
    progress() {
      return window.ProgressBarEventBus.PROGRESS_BAR;
    }
  }
};
</script>

<style lang="scss">
.progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  height: 2px;
  background-color: $themeColor;
  z-index: 10000;
}
</style>
