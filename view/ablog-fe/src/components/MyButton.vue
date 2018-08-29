<template>
  <button :class="className" type="button" @click="click">
    <icon v-if="icon" :name="icon" />
    <slot></slot>
  </button>
</template>
<script>
import Icon from './Icon.vue';

export default {
  props: {
    click: {
      type: Function,
      default(e) {
        console.log(`mybutton: click -> ${e}`);
      }
    },
    icon: {
      type: String
    },
    type: {
      type: String
    },
    mode: {
      type: String,
      default: 'default'
    },
    size: {
      type: String,
      default: 's'
    }
  },
  computed: {
    className() {
      let cn = { button: true };

      this.type && (cn['button-' + this.type] = true);
      this.mode && (cn['button-' + this.mode] = true);
      this.size && (cn['button-' + this.size] = true);

      return cn;
    }
  },
  components: {
    Icon
  }
};
</script>
<style lang="scss">
.button {
  outline: none;
  border: none;
  color: #333;
  border-radius: 4px;
  height: 24px;
  font-size: 11px;

  transition: all 0.3s;

  &.button-m {
    height: 36px;
    font-size: 16px;
    &.button-round {
      border-radius: 18px;
    }
  }

  &.button-l {
    height: 48px;
    font-size: 24px;
    &.button-round {
      border-radius: 24px;
    }
  }

  &.button-blank {
    background-color: transparent;
    &:hover {
      color: #666;
    }

    &:active {
      color: #000;
    }
  }

  &.button-primary {
    background-color: #1e90ff;
    color: white;

    &:hover {
      background-color: #2fa1ff;
    }

    &:active {
      background-color: #0d80ee;
    }
  }

  &.button-default {
    background-color: #ccc;
    &:hover {
      background-color: #ddd;
    }

    &:active {
      background-color: #bbb;
    }
  }

  &.button-warn {
    background-color: #ffa500;
    color: white;

    &:hover {
      background-color: #ffb611;
    }

    &:active {
      background-color: #ee9400;
    }
  }

  &.button-danger {
    background-color: #dc143c;
    color: white;

    &:hover {
      background-color: #ed254d;
    }

    &:active {
      background-color: #cb032b;
    }
  }

  &.button-rect {
    border-radius: 0;
  }

  &.button-round {
    border-radius: 12px;
  }
}
</style>
