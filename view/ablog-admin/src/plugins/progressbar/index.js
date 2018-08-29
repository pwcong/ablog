import ProgressBar from './ProgressBar.vue';

const DEFAULT_OPTIONS = {
  show: false
};

function install(Vue, options = {}) {
  let Progress = {
    $vm: null,
    init(vm) {
      this.$vm = vm;
    },
    start() {
      if (!this.$vm) {
        return;
      }

      this.clearProgressTimer();
      this.clearHideTimer();

      this.$vm.PROGRESS_BAR.percent = 0;
      this.$vm.PROGRESS_BAR.options.show = true;

      this.progressTimer = setInterval(() => {
        this.$vm.PROGRESS_BAR.percent +=
          (100 - this.$vm.PROGRESS_BAR.percent) * 0.1;
      }, 200);
    },
    stop() {
      this.clearProgressTimer();
      this.clearHideTimer();

      this.$vm.PROGRESS_BAR.options.show = false;
    },
    finish() {
      if (!this.$vm) {
        return;
      }
      this.clearProgressTimer();
      this.$vm.PROGRESS_BAR.percent = 100;
      this.hideTimer = setTimeout(() => {
        this.$vm.PROGRESS_BAR.options.show = false;
        this.hideTimer = null;
      }, 400);
    },
    clearProgressTimer() {
      if (this.progressTimer) {
        clearInterval(this.progressTimer);
        this.progressTimer = null;
      }
    },
    clearHideTimer() {
      if (this.hideTimer) {
        clearTimeout(this.hideTimer);
        this.hideTimer = null;
      }
    }
  };

  window.ProgressBarEventBus = new Vue({
    data: {
      PROGRESS_BAR: {
        percent: 0,
        options: Object.assign({}, DEFAULT_OPTIONS, options)
      }
    }
  });
  Progress.init(window.ProgressBarEventBus);

  Vue.component('progress-bar', ProgressBar);
  Vue.prototype.$progress = Progress;
}

export default {
  install
};
