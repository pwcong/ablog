<template>
  <div class="container">
    <toolbar style="padding: 32px 8px;" :app-name="appName" :navs="blogNavs" />

    <div class="tip" v-if="showTip">
      <div class="tip-content">
        {{tip}}
      </div>
      <my-button :click="handleBack" size="s" style="width: 64px;">返回</my-button>
    </div>

    <ul class="articles">

      <li class="article" v-for="(article, index) in articles" :key="'article-item-' + index">
        <span class="created_at">{{article.created_at}}</span>
        <h2 class="title">
          <router-link :to="'/blog/article/' + article.id">
            {{article.title}}
          </router-link>
        </h2>
      </li>

    </ul>
    <pagination style="margin: 16px 0;" :total="total" :current="current" :size="size" :on-page-click="handlePageClick" :on-page-arrow-click="handlePageArrowClick" />
  </div>
</template>

<script>
import { ARTICLE_ACTION_GET_ARTICLES } from '@/store/types';

import { APP_NAME } from '@/config';
import { BLOG_NAVS } from '@/const';

import { getCategory } from '@/network/api/category';
import { getTag } from '@/network/api/tag';

const PAGE_SIZE = 12;

export default {
  data() {
    return {
      total: 0,
      current: 1,
      size: PAGE_SIZE,
      hasMore: false,
      showTip: false,
      tip: '',
      appName: APP_NAME,
      blogNavs: BLOG_NAVS.map(nav => ({
        ...nav,
        active: nav.link === '/'
      }))
    };
  },
  methods: {
    loadArticles(search, filter, value, page_no) {
      const ctx = this;

      this.$store
        .dispatch(ARTICLE_ACTION_GET_ARTICLES, {
          search,
          filter,
          value,
          page_no,
          page_size: PAGE_SIZE
        })
        .then(res => {
          const { total_size, page_no, data } = res;
          ctx.total = total_size;
          ctx.current = page_no;

          const noData = (data || []).length <= 0;
          const { filter, search, value } = ctx.$route.params;

          if (filter) {
            ctx.showTip = true;
            if (noData) {
              ctx.tip = '没有啦/(ㄒoㄒ)/~~';
            } else if (search) {
              ctx.tip = `搜索结果如下`;
            } else {
              switch (filter) {
                case 'category':
                  getCategory(value)
                    .then(res => {
                      ctx.tip = `已选择分类“${res.name}”`;
                    })
                    .catch(err => {});
                  break;
                case 'tag':
                  getTag(value)
                    .then(res => {
                      ctx.tip = `已选择标签“${res.name}”`;
                    })
                    .catch(err => {});
                  break;
                default:
                  break;
              }
            }
          } else {
            ctx.showTip = false;
          }
        })
        .catch(err => {
          ctx.showTip = true;
          ctx.tip = err.message;
        });
    },
    handlePageClick(page) {
      const { search, filter, value } = this.$route.params;
      this.loadArticles(search, filter, value, page);
    },
    handlePageArrowClick(arrow) {
      const { search, filter, value } = this.$route.params;
      let t = this.current;
      if (arrow === 'prev') {
        this.loadArticles(search, filter, value, t - 1);
      } else if (arrow) {
        this.loadArticles(search, filter, value, t + 1);
      }
    },
    handleBack() {
      this.$router.back();
    }
  },
  computed: {
    articles() {
      return this.$store.getters.articles;
    }
  },

  mounted() {
    const { search, filter, value } = this.$route.params;
    this.loadArticles(search, filter, value, 1);
  }
};
</script>

<style lang="scss" scoped>
.container {
  .tip {
    max-width: 360px;
    padding: 24px 16px;
    box-sizing: border-box;
    margin: 0 auto;
    margin-bottom: 16px;
    border-radius: 4px;
    background-color: #f5f5f5;
    text-align: center;

    .tip-content {
      font-size: 14px;
      color: #666;
    }

    .button {
      font-size: 13px !important;
      margin-top: 8px;
      font-size: 18px;
      color: white;
      background-color: $themeColor;
    }
  }
  .articles {
    min-height: 66vh;
    max-width: 600px;
    padding: 0 16px;
    margin: 0 auto;
    list-style: none;

    .article {
      display: flex;
      align-items: center;
      flex-flow: row wrap;
      border-bottom: 1px solid #d5d5d5;
      .created_at {
        color: #999;
        font-size: 13px;
        margin-right: 12px;
      }

      .title {
        min-width: 250px;
        font-size: 22px;
        font-weight: normal;
        flex: 1;
        a {
          color: #333;
          text-decoration: none;
          &:hover {
            color: $themeColor;
          }
        }
      }
    }
  }
}
</style>
