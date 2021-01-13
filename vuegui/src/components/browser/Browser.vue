<template>
  <Loading :condition="!chartData || chartData.length == 0">
    <div slot="content" class="px-2 mt-4">
      <div class="m-12 inline-flex">
        <h1 class="text-2xl font-bold uppercase">Browse the catalog</h1>
        <vue-json-to-csv class="ml-3 text-sm" :json-data="csvData" csv-title="Catalog_Backends">
          <button
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded inline-flex items-center"
          >
            <svg
              class="fill-current w-4 h-4 mr-2"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            >
              <path d="M13 8V2H7v6H2l8 8 8-8h-5zM0 18h20v2H0v-2z" />
            </svg>
            <span>Download CSV</span>
          </button>
        </vue-json-to-csv>
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
            v-bind:stars="getStaCount(value.name)"
            v-bind:userStared="userInfo.starred.includes(value.name)"
            @star="starredArtefact"
          >
            <div slot="content" class="flex flex-col h-64 overflow-auto">
              <table class="table-auto text-sm border-0 overflow-auto h-auto w-auto">
                <thead>
                  <tr class="border-b-2 border-gray-500">
                    <th
                      v-for="(column, index) in columns"
                      :key="index+column.label"
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
                          :key="index"
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
                                :key="index+value.name"
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
                    class="tab-content items-baseline self-end flex flex-col pl-2 overflow-hidden border-l-2 border-indigo-500 leading-normal"
                  >
                    <div
                      v-for="(otherv, index) in getSecondaryDataset(value.otherv)"
                      :key="index+otherv.tag"
                    >
                      <label class="md:w-2/3 block text-gray-500 font-bold">
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
import Cards from "../dashboard/cards/Cards_new_full";
import Loading from "../miscellaneous/Loading";
import VueJsonToCsv from "vue-json-to-csv";
export default {
  name: "Browser",
  components: {
    Cards,
    Loading,
    VueJsonToCsv,
  },
  props: {
    project: String,
  },
  data() {
    return {
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
    getStaCount(name){
      let count = 0
      console.log(this.starredArtifacts)
      if(!this.starredArtifacts){
        return 0
      }
      this.starredArtifacts.forEach(s => {
        console.log("looping artiefact star")
        if(s.name === name){
          count = s.count
        }
      })
      return count
    },
    starredArtefact(name){
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/starArtifact",
          JSON.stringify(name)
        ).then(()=>{
          this.$store.dispatch("datastore/fetchStarredArtifacts");
          this.$store.dispatch("datastore/fetchUserInfo");
        })
        .catch(error => {
          this.$toast.open({
            message: error.response.data,
            type: "error"
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
  },
  computed: {
    chartData() {
      return this.$store.state.datastore.chartsFull;
    },
    starredArtifacts(){
      return this.$store.state.datastore.starredArtifacts;
    },
    userInfo(){
      return this.$store.state.datastore.userInfo;
    },
    csvData() {
      let data = [];
      this.$store.state.datastore.chartsFull.forEach((chartdata) => {
        let c = {
          name: chartdata.name,
          appVersion: "",
          dependencies: "",
          otherVersion: "",
        };
        chartdata.charts.forEach((chart) => {
          c.appVersion = c.appVersion + chart.metadata.appVersion + "\n";
          if (chart.dependencies){
            chart.dependencies.forEach((d) => {
              c.dependencies = c.dependencies + d.name + ":" + d.version + "\n";
            });
          }
        });
        chartdata.otherv.forEach((ov) => {
          c.otherVersion =
            c.otherVersion +
            ov.split("/").pop() +
            "\n";
        });
        data.push(c);
      });
      return data;
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
</style>