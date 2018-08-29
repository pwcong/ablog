<template>
  <ul class="toctree">
    <toctree-item v-for="(model, index) in treeData.children" :key="'toctree-item-' + index" :model="model"></toctree-item>
  </ul>
</template>
<script>
export default {
  props: {
    treeData: {
      type: Object,
      default: {}
    }
  }
};
</script>

<style lang="scss">
.toctree {
  .toctree-item-name {
    cursor: pointer;
    font-size: 13px;
    display: inline-block;
    padding: 4px;
    border-bottom: 1px solid transparent;
    transition: all 0.3s;
    &:hover {
      color: $themeColor;
      border-bottom-color: $themeColor;
    }
  }

  .toctree-item-children {
    padding-left: 16px;
  }
}
</style>

