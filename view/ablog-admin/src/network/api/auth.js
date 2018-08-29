import request from '../request';

import api from './config';

export function login(username, password) {
  return request(
    api.auth.login.url(),
    api.auth.login.method,
    api.auth.login.data(username, password),
    {}
  );
}
