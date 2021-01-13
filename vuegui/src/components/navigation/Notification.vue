<template>
   <div  :class="severity" class=" border-t-4 rounded-b p-2 shadow-md" role="alert">
    <div class="flex">
        <div class="py-1" v-if="notification.type === 'project'" slot="logo">
            <svg v-if="notification.severity === 'Success'" class="fill-current h-6 w-6 text-teal-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z"/></svg>
            <svg v-if="notification.severity === 'Warning'" class="fill-current h-6 w-6 text-orange-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 5h2v6H9V5zm0 8h2v2H9v-2z"/></svg>
            <svg v-if="notification.severity === 'Error'" class="fill-current h-6 w-6 text-red-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M17 16a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4.01V4a1 1 0 0 1 1-1 1 1 0 0 1 1 1v6h1V2a1 1 0 0 1 1-1 1 1 0 0 1 1 1v8h1V1a1 1 0 1 1 2 0v9h1V2a1 1 0 0 1 1-1 1 1 0 0 1 1 1v13h1V9a1 1 0 0 1 1-1h1v8z"/></svg>
            <svg v-if="notification.severity === 'Question'" class="fill-current h-6 w-6 text-blue-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 20a10 10 0 1 1 0-20 10 10 0 0 1 0 20zm2-13c0 .28-.21.8-.42 1L10 9.58c-.57.58-1 1.6-1 2.42v1h2v-1c0-.29.21-.8.42-1L13 9.42c.57-.58 1-1.6 1-2.42a4 4 0 1 0-8 0h2a2 2 0 1 1 4 0zm-3 8v2h2v-2H9z"/></svg>
        </div>
        <div class="py-1" v-if="notification.type === 'progress'">
            <loading-progress
                :progress="notification.progress"
                size="45"
                rotate
                fillDuration="2"
                rotationDuration="1"
            />
        </div>
        <div class="flex-1 md:text-left" slot="content">
            <p class="font-bold">{{notification.title}}</p>
            <p class="text-sm" v-if="notification.payload.startsWith('<')" >{{notification.owner}}, <span v-html="notification.payload"></span></p>
            <p class="text-sm" v-else >{{notification.owner}},{{notification.payload}}</p>
        </div>
         <span class="" v-on:click="$emit('dismiss', notification.id)">
        <svg class="fill-current h-6 w-6 text-red-500" role="button" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><title>Close</title><path d="M14.348 14.849a1.2 1.2 0 0 1-1.697 0L10 11.819l-2.651 3.029a1.2 1.2 0 1 1-1.697-1.697l2.758-3.15-2.759-3.152a1.2 1.2 0 1 1 1.697-1.697L10 8.183l2.651-3.031a1.2 1.2 0 1 1 1.697 1.697l-2.758 3.152 2.758 3.15a1.2 1.2 0 0 1 0 1.698z"/></svg>
    </span>
    </div>
</div>
</template>

<script>
export default {
  name: "Notification",
  props: {
    notification: Object
  },
  data(){
      return ({
          indeterminate: false,
          counterClockwise: false,
          hideBackground: false
      })
  },
  computed: {
    severity() {
        switch (this.notification.severity) {
            case "Error":
                return 'bg-red-100 border-red-500 text-red-900'   
            case "Warning":
                return 'bg-orange-100 border-orange-500 text-orange-900'   
            case "Question":
                return 'bg-blue-100 border-blue-500 text-blue-900'   
            default:
                return 'bg-teal-100 border-teal-500 text-teal-900'
        }
    }
}
}
</script>

<style >
.vue-progress-path{
    display:inline-block;
    height: 60px;
    margin-top: -20px;
}
.vue-progress-path path{fill:none;stroke-width:6;transition:stroke-dashoffset .3s ease-in-out}
.vue-progress-path .background{stroke:#ddd}
.vue-progress-path .progress{stroke:#40b883}
.vue-progress-path.indeterminate path{transition:none}

</style>