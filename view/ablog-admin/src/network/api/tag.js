import request from '../request';

import api from './config';

export function getTags() {
  return request(api.tag.getTags.url(), api.tag.getTags.method, {}, {});
}

export function getTag(id) {
  return request(api.tag.getTag.url(id), api.tag.getTag.method, {}, {});
}
