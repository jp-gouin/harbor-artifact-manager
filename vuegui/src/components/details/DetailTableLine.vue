<template>
  <tr>
    <td class="border-0 px-4 py-2">
      <ToogleValidator 
        v-bind:data=value
        :index="index" 
        onText 
        offText 
        :config="config"
        @toogleValidation="toogleValidation"/>
    </td>
    <td class="border-0 px-4 py-2">{{value.repository}}</td>
    <td class="border-0 px-4 py-2">{{value.tag}}</td>
    <td class="border-0 px-4 py-2">{{value.created | formatDate}}</td>
    <td
      class="border-0 px-4 py-2"
    >{{value.scan_overview?value.scan_overview["application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"].severity:''}}</td>
    <td class="border-0 px-4 py-2">
      <div class="inline-block relative w-64" v-if="config">
          <PackageSelect 
            v-model="value"
            v-bind:config="config"
            @removeLabel="removelabel"
            @selectLabel="selectlabel"
          />
      </div>
    </td>
  </tr>
</template>

<script>
import PackageSelect from "./PackageSelect.vue"
import ToogleValidator from "./ToogleValidator.vue"
export default {
  name: "DetailTableLine",
  components: { PackageSelect,ToogleValidator },
  props: {
    value: Object,
    config: Object,
    index: String,
    chart: Object
  },
  data() {
    return {};
  },
  methods: {
    toogleValidation(){
      console.log(this.value)
        let data = {
          name: this.chart.name,
          allDockerImages: [this.value],
          charts: [],
          project: this.chart.project
        };
        this.$emit("sendPayload", { data: data, url: "/v1/postChartData" });
        this.$emit("notificateValidation",this.value.repository+":"+this.value.tag)
        this.$emit("input", this.value);
    },
    removelabel(option) {
      let data = {
        name: this.chart.name,
        allDockerImages: [this.value],
        charts: [],
        project: this.chart.project,
        projectLab: option.name
      };
      this.$emit("sendPayload", {
        data: data,
        url: "/v1/removeProjectToArtifact"
      });
      this.$emit("input", this.value);
    },
    selectlabel(option) {
      console.log(option)
      let data = {
        name: this.chart.name,
        allDockerImages: [this.value],
        charts: [],
        project: this.chart.project,
        projectLab: option.name
      };
      console.log(JSON.stringify(data))
      this.$emit("sendPayload", {
        data: data,
        url: "/v1/addProjectToArtifact"
      });
      this.$emit("input", this.value);
    }
  },
};
</script>