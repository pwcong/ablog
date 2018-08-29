import { TAG_MUTATION_SET_TAGS, TAG_ACTION_GET_TAGS } from '../types';

import {} from '@/network/api/article';

import { getTags } from '@/network/api/tag';

const store = {
  state: {
    tags: []
  },
  getters: {
    tags: state => state.tags
  },
  mutations: {
    [TAG_MUTATION_SET_TAGS]: (state, payload) => {
      state.tags = (payload || []).map(tag => ({
        ...tag,
        articlesCount: (tag.articles || []).length
      }));
    }
  },
  actions: {
    [TAG_ACTION_GET_TAGS]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        getTags()
          .then(res => {
            commit(TAG_MUTATION_SET_TAGS, res.data);
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
