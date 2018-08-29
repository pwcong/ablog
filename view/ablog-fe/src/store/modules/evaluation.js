import {
  EVALUATION_ACTION_GET_EVALUATIONS,
  EVALUATION_ACTION_PUBLISH_EVALUATIONS,
  EVALUATION_MUTATION_APPEND_EVALUATIONS,
  EVALUATION_MUTATION_CLEAR_EVALUATIONS
} from '../types';

import dayjs from 'dayjs';

import { evaluateArticle } from '@/network/api/article';

import { getEvaluations } from '@/network/api/evaluation';

const store = {
  state: {
    evaluations: []
  },
  getters: {
    evaluations: state => state.evaluations
  },
  mutations: {
    [EVALUATION_MUTATION_APPEND_EVALUATIONS]: (state, payload) => {
      state.evaluations = (state.evaluations || []).concat(
        (payload || []).map(evaluation => ({
          ...evaluation,
          created_at: dayjs(evaluation.created_at).format('YYYY-MM-DD hh:mm')
        }))
      );
    },
    [EVALUATION_MUTATION_CLEAR_EVALUATIONS]: (state, payload) => {
      state.evaluations = [];
    }
  },
  actions: {
    [EVALUATION_ACTION_GET_EVALUATIONS]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        const { id, page_no, page_size } = payload;

        getEvaluations(id, page_no, page_size)
          .then(res => {
            commit(EVALUATION_MUTATION_APPEND_EVALUATIONS, res.data);
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    },
    [EVALUATION_ACTION_PUBLISH_EVALUATIONS]: (
      { commit, state, dispatch },
      payload = {}
    ) => {
      return new Promise((resolve, reject) => {
        const { id, score, content } = payload;

        evaluateArticle(id, score, content)
          .then(res => {
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
