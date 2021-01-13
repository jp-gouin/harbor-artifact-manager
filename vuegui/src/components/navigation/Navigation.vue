<template>
  <!--Nav-->
    <nav class="bg-gray-900 pt-2 pb-1 px-1 mt-0 h-auto fixed w-full z-40 top-0 shadow-navbar">
      <div class="flex flex-wrap items-center">
        <div class="flex flex-shrink lg:w-1/3 justify-center lg:justify-start text-white">
          <a href="#">
            <span class="text-xl pl-2">
              <i class="em em-grinning"></i>
            </span>
          </a>
        </div>

        <div class="flex flex-1 lg:w-1/3 justify-center lg:justify-start text-white px-2">
          <span class="relative w-full">
            <input
              type="search"
              placeholder="Search"
              class="w-full bg-gray-800 text-sm text-white transition border border-transparent focus:outline-none focus:border-gray-700 rounded py-1 px-2 pl-10 appearance-none leading-normal"
            />
            <div class="absolute search-icon" style="top: .5rem; left: .8rem;">
              <svg
                class="fill-current pointer-events-none text-white w-4 h-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
              >
                <path
                  d="M12.9 14.32a8 8 0 1 1 1.41-1.41l5.35 5.33-1.42 1.42-5.33-5.34zM8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12z"
                />
              </svg>
            </div>
          </span>
        </div>

        <div class="flex w-full pt-2 content-center justify-between lg:w-1/3 lg:justify-end">
          <ul class="list-reset flex justify-between flex-1 lg:flex-none items-center">
            <li class="flex-1 lg:flex-none lg:mr-3">
              <button  v-on:click="resetTour" title="Refresh the page to lauch the help"  class="hover:bg-gray text-white hover:text-blue-500  font-bold text-xs inline-flex items-center ">
                <svg class="fill-current   w-6 h-6" aria-hidden="true" focusable="false" data-prefix="far" data-icon="question-circle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 448c-110.532 0-200-89.431-200-200 0-110.495 89.472-200 200-200 110.491 0 200 89.471 200 200 0 110.53-89.431 200-200 200zm107.244-255.2c0 67.052-72.421 68.084-72.421 92.863V300c0 6.627-5.373 12-12 12h-45.647c-6.627 0-12-5.373-12-12v-8.659c0-35.745 27.1-50.034 47.579-61.516 17.561-9.845 28.324-16.541 28.324-29.579 0-17.246-21.999-28.693-39.784-28.693-23.189 0-33.894 10.977-48.942 29.969-4.057 5.12-11.46 6.071-16.666 2.124l-27.824-21.098c-5.107-3.872-6.251-11.066-2.644-16.363C184.846 131.491 214.94 112 261.794 112c49.071 0 101.45 38.304 101.45 88.8zM298 368c0 23.159-18.841 42-42 42s-42-18.841-42-42 18.841-42 42-42 42 18.841 42 42z"></path></svg>
            </button>
             </li>
             <li class="flex-1 lg:flex-none lg:mr-3">
              <button  v-on:click="showNotification = !showNotification" v-bind:class="{ 'pulsate-bck': notifications.length > 0 }" class="v-step-8 relative text-white p-2 bg-gray-900 text-gray-200 hover:bg-gray  font-bold text-xs inline-flex items-center ">
                <svg class="fill-current text-gray-200  w-4 h-4"  aria-hidden="true" focusable="false" data-prefix="fas" data-icon="bell" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M224 512c35.32 0 63.97-28.65 63.97-64H160.03c0 35.35 28.65 64 63.97 64zm215.39-149.71c-19.32-20.76-55.47-51.99-55.47-154.29 0-77.7-54.48-139.9-127.94-155.16V32c0-17.67-14.32-32-31.98-32s-31.98 14.33-31.98 32v20.84C118.56 68.1 64.08 130.3 64.08 208c0 102.3-36.15 133.53-55.47 154.29-6 6.45-8.66 14.16-8.61 21.71.11 16.4 12.98 32 32.1 32h383.8c19.12 0 32-15.6 32.1-32 .05-7.55-2.61-15.27-8.61-21.71z"></path></svg>
                <span class="flex absolute top-0 right-0 rounded-full bg-indigo-500 uppercase px-1  text-xs font-bold">{{notifications.length}}</span>
            </button>
             </li>
            <li class="flex-1 lg:flex-none lg:mr-3">
              <div class="relative inline-block">
                <button
                  class="drop-button text-white focus:outline-none"
                >
                  <span class="pr-2">
                    <i class="em em-robot_face"></i>
                  </span> Hi, {{user}}
                  <svg
                    class="h-3 fill-current inline"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                  >
                    <path
                      d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                    />
                  </svg>
                </button>
              </div>
            </li>
            <li class="flex-1 lg:flex-none lg:mr-3">
              <p>
                <router-link to="/login" class="text-white">Logout</router-link>
              </p>
            </li>
          </ul>
        </div>
         <div v-if="showNotification" v-click-outside="hide" class="rounded  absolute pin-t pin-l mt-12 right-0 lg:w-1/4 w-full top-0 slide-in-right">
              <ul>
                <li v-for="(notification, index) in notifications" :key="index" class="mb-2 animate">
                  <Notification
                    :notification="notification"
                    @dismiss="hideNotification"
                   />
                </li>
                <li>
                  <div class="bg-gray-800 shadow-navbar px-1 py-1">
                    <router-link to="/notification-center" class="text-white underline">Show all notifications</router-link>
                  </div>
                </li>
              </ul>
          </div>
      </div>
    </nav>

</template>

<script>
import Notification from './Notification.vue'
import ClickOutside from 'vue-click-outside'
export default {
  name: "Naviagtion",
  components:{ Notification},
  props: {
    msg: String,
    user: String,
  },
  data(){
    return {
      showNotification: false,
    }
  },
  computed:{
    notifications(){
      return this.$store.state.socketstore.quickNotifications;
    }
  },
  methods:{
    hideNotification(notificationID){
      this.$store.commit('HIDE_NOTIFICATION', notificationID)
    },
    hide () {
      this.showNotification = false
    },
    resetTour(){
      localStorage.removeItem('projectTour')
      //add other tour here
      location.reload();
    }
  },
  mounted () {
    // prevent click outside event with popupItem.
    this.popupItem = this.$el
  },
  // do not forget this section
  directives: {
    ClickOutside
  }
};
</script>
<style scoped>
.animate{
    right: -200px;
    -webkit-animation: slide 0.5s forwards;
    -webkit-animation-delay: 2s;
    animation: slide 0.5s forwards;
    animation-delay: 2s;
}
.pulsate-bck {
  -webkit-animation: pulsate-bck 0.5s ease-in-out infinite both;
	animation: pulsate-bck 0.5s ease-in-out infinite both;
}
/**
 * ----------------------------------------
 * animation pulsate-bck
 * ----------------------------------------
 */
@-webkit-keyframes pulsate-bck {
  0% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    -webkit-transform: scale(0.9);
            transform: scale(0.9);
  }
  100% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}
@keyframes pulsate-bck {
  0% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    -webkit-transform: scale(0.9);
            transform: scale(0.9);
  }
  100% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
}

.slide-in-right {
	-webkit-animation: slide-in-right 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
	animation: slide-in-right 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}

/**
 * ----------------------------------------
 * animation slide-in-right
 * ----------------------------------------
 */
@-webkit-keyframes slide-in-right {
  0% {
    -webkit-transform: translateX(1000px);
            transform: translateX(1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
            transform: translateX(0);
    opacity: 1;
  }
}
@keyframes slide-in-right {
  0% {
    -webkit-transform: translateX(1000px);
            transform: translateX(1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
            transform: translateX(0);
    opacity: 1;
  }
}

</style>
