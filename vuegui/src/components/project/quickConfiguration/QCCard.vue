<template>
<div>
    <Cards
      v-bind:name="artifact.name"
      v-bind:icon="artifact.icon"
      v-bind:latest_version="artifact.latest_version"
      v-bind:docker_images="artifact.allDockerImages?artifact.allDockerImages.length:0"
      v-bind:description="artifact.charts[0].metadata.description"
      v-bind:date="artifact.charts[0].metadata.created"
      v-bind:repo="artifact.project"
      v-bind:showselected="true"
      v-bind:packageStatus="getPackageStatus"
    >
    <div slot="header2">
      <div class="flex flex-row justify-center">
        <div class="m-1">
          <span class="text-xs text-gray-600">Latest App : </span>
          <span>{{artifact.charts[0].metadata.appVersion}}</span>
        </div>
        <div class="m-1">
          <span class="text-xs text-gray-600">Lastest Chart : </span>
          <span>{{artifact.charts[0].version}}</span>
        </div>
      </div>
    </div>
      <div slot="content" class="flex-auto w-full">
        <div v-if="artifact.charts[0].dependencies && artifact.charts[0].dependencies.length > 0">
          <span class="text-gray-600">Dependencies</span>
          <div class="flex justify-center flex-row w-full">
          <div 
            class="bg-gray-200 border border-gray-400 m-1 text-gray-800  font-bold text-xs p-1 rounded-full"
            v-for="(value, index) in artifact.charts[0].dependencies "
            :key="index+value.version">
              {{value.name}}
          </div>
          </div>
        </div>
        <div class="flex justify-between mt-2">
          <button v-if="!value.charts.includes(artifact.charts[0].name+artifact.charts[0].version)" class="border border-gray-400 hover:bg-gray-400 text-gray-800 font-bold py-1 px-2 rounded inline-flex items-center" v-on:click="datachange(artifact.charts[0],artifact,true)">
            <svg class="fill-current w-4 h-4 mr-2" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="cart-plus" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M504.717 320H211.572l6.545 32h268.418c15.401 0 26.816 14.301 23.403 29.319l-5.517 24.276C523.112 414.668 536 433.828 536 456c0 31.202-25.519 56.444-56.824 55.994-29.823-.429-54.35-24.631-55.155-54.447-.44-16.287 6.085-31.049 16.803-41.548H231.176C241.553 426.165 248 440.326 248 456c0 31.813-26.528 57.431-58.67 55.938-28.54-1.325-51.751-24.385-53.251-52.917-1.158-22.034 10.436-41.455 28.051-51.586L93.883 64H24C10.745 64 0 53.255 0 40V24C0 10.745 10.745 0 24 0h102.529c11.401 0 21.228 8.021 23.513 19.19L159.208 64H551.99c15.401 0 26.816 14.301 23.403 29.319l-47.273 208C525.637 312.246 515.923 320 504.717 320zM408 168h-48v-40c0-8.837-7.163-16-16-16h-16c-8.837 0-16 7.163-16 16v40h-48c-8.837 0-16 7.163-16 16v16c0 8.837 7.163 16 16 16h48v40c0 8.837 7.163 16 16 16h16c8.837 0 16-7.163 16-16v-40h48c8.837 0 16-7.163 16-16v-16c0-8.837-7.163-16-16-16z"></path></svg>
            <span class="text-sm">Add to cart</span>
          </button>
          <button v-else class="border border-red-400 hover:bg-red-400 bg-red-500  text-white font-bold py-1 px-2 rounded inline-flex items-center" v-on:click="datachange(artifact.charts[0],artifact,true)">
            <svg class="fill-current w-4 h-4 mr-2" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="cart-plus" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M504.717 320H211.572l6.545 32h268.418c15.401 0 26.816 14.301 23.403 29.319l-5.517 24.276C523.112 414.668 536 433.828 536 456c0 31.202-25.519 56.444-56.824 55.994-29.823-.429-54.35-24.631-55.155-54.447-.44-16.287 6.085-31.049 16.803-41.548H231.176C241.553 426.165 248 440.326 248 456c0 31.813-26.528 57.431-58.67 55.938-28.54-1.325-51.751-24.385-53.251-52.917-1.158-22.034 10.436-41.455 28.051-51.586L93.883 64H24C10.745 64 0 53.255 0 40V24C0 10.745 10.745 0 24 0h102.529c11.401 0 21.228 8.021 23.513 19.19L159.208 64H551.99c15.401 0 26.816 14.301 23.403 29.319l-47.273 208C525.637 312.246 515.923 320 504.717 320zM408 168h-48v-40c0-8.837-7.163-16-16-16h-16c-8.837 0-16 7.163-16 16v40h-48c-8.837 0-16 7.163-16 16v16c0 8.837 7.163 16 16 16h48v40c0 8.837 7.163 16 16 16h16c8.837 0 16-7.163 16-16v-40h48c8.837 0 16-7.163 16-16v-16c0-8.837-7.163-16-16-16z"></path></svg>
            <span class="text-sm">Remove lastest</span>
          </button>
          <button class="border border-gray-400 hover:bg-gray-400 text-gray-800 font-bold py-1 px-2 rounded inline-flex items-center" v-on:click="showModal=true">
            <svg class="fill-current w-4 h-4 mr-2" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="external-link-alt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z"></path></svg>
            <span class="text-sm">More</span>
          </button>
        </div>
      </div>
    </Cards>
  <modal v-bind:large="true" v-if="showModal" @close="showModal = false">
    <h1 slot="header" class="text-xl m-4 font-bold uppercase">Advance configuration</h1>
    <div slot="body" >
      <div class="limit-table overflow-auto">
      <table class="table-auto m-auto text-sm border-0 overflow-auto h-1/2 w-auto">
        <thead>
          <tr class="border-b-2 border-gray-500">
            <th
              v-for="(column, index) in columns"
              :key="index"
              class="px-2 py-1"
            >{{ column.label }}</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(chart, index) in artifact.charts"
            :key="index+chart.version"
          >
            <td class="border-0">
              <label class="block text-gray-500 font-bold">
                <input
                  class="mr-2 leading-tight"
                  type="checkbox"
                  :checked="value.charts.includes(chart.name+chart.version)"
                  v-on:change="datachange(chart, artifact, true)"
                />
                <span class="text-sm">{{chart.metadata.appVersion}}</span>
              </label>
            </td>
            <td class="border-0">
              <span class="text-sm">{{chart.version}}</span>
            </td>
            <td class="border-0">
              <div class="md:inline-flex flex-wrap">
                <div
                  class="bg-blue-500 m-1 text-white font-bold text-xs p-1 rounded-full"
                  v-for="(value, index) in chart.dependencies "
                  :key="index+value.version"
                >{{value.name}}:{{value.version}}</div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      </div>
      <div class="" v-if="artifact.otherv && artifact.otherv.length > 0">
        <h2 class="text-sm font-bold uppercase m-4">Other Version</h2>
        <div class="limit-table overflow-auto flex flex-col flex-wrap whitespace-no-wrap align-baseline">
          <div
            v-for="(otherv, index) in artifact.otherv"
            :key="index+otherv.tag"
            style="align-self: end;"
          >
            <label class="md:w-2/3 block text-gray-500 font-bold">
              <input
                class="mr-2 leading-tight"
                type="checkbox"
                :checked="value.dockerImages.includes(otherv)"
                v-on:change="datachange(otherv, artifact, false)"
              />
              <span class="text-sm">{{otherv.split('/').pop()}}</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </modal>
