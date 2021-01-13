  
import Vue from 'vue';
import Router from 'vue-router';

import LoginPage from '../components/LoginPage.vue'
import Dashboard from '../components/dashboard/myTemplate.vue'
import ProjectDashboard from '../components/project/ProjectDashboard.vue'
import Details from '../components/details/Details.vue'
import QuickConfiguration from '../components/project/quickConfiguration/QuickConfiguration_new.vue'
import AdminDashboard from '../components/admin/AdminDashboard.vue'
import NotificationCenter from '../components/notification-center/NotificationCenter.vue' 
import Browser from '../components/browser/Browser.vue'
Vue.use(Router);

export const router = new Router({
  mode: 'history',
  base: '/',
  routes: [
    { path: '/', component: ProjectDashboard,props: true },
    { path: '/dashboard', component: Dashboard,props: true },
    { path: '/details/:chart', component: Details, props: true},
    { path: '/login', component: LoginPage },
    { path: '/quick/:project', component: QuickConfiguration , props: true},
    { path: '/admin', component: AdminDashboard },
    { path: '/notification-center', component: NotificationCenter },
    { path: '/browse', component: Browser},
    // otherwise redirect to home
    { path: '*', redirect: '/' }
  ]
});

router.beforeEach((to, from, next) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = JSON.parse(localStorage.getItem('user'));
  const configid = localStorage.getItem('configid');
  if (authRequired && !loggedIn) {
    return next('/login');
  }
  if( to.path === '/admin' || to.path === "/dashboard"){
    console.log(loggedIn.id)
    if (!loggedIn.id || !configid || loggedIn.id != configid ){
      return next('/login');
    }
  }
  next();
})