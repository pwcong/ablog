<template>
  <div class="container">

    <my-header :class="{
      hide: hideNavs
    }">
      <navbar :navs="navs"></navbar>
      <searchbar placeholder="搜索标题" :search="handleSearch"></searchbar>
    </my-header>

    <div class="main">

      <div v-if="banner" class="banner" :style="{
          'background-image': 'url(' + banner + ')'
        }"></div>

      <div class="article">
        <div class="tip" v-if="causeError">
          <icon name="error"></icon>
          <span>{{errorMsg}}</span>
        </div>
        <div class="head" v-if="title">
          <h1>{{title}}</h1>
          <span v-if="count">阅读次数：{{count}}</span>
          <span v-if="created_at">发布时间：{{created_at}}</span>
        </div>

        <div ref="article" v-if="content" class="markdown-section content" v-html="content"></div>

      </div>

    </div>

    <div class="evaluations" v-if="!causeError">

      <div class="evaluations-head">
        评论列表（{{evaluationsLength}}）
        <my-button icon="sync" style="float: right; font-size: 18px;" mode="blank" :click="reloadEvaluations"></my-button>
      </div>

      <div class="evaluations-wrapper">

        <div class="evaluations-item" v-for="(evaluation, idx) in evaluations" :key="'evaluations-item-' + idx + '-' + evaluation.created_at">

          <div class="left">
            <canvas width="48" height="48" :data-jdenticon-value="evaluation.ip"></canvas>
          </div>
          <div class="right">
            <div class="content">
              {{evaluation.content}}
            </div>
            <div class="others">
              <span class="time">{{evaluation.created_at}}</span>
              <score :editable="false" :default-value="evaluation.score" />
            </div>
          </div>

        </div>
        <load-more style="padding: 16px;" :has-more="hasMoreEvaluations" :tip="loadMoreEvaluationsTip" :more="loadMoreEvaluations"></load-more>
      </div>

      <evaluate-box :submit="handleSubmitEvaluation"></evaluate-box>
    </div>

    <sidebar align="right">
      <toctree :tree-data="treeData" />
    </sidebar>

    <my-footer></my-footer>

  </div>
</template>

<script>
import {
  ARTICLE_ACTION_GET_ARTICLE,
  ARTICLE_ACTION_PLUS_ARTICLE_COUNTER,
  EVALUATION_MUTATION_CLEAR_EVALUATIONS,
  EVALUATION_ACTION_GET_EVALUATIONS,
  EVALUATION_ACTION_PUBLISH_EVALUATIONS
} from '@/store/types';

import { BASE_URL } from '@/config';
import { BLOG_NAVS } from '@/const';

import m2h from 'm2h';
import dayjs from 'dayjs';
import jdenticon from 'jdenticon';

import mdtoc from 'mdtoc.js';

const EVALUATIONS_PAGE_SIZE = 6;

