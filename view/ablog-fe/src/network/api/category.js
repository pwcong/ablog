import request from '../request';

import api from './config';

export function getCategories() {
  return request(
    api.category.getCategories.url(),
    api.category.getCategories.method,
    {},
    {}
  );
}
export function getCategory(id) {
  return request(
    api.category.getCategory.url(id),
    api.category.getCategory.method,
    {},
    {}
  );
}
