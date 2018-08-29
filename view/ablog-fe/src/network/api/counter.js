import request from '../request';

import api from './config';

export function getCounter(key, id) {
  return request(
    api.counter.getCounter.url(key, id),
    api.counter.getCounter.method,
    {},
    {}
  );
}

export function plusCounter(key, id) {
  return request(
    api.counter.plusCounter.url(),
    api.counter.plusCounter.method,
    api.counter.plusCounter.data(key, id),
    {}
  );
}
