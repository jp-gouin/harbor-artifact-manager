<template>
  <div id="app" class="bg-gray-100 font-sans leading-normal tracking-normal" v-bind:class="loggingIn?'mt-12':'mt-0'">
    <Navigation v-if="loggingIn" :user="user.username" />
    <div class="flex flex-col md:flex-row">
      <div v-if="loggingIn" class="bg-gray-900 h-16 fixed bottom-0 mt-12 md:h-screen z-30 w-full md:w-48">
        <div
          class="md:mt-12 md:w-48 md:fixed md:left-0 md:top-0 content-center md:content-start text-left justify-between"
        >
          <ul
            class="list-reset flex flex-row md:flex-col py-0 md:py-3 px-1 md:px-2 text-center md:text-left"
          >
          <li class="mr-3 flex-1">
              <router-link
                to="/browse"
                class="block py-1 md:py-3 pl-1 align-middle text-white no-underline hover:text-white border-b-2 border-gray-800 hover:border-pink-500"
              >
                <i class="fas fa-tasks pr-0 md:pr-3"></i>
                <span
                  class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
                >Browse the catalog</span>
              </router-link>
            </li>
            <li class="mr-3 flex-1">
              <router-link v-if="user && user.id && config.id && user.id == config.id"
                to="/dashboard"
                class="block py-1 md:py-3 pl-1 align-middle text-white no-underline hover:text-white border-b-2 border-gray-800 hover:border-pink-500"
              >
                <i class="fas fa-tasks pr-0 md:pr-3"></i>
                <span
                  class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
                >Dashboard</span>
              </router-link>
            </li>
            <li class="mr-3 flex-1">
              <router-link
                to="/"
                class="block py-1 md:py-3 pl-1 align-middle text-white no-underline hover:text-white border-b-2 border-gray-800 hover:border-purple-500"
              >
                <i class="fa fa-envelope pr-0 md:pr-3"></i>
                <span
                  class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
                >Projects</span>
              </router-link>
            </li>
            <li class="mr-3 flex-1">
              <router-link v-if="user && user.id && config.id && user.id == config.id"
                to="/admin"
                class="block py-1 md:py-3 pl-1 align-middle text-white no-underline hover:text-white border-b-2 border-gray-800 hover:border-purple-500"
              >
                <i class="fa fa-envelope pr-0 md:pr-3"></i>
                <span
                  class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
                >Administration</span>
              </router-link>
            </li>
            <li class="mr-3 flex-1">
              <router-link
                to="/notification-center"
                class="block py-1 md:py-3 pl-1 align-middle text-white no-underline hover:text-white border-b-2 border-gray-800 hover:border-purple-500"
              >
                <i class="fa fa-envelope pr-0 md:pr-3"></i>
                <span
                  class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
                >Notifications</span>
              </router-link>
            </li>
          </ul>
        </div>
      </div>
      <div class="main-content w-full h-full flex-1 pb-24 md:pb-5 " v-bind:class="loggingIn?'md:ml-48':''">
         <transition name="fade">
          <router-view></router-view>
         </transition>
      </div>
    </div>
  </div>
</template>

<script>
import Navigation from "./components/navigation/Navigation.vue";
import "./assets/styles/tailwind.css";
import localConfig from './assets/config.json'

