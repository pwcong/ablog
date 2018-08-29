<template>
  <div :class="{
    sidebar: true,
    ['sidebar-align-' + align]: true,
    active: isActive
  }">

    <span class="sidebar-toggle" @click="handleToggle">
      <icon :name="isActive ? 'close' : (align === 'left' ? 'right' : 'left')" />
    </span>

    <slot></slot>
  </div>
</template>

<script>
export default {
  props: {
    active: {
      type: Boolean,
      default: false
    },
    align: {
      type: String,
      default: 'right'
    }
  },
  data() {
    return {
      isActive: this.active
    };
  },
  methods: {
    handleToggle() {
      this.isActive = !this.isActive;
    }
  },
  mounted() {
    this.$on('sidebar-toggle', toggle => {
      this.isActive = toggle;
    });
  }
};
</script>

<style lang="scss">
.sidebar {
  position: fixed;
  top: 0;
  bottom: 0;

  background-color: rgba(0, 0, 0, 0.8);
  transition: all 0.3s;

  .sidebar-toggle {
    position: absolute;
    font-size: 24px;
    color: white;
    top: 50%;
    transform: translateY(-50%);
    display: inline-block;
    width: 42px;
    height: 52px;
    text-align: center;
    line-height: 52px;
    cursor: pointer;
    background-color: rgba(0, 0, 0, 0.2);
    transition: background-color 0.3s;
    &:hover {
      background-color: rgba(0, 0, 0, 0.8);
    }
  }

  &.sidebar-align-left {
    left: 0;
    transform: translate(-101%);

    .sidebar-toggle {
      border-top-right-radius: 50%;
      border-bottom-right-radius: 50%;
      right: -42px;
    }
  }

  &.sidebar-align-right {
    right: 0;
    transform: translate(101%);

    .sidebar-toggle {
      border-top-left-radius: 50%;
      border-bottom-left-radius: 50%;
      left: -42px;
    }
  }

  &.active {
    transform: translate(0);

    .sidebar-toggle {
      background-color: rgba(0, 0, 0, 0.8);
    }
  }
}
</style>
