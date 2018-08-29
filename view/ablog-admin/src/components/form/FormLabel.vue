<template>
  <div class="form-label">
    <icon :name="icon" v-if="icon" />{{label}}
  </div>
</template>

<script>
import Icon from '../Icon.vue';

export default {
  props: {
    label: {
      type: String
    },
    icon: {
      type: String
    }
  },
  components: {
    Icon
  }
};
</script>
<style lang="scss">
</style>
