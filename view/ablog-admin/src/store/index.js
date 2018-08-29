import Vue from 'vue';
import Vuex from 'vuex';

import articleModule from './modules/article';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    article: articleModule
  }
});

export default store;
