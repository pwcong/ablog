<template>
  <div class="score">

    <div class="score-item" v-for="(score, idx) in options" :key="'score-item-' + idx" @click="handleChange(score)">
      <icon :name="score.active ? icon + '-fill' : icon"></icon>
    </div>

  </div>
</template>

<script>
export default {
  props: {
    editable: {
      type: Boolean,
      default: true
    },
    defaultValue: {
      type: Number,
      default: 0
    },
    max: {
      type: Number,
      default: 5
    },
    min: {
      type: Number,
      default: 1
    },
    step: {
      type: Number,
      default: 1
    },
    icon: {
      type: String,
      default: 'star'
    },
    change: {
      type: Function,
      default(value) {
        console.log('score: ' + value);
      }
    }
  },
  data() {
    return {
      value: this.defaultValue
    };
  },
  methods: {
    handleChange(score) {
      if (!this.editable) {
        return;
      }

      this.value = score.value;
      this.change(this.value);
    }
  },
  computed: {
    options() {
      var opts = [];

      for (let i = this.min; i <= this.max; i += this.step) {
        opts.push({
          active: this.value >= i,
          value: i
        });
      }

      return opts;
    }
  }
};
</script>

<style lang="scss">
.score {
  .score-item {
    cursor: pointer;
    display: inline-block;
  }
}
</style>
