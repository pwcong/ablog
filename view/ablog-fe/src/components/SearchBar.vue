<template>
  <div :class="className">
    <input :placeholder="placeholder" @keyup.enter="handleSearch" v-model="text" class="searchbar-input" type="text" @focus="active = true" @blur="active = false">
    <my-button mode="blank" class="searchbar-submit" icon="search" size="l" :click="handleSearch"></my-button>
  </div>
</template>

<script>
import MyButton from './MyButton.vue';

export default {
  props: {
    search: {
      type: Function,
      default(text) {
        console.log(`searchbar: ${text}`);
      }
    },
    placeholder: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      active: false,
      text: ''
    };
  },
  methods: {
    handleSearch() {
      this.search(this.text);
    }
  },
  computed: {
    className() {
      return {
        searchbar: true,
        active: this.active
      };
    }
  },
  components: {
    MyButton
  }
};
</script>

<style lang="scss">
.searchbar {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  font-size: 13px;
  .searchbar-input {
    padding: 4px 0;
    box-sizing: border-box;
    background-color: transparent;
    width: 0;
    border: none;
    outline: none;
    border-bottom: 1px solid #999;
    transition: all 0.3s;
  }

  &.active,
  &:hover {
    .searchbar-input {
      width: 150px;
      border-color: $themeColor;
    }
    .searchbar-submit {
      color: $themeColor;
    }
  }
}
</style>
