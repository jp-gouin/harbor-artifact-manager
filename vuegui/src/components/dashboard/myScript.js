import Cards from './cards/Cards_new.vue'
import Loading from '../miscellaneous/Loading.vue'
export default {
    name: 'Dashboard',
    components: {
      Cards,
      Loading
    },
    props: {
      msg: String
    },
    computed: {
      chartData () {
        return this.$store.state.datastore.charts
      },
      config() {
        return this.$store.state.datastore.config
      }
    }
  }