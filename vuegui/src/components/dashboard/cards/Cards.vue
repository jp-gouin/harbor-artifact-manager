<template>
  <div  class="animate"> 
    <div class="flex flex-col relative">
      <div class="flex flex-col p-1  self-center  w-1/2 z-10 -mt-6">
      <img class="h-12 w-12  mx-auto  top-0 -ml-6 -mt-12 absolute z-20" :src="icon">
        <div class="bg-gray-700 p-1 rounded-lg shadow">
          <h2 class="text-lg capitalize text-white whitespace-no-wrap overflow-auto">{{name}}</h2>
        </div>
      </div>
      <div class="flex flex-col bg-white rounded-lg p-6 shadow-md -mt-4">
        <div class="flex m-auto ">
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
  </div>
</template>
<script>
export default {
    name: 'Cards',
    props: {
      icon: String,
      name: String,
      latest_version: String,
      docker_images: Number,
      charts: Number
    }
  }
</script>
<style scoped>
path {
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