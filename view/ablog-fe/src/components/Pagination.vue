<template>
  <div class="pagination">
    <div class="pages">
      <span class="page-arrow page-arrow-prev" @click="handlePageArrowClick('prev')" v-if="current > 1">&lt;</span>
      <span :class="{
          'page-item': true,
          'active': page.active,
          'page-item-ellipsis': page.isEllipsis
        }" v-for="(page, index) in pages" :key="'page-item-' + index" @click="handlePageClick(page.index)">
        {{page.isEllipsis ? '...' : page.index}}
      </span>
      <span class="page-arrow page-arrow-next" @click="handlePageArrowClick('next')" v-if="current < totalNo">&gt;</span>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    total: {
      type: Number,
      default: 0
    },
    current: {
      type: Number,
      default: 1
    },
    size: {
      type: Number,
      default: 8
    },
    span: {
      type: Number,
      default: 1
    },
    onPageArrowClick: {
      type: Function,
      default(arrow) {
        console.log(arrow);
      }
    },
    onPageClick: {
      type: Function,
      default(page) {
        console.log(page);
      }
    }
  },
  data() {
    return {};
  },
  methods: {
    handlePageClick(page) {
      this.onPageClick(page);
    },
    handlePageArrowClick(arrow) {
      this.onPageArrowClick(arrow);
    }
  },
  computed: {
    totalNo() {
      const _totalNo = this.total === 0 ? 1 : parseInt(this.total / this.size);
      const totalNo = this.total % this.size === 0 ? _totalNo : _totalNo + 1;
      return totalNo;
    },

    pages() {
      const totalNo = this.totalNo

      let res = [];
      for (let i = 1; i <= totalNo; i++) {
        var t = {
          active: i === this.current,
          index: i
        };

        if (
          i === 1 ||
          i === totalNo ||
          (i >= this.current - this.span && i <= this.current + this.span)
        ) {
          res.push(t);
        } else if (
          (i === this.current - this.span - 1 && i !== 1) ||
          (i === this.current + this.span + 1 && i !== totalNo)
        ) {
          t.isEllipsis = true;
          res.push(t);
        }
      }

      return res;
    }
  }
};
</script>
<style lang="scss">
.pagination {
  .pages {
    display: flex;
    flex-flow: row wrap;
    align-items: center;
    justify-content: center;

    span {
      color: #333;
      display: inline-block;
      cursor: pointer;
      margin: 0 4px;
      width: 24px;
      height: 24px;
      line-height: 24px;
      text-align: center;
      border-top: 1px solid transparent;

      &.active,
      &:hover {
        color: $themeColor;
        border-color: $themeColor;
      }

      &.page-arrow {
        border: none;
      }

      &.page-item-ellipsis {
        border: none;
        cursor: default;
      }
    }
  }
}
</style>
