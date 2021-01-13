<template>
  <Loading :condition="!chartData || chartData.length == 0">
    <div slot="content" class="px-2 mt-4">
      <h1 class="text-2xl m-12 font-bold uppercase">Artifacts</h1>
      <div class="flex flex-wrap -mx-2 justify-center">
        <div
          v-for="(value, index) in chartData"
          :key="index"
          class="m-4 sm:w-full md:w-full lg:w-1/3 xl:w-1/4 mb-12">
          <Cards
            v-bind:name="value.name"
            v-bind:icon="value.icon"
            v-bind:latest_version="value.latest_version"
            v-bind:docker_images="value.allDockerImages?value.allDockerImages.length:0"
            v-bind:charts="value.charts.length"
            v-bind:repo="value.project"/>
        </div>
      </div>
    </div>
  </Loading>
</template>


<script src="./myScript.js"></script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style src="./myCss.css" scoped></style>

