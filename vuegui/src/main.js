import Vue from 'vue'
import App from './App.vue'
import Axios from 'axios'
import VueLodash from 'vue-lodash'
import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/index.css';
import 'vue-tour/dist/vue-tour.css'
import { store } from './store'
import { router } from './services/router'
import VueJwtDecode from 'vue-jwt-decode'
import { authHeader } from './services/auth.header';
import moment from 'moment'
import VueNativeSock from 'vue-native-websocket'
import VueProgress from 'vue-progress-path'
import VueTour from 'vue-tour'
import GAuth from 'vue-google-oauth2'
import VueBlobJsonCsv from 'vue-blob-json-csv';

Vue.use(VueBlobJsonCsv)

const gauthOption = {
  clientId: '24121685362-5npi93os3dg3ufl9e0j3tnfoistippp3',
  scope: 'profile email',
  prompt: 'select_account'
}
Vue.use(GAuth, gauthOption)
Vue.use(VueTour)
Vue.use(require('vue-moment'));
Vue.use(VueProgress)
Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('MM/DD/YYYY hh:mm')
  }
});
Vue.use(VueNativeSock, 'wss://localhost:9090', {
  connectManually: true,
  store: store,
  format: 'json',
  reconnection: true, // (Boolean) whether to reconnect automatically (false)
  reconnectionDelay: 3000, // (Number) how long to initially wait before attempting a new (1000)
 }
)
Vue.use(VueJwtDecode)
Vue.use(VueToast, {
  // One of options
  position: 'top-right',
})
const options = { name: 'lodash' } // customize the way you want to call it


Vue.use(VueLodash, options) // options is optional
Vue.config.productionTip = false
Vue.prototype.$http = Axios;

// Add a request interceptor
Vue.prototype.$http.interceptors.request.use(function (config) {
  config.headers.Authorization = authHeader();
  let user = JSON.parse(localStorage.getItem('user'));
  if (user.signIn) {
    config.headers.SignIn = user.signIn
  }
  return config;
});
Vue.prototype.$http.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    // token expired
    // TODO handle refresh token, however this solution is still apreciable since the token has a lifespan of 1h , it allow auto logout
    if (error.response.status === 401) {
        router.push('/login')
    } else {
      return Promise.reject(error);
  }
  }
);

new Vue({
  router,
  store,
  mounted() {

  },
  render: h => h(App),
}).$mount('#app')
