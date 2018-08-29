import Vue from 'vue';
import Vuex from 'vuex';

import articleModule from './modules/article';
import categoryModule from './modules/category';
import tagModule from './modules/tag';
import evaluationModule from './modules/evaluation';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    article: articleModule,
    tag: tagModule,
    category: categoryModule,
    evaluation: evaluationModule
  }
});

export default store;
