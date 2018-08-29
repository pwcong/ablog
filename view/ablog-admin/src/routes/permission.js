import Cookies from 'js-cookie';
import { AUTH_COOKIE_NAME } from '@/config';

export default function(router) {
  router.beforeEach((to, from, next) => {
    if (to.path === '/login') {
      next();
    } else {
      const token = Cookies.get(AUTH_COOKIE_NAME);
      if (!token) {
        next({
          path: '/login',
          query: {
            redirect: to.path
          }
        });
      } else {
        next();
      }
    }
  });
}
