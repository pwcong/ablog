import request from '../request';

import api from './config';

export function getEvaluations(id, page_no, page_size) {
  return request(
    api.evaluation.getEvaluations.url(id, page_no, page_size),
    api.evaluation.getEvaluations.method,
    {},
    {}
  );
}