export default {
  data() {
    return {
      navs: BLOG_NAVS,
      hideNavs: false,

      causeError: false,
      errorMsg: '',
      countDown: 3,

      evaluationsLength: 0,
      loadMoreEvaluationsPageNo: 1,
      hasMoreEvaluations: false,
      loadMoreEvaluationsTip: '',
      treeData: {}
    };
  },

  methods: {
    handleSubmitEvaluation(form) {
      const { id } = this.$route.params;
      const { content, score } = form;
      if (!content) {
        this.$toast.tip('评论内容不能为空');
        return;
      }

      this.$toast.process('提交中');
      this.$store
        .dispatch(EVALUATION_ACTION_PUBLISH_EVALUATIONS, {
          id,
          content,
          score
        })
        .then(res => {
          this.$toast.hide();
          this.reloadEvaluations();
        })
        .catch(err => {
          this.$toast.hide();
        });
    },
    loadArticle(id) {
      this.$store
        .dispatch(ARTICLE_ACTION_GET_ARTICLE, {
          id
        })
        .then(article => {
          this.plusCounter(id);

          this.treeData = mdtoc(this.$refs.article).get();
        })
        .catch(err => {
          this.causeError = true;
          this.errorMsg = '文章不存在或已被删除(っ °Д °;)っ';
        });
    },
    reloadEvaluations() {
      this.loadMoreEvaluationsPageNo = 1;
      const { id } = this.$route.params;
      this.$store.commit(EVALUATION_MUTATION_CLEAR_EVALUATIONS);
      this.loadEvaluations(id);
    },
    loadEvaluations(id) {
      this.$store
        .dispatch(EVALUATION_ACTION_GET_EVALUATIONS, {
          id,
          page_no: this.loadMoreEvaluationsPageNo,
          page_size: EVALUATIONS_PAGE_SIZE
        })
        .then(res => {
          const { data, page_size, total_size } = res;

          this.evaluationsLength = total_size;

          if (total_size <= 0) {
            this.hasMoreEvaluations = false;
            this.loadMoreEvaluationsTip = '暂无评论';
          } else if ((data || []).length < page_size) {
            this.hasMoreEvaluations = false;
            this.loadMoreEvaluationsTip = '没有啦(。・∀・)ノ';
          } else {
            this.hasMoreEvaluations = true;
          }
        })
        .catch(err => {
          this.hasMoreEvaluations = false;
          this.loadMoreEvaluationsTip = '评论加载失败 (っ °Д °;)っ';
        });
    },
    loadMoreEvaluations() {
      this.loadMoreEvaluationsPageNo++;
      const { id } = this.$route.params;
      this.loadEvaluations(id);
    },
    plusCounter(id) {
      this.$store
        .dispatch(ARTICLE_ACTION_PLUS_ARTICLE_COUNTER, {
          id
        })
        .then(counter => {})
        .catch(err => {});
    },
    handleSearch(text) {
      if (!text) {
        return;
      }
      this.$router.push(`/blog/articles/title/${text}/search`);
    }
  },

  computed: {
    banner() {
      const article = this.$store.getters.article;
      const banner = article.banner || '';
      return banner
        ? /$(http:\/\/|https:\/\/)/.test(banner) ? banner : BASE_URL + banner
        : '';
    },
    title() {
      const article = this.$store.getters.article;
      return article.title || '';
    },
    created_at() {
      const article = this.$store.getters.article;
      return article.created_at
        ? dayjs(article.created_at || '').format('YYYY-MM-DD hh:mm:ss')
        : '';
    },
    content() {
      const article = this.$store.getters.article;
      return article.content ? m2h(article.content) : '';
    },
    count() {
      const counter = this.$store.getters.articleCounter;
      return counter.count || '';
    },
    evaluations() {
      return this.$store.getters.evaluations || [];
    },
    articleTree() {
      return null;
    }
  },

  components: {},

  mounted() {
    const { id } = this.$route.params;
    this.loadArticle(id);
    this.reloadEvaluations(id);

    this.$on('root-scroll', res => {
      if (res.scrollY > 88) {
        if (res.direction === 'up') {
          this.hideNavs = false;
        } else {
          this.hideNavs = true;
        }
      } else {
        this.hideNavs = false;
      }
    });

    setInterval(() => {
      jdenticon();
    }, 1000);
  }
};
</script>

<style lang="scss" scoped>
@media screen and(min-width: 980px) {
  .main,
  .evaluations {
    max-width: 980px;
  }
}

.container {
  background-color: #eee;
  min-height: 100vh;
  padding-top: 88px;
  box-sizing: border-box;

  .header {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    z-index: 9999;
    transition: all 0.3s;
    transition-delay: 0.1s;

    .header-main {
      .navbar {
        height: 100%;
        float: left;
      }

      .searchbar {
        height: 100%;
        float: right;
      }
    }

    &.hide {
      top: -88px;
    }
  }

  .main {
    margin: 0 auto;
    background-color: white;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    .banner {
      height: 256px;
      background-size: cover;
      background-position: center;
      background-repeat: no-repeat;
    }

    .article {
      box-sizing: border-box;
      padding: 16px;
      max-width: 800px;
      margin: 0 auto;
      min-height: 72vh;

      .tip {
        text-align: center;
        padding-top: 10vh;
        font-size: 18px;
        color: #666;
      }

      .head {
        position: relative;
        padding-bottom: 24px;
        h1 {
          color: #333;
          margin-bottom: 12px;
        }
        span {
          font-size: 11px;
          color: #666;
          &:not(:last-child) {
            margin-right: 16px;
          }
        }

        &::after {
          content: '';
          position: absolute;
          bottom: 0;
          left: 0;
          width: 100%;
        }
      }
      .content {
        border-top: 1px solid #666;
        font-size: 85%;
      }
    }
  }
  .evaluations {
    box-sizing: border-box;
    background-color: white;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
    margin: 0 auto;
    margin-top: 16px;
    padding: 16px;

    .evaluations-head {
      color: #333;
      padding-bottom: 12px;
      border-bottom: 1px solid #aaa;
    }

    .evaluations-wrapper {
      padding: 16px;

      .evaluations-item {
        display: flex;
        flex-flow: row nowrap;
        align-items: flex-start;
        padding: 8px 0;
        .left {
          width: 64px;
          text-align: center;
        }

        .right {
          flex: 1;

          .content {
            padding-top: 6px;
            font-size: 13px;
            color: #333;
          }

          .others {
            display: flex;
            flex-flow: row nowrap;
            align-items: center;
            justify-content: flex-end;
            .time {
              margin-right: 8px;
              font-size: 13px;
              color: #666;
            }
            .score {
              color: $themeColor;
            }
          }
        }

        &:not(:last-child) {
          border-bottom: 1px dashed #ccc;
        }
      }
    }
  }

  .sidebar {
    padding: 12px;
    display: flex;
    flex-flow: column nowrap;
    align-items: center;
    justify-content: center;
    color: white;
    width: 240px;
  }

  .footer {
    margin-top: 24px;
  }
}
</style>
