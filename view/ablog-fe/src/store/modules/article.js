import {
  ARTICLE_MUTATION_SET_ARTICLES,
  ARTICLE_ACTION_GET_ARTICLES,
  ARTICLE_MUTATION_SET_ARTICLE,
  ARTICLE_ACTION_GET_ARTICLE,
  ARTICLE_MUTATION_SET_ARTICLES_COUNTER,
  ARTICLE_ACTION_PLUS_ARTICLE_COUNTER
} from '../types';

import dayjs from 'dayjs';

import {
  getArticles,
  getArticlesByFilter,
  getArticle,
  searchArticles
} from '@/network/api/article';

import { plusCounter } from '@/network/api/counter';

const store = {
  state: {
    articles: [],
    article: {},
    counter: {}
  },
  getters: {
    articles: state => state.articles,
    article: state => state.article,
    articleCounter: state => state.counter
  },
  mutations: {
    [ARTICLE_MUTATION_SET_ARTICLES]: (state, payload) => {
      state.articles = (payload || []).map(article => ({
        ...article,
        created_at: dayjs(article.created_at).format('MMM DD, YYYY')
      }));
    },
    [ARTICLE_MUTATION_SET_ARTICLE]: (state, payload) => {
      state.article = {
        ...payload,
        created_at: dayjs(payload.created_at).format('YYYY-MM-DD hh:mm:ss')
      };
    },
    [ARTICLE_MUTATION_SET_ARTICLES_COUNTER]: (state, payload) => {
      state.counter = payload;
    }
  },
  actions: {
    [ARTICLE_ACTION_GET_ARTICLES]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        const { filter, value, page_no, page_size, search } = payload;

        let req = null;
        if (filter) {
          if (search) {
            req = searchArticles(filter, value, page_no, page_size);
          } else {
            req = getArticlesByFilter(filter, value, page_no, page_size);
          }
        } else {
          req = getArticles(page_no, page_size);
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
    },
    [ARTICLE_ACTION_GET_ARTICLE]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        const { id } = payload;

        getArticle(id)
          .then(res => {
            commit(ARTICLE_MUTATION_SET_ARTICLE, res);
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    },
    [ARTICLE_ACTION_PLUS_ARTICLE_COUNTER]: (
      { commit, state },
      payload = {}
    ) => {
      return new Promise((resolve, reject) => {
        const { id } = payload;

        plusCounter('article', parseInt(id))
          .then(res => {
            commit(ARTICLE_MUTATION_SET_ARTICLES_COUNTER, res);
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
