import request from '../request';

import api from './config';

export function getArticles(page_no, page_size) {
  return request(
    api.article.getArticles.url(page_no, page_size),
    api.article.getArticles.method,
    {},
    {}
  );
}

export function getArticlesByFilter(filter, id, page_no, page_size) {
  return request(
    api.article.getArticlesByFilter.url(filter, id, page_no, page_size),
    api.article.getArticlesByFilter.method,
    {},
    {}
  );
}

export function getArticle(id) {
  return request(
    api.article.getArticle.url(id),
    api.article.getArticle.method,
    {},
    {}
  );
}

export function evaluateArticle(id, score, content) {
  return request(
    api.article.evaluateArticle.url(id),
    api.article.evaluateArticle.method,
    api.article.evaluateArticle.data(score, content),
    {}
  );
}

export function modifyArticle(
  id,
  title,
  content,
  banner,
  category_id,
  tag_ids
) {
  return request(
    api.article.modifyArticle.url(id),
    api.article.modifyArticle.method,
    api.article.modifyArticle.data(
      title,
      content,
      banner,
      category_id,
      tag_ids
    ),
    {}
  );
}

export function addArticle(title, content, banner, category_id, tag_ids) {
  return request(
    api.article.addArticle.url(),
    api.article.addArticle.method,
    api.article.addArticle.data(title, content, banner, category_id, tag_ids),
    {}
  );
}

export function delArticle(id) {
  return request(
    api.article.delArticle.url(id),
    api.article.delArticle.method,
    {},
    {}
  );
}
