import { BASE_URL } from '@/config';

export function resolveURL(url) {
  return BASE_URL + url;
}

export const DEFAULT_PAGE_NO = -1;
export const DEFAULT_PAGE_SIZE = -1;

export default {
  article: {
    getArticles: {
      url: (page_no = DEFAULT_PAGE_NO, page_size = DEFAULT_PAGE_SIZE) =>
        resolveURL(`/api/articles?page_no=${page_no}&page_size=${page_size}`),
      method: 'GET'
    },
    getArticlesByFilter: {
      url: (
        filter,
        value,
        page_no = DEFAULT_PAGE_NO,
        page_size = DEFAULT_PAGE_SIZE
      ) =>
        resolveURL(
          `/api/articles/${filter}/${value}?page_no=${page_no}&page_size=${page_size}`
        ),
      method: 'GET'
    },
    searchArticles: {
      url: (
        filter,
        value,
        page_no = DEFAULT_PAGE_NO,
        page_size = DEFAULT_PAGE_SIZE
      ) =>
        resolveURL(
          `/api/articles/search/${filter}/${value}?page_no=${page_no}&page_size=${page_size}`
        ),
      method: 'GET'
    },
    getArticle: {
      url: id => resolveURL(`/api/article/${id}`),
      method: 'GET'
    },
    evaluateArticle: {
      url: id => resolveURL(`/api/article/${id}/evaluate`),
      method: 'POST',
      data: (score, content) => ({
        score,
        content
      })
    },
    modifyArticle: {
      url: id => resolveURL(`/api/article/${id}`),
      method: 'POST',
      data: (title, content, banner, category_id, tag_ids) => ({
        title,
        content,
        banner,
        category_id,
        tag_ids
      })
    },
    addArticle: {
      url: () => resolveURL('/api/article'),
      method: 'POST',
      data: (title, content, banner, category_id, tag_ids) => ({
        title,
        content,
        banner,
        category_id,
        tag_ids
      })
    },
    delArticle: {
      url: id => resolveURL(`/api/article/${id}`),
      method: 'DELETE'
    }
  },
  tag: {
    getTags: {
      url: () => resolveURL(`/api/tags`),
      method: 'GET'
    },
    getTag: {
      url: id => resolveURL(`/api/tag/${id}`),
      method: 'GET'
    }
  },
  category: {
    getCategories: {
      url: () => resolveURL(`/api/categories`),
      method: 'GET'
    },
    getCategory: {
      url: id => resolveURL(`/api/category/${id}`),
      method: 'GET'
    }
  },
  counter: {
    getCounter: {
      url: (key, id) => resolveURL(`/api/counter/${key}/${id}`),
      method: 'GET'
    },
    plusCounter: {
      url: () => resolveURL(`/api/counter`),
      method: 'POST',
      data: (key, id) => ({
        key,
        id
      })
    }
  },
  evaluation: {
    getEvaluations: {
      url: (id, page_no, page_size) =>
        resolveURL(
          `/api/evaluations/${id}?page_no=${page_no}&page_size=${page_size}`
        ),
      method: 'GET'
    }
  }
};
