<template>
  <div class="evaluate-box">

    <div class="evaluate-box-input">
      <textarea v-model="content" :rows="rows" placeholder="请输入评论内容 ( •̀ ω •́ )y"></textarea>

      <span class="limit">{{content.length}} / {{limit}}</span>
    </div>
    <div class="evaluate-box-tools">

      <score ref="score" />
      <my-button :click="handleSubmit" mode="primary" size="m">提交</my-button>

    </div>
  </div>
</template>

<script>
export default {
  props: {
    rows: {
      type: Number,
      default: 5
    },
    limit: {
      type: Number,
      default: 256
    },
    submit: {
      type: Function,
      default(form) {
        console.log('evaluate-box: ' + JSON.stringify(form));
      }
    }
  },
  data() {
    return {
      content: ''
    };
  },
  methods: {
    handleSubmit(e) {
      if (this.content.length > this.limit) {
        return;
      }

      this.submit({
        score: this.$refs.score.value,
        content: this.content
      });
    }
  }
};
</script>


<style lang="scss">
.evaluate-box {
  .evaluate-box-input {
    position: relative;
    textarea {
      border-radius: 4px;
      width: 100%;
      padding: 8px;
      margin: 0;
      box-sizing: border-box;
      outline: none;
      font-size: 14px;
      resize: none;
      color: #333;
      border: 1px solid #ccc;

      transition: all 0.3s;
      &::-webkit-scrollbar {
        width: 4px;
      }
      &::-webkit-scrollbar-thumb {
        background-color: #aaa;
      }
      &::-webkit-scrollbar-track {
        background-color: #ccc;
      }

      &:focus {
        border-color: $themeColor;
      }
    }

    .limit {
      position: absolute;
      right: 0;
      bottom: 0;
      color: #aaa;
      padding: 8px;
    }
  }

  .evaluate-box-tools {
    margin-top: 8px;
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: flex-end;
    .score {
      margin-right: 16px;
      font-size: 24px;
      color: $themeColor;
    }

    .button {
      padding-left: 12px;
      padding-right: 12px;
    }
  }
}
</style>
