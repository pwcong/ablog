<template>
  <footer class="footer">
    &copy; {{duration}} . {{sign}}
  </footer>
</template>
<script>
export default {
  props: {
    from: {
      type: Number,
      default: 2018
    },
    sign: {
      type: String,
      default: 'Powered By Pwcong'
    }
  },

  data() {
    const now = new Date();
    let years = [this.from];

    if (now.getFullYear() > this.from) {
      years.push(now.getFullYear());
    }

    return {
      duration: years.join(' - ')
    };
  }
};
</script>
<style lang="scss">
.footer {
  padding: 24px 16px;
  text-align: center;
  font-size: 14px;
  color: #999;
  cursor: default;
}
</style>
