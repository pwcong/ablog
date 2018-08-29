<template>
  <div class="login">
    <div class="login-box">
      <div class="login-box-logo">
        <img :src="logo" alt="LOGO">
      </div>
      <h1 class="login-box-title">登录至 {{appName}}</h1>
      <div class="login-box-form">
        <form-field>
          <form-label label="用户名" icon="user" />
          <form-input v-model="username" />
        </form-field>
        <form-field>
          <form-label label="密码" icon="lock" />
          <form-input v-model="password" type="password" />
        </form-field>
        <form-field>
          <my-button class="btn-submit" :disabled="showSync" :icon="showSync ? 'sync' : ''" icon-anim="anim-rotate 1.2s linear infinite">
            {{showSync ? '登陆中' : '登录'}}
          </my-button>
        </form-field>
      </div>
    </div>
  </div>
</template>

<script>
import { APP_NAME, APP_DESCRIPTION } from '@/config';
import logo from '@/assets/imgs/logo.png';

export default {
  data() {
    return {
      appName: APP_NAME,
      appDescription: APP_DESCRIPTION,
      logo,
      showSync: false,
      username: '',
      password: ''
    };
  }
};
</script>

<style lang="scss">
.login {
  min-height: 100vh;
  background-color: #eee;
  display: flex;
  align-items: center;
  justify-content: center;

  .login-box {
    position: relative;
    top: -48px;

    .login-box-logo {
      text-align: center;
    }
    .login-box-title {
      font-weight: 300;
      font-size: 22px;
      text-align: center;
    }

    .login-box-form {
      box-sizing: border-box;
      padding: 24px;
      min-width: 320px;
      background-color: white;
      border: 1px solid #ccc;
      border-radius: 4px;

      .form-field {
        &:not(:last-child) {
          margin-bottom: 8px;
        }
        .form-label {
          padding: 4px 0;
          i {
            margin-right: 4px;
          }
        }

        .form-input {
          width: 100%;
        }

        .btn-submit {
          background-color: $themeColor;
          color: white;
          width: 100%;
          margin-top: 16px;
          padding: 8px;
          height: auto;

          &:hover {
            background-color: #ff5611;
          }
        }
      }
    }
  }
}
</style>