//import Vue from 'vue'
//import VueNativeSock from 'vue-native-websocket'
//import store from './store'
export default {
  name: "app",
  components: {
    Navigation
  },
  computed:{
    loggingIn () {
      if(this.$store.state.authentication.status.loggedIn){
       let signin = this.$store.state.authentication.user.signIn?"google":"nosignin"
       this.$connect('wss://'+window.location.host +'/sapi/v1/ws?token='+this.$store.state.authentication.user.token+'&signIn='+signin)
       this.$store.dispatch("datastore/fetchChartData",{filter: "", quick: true})
       this.$store.dispatch("datastore/fetchFullChartData",{filter: "", quick: false})
       this.$store.dispatch("datastore/fetchConfiguration")
       this.$store.dispatch("datastore/fetchProjects")
       this.$store.dispatch("datastore/fetchStarredArtifacts");
       this.$store.dispatch("datastore/fetchUserInfo");
      }
      return this.$store.state.authentication.status.loggedIn;
    },
     user () {
      return this.$store.state.authentication.user;
    },
    config(){
      return this.$store.state.datastore.config;
    }
  },
  created() {
    this.$store.commit("datastore/setLocalConfig",localConfig)
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.notices .toast {
  font-size: small;
}
body.v-tour--active{pointer-events:none}.v-tour{pointer-events:auto}.v-tour__target--highlighted{-webkit-box-shadow:0 0 0 4px rgba(0,0,0,.4);box-shadow:0 0 0 4px rgba(0,0,0,.4);pointer-events:auto;z-index:9999}.v-tour__target--relative{position:relative}.v-step[data-v-7c9c03f0]{background:#50596c;color:#fff;max-width:320px;border-radius:3px;-webkit-filter:drop-shadow(0 0 2px rgba(0,0,0,.5));filter:drop-shadow(0 0 2px rgba(0,0,0,.5));padding:1rem;text-align:center;z-index:10000}.v-step .v-step__arrow[data-v-7c9c03f0]{width:0;height:0;border-style:solid;position:absolute;margin:.5rem;border-color:#50596c}.v-step .v-step__arrow--dark[data-v-7c9c03f0]{border-color:#454d5d}.v-step[x-placement^=top][data-v-7c9c03f0]{margin-bottom:.5rem}.v-step[x-placement^=top] .v-step__arrow[data-v-7c9c03f0]{border-width:.5rem .5rem 0 .5rem;border-left-color:transparent;border-right-color:transparent;border-bottom-color:transparent;bottom:-.5rem;left:calc(50% - 1rem);margin-top:0;margin-bottom:0}.v-step[x-placement^=bottom][data-v-7c9c03f0]{margin-top:.5rem}.v-step[x-placement^=bottom] .v-step__arrow[data-v-7c9c03f0]{border-width:0 .5rem .5rem .5rem;border-left-color:transparent;border-right-color:transparent;border-top-color:transparent;top:-.5rem;left:calc(50% - 1rem);margin-top:0;margin-bottom:0}.v-step[x-placement^=right][data-v-7c9c03f0]{margin-left:.5rem}.v-step[x-placement^=right] .v-step__arrow[data-v-7c9c03f0]{border-width:.5rem .5rem .5rem 0;border-left-color:transparent;border-top-color:transparent;border-bottom-color:transparent;left:-.5rem;top:calc(50% - 1rem);margin-left:0;margin-right:0}.v-step[x-placement^=left][data-v-7c9c03f0]{margin-right:.5rem}.v-step[x-placement^=left] .v-step__arrow[data-v-7c9c03f0]{border-width:.5rem 0 .5rem .5rem;border-top-color:transparent;border-right-color:transparent;border-bottom-color:transparent;right:-.5rem;top:calc(50% - 1rem);margin-left:0;margin-right:0}.v-step__header[data-v-7c9c03f0]{margin:-1rem -1rem .5rem;padding:.5rem;background-color:#454d5d;border-top-left-radius:3px;border-top-right-radius:3px}.v-step__content[data-v-7c9c03f0]{margin:0 0 1rem 0}.v-step__button[data-v-7c9c03f0]{background:transparent;border:.05rem solid #fff;border-radius:.1rem;color:#fff;cursor:pointer;display:inline-block;font-size:.8rem;height:1.8rem;line-height:1rem;outline:none;margin:0 .2rem;padding:.35rem .4rem;text-align:center;text-decoration:none;-webkit-transition:all .2s ease;transition:all .2s ease;vertical-align:middle;white-space:nowrap}.v-step__button[data-v-7c9c03f0]:hover{background-color:hsla(0,0%,100%,.95);color:#50596c}
@keyframes fadeOut{from{opacity:1}to{opacity:0}}.fadeOut{animation-name:fadeOut}@keyframes fadeInDown{from{opacity:0;transform:translate3d(0, -100%, 0)}to{opacity:1;transform:none}}.fadeInDown{animation-name:fadeInDown}@keyframes fadeInUp{from{opacity:0;transform:translate3d(0, 100%, 0)}to{opacity:1;transform:none}}.fadeInUp{animation-name:fadeInUp}.fade-enter-active,.fade-leave-active{transition:opacity 150ms ease-out}.fade-enter,.fade-leave-to{opacity:0}.notices{position:fixed;display:flex;top:0;bottom:0;left:0;right:0;padding:2em;overflow:hidden;z-index:1052;pointer-events:none}.notices .toast{display:inline-flex;align-items:center;animation-duration:150ms;margin:.5em 0;box-shadow:0 1px 4px rgba(0,0,0,.12),0 0 6px rgba(0,0,0,.04);border-radius:.25em;pointer-events:auto;opacity:.92;color:#fff;min-height:3em;cursor:pointer}.notices .toast .toast-text{margin:0;padding:.5em 1em}.notices .toast-success{background-color:#28a745}.notices .toast-info{background-color:#17a2b8}.notices .toast-warning{background-color:#ffc107}.notices .toast-error{background-color:#dc3545}.notices .toast-default{background-color:#343a40}.notices .toast.is-top,.notices .toast.is-bottom{align-self:center}.notices .toast.is-top-right,.notices .toast.is-bottom-right{align-self:flex-end}.notices .toast.is-top-left,.notices .toast.is-bottom-left{align-self:flex-start}.notices.is-top{flex-direction:column}.notices.is-bottom{flex-direction:column-reverse}.notices.is-custom-parent{position:absolute}@media screen and (max-width: 768px){.notices{padding:0;position:fixed !important}}
fieldset[disabled] .multiselect{pointer-events:none}.multiselect__spinner{position:absolute;right:1px;top:1px;width:48px;height:35px;background:#fff;display:block}.multiselect__spinner:after,.multiselect__spinner:before{position:absolute;content:"";top:50%;left:50%;margin:-8px 0 0 -8px;width:16px;height:16px;border-radius:100%;border:2px solid transparent;border-top-color:#41b883;box-shadow:0 0 0 1px transparent}.multiselect__spinner:before{animation:spinning 2.4s cubic-bezier(.41,.26,.2,.62);animation-iteration-count:infinite}.multiselect__spinner:after{animation:spinning 2.4s cubic-bezier(.51,.09,.21,.8);animation-iteration-count:infinite}.multiselect__loading-enter-active,.multiselect__loading-leave-active{transition:opacity .4s ease-in-out;opacity:1}.multiselect__loading-enter,.multiselect__loading-leave-active{opacity:0}.multiselect,.multiselect__input,.multiselect__single{font-family:inherit;font-size:16px;-ms-touch-action:manipulation;touch-action:manipulation}.multiselect{box-sizing:content-box;display:block;position:relative;width:100%;min-height:40px;text-align:left;color:#35495e}.multiselect *{box-sizing:border-box}.multiselect:focus{outline:none}.multiselect--disabled{background:#ededed;pointer-events:none;opacity:.6}.multiselect--active{z-index:50}.multiselect--active:not(.multiselect--above) .multiselect__current,.multiselect--active:not(.multiselect--above) .multiselect__input,.multiselect--active:not(.multiselect--above) .multiselect__tags{border-bottom-left-radius:0;border-bottom-right-radius:0}.multiselect--active .multiselect__select{transform:rotate(180deg)}.multiselect--above.multiselect--active .multiselect__current,.multiselect--above.multiselect--active .multiselect__input,.multiselect--above.multiselect--active .multiselect__tags{border-top-left-radius:0;border-top-right-radius:0}.multiselect__input,.multiselect__single{position:relative;display:inline-block;min-height:20px;line-height:20px;border:none;border-radius:5px;background:#fff;padding:0 0 0 5px;width:100%;transition:border .1s ease;box-sizing:border-box;margin-bottom:8px;vertical-align:top}.multiselect__input:-ms-input-placeholder{color:#35495e}.multiselect__input::placeholder{color:#35495e}.multiselect__tag~.multiselect__input,.multiselect__tag~.multiselect__single{width:auto}.multiselect__input:hover,.multiselect__single:hover{border-color:#cfcfcf}.multiselect__input:focus,.multiselect__single:focus{border-color:#a8a8a8;outline:none}.multiselect__single{padding-left:5px;margin-bottom:8px}.multiselect__tags-wrap{display:inline}.multiselect__tags{min-height:40px;display:block;padding:8px 40px 0 8px;border-radius:5px;border:1px solid #e8e8e8;background:#fff;font-size:14px}.multiselect__tag{position:relative;display:inline-block;padding:4px 26px 4px 10px;border-radius:5px;margin-right:10px;color:#fff;line-height:1;background:#41b883;margin-bottom:5px;white-space:nowrap;overflow:hidden;max-width:100%;text-overflow:ellipsis}.multiselect__tag-icon{cursor:pointer;margin-left:7px;position:absolute;right:0;top:0;bottom:0;font-weight:700;font-style:normal;width:22px;text-align:center;line-height:22px;transition:all .2s ease;border-radius:5px}.multiselect__tag-icon:after{content:"\D7";color:#266d4d;font-size:14px}.multiselect__tag-icon:focus,.multiselect__tag-icon:hover{background:#369a6e}.multiselect__tag-icon:focus:after,.multiselect__tag-icon:hover:after{color:#fff}.multiselect__current{min-height:40px;overflow:hidden;padding:8px 30px 0 12px;white-space:nowrap;border-radius:5px;border:1px solid #e8e8e8}.multiselect__current,.multiselect__select{line-height:16px;box-sizing:border-box;display:block;margin:0;text-decoration:none;cursor:pointer}.multiselect__select{position:absolute;width:40px;height:38px;right:1px;top:1px;padding:4px 8px;text-align:center;transition:transform .2s ease}.multiselect__select:before{position:relative;right:0;top:65%;color:#999;margin-top:4px;border-color:#999 transparent transparent;border-style:solid;border-width:5px 5px 0;content:""}.multiselect__placeholder{color:#adadad;display:inline-block;margin-bottom:10px;padding-top:2px}.multiselect--active .multiselect__placeholder{display:none}.multiselect__content-wrapper{position:absolute;display:block;background:#fff;width:100%;max-height:240px;overflow:auto;border:1px solid #e8e8e8;border-top:none;border-bottom-left-radius:5px;border-bottom-right-radius:5px;z-index:50;-webkit-overflow-scrolling:touch}.multiselect__content{list-style:none;display:inline-block;padding:0;margin:0;min-width:100%;vertical-align:top}.multiselect--above .multiselect__content-wrapper{bottom:100%;border-bottom-left-radius:0;border-bottom-right-radius:0;border-top-left-radius:5px;border-top-right-radius:5px;border-bottom:none;border-top:1px solid #e8e8e8}.multiselect__content::webkit-scrollbar{display:none}.multiselect__element{display:block}.multiselect__option{display:block;padding:12px;min-height:40px;line-height:16px;text-decoration:none;text-transform:none;vertical-align:middle;position:relative;cursor:pointer;white-space:nowrap}.multiselect__option:after{top:0;right:0;position:absolute;line-height:40px;padding-right:12px;padding-left:20px;font-size:13px}.multiselect__option--highlight{background:#41b883;outline:none;color:#fff}.multiselect__option--highlight:after{content:attr(data-select);background:#41b883;color:#fff}.multiselect__option--selected{background:#f3f3f3;color:#35495e;font-weight:700}.multiselect__option--selected:after{content:attr(data-selected);color:silver}.multiselect__option--selected.multiselect__option--highlight{background:#ff6a6a;color:#fff}.multiselect__option--selected.multiselect__option--highlight:after{background:#ff6a6a;content:attr(data-deselect);color:#fff}.multiselect--disabled .multiselect__current,.multiselect--disabled .multiselect__select{background:#ededed;color:#a6a6a6}.multiselect__option--disabled{background:#ededed!important;color:#a6a6a6!important;cursor:text;pointer-events:none}.multiselect__option--group{background:#ededed;color:#35495e}.multiselect__option--group.multiselect__option--highlight{background:#35495e;color:#fff}.multiselect__option--group.multiselect__option--highlight:after{background:#35495e}.multiselect__option--disabled.multiselect__option--highlight{background:#dedede}.multiselect__option--group-selected.multiselect__option--highlight{background:#ff6a6a;color:#fff}.multiselect__option--group-selected.multiselect__option--highlight:after{background:#ff6a6a;content:attr(data-deselect);color:#fff}.multiselect-enter-active,.multiselect-leave-active{transition:all .15s ease}.multiselect-enter,.multiselect-leave-active{opacity:0}.multiselect__strong{margin-bottom:8px;line-height:20px;display:inline-block;vertical-align:top}[dir=rtl] .multiselect{text-align:right}[dir=rtl] .multiselect__select{right:auto;left:1px}[dir=rtl] .multiselect__tags{padding:8px 8px 0 40px}[dir=rtl] .multiselect__content{text-align:right}[dir=rtl] .multiselect__option:after{right:auto;left:0}[dir=rtl] .multiselect__clear{right:auto;left:12px}[dir=rtl] .multiselect__spinner{right:auto;left:1px}@keyframes spinning{0%{transform:rotate(0)}to{transform:rotate(2turn)}}
</style>
