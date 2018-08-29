import {
  CATEGORY_MUTATION_SET_CATEGORIES,
  CATEGORY_ACTION_GET_CATEGORIES
} from '../types';

import { getCategories } from '@/network/api/category';

const store = {
  state: {
    categories: []
  },
  getters: {
    categories: state => state.categories
  },
  mutations: {
    [CATEGORY_MUTATION_SET_CATEGORIES]: (state, payload) => {
      state.categories = (payload || []).map(category => ({
        ...category,
        articlesCount: (category.articles || []).length
      }));
    }
  },
  actions: {
    [CATEGORY_ACTION_GET_CATEGORIES]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        getCategories()
          .then(res => {
            commit(CATEGORY_MUTATION_SET_CATEGORIES, res.data);
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    }
  }
};

export default store;
