<template>
  <div  class="animate"> 
      <div class="border bg-white rounded lg:rounded-none p-4 flex flex-col justify-between leading-normal">
        <div class="flex flex-row">
          <div class="rounded-full flex h-16 w-16 shadow-md border-gray-400 border">
            <img class="h-12 w-12 m-auto" :src="icon">
          </div>
          <slot name="header" class="flex flex-col flex-wrap">
            <div class="flex flex-col mt-1 ml-4 items-baseline ">
              <div v-if="!editmode" class="whitespace-no-wrap capitalize font-medium ">
                {{title}}
              </div>
              <div v-else class="whitespace-no-wrap capitalize font-medium ">
                <div
                  class="flex items-center border-b-2 border-teal-500 py-2"
                >
                <span class="text-white">Project_</span>
                  <input
                    v-model="editName"
                    class="text-gray-800 appearance-none bg-transparent border-none w-full mr-3 py-1 px-2 leading-tight focus:outline-none"
                    type="text"
                    placeholder="Name"
                    aria-label="Name"
                  />
                  <button
                    class="flex-shrink-0 mx-1 btn-secondary text-white py-1 px-2 rounded"
                    type="button"
                    v-on:click="$emit('title',editName)"
                  >Validate</button>
                  <button
                    class="flex-shrink-0 mx-1 btn-secondary text-sm py-1 px-2 rounded"
                    type="button"
                    v-on:click="$emit('cancel')"
                  >Cancel</button>
                </div>
              </div>
              <slot name="subheader" class="flex flex-col flex-wrap">
              </slot>
            </div>
          </slot>
          <div class="flex-1 flex flex-col flex-wrap items-end" >
            <slot name="rightheader"></slot>
          </div>
        </div>
        <slot name="header2" class="flex flex-col flex-wrap"></slot>
        <div class="flex m-auto w-full">
            <slot name="content" class="flex flex-col flex-wrap">
            </slot>
        </div>
      </div>
    </div>
</template>
<script>
export default {
    name: 'Cardbase',
    props: {
      icon: String,
      title: String,
      editmode: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        editName: "",
      }
    }
}
</script>
<style scoped>
.btn-secondary{
  color: #fff;
    background-color: #39596c;
    border-color: #39596c;
    border: 1px solid transparent;
    padding: .375rem .75rem;
    font-size: 1rem;
    line-height: 1.5;
    border-radius: .25rem;
    transition: color .15s ease-in-out,background-color .15s ease-in-out,border-color .15s ease-in-out,box-shadow .15s ease-in-out;
}
.btn-secondary:hover {
    color: #fff;
    background-color: #2c4453;
    border-color: #273d4b;
}
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