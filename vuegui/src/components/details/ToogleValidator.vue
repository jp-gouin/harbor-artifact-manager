<template>
  <ToogleSwitch v-model="isElementValid" :index="index" :onText="onText" :offText="offText" />
</template>

<script>
import ToogleSwitch from "../ToogleSwitch.vue";
export default {
  name: "ToogleValidator",
  components: { ToogleSwitch },
  props: {
    data: Object,
    config: Object,
    index: String,
    onText: String,
    offText: String
  },
  computed: {
    isElementValid: {
      get: function() {
        let valid = false;
        this.data.labels.forEach(label => {
          if (
            label.name === this.config.configlabel.name &&
            label.id === this.config.configlabel.id &&
            !label.deleted
          ) {
            valid = true;
          }
        });
        return valid;
      },
      set: function() {
        let found = false;
        this.data.labels.forEach(label => {
          if (
            label.name === this.config.configlabel.name &&
            label.id === this.config.configlabel.id
          ) {
            found = true;
            label.deleted = !label.deleted;
          }
        });
        if (!found) {
          this.data.labels.push(this.config.configlabel);
        }
        this.$emit("toogleValidation");
      }
    }
  }
};
</script>