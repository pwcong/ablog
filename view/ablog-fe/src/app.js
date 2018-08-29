import 'normalize-css';
import 'animate.css';

import './assets/css/google-font.css';
import './assets/css/markdown.css';

import Vue from 'vue';

import store from './store';
import router from './routes';

import components from './components';
components.install(Vue);

import plugins from './plugins';
plugins.install(Vue);

import App from './pages/App.vue';

window.app = new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
});
