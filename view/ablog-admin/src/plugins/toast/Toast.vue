<template>
  <div class="toast" :style="style">
    <div class="toast-wrapper">
      <div class="process" v-if="options.isProcess">
        <icon name="sync" />
      </div>
      {{options.text}}
    </div>
  </div>
</template>
<script>
export default {
  computed: {
    style() {
      const options = this.options;
      return {
        display: options.active ? 'flex' : 'none'
      };
    },
    options() {
      return this.toast.options;
    },
    toast() {
      return window.ToastEventBus.TOAST;
    }
  }
};
</script>
<style lang="scss">
.toast {
  position: fixed;
  z-index: 10000;
  bottom: 0;
  left: 0;
  right: 0;
  top: 0;
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  justify-content: center;
  .toast-wrapper {
    color: white;
    padding: 16px;
    border-radius: 4px;
    font-size: 13px;
    display: inline-block;
    background-color: rgba(0, 0, 0, 0.8);

    .process {
      text-align: center;
      font-size: 26px;
      @keyframes anim-ring {
        0% {
          transform: rotate(0);
        }

        100% {
          transform: rotate(359deg);
        }
      }
      .icon {
        animation: anim-ring 2s linear infinite;
      }
    }
  }
}
</style>
