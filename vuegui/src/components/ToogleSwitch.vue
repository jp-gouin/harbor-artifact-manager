<template>
  <div class="flex items-center justify-center">
  
  <!-- Toggle Button -->
  <label 
    :for="index"
    class="flex items-center cursor-pointer"
  >
    <!-- toggle -->
    <div class="relative">
      <!-- input -->
      <input v-bind:value="value" v-model="value" :id="index" type="checkbox" class="hidden" v-on:change="approveChart" />
      <!-- line -->
      <div
        class="toggle__line w-10 h-4 bg-gray-400 rounded-full shadow-inner"
      ></div>
      <!-- dot -->
      <div
        class="toggle__dot absolute w-6 h-6 bg-white rounded-full shadow inset-y-0 left-0"
      ></div>
    </div>
    <!-- label -->
    <div
      class="ml-3 text-sm"
    >
      {{ `${ isToggleOn ? onText : offText }` }}
    </div>
  </label>
  
</div>
</template>

<script>
export default {
    name: "ToogleSwitch",
  mounted () {
    this.isToggleOn = this.value
  },
  props: ['value', 'onText', 'offText', 'index'],
  data () {
    return {
      isToggleOn: false
    }
  },
  methods: {
    approveChart (event) {
        this.$emit('approveChart', event.target.value)
        this.$emit('input',event.target.value)
    },
  }
}
</script>

<style>
  .toggle__dot {
  top: -.25rem;
  left: -.25rem;
  transition: all 0.3s ease-in-out;
}

input:checked ~ .toggle__dot {
  transform: translateX(100%);
  background-color: #48bb78;
}
</style>