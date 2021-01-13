<template>
  <loading :condition="!chartData || chartData.length == 0">
    <div slot="content">
    <h1 class="text-2xl m-12 font-bold uppercase">{{chart}}</h1>
    <div class="flex flex-col lg:flex-row justify-around" v-if="chartData">
        <div class="w-full lg:w-1/3">
            <h2 class="text-l m-3 font-bold uppercase">Chart List</h2>
            <div class="flex flex-col flex-wrap -mb-4 -mx-2 content-center" >
            <DetailCard class="ml-4 mb-2" v-for="(value, index) in chartData.charts" :key="index"
             v-bind:chart=value
             v-bind:config=config
             v-bind:project=chartData.project
             v-bind:index="'chart'+index"
             @filterChange="filterChange"
             @sendPayload="sendPayload"
             @notificateValidation="notificateValidation"
            />
            </div>
        </div>
        <div class="w-full lg:w-3/5 overflow" >
        <h2 class="text-l m-3 font-bold uppercase">Docker Images</h2>
        <div class="mb-4">
      <input v-model="searchQuery" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Search">
      <div class="flex m-2 flex-wrap" >
          <h2>Quick filter</h2>
          <button v-for="(value, index) in chartData.charts" :key="index" v-on:click="quickFilter=value" v-bind:class="quickFilter==value ? 'bg-blue-600 text-white':'text-gray-800'" class="bg-gray-300 hover:bg-gray-400 active:bg-blue-600 active:text-white  font-bold text-xs p-1 m-1 rounded-full inline-flex items-center">
                <svg class="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20"><path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
                <span>{{value.version}}-{{value.metadata.appVersion}}</span>
            </button> 
          <button v-bind:class="quickFilter=='All' ? ['bg-blue-600','text-white']:'text-gray-800'" v-on:click="quickFilter='All'" class="bg-gray-300 hover:bg-gray-400  font-bold text-xs p-1 m-1 rounded-full inline-flex items-center">
                <svg class="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20"><path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
                <span>All</span>
            </button>
          <button v-bind:class="quickFilter=='latest' ? ['bg-blue-600','text-white']:'text-gray-800'" v-on:click="quickFilter='latest'" class="bg-gray-300 hover:bg-gray-400  font-bold text-xs p-1 m-1 rounded-full inline-flex items-center">
                <svg class="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20"><path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
                <span>Latest</span>
            </button>
      </div>
    </div>
        <table class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-full">
            <thead>
            <tr class="border-b-2 border-gray-500">
                <th v-for="(column, index) in columns" :key="index"  class="px-4 py-2">
                    <a class="hover:underline"
                      v-on:click="sortBy(column.jsonId)"
                      v-bind:class="{ underline: sortKey == column.jsonId}">
                         {{ column.label  }}
                    </a>
                </th>
            </tr>
            </thead>
            <tbody>
                <DetailTableLine  
                    v-for="(value, index) in filteredDockerImages " :key="index" 
                    class="border-b border-gray-300"
                    v-model="filteredDockerImages[index]"
                    v-bind:chart=chartData
                    v-bind:data=value
                    v-bind:config=config
                    v-bind:index="'di'+index"
                    @sendPayload="sendPayload"
                    @notificateValidation="notificateValidation" />
            </tbody>
        </table>
        </div>
    </div>
  </div>
  </loading>
</template>

<script>
import DetailCard from './DetailCards.vue'
import DetailTableLine from './DetailTableLine.vue'
import Loading from '../miscellaneous/Loading.vue'

export default {
  name: "Cards",
  components: { DetailCard, DetailTableLine, Loading },
  props: {
    chart: String
  },
  data() {
    return {
      loading: true,
      errored: false,
      sortKey: '',
      chartData: null,
      searchQuery: '',
      reverse: false,
      isToggleOn: false,
      quickFilter: 'All',
      columns: [{
          label:'Image Valid',
          jsonId:''
      },{
          label: 'Repo',
          jsonId: 'repository'
      },{
          label: 'Tags',
          jsonId: 'tag'
      },{
          label: 'Created',
          jsonId: 'created'
      },{
          label: 'Severity',
          jsonId: 'scan_overview'
      },{
          label: 'Versions',
          jsonId: 'labels'
      }]
    };
  },
  computed: {
     /* chartData () {
        let data;
        this.$store.dispatch("datastore/fetchChartData",{filter: "", quick: true})
        this.$store.state.datastore.charts.forEach(element => {
          if (element.name == this.chart) {
            data = element;
             data.allDockerImages.forEach(element => {
                if(!element.labels){
                    element.labels=[];
                }
            });
          }
        });
        return data;
      },*/
      config() {
        return this.$store.state.datastore.config
      },
      filteredDockerImages(){
          console.log('filter')
          console.log(this.chart)
       let data =  this.chartData.allDockerImages.filter((value)=>{
            var searchRegex = new RegExp(this.searchQuery, 'i')
            return (
                this.applyQuickFilter(value) &&
               ( searchRegex.test(value.repository) ||
                searchRegex.test(value.tag) ||
                searchRegex.test(value.scan_overview ? value.scan_overview["application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"].severity: ''))
            )
        })
        return this._.orderBy(data, this.sortKey,this.reverse?'desc':'asc')
      }
  },
  created() {
      this.$http.get(this.$store.state.datastore.backendUrl+"/v1/getChartList?filter="+this.chart+"&quick=false")
        .then((response)=>{
          console.log(response.data)
          this.chartData = response.data[0]
          this.chartData.charts.forEach((c)=>{
            if (!c.labels){
              c.labels=[]
            }
          })
          this.chartData.allDockerImages.forEach((d)=>{
            if (!d.labels){
              d.labels = []
            }
          })
        });
  },
  methods: {
    getAllDockerImages() {
      let data = { data: this.chartData.allDockerImages };
      return data;
    },
    sortBy(sortKey) {
      this.reverse = (this.sortKey == sortKey) ? ! this.reverse : false;
      this.sortKey = sortKey;
    },
    filterChange(value) {
      this.quickFilter=value
    },
    sendPayload(payload){
        this.$http.post(this.$store.state.datastore.backendUrl+payload.url, JSON.stringify(payload.data))
        .then((response)=> {
            this.$toast.open({
            message: ""+response.status,
            type: 'success',
        });
        })
        .catch((error) =>{
            this.$toast.open({
            message: error,
            type: 'error',
        });
        });
    },
    notificateValidation(data){
      this.$http.post(this.$store.state.datastore.backendUrl+"/v1/notificateValidation", JSON.stringify(data))
    },
    applyQuickFilter(di){
        if(!this.quickFilter || this.quickFilter === 'All'){
            return true;
        }
        if(this.quickFilter === 'latest'){
          let found = false;
          this.chartData.latestDockerImages.forEach(element => {
              if(element.repository === di.repository && element.tag === di.tag){
                  found = true
              }
          })
          return found;
        }else{
          let found = false;
          console.log(this.quickFilter.currentDockerImages)
          this.quickFilter.currentDockerImages.forEach(element => {
              if(element.repository === di.repository && element.tag === di.tag){
                  found = true
              }
          })
          return found;
        }
    }
  }
};
</script>
<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style scoped>
</style>