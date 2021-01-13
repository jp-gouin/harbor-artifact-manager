<template>
  <div  class="animate"> 
      <div v-bind:class="[ packageStatus=='outdated' ? 'border-orange-400 shadow-lg': packageStatus=='uptodate' ? 'border-green-400 shadow-lg':'border-gray-400' ]" class="border bg-white rounded lg:rounded-none p-4 flex flex-col justify-between leading-normal">
        <div class="flex flex-row">
          <div class="rounded-full flex h-16 w-16 shadow-md border-gray-400 border">
            <img class="h-12 w-12 m-auto" :src="icon">
          </div>
          <slot name="header" class="flex flex-col flex-wrap">
            <div class="flex flex-col mt-1 ml-4 items-baseline ">
              <div class="whitespace-no-wrap capitalize font-medium ">
                {{name}}
              </div>
              <div v-if="repo">
                <span class="uppercase text-xs text-gray-600">Repo : </span>
                <span>{{repo}}</span>
              </div>
            </div>
          </slot>
          <div v-if="date" class="flex-1 flex flex-col flex-wrap items-end" >
            <span class="text-xs">{{ date | moment("from", "now") }}</span>
            <div v-if="showselected">
              <div v-if="packageStatus=='uptodate'">
                <button class=" m-auto flex rounded-full border border-green-400 text-gray-600 pl-2 items-center text-xs bg-white" disabled>
                <span class="">Up to date</span>
                <span class="rounded-full color text-green-400 uppercase ml-1 px-1 py-1 text-xs font-bold shadow bg-white">
                  <svg class="fill-current w-3 h-3" aria-hidden="true" focusable="false" data-prefix="far" data-icon="check-circle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M256 8C119.033 8 8 119.033 8 256s111.033 248 248 248 248-111.033 248-248S392.967 8 256 8zm0 48c110.532 0 200 89.451 200 200 0 110.532-89.451 200-200 200-110.532 0-200-89.451-200-200 0-110.532 89.451-200 200-200m140.204 130.267l-22.536-22.718c-4.667-4.705-12.265-4.736-16.97-.068L215.346 303.697l-59.792-60.277c-4.667-4.705-12.265-4.736-16.97-.069l-22.719 22.536c-4.705 4.667-4.736 12.265-.068 16.971l90.781 91.516c4.667 4.705 12.265 4.736 16.97.068l172.589-171.204c4.704-4.668 4.734-12.266.067-16.971z"></path></svg>
                </span>
              </button>
              </div>
              <div v-else-if="packageStatus=='outdated'">
                <button class=" m-auto flex rounded-full border border-orange-400 text-gray-600 pl-2 items-center text-xs bg-white" disabled>
                <span class="mr-2 ">Outdated</span>
                <span class="rounded-full color text-orange-600 uppercase ml-1 px-1 py-1 text-xs font-bold shadow bg-white">
                  <svg class="fill-current w-3 h-3" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="exclamation-triangle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z"></path></svg>
                </span>
              </button>
              </div>
            </div>
            <button v-else class="mr-1 m-auto flex rounded-full items-center uppercase  text-xs bg-gray-200 border border-gray-400" v-on:click="$emit('star',name)">
             <span v-bind:class="[userStared ? 'text-blue-700':'text-gray-800']" class="flex rounded-full mr-1 color uppercase px-1 py-1 text-xs font-bold">
                <svg class="fill-current w-3 h-3" fill="currentColor" stroke="currentColor" stroke-width="0" viewBox="0 0 576 512" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path fill="currentColor"  d="M259.3 17.8L194 150.2 47.9 171.5c-26.2 3.8-36.7 36.1-17.7 54.6l105.7 103-25 145.5c-4.5 26.3 23.2 46 46.4 33.7L288 439.6l130.7 68.7c23.2 12.2 50.9-7.4 46.4-33.7l-25-145.5 105.7-103c19-18.5 8.5-50.8-17.7-54.6L382 150.2 316.7 17.8c-11.7-23.6-45.6-23.9-57.4 0z"></path></svg>
              </span>
              <span class="mr-2 ">{{stars}}</span>
            </button>
          </div>
        </div>
        <slot name="header2" class="flex flex-col flex-wrap"></slot>
        <span class="whitespace-normal  overflow-auto mt-2 mb-2 break-words h-16 max-h-xs text-xs text-gray-600">{{description}}</span>
        <div class="flex m-auto w-full">
            <slot name="content" class="flex flex-col flex-wrap">
            <div class="text-center md:text-left flex flex-col">
                <div  class="text-purple-500 inline-flex flex-1 m-1">
                  <svg class="h-6 w-6 last overflow-visible mr-3" viewBox="0 0 512 512">
                     <defs>
                      <linearGradient id="brand-gradient">
                        <stop offset="0%" stop-color="#beb500"/>
                        <stop offset="50%" stop-color="#EA118D"/>
                        <stop offset="100%" stop-color="#2e368f"/>
                      </linearGradient>
                    </defs>
                    <path stroke="url(#brand-gradient)" d="M502.285 159.704l-234-156c-7.987-4.915-16.511-4.96-24.571 0l-234 156C3.714 163.703 0 170.847 0 177.989v155.999c0 7.143 3.714 14.286 9.715 18.286l234 156.022c7.987 4.915 16.511 4.96 24.571 0l234-156.022c6-3.999 9.715-11.143 9.715-18.286V177.989c-.001-7.142-3.715-14.286-9.716-18.285zM278 63.131l172.286 114.858-76.857 51.429L278 165.703V63.131zm-44 0v102.572l-95.429 63.715-76.857-51.429L234 63.131zM44 219.132l55.143 36.857L44 292.846v-73.714zm190 229.715L61.714 333.989l76.857-51.429L234 346.275v102.572zm22-140.858l-77.715-52 77.715-52 77.715 52-77.715 52zm22 140.858V346.275l95.429-63.715 76.857 51.429L278 448.847zm190-156.001l-55.143-36.857L468 219.132v73.714z"></path>
                  </svg>
                  Latest version : {{latest_version}}
                </div>
                <div class="text-gray-600 inline-flex flex-1 m-1">
                  <svg class="h-6 w-6 hash overflow-visible mr-3" viewBox="0 0 448 512">
                    <defs>
                      <linearGradient id="hashchart-gradient">
                        <stop offset="0%" stop-color="#1177eb"/>
                        <stop offset="100%" stop-color="#001eca"/>
                      </linearGradient>
                    </defs>
                      <path stroke="url(#hashchart-gradient)" d="M440.667 182.109l7.143-40c1.313-7.355-4.342-14.109-11.813-14.109h-74.81l14.623-81.891C377.123 38.754 371.468 32 363.997 32h-40.632a12 12 0 0 0-11.813 9.891L296.175 128H197.54l14.623-81.891C213.477 38.754 207.822 32 200.35 32h-40.632a12 12 0 0 0-11.813 9.891L132.528 128H53.432a12 12 0 0 0-11.813 9.891l-7.143 40C33.163 185.246 38.818 192 46.289 192h74.81L98.242 320H19.146a12 12 0 0 0-11.813 9.891l-7.143 40C-1.123 377.246 4.532 384 12.003 384h74.81L72.19 465.891C70.877 473.246 76.532 480 84.003 480h40.632a12 12 0 0 0 11.813-9.891L151.826 384h98.634l-14.623 81.891C234.523 473.246 240.178 480 247.65 480h40.632a12 12 0 0 0 11.813-9.891L315.472 384h79.096a12 12 0 0 0 11.813-9.891l7.143-40c1.313-7.355-4.342-14.109-11.813-14.109h-74.81l22.857-128h79.096a12 12 0 0 0 11.813-9.891zM261.889 320h-98.634l22.857-128h98.634l-22.857 128z"></path>
                    </svg>
                    Docker images : {{docker_images}}
                  </div>
                  <div class="text-gray-600 inline-flex flex-1 m-1">
                  <svg class="h-6 w-6 hash overflow-visible mr-3" viewBox="0 0 448 512">
                    
                      <path stroke="url(#hashchart-gradient)" d="M440.667 182.109l7.143-40c1.313-7.355-4.342-14.109-11.813-14.109h-74.81l14.623-81.891C377.123 38.754 371.468 32 363.997 32h-40.632a12 12 0 0 0-11.813 9.891L296.175 128H197.54l14.623-81.891C213.477 38.754 207.822 32 200.35 32h-40.632a12 12 0 0 0-11.813 9.891L132.528 128H53.432a12 12 0 0 0-11.813 9.891l-7.143 40C33.163 185.246 38.818 192 46.289 192h74.81L98.242 320H19.146a12 12 0 0 0-11.813 9.891l-7.143 40C-1.123 377.246 4.532 384 12.003 384h74.81L72.19 465.891C70.877 473.246 76.532 480 84.003 480h40.632a12 12 0 0 0 11.813-9.891L151.826 384h98.634l-14.623 81.891C234.523 473.246 240.178 480 247.65 480h40.632a12 12 0 0 0 11.813-9.891L315.472 384h79.096a12 12 0 0 0 11.813-9.891l7.143-40c1.313-7.355-4.342-14.109-11.813-14.109h-74.81l22.857-128h79.096a12 12 0 0 0 11.813-9.891zM261.889 320h-98.634l22.857-128h98.634l-22.857 128z"></path>
                    </svg>
                    Charts : {{charts}}
                  </div>
                 <router-link class="self-end m-auto flex justify-center" :to="'/details/'+name">
                  <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full">
                      More
                  </button>
                </router-link>
            </div>
            </slot>
        </div>
      </div>
    </div>
</template>
<script>
export default {
    name: 'Cardscp',
    props: {
      icon: String,
      repo: String,
      date: String,
      showselected: Boolean,
      packageStatus: String,
      description: String,
      name: String,
      latest_version: String,
      docker_images: Number,
      charts: Number,
      stars: Number,
      userStared: Boolean,
    }
  }
</script>
<style scoped>
.max-h-xs{
  max-height: 4rem;
}
.hash path {
    fill: none;
    stroke-width: 30
  } 
.last path {
    fill: none;
    stroke-width: 30
  } 
.hash path {
      stroke-dasharray: 3000;
      stroke-dashoffset: 3000;
      animation: animatePath 2s 0.5s forwards;
    }
.last path {
      stroke-dasharray: 2000;
      stroke-dashoffset: 2000;
      animation: animatePath 2s 0.5s forwards;
    }
@keyframes animatePath {
  to {
    stroke-dashoffset: 0;
  }
}
</style>