import {
  ARTICLE_MUTATION_SET_ARTICLES,
  ARTICLE_ACTION_GET_ARTICLES
} from '../types';

import dayjs from 'dayjs';

import { getArticles, getArticlesByFilter } from '@/network/api/article';

const store = {
  state: {
    articles: []
  },
  getters: {
    articles: state => state.articles
  },
  mutations: {
    [ARTICLE_MUTATION_SET_ARTICLES]: (state, payload) => {
      state.articles = (payload || []).map(article => ({
        ...article,
        created_at: dayjs(article.created_at).format('MMM DD, YYYY')
      }));
    }
  },
  actions: {
    [ARTICLE_ACTION_GET_ARTICLES]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        let req = null;
        if (payload.filter) {
          req = getArticlesByFilter(
            payload.filter,
            payload.id,
            payload.page_no,
            payload.page_size
          );
        } else {
          req = getArticles(payload.page_no, payload.page_size);
        }

        req
          .then(res => {
            commit(ARTICLE_MUTATION_SET_ARTICLES, res.data);
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
