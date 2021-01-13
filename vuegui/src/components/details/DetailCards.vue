<template>
    <div class="flex flex-col">
        <div class="flex p-1 bg-gray-700 rounded-lg shadow self-center w-1/2 z-10">
        <button v-on:click="$emit('filterChange',chart)" class="bg-transparent w-1/4 hover:bg-gray-400 text-gray-800 font-bold text-m p-1 rounded-full items-center">
            <svg class="fill-current text-white w-4 h-4 " xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20"><path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
        </button>  
        <h2 class="text-lg w-1/2 capitalize text-white">{{chart.metadata.appVersion}}</h2>
      </div>
      <div class="flex flex-wrap flex-col bg-white rounded-lg shadow-md p-2 content-center -mt-5">
        <div class="flex flex-1 items-center mt-2">
            <div class="flex p-2">
                <div class="text-center md:text-left whitespace-no-wrap">  
                    <div class="text-purple-500">Chart version : {{chart.version}}</div>
                    <div class="text-gray-600">App version : {{chart.metadata.appVersion}}</div>
                </div>
            </div>
            <div class="flex p-2">
                <div class="text-center md:text-left whitespace-no-wrap">
                    <ToogleValidator 
                        v-bind:data=chart
                        :index="index" 
                        onText="Chart Valid"
                        offText="Chart not valid"
                        :config="config"
                        @toogleValidation="toogleValidation"/>
                </div>
            </div>
        </div>
     <div class="inline-block relative flex-1" v-if="config">
            <PackageSelect 
            v-model="chart"
            v-bind:config="config"
            @removeLabel="removelabel"
            @selectLabel="selectlabel"
          />
    </div>
    <div class="inline-block flex-wrap relative flex-1 content-center">
        <h3 class="text-gray-800 text-center font-bold">Dependencies</h3>
        <div class="md:inline-flex flex-wrap">
            <div class="bg-blue-500 m-1 text-white font-bold text-xs p-1 rounded-full" v-for="(value, index) in chart.dependencies " :key="index">
                {{value.name}}:{{value.version}}
            </div>  
        </div>
    </div>
    </div>
  </div>
</template>

<script>``
import PackageSelect from "./PackageSelect.vue"
import ToogleValidator from "./ToogleValidator.vue"
export default {
  name: "DetailCard",
  components: { PackageSelect,ToogleValidator },
  props: {
    chart: Object,
    config: Object,
    index: String,
    project: String
  },
  methods: {
    setSelectVersion () {
      this.$emit('setSelectVersion', this.chart_version)
    },
    removelabel(option) {
      let data = {
        name: this.chart.metadata.name,
        allDockerImages: [],
        charts: [this.chart],
        project: this.project,
        projectLab: option.name
      };
      this.$emit("sendPayload", {
        data: data,
        url: "/v1/removeProjectToArtifact"
      });
      this.$emit("input", this.chart);
    },
    selectlabel(option) {
      console.log(option)
      let data = {
        name: this.chart.metadata.name,
        allDockerImages: [],
        charts: [this.chart],
        project: this.project,
        projectLab: option.name
      };
      this.$emit("sendPayload", {
        data: data,
        url: "/v1/addProjectToArtifact"
      });
      this.$emit("input", this.chart);
    },
     toogleValidation(){
        let data = {
          name: this.chart.metadata.name,
          allDockerImages: [],
          charts: [this.chart],
          project: this.project
        };
        this.$emit("sendPayload", { data: data, url: "/v1/postChartData" });
        this.$emit("notificateValidation", this.chart.name+":"+this.chart.version);
        this.$emit("input", this.value);
    },
  }
}
</script>