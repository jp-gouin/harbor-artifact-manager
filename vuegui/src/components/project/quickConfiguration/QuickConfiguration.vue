<template>
  <Loading :condition="!chartData || chartData.length == 0">
    <div slot="content" class="px-2 mt-4">
      <h1 class="text-2xl mt-12 font-bold uppercase">{{project}}</h1>
       <Annoucement/>
       <div class="flex sticky top-8 z-40 flex-wrap  m-4 justify-left">
         <div class="bg-blue-100 border-t bg-opacity-75 border-b border-blue-500 text-blue-700 px-4 py-3 w-full flex" >
          <ToogleSwitch class="mr-2" v-model="autoDependencies" index="autoDependencies" onText="Automatic dependencies" />
          <button
            v-on:click="validate()"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >Validate</button>
        </div>
      </div>
      <div class="flex flex-wrap -mx-2 justify-center">
        <div
          v-for="(value, index) in chartData"
          :key="index+value.name"
          class="m-4 sm:w-full md:w-full lg:w-1/3 xl:w-1/4 mb-4 mt-12"
        >
          <Cards
            v-bind:name="value.name"
            v-bind:icon="value.icon"
            v-bind:latest_version="value.latest_version"
            v-bind:docker_images="value.allDockerImages?value.allDockerImages.length:0"
            v-bind:description="value.charts[0].metadata.description"
            v-bind:date="value.charts[0].metadata.created"
            v-bind:repo="value.project"
            v-bind:showselected="true"
            v-bind:packageStatus="'uptodate'"
          >
            <div slot="content" class="flex flex-col">
              <table class="table-auto text-sm border-0 overflow-auto h-auto w-auto">
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
                    v-for="(chart, index) in getPrimaryDataset(value.charts)"
                    :key="index+chart.version"
                  >
                    <td class="border-0">
                      <label class="block text-gray-500 font-bold">
                        <input
                          class="mr-2 leading-tight"
                          type="checkbox"
                          :checked="labelProject.charts.includes(chart.name+chart.version)"
                          v-on:change="datachange($event.target.checked,chart, value, true)"
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
              <div class="flex flex-col flex-wrap justify-around whitespace-no-wrap align-baseline">
                <div class="tab border-t" v-if="getSecondaryDataset(value.charts)">
                  <input
                    class="absolute opacity-0"
                    :id="'tab-multi-one'+index"
                    type="checkbox"
                    name="tabs"
                  />
                  <label
                    class="block p-1 leading-normal cursor-pointer"
                    :for="'tab-multi-one'+index"
                  >Show more...</label>
                  <div
                    class="tab-content items-baseline self-end flex flex-col pl-2 overflow-auto border-l-2 border-indigo-500 leading-normal"
                  >
                    <table class="table-auto text-sm border-0 overflow-auto h-auto w-32">
                      <tbody>
                        <tr
                          v-for="(chart, index) in getSecondaryDataset(value.charts)"
                          :key="index+chart.version"
                        >
                          <td class="border-0">
                            <label class="md:w-2/3 block text-gray-500 font-bold">
                              <input
                                class="mr-2 leading-tight"
                                type="checkbox"
                                :checked="labelProject.charts.includes(chart.name+chart.version)"
                                v-on:change="datachange($event.target.checked,chart, value, true)"
                              />
                              <span class="text-sm">{{chart.metadata.appVersion}}</span>
                            </label>
                          </td>
                          <td>
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
                </div>
              </div>
              <h2 class="text-sm font-bold uppercase mt-1">Other Version</h2>
              <div class="flex flex-col flex-wrap whitespace-no-wrap align-baseline">
                <div
                  v-for="(otherv, index) in getPrimaryDataset(value.otherv)"
                  :key="index+otherv.tag"
                  style="align-self: end;"
                >
                  <label class="md:w-2/3 block text-gray-500 font-bold">
                    <input
                      class="mr-2 leading-tight"
                      type="checkbox"
                      :checked="labelProject.dockerImages.includes(otherv)"
                      v-on:change="datachange($event.target.checked,otherv, value, false)"
                    />
                    <span class="text-sm">{{otherv.split('/').pop()}}</span>
                  </label>
                </div>
                <div class="tab overflow-hidden border-t" v-if="getSecondaryDataset(value.otherv)">
                  <input
                    class="absolute opacity-0"
                    :id="'tab-multi-other'+index"
                    type="checkbox"
                    name="tabs"
                  />
                  <label
                    class="block p-1 leading-normal cursor-pointer"
                    :for="'tab-multi-other'+index"
                  >Show more...</label>
                  <div
                    class="tab-content items-baseline self-end  flex flex-col pl-2 overflow-hidden border-l-2 border-indigo-500 leading-normal"
                  >
                    <div
                      v-for="(otherv, index) in getSecondaryDataset(value.otherv)"
                      :key="index+otherv"
                    >
                      <label class="md:w-2/3 block text-gray-500 font-bold">
                        <input
                          class="mr-2 leading-tight"
                          type="checkbox"
                          :checked="labelProject.dockerImages.includes(otherv)"
                          v-on:change="datachange($event.target.checked,otherv,value, false)"
                        />
                        <span class="text-sm">{{otherv.split('/').pop()}}</span>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </Cards>
        </div>
      </div>
    </div>
  </Loading>
</template>

<script>
import Cards from "../../dashboard/cards/Cards_new";
import Loading from "../../miscellaneous/Loading";
import ToogleSwitch from "../../ToogleSwitch.vue";
import Annoucement from "../../miscellaneous/Annoucement.vue"

export default {
  name: "QuickConfiguration",
  components: {
    Cards,
    Loading,
    ToogleSwitch,
    Annoucement
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
    labelProject() {
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
  },
};
</script>

<style scoped>
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