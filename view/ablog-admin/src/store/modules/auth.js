import { AUTO_ACTION_LOGIN } from '../types';

import Cookies from 'js-cookie';
import { AUTH_COOKIE_NAME } from '@/config';
import { login } from '@/network/api/auth';

const store = {
  state: {},
  getters: {},
  mutations: {},
  actions: {
    [AUTO_ACTION_LOGIN]: ({ commit, state }, payload = {}) => {
      return new Promise((resolve, reject) => {
        const { username, password } = payload;

        login(username, password)
          .then(res => {
            Cookie.set(AUTH_COOKIE_NAME, res.token);
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
