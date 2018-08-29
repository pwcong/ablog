<template>
  <div class="container" ref="container">
    <canvas class="bg" ref="bg"></canvas>

    <div class="box">

      <div class="box-wrapper">
        <h1>{{appName}}</h1>
        <p>{{appDescription}}</p>
        <ul>
          <li>
            <router-link to="/blog">博客</router-link>
          </li>
          <li v-for="(nav, index) in indexNavs" :key="'index-nav-' + index">
            <a :href="nav.link">{{nav.name}}</a>
          </li>
        </ul>
      </div>

    </div>

  </div>
</template>

<script>
import { APP_NAME, APP_DESCRIPTION } from '@/config';
import { INDEX_NAVS } from '@/const';

export default {
  data() {
    return {
      indexNavs: INDEX_NAVS,
      appName: APP_NAME,
      appDescription: APP_DESCRIPTION
    };
  },

  mounted() {
    this.$refs.container.addEventListener('touchmove', function(e) {
      e.preventDefault();
    });

    var c = this.$refs.bg,
      x = c.getContext('2d'),
      pr = window.devicePixelRatio || 1,
      w = window.innerWidth,
      h = window.innerHeight,
      f = 90,
      q,
      m = Math,
      r = 0,
      u = m.PI * 2,
      v = m.cos,
      z = m.random;
    c.width = w * pr;
    c.height = h * pr;
    x.scale(pr, pr);
    x.globalAlpha = 0.6;
    function i() {
      x.clearRect(0, 0, w, h);
      q = [{ x: 0, y: h * 0.7 + f }, { x: 0, y: h * 0.7 - f }];
      while (q[1].x < w + f) d(q[0], q[1]);
    }
    function d(i, j) {
      x.beginPath();
      x.moveTo(i.x, i.y);
      x.lineTo(j.x, j.y);
      var k = j.x + (z() * 2 - 0.25) * f,
        n = y(j.y);
      x.lineTo(k, n);
      x.closePath();
      r -= u / -50;
      x.fillStyle =
        '#' +
        (
          ((v(r) * 127 + 128) << 16) |
          ((v(r + u / 3) * 127 + 128) << 8) |
          (v(r + u / 3 * 2) * 127 + 128)
        ).toString(16);
      x.fill();
      q[0] = q[1];
      q[1] = { x: k, y: n };
    }
    function y(p) {
      var t = p + (z() * 2 - 1.1) * f;
      return t > h || t < 0 ? y(p) : t;
    }
    document.onclick = i;
    document.ontouchstart = i;
    i();
  }
};
</script>

<style lang="scss" scoped>
.container {
  height: 100vh;
  overflow: hidden;
  box-sizing: border-box;
  position: relative;

  .bg,
  .box {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
  }

  .box {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;

    .box-wrapper {
      width: 400px;
      text-align: center;

      h1 {
        color: #333;
        font-size: 36px;
        margin: 0;
      }

      p {
        color: #999;
        margin-bottom: 36px;
      }

      ul {
        list-style: none;
        padding: 0;

        li {
          margin: 6px 0;
          a {
            padding: 2px;
            display: inline-block;
            color: #666;
            text-decoration: none;
            border-bottom: 1px solid transparent;
            &:hover {
              color: $themeColor;
              border-color: $themeColor;
            }
          }
        }
      }
    }
  }
}
</style>