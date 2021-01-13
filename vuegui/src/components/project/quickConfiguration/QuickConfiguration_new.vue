<template>
  <Loading :condition="!chartData || chartData.length == 0">
    <div slot="content" class="px-2 mt-4">
      <h1 class="text-2xl mt-12 font-bold uppercase">{{project}}</h1>
       <Annoucement/>
       <div class="flex sticky top-8 z-40 flex-wrap  m-4 justify-left">
         <div class="bg-blue-100 border-t bg-opacity-75 border-b border-blue-500 text-blue-700 px-4 py-3 w-full flex" >
          <ToogleSwitch class="mr-2" v-model="autoDependencies" index="autoDependencies" onText="Automatic dependencies" />
          <button  class="btn-secondary font-bold py-1 px-2 rounded inline-flex items-center" v-on:click="validate()">
            <svg class="fill-current w-6 h-6 mr-2" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="cart-arrow-down" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M504.717 320H211.572l6.545 32h268.418c15.401 0 26.816 14.301 23.403 29.319l-5.517 24.276C523.112 414.668 536 433.828 536 456c0 31.202-25.519 56.444-56.824 55.994-29.823-.429-54.35-24.631-55.155-54.447-.44-16.287 6.085-31.049 16.803-41.548H231.176C241.553 426.165 248 440.326 248 456c0 31.813-26.528 57.431-58.67 55.938-28.54-1.325-51.751-24.385-53.251-52.917-1.158-22.034 10.436-41.455 28.051-51.586L93.883 64H24C10.745 64 0 53.255 0 40V24C0 10.745 10.745 0 24 0h102.529c11.401 0 21.228 8.021 23.513 19.19L159.208 64H551.99c15.401 0 26.816 14.301 23.403 29.319l-47.273 208C525.637 312.246 515.923 320 504.717 320zM403.029 192H360v-60c0-6.627-5.373-12-12-12h-24c-6.627 0-12 5.373-12 12v60h-43.029c-10.691 0-16.045 12.926-8.485 20.485l67.029 67.029c4.686 4.686 12.284 4.686 16.971 0l67.029-67.029c7.559-7.559 2.205-20.485-8.486-20.485z"></path></svg>
            <span class="text-sm">Save shopping cart</span>
          </button>
        </div>
      </div>
      <div class="flex flex-wrap -mx-2 justify-center">
        <div
          v-for="(value, index) in chartData"
          :key="index+value.name"
          class="m-4 sm:w-full md:w-full lg:w-1/3 xl:w-1/4 mb-4 mt-12">
            <QCCard 
            v-bind:artifact="value" 
            v-model="labelProject" 
            v-bind:autoDependencies="autoDependencies"
            v-bind:chartData="chartData"></QCCard>
        </div>
      </div>
    </div>
  </Loading>
</template>

<script>
import Loading from "../../miscellaneous/Loading";
import ToogleSwitch from "../../ToogleSwitch.vue";
import Annoucement from "../../miscellaneous/Annoucement.vue"
import QCCard from "./QCCard.vue"

export default {
  name: "QuickConfiguration",
  components: {
    Loading,
    ToogleSwitch,
    Annoucement,
    QCCard,
  },
  props: {
    project: String,
  },
  data() {
    return {
      autoDependencies: true,
      dataSet: [],
      chart_payload: new Map(),
      chart_payload_to_delete: new Map(),
      columns: [
        {
          label: "App Version",
          jsonId: "",
        },
        {
          label: "Chart Version",
          jsonId: "",
        },
        {
          label: "Dependencies",
          jsonId: "",
        },
      ],
    };
  },
  methods: {
    validate() {
      this.$toast.open({
        message: "Send data",
        type: "info",
      });
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/addProjectsToArtifact",
          JSON.stringify(this.labelProject)
        )
        .catch((error) => {
          this.$toast.open({
            message: error,
            type: "error",
          });
        });
    },
    getPrimaryDataset(charts) {
      if (charts.length > 4) {
        return charts.slice(0, 4);
      }
      return charts;
    },
    getSecondaryDataset(charts) {
      if (charts.length > 4) {
        return charts.slice(5, charts.length);
      }
    },
    addDependencies(chart){
      chart.dependencies.forEach((dependency) =>{
        console.log("iterate on chart dep")
        let depVersion = dependency.version.split('x')[0]
        console.log(depVersion)
        this.chartData.some((cdata)=>{
          if (dependency.name.toUpperCase() === cdata.name.toUpperCase()){
            return cdata.charts.some((chart)=>{
              if (chart.version.includes(depVersion)){
                if(!this.labelProject.charts.includes(chart.name+chart.version)){
                  this.labelProject.charts.push(chart.name+chart.version);
               }
               return true
              }
            });
          }
        });
      });
      console.log(this.labelProject)
    },
    // TODO Find a way to better managed the add and removal of chart/artifact
    // For the UI create a component QuickCOnfigurationCard to lighten this section
    datachange(select, data, chart, isChart) {
      console.log(chart);
      console.log(data);
      //init the value since it only extist on the master object and we use the child chart object after
      if(isChart){
        data.project = chart.project;
        data.name = chart.name;
        if(this.labelProject.charts.includes(data.name+data.version)){
          const index = this.labelProject.charts.indexOf(data.name+data.version);
          if (index > -1) {
            this.labelProject.charts.splice(index, 1);
          }
        }else{
          this.labelProject.charts.push(data.name+data.version)
          if(this.autoDependencies){
            console.log("automatic dependencies resolver")
            this.addDependencies(data)
          }
        }
      }
      else{
        if(this.labelProject.dockerImages.includes(data)){
          const index = this.labelProject.dockerImages.indexOf(data);
          if (index > -1) {
            this.labelProject.dockerImages.splice(index, 1);
          }
        }else{
          this.labelProject.dockerImages.push(data)
        }
      }
    },
  },
  computed: {
    chartData() {
      return this.$store.state.datastore.chartsFull;
    },
    config() {
      return this.$store.state.datastore.config;
    },
    labelProject: {
      get(){
        let p;
        this.$store.state.datastore.projects.forEach((element) => {
          element.versions.forEach((v)=>{
            if (v.name === this.project) {
              console.log("found")
                p = v;
            }
        });
      });
      return p;
      },
      set(project){
        this.$store.commit("datastore/SET_PROJECT",{project})
      }
    },
  },
};
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
/* Tab content - closed */
.tab-content {
  max-height: 0;
}
/* :checked - resize to full height */
.tab input:checked ~ .tab-content {
  max-height: 100%;
}
/* Label formatting when open */
.tab input:checked + label {
  /*@apply text-xl p-5 border-l-2 border-indigo-500 bg-gray-100 text-indigo*/
  font-size: 1.25rem; /*.text-xl*/
  border-left-width: 2px; /*.border-l-2*/
  border-color: #6574cd; /*.border-indigo*/
  background-color: #f7fafc54; /*.bg-gray-100 */
  color: #6574cd; /*.text-indigo*/
}
/* Icon */
.tab label::after {
  float: right;
  right: 0;
  top: 0;
  display: block;
  width: 1.5em;
  height: 1.5em;
  text-align: center;
  -webkit-transition: all 0.35s;
  -o-transition: all 0.35s;
  transition: all 0.35s;
}
/* Icon formatting - closed */
.tab input[type="checkbox"] + label::after {
  content: "+";
  font-weight: bold; /*.font-bold*/
  border-radius: 9999px; /*.rounded-full */
  border-color: #313131; /*.border-grey*/
}
.tab input[type="radio"] + label::after {
  content: "\25BE";
  font-weight: bold; /*.font-bold*/
  border-width: 1px; /*.border*/
  border-radius: 9999px; /*.rounded-full */
  border-color: #313131; /*.border-grey*/
}
/* Icon formatting - open */
.tab input[type="checkbox"]:checked + label::after {
  transform: rotate(315deg);
  background-color: #6574cd; /*.bg-indigo*/
  color: #f8fafc; /*.text-grey-lightest*/
  border-color: #f8fafc;
}
.tab input[type="radio"]:checked + label::after {
  transform: rotateX(180deg);
  background-color: #6574cd; /*.bg-indigo*/
  color: #f8fafc; /*.text-grey-lightest*/
  border-color: #f8fafc;
}
.top-8{
  top: 4rem;
}
</style>