import Icon from './Icon.vue';
import Pagination from './Pagination.vue';
import MyButton from './MyButton.vue';

import FormField from './form/FormField.vue';
import FormLabel from './form/FormLabel.vue';
import FormInput from './form/FormInput.vue';

function install(Vue) {
  Vue.component('icon', Icon);
  Vue.component('pagination', Pagination);
  Vue.component('my-button', MyButton);
  Vue.component('form-field', FormField);
  Vue.component('form-label', FormLabel);
  Vue.component('form-input', FormInput);
}

export default {
  install
};