</div>
</template>
<script>
import Cards from "../../dashboard/cards/Cards_new_full";
import Modal from "../Modal.vue";
export default {
  components: {
    Cards,
    Modal,
  },
    props:{
     artifact: Object,
     value: Object, 
     autoDependencies: Boolean,
     chartData: Array,
    },
    data() {
      return {
        showModal: false,
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
      }
    },
    methods: {
      addDependencies(chart){
        chart.dependencies.forEach((dependency) =>{
          console.log("iterate on chart dep")
          let depVersion = dependency.version.split('x')[0]
          console.log(depVersion)
          this.chartData.some((cdata)=>{
            if (dependency.name.toUpperCase() === cdata.name.toUpperCase()){
              return cdata.charts.some((chart)=>{
                if (chart.version.includes(depVersion)){
                  if(!this.value.charts.includes(chart.name+chart.version)){
                    this.value.charts.push(chart.name+chart.version);
                }
                return true
                }
              });
            }
          });
        });
      },
      // TODO Find a way to better managed the add and removal of chart/artifact
      // For the UI create a component QuickCOnfigurationCard to lighten this section
      datachange(data, chart, isChart) {
        //init the value since it only extist on the master object and we use the child chart object after
        if(isChart){
          data.project = chart.project;
          data.name = chart.name;
          if(this.value.charts.includes(data.name+data.version)){
            const index = this.value.charts.indexOf(data.name+data.version);
            if (index > -1) {
              this.value.charts.splice(index, 1);
            }
          }else{
            this.value.charts.push(data.name+data.version)
            if(this.autoDependencies){
              console.log("automatic dependencies resolver")
              this.addDependencies(data)
            }
          }
        }
        else{
          if(this.value.dockerImages.includes(data)){
            const index = this.value.dockerImages.indexOf(data);
            if (index > -1) {
              this.value.dockerImages.splice(index, 1);
            }
          }else{
            this.value.dockerImages.push(data)
          }
        }
        this.$emit('input',this.value)
      },
    },
    computed: {
      getPackageStatus() {
        if (this.value.charts.includes(this.artifact.charts[0].name+this.artifact.charts[0].version)){
          return 'uptodate'
        }
        let result= this.value.charts.find((e)=>{
          if (e.startsWith(this.artifact.charts[0].name)){
            return true
          }
        })
        return result ? 'outdated': 'none'
      }
    },
}
</script>
<style scoped>
.limit-table {
  height: 40vh;
}
.modal-container{
  width: 50%!important;
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