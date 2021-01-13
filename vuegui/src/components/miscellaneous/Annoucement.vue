<template>
    <div class="w-3/4 m-auto p-2">
      <div v-for="(annoucement, index) in annoucements" :key="index+annoucements.id" 
        class="bg-teal-100 border-t-4 border-teal-500 rounded-b text-teal-900 p-2 shadow-md">
          <div class="flex">
            <div class="py-1"><svg class="fill-current h-6 w-6 text-teal-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z"/></svg></div>
            <div class="text-left">
              <p class="font-bold">Annoucement</p>
              <span v-html="annoucement.payload"></span>
            </div>
          </div>
      </div>
    </div>
</template>
<script>
export default {
  mounted: function () {
    this.$store.dispatch("datastore/fetchAnnoucements")
  },
  computed: {
    annoucements() {
      return this.$store.state.datastore.annoucements;
    }
  }
}
</script>
<style scoped>

</style>