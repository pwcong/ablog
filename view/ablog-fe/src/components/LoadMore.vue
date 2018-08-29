<template>
  <div class="load-more">
    <span class="load-more-tip" v-if="!hasMore">{{tip}}</span>
    <span class="load-more-action" v-if="hasMore" @click="more">{{actionTip}}</span>
  </div>
</template>
<script>
export default {
  props: {
    tip: {
      type: String,
      default: ''
    },
    loadingMore: {
      type: Boolean,
      default: false
    },
    hasMore: {
      type: Boolean,
      default: true
    },
    more: {
      type: Function,
      default() {
        console.log('load more');
      }
    }
  },

  action: {
    handleLoadMore() {
      if (this.hasMore && !this.loadingMore) {
        this.more();
      }
    }
  },

  computed: {
    actionTip() {
      if (this.loadingMore) {
        return '加载中';
      }

      return '加载更多';
    }
  }
};
</script>
<style lang="scss">
.load-more {
  text-align: center;
  font-size: 13px;
  .load-more-tip {
    color: #999;
  }
  .load-more-action {
    cursor: pointer;
    padding: 4px 8px;
    color: #666;
    transition: all 0.3s;
    &:active {
      color: $themeColor;
    }
  }
}
</style>
