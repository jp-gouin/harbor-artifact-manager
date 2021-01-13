<template>
    <multiselect
          v-model="getFilteredLabels"
          :options="config.projects"
          :multiple="true"
          :close-on-select="false"
          :clear-on-select="false"
          :preserve-search="true"
          placeholder="Pick some"
          @remove="removelabel"
          @select="selectlabel"
          label="name"
          track-by="name"
        ></multiselect>
</template>

<script>
import Multiselect from "vue-multiselect";
export default {
  name: "PackageSelect",
  components: {  Multiselect },
  props: {
    value: Object,
    config: Object,
  },
  methods: {
    removelabel(option) {
      this.value.labels.forEach(label => {
        if (label.name === option.name && label.id === option.id) {
          label.deleted = true;
        }
      });
      this.$emit("input", this.value);
      this.$emit("removeLabel", option);
      
    },
    selectlabel(select) {
      select.deleted=false;
      if(!this.value.labels){
        this.value.labels=[]
        this.value.labels.push(select)
      }
      let updated = false
      this.value.labels.forEach(label => {
        if (label.name === select.name && label.id === select.id) {
          label.deleted = false
          updated=true
        }
      });
      if(!updated){
        this.value.labels.push(select)
      }
      this.$emit("input", this.value);
      this.$emit("selectLabel", select);
    }
  },
  computed: {
    getFilteredLabels: {
      get() {
        return this.value.labels.filter(label => {
          return (
            label.name != this.config.configlabel.name &&
            label.id != this.config.configlabel.id &&
            !label.deleted
          );
        });
      },
      set(newValue) {
          newValue.forEach((newLabel)=>{
              this.value.labels.forEach((oldLabel,index,object) => {
                  if(oldLabel.name === newLabel.name && oldLabel.id === newLabel.id){
                      object.splice(index, 1);
                  }
              });
          });
        this.value.labels.push(...newValue);
      }
    }
  }
}
</script>