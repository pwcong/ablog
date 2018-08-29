<template>
  <li class="toctree-item">
    <div class="toctree-item-name" @click="onItemClick(model.link)">{{model.name}}</div>
    <ul class="toctree-item-children" v-if="model.children && model.children.length > 0">
      <toctree-item v-for="(model, index) in model.children" :key="'toctree-item-' + index" :model="model" />
    </ul>
  </li>
</template>
<script>
export default {
  props: {
    model: {
      type: Object
    },
    onItemClick: {
      type: Function,
      default(link) {
        console.log(`onItemClick: ${link}`);
      }
    }
  }
};
</script>

<style>
</style>
