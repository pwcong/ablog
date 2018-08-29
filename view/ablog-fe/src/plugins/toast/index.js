import ToastWidget from './Toast.vue';

function install(Vue, options = {}) {
  const DEFAULT_OPTIONS = {
    active: false,
    text: '',
    isProcess: false
  };

  let Toast = {
    $vm: null,
    init(vm) {
      this.$vm = vm;
    },
    tip(text) {
      if (!this.$vm) {
        return;
      }

      this.reset();

      this.$vm.TOAST.options.active = true;
      this.$vm.TOAST.options.text = text;
      this.timer = setTimeout(() => {
        this.$vm.TOAST.options.active = false;
        this.$vm.TOAST.options.text = '';
      }, 2000);
    },
    process(text) {
      if (!this.$vm) {
        return;
      }

      this.reset();

      this.$vm.TOAST.options.active = true;
      this.$vm.TOAST.options.text = text;
      this.$vm.TOAST.options.isProcess = true;
    },

    hide() {
      this.reset();
    },

    reset() {
      this.clearTimer();
      this.$vm.TOAST.options.active = false;
      this.$vm.TOAST.options.text = '';
      this.isProcess = false;
    },
    clearTimer() {
      if (this.timer) {
        clearTimeout(this.timer);
        this.timer = null;
      }
    }
  };

  window.ToastEventBus = new Vue({
    data: {
      TOAST: {
        options: Object.assign({}, DEFAULT_OPTIONS, options)
      }
    }
  });
  Toast.init(window.ToastEventBus);

  Vue.component('toast', ToastWidget);
  Vue.prototype.$toast = Toast;
}

export default {
  install
};
