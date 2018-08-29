import NotificationWidget from './Notification.vue';

const DEFAULT_OPTIONS = {
  active: false,
  text: '',
  icon: ''
};

function install(Vue, options = {}) {
  let Notification = {
    $vm: null,
    _init(vm) {
      this.$vm = vm;
    },
    notify(content) {
      const type = Object.prototype.toString.call(content);

      let options = Object.assign({}, this.$vm.NOTIFICATION.options);
      if (type === '[object String]') {
        options.text = content;
      } else if (type === '[object Object]') {
        options = Object.assign(options, content);
      } else {
        throw new Error('invalid type of content');
      }
    }
  };

  window.NotificationEventBus = new Vue({
    data: {
      NOTIFICATION: {
        options: Object.assign({}, DEFAULT_OPTIONS, options)
      }
    }
  });

  Notification._init(window.NotificationEventBus);

  Vue.component('notification', NotificationWidget);
  Vue.prototype.$notificate = Notification;
}

export default {
  install
};
