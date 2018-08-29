import Vue from 'vue';
import axios from 'axios';
import Cookies from 'js-cookie';

import router from '@/routes';
import { AUTH_COOKIE_NAME } from '@/config';

const HEADERS = {};

export default function(url, method, data, headers) {
  Vue.prototype.$progress.start();

  return new Promise((resolve, reject) => {
    axios({
      url,
      method,
      data,
      headers: Object.assign({}, HEADERS, headers)
    })
      .then(res => {
        if (res.status === 200 && res.data && res.data.code === 20000) {
          Vue.prototype.$progress.finish();
          resolve(res.data.payload);
        } else {
          if (res.status === 401) {
            Cookies.remove(AUTH_COOKIE_NAME);
            console.log(router);
          }
          Vue.prototype.$progress.stop();
          reject(res);
        }
      })
      .catch(err => {
        Vue.prototype.$progress.stop();
        reject(err);
      });
  });
}
