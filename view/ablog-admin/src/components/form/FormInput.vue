<template>
  <div :class="{
    'form-input': true,
    'error': showError
  }">
    <input @input="handleInput" :placeholder="placeholder" :type="type">
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      required: true
    },
    required: {
      type: Boolean,
      default: true
    },
    placeholder: {
      type: String
    },
    type: {
      type: String,
      default: 'text'
    }
  },
  data() {
    return {
      showError: false
    };
  },
  methods: {
    handleInput(e) {
      const value = e.target.value;

      if (this.required && !value) {
        this.showError = true;
      } else {
        this.showError = false;
      }

      this.$emit('input', value);
    }
  }
};
</script>

<style lang="scss">
.form-input {
  display: inline-block;

  input {
    font-size: 14px;
    box-sizing: border-box;
    border: 1px solid #ccc;
    outline: none;
    width: 100%;
    border-radius: 4px;
    padding: 4px;
    transition: all 0.3s;

    &:focus {
      border-color: #aaa;
    }
  }

  &.error {
    input {
      border-color: $themeColor;
    }
  }
}
</style>
