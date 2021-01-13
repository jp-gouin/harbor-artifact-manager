import Vue from 'vue'
import Vuex from 'vuex'
 
Vue.use(Vuex);
 
export const socketstore ={
  state: {
    socket: {
      isConnected: false,
      message: '',
      reconnectError: false,
    },
    notifications: [],
    quickNotifications: [],
    logs:[],
    health: {},
  },
  mutations:{
    SOCKET_ONOPEN (state, event)  {
      Vue.prototype.$socket = event.currentTarget
      state.socket.isConnected = true
    },
    SOCKET_ONCLOSE (state)  {
      state.socket.isConnected = false
    },
    SOCKET_ONERROR (state, event)  {
      console.error(state, event)
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE (state, message)  {
      state.socket.message = message
    },
    LOG(state, message){
      state.logs.push(message);
    },
    ADMHEALTHCHECK(state, message){
      state.health = message
    },
    NOTIFICATION(state, message){
      // update by id notification
      let index = state.notifications.findIndex(elem =>  elem.id === message.id);
      if(index === -1){
        state.notifications.unshift(message)
      }else {
        if(message.type == 'progress'){
          state.notifications[index] = message
        }else{
          state.notifications.splice(index, 1);
          state.notifications.unshift(message)
        }
       
      }
      // update quick notification 
      index = state.quickNotifications.findIndex(elem =>  elem.id === message.id);
      if(index === -1){
        state.quickNotifications.unshift(message)
      }else {
        if(message.type == 'progress'){
          state.quickNotifications.splice(index, 1);
         // state.notifications[index].progress = message.progress
          state.quickNotifications.splice(index, 0, message);
        }else{
          state.quickNotifications.splice(index, 1);
          state.quickNotifications.unshift(message)
        }
      }
      if(message.type === 'project'){
        this.dispatch("datastore/fetchProjects")
        this.dispatch("datastore/fetchConfiguration")
      }
    },
    HIDE_NOTIFICATION(state, notificationID){
      state.quickNotifications = state.quickNotifications.filter((value)=>{
        if(value.id != notificationID){
          return value;
        }
      });
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT(state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true;
    },
  },
  actions: {
    sendMessage (context, message) {
      Vue.prototype.$socket.send(message)
    }
  }
}