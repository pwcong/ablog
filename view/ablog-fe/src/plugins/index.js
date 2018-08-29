import ProgressBar from './progressbar';
import Toast from './toast';
function install(Vue) {
  Vue.use(ProgressBar);
  Vue.use(Toast);
}

export default {
  install
};
