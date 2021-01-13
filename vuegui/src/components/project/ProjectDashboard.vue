<template>
  <div >
    <h1 class="v-step-0 v-step-9 text-2xl mt-12 font-bold uppercase">Projects</h1>
    <Loading v-if="!projects || projects.length == 0"/>
    <Annoucement/>
    <div class="md:flex flex-wrap -mb-4">
      <ProjectCards
        v-if="showTour"
        class="p-2 h-full content-center lg:w-1/3 xl:w-1/3"
        :key="'projectTour'"
        v-model="dummyValue"
        v-bind:icon="iconlist[0]"
        v-bind:project="dummyValue"
      />
      <ProjectCards v-else
        class=" p-2 h-full content-center lg:w-1/3 xl:w-1/3"
        v-for="(project, index) in projects"
        :key="index+project.name"
        v-model="projects[index]"
        v-bind:icon="iconlist[index % iconlist.length]"
        v-bind:project="project"
        @createProject="createProject"
        @detail="detail"
        @delete="removeProject"
        @updateMemberlist="updateMemberlist"
      />
      <modal v-if="showModal" @close="showModal = false">
        <h1 slot="header" class="text-xl m-4 font-bold uppercase">{{version}}</h1>
        <table
          v-if="details"
          slot="body"
          class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-32 m-auto"
        >
          <thead>
            <tr class="border-b-2 border-gray-500">
              <th
                v-for="(column, index) in columns"
                :key="index"
                class="px-4 py-2"
              >{{ column.label }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(app, index) in details.charts" :key="index">
              <td class="border-0 px-4 py-2">{{app.metadata.name}}</td>
              <td class="border-0 px-4 py-2">{{app.version}}</td>
              <td class="border-0 px-4 py-2">{{app.metadata.appVersion}}</td>
              <!--  <td class="border-0 px-4 py-2"><li v-for="(di, indexdi) in app.dockerImages" :key="indexdi">
                    
              </li></td>-->
            </tr>
          </tbody>
        </table>
      </modal>
    </div>
    <button
      v-on:click="addProject()"
      class="v-step-6 p-2 mt-4 text-2xl h-full w-1/2 text-center items-center content-center border-gray-500 border-4 border-dashed bg-gray-300 hover:border-gray-700 hover:bg-gray-400 text-gray-800 font-bold rounded-lg flex-1 justify-center inline-flex items-center"
    >
      <svg
        class="fill-current w-10 h-10 mr-2"
        aria-hidden="true"
        focusable="false"
        data-prefix="fas"
        data-icon="plus"
        role="img"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 448 512"
      >
        <path
          fill="currentColor"
          d="M416 208H272V64c0-17.67-14.33-32-32-32h-32c-17.67 0-32 14.33-32 32v144H32c-17.67 0-32 14.33-32 32v32c0 17.67 14.33 32 32 32h144v144c0 17.67 14.33 32 32 32h32c17.67 0 32-14.33 32-32V304h144c17.67 0 32-14.33 32-32v-32c0-17.67-14.33-32-32-32z"
        />
      </svg>
      <span>
        Add
        <br />New
        <br />Project
      </span>
    </button>
    <v-tour name="projectTour" :steps="steps" :onSkip="endTour" :onFinish="endTour" :callbacks="myCallbacks"></v-tour>
  </div>
</template>
<script>
import ProjectCards from "./ProjectCard_new.vue";
import Modal from "./Modal.vue";
import Loading from "../miscellaneous/Loading.vue"
import Annoucement from "../miscellaneous/Annoucement.vue"

export default {
  name: "ProjectDashboard",
  components: { ProjectCards, Modal, Loading, Annoucement },
  methods: {
    detail(version) {
      this.version = version;
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/getProjectDetail",
          JSON.stringify(version)
        )
        .then((response) => {
            this.showModal = true;
            this.details = response.data;
        })
        .catch(() => {
          this.showModal = true;
          this.details = {
            charts: []
          };
        });
    },
    removeProject(data) {
      //this.projects.splice(this.projects.indexOf(data.label), 1);
      console.log(data)
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/removeProject",
          JSON.stringify(data.label)
        )
        .then(response => {
          // this.$store.commit("datastore/removeProject",data.name,this.projects[data.name]);
          this.$toast.open({
            message: response.statusText,
            type: "success"
          });
          this.$store.dispatch("datastore/fetchProjects");
        })
        .catch(error => {
          this.$toast.open({
            message: error.response.data,
            type: "error"
          });
        });
    },
    addProject() {
      this.$store.commit("datastore/addEditProject");
    },
    createProject(projectName) {
      this.$toast.open({
        message: "Creation",
        type: "info"
      });
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/postProject",
          JSON.stringify(projectName)
        ).then(()=>{
          this.$store.dispatch("datastore/fetchProjects");
        })
        .catch(error => {
          this.$toast.open({
            message: error.response.data,
            type: "error"
          });
        });
    },
    updateMemberlist(project){
      this.$http
        .post(
          this.$store.state.datastore.backendUrl + "/v1/updateProjectMembers",
          JSON.stringify(project)
        )
        .catch(error => {
          this.$toast.open({
            message: error.response.data,
            type: "error"
          });
        });
    },
    endTour(){
      this.showTour=false;
      localStorage.setItem('projectTour', true)
    }
  },
  data() {
    return {
      iconlist: [
        "https://image.flaticon.com/icons/svg/86/86572.svg",
        "https://image.flaticon.com/icons/svg/86/86580.svg",
        "https://image.flaticon.com/icons/svg/813/813488.svg",
        "https://image.flaticon.com/icons/svg/813/813430.svg",
        "https://image.flaticon.com/icons/svg/813/813466.svg",
        "https://image.flaticon.com/icons/svg/813/813531.svg",
        "https://image.flaticon.com/icons/svg/2034/2034786.svg",
        "https://image.flaticon.com/icons/svg/813/813491.svg",
        "https://image.flaticon.com/icons/svg/81/81089.svg"
      ],
      showModal: false,
      showTour: false,
      dummyValue: {
          "name": "Example",
          "versions": [{
            "name": "Project_Example_Version_0",
            "charts": ["etcd4.4.9", "etcd4.4.0"],
            "dockerImages": [],
            "lastLink": "",
            "label": {
              "name": "Project_Example_Version_0",
            }
          }, {
            "name": "Project_Example_Version_1",
            "charts": [],
            "dockerImages": [],
            "lastLink": "",
            "label": {
              "name": "Project_Example_Version_1",
            }
          }],
          "owners": ["exowner"],
          "members": ["exmember"]
        },
      myCallbacks: {
        onSkip: this.endTour,
        onFinish: this.endTour
      },
      steps: [
          {
            target: '.v-step-0',  // We're using document.querySelector() under the hood
            content: 'Welcome to the <strong>Tour</strong>!, <br><br>You can manage all your projects here <br>Creation, Configuration, Build and more <br><br>Follow the guide to discover all functionnalities :) <br><br>If you have any question contact Jean-Philippe Gouin',
            params: {
              highlight: true,
            }
          },
          {
            target: '.v-step-1',
            content: '<h2 class="font-bold">Versionning</h2>All projects are versionned , and for each version you can perform some actions. <br>You can manage all your projects here',
            params: {
              placement: 'top' // Any valid Popper.js placement. See https://popper.js.org/popper-documentation.html#Popper.placements
            }
          },
          {
            target: '.v-step-2',
            content: '<h2 class="font-bold">Quick configuration </h2>This allows you to add <strong>Artifacts</strong> to your project easily',
            params: {
              placement: 'top' // Any valid Popper.js placement. See https://popper.js.org/popper-documentation.html#Popper.placements
            }
          },{
            target: '.v-step-3',  // We're using document.querySelector() under the hood
            content: '<h2 class="font-bold">Build allows you to build your version </h2> <ul class="list-disc mx-1"> <li> Using GCP and store the package in GCS (need access to G3S organisation , please contact an administrator ) </li> <li> Downlading the convenience script </li> </ul>',
            params: {
              placement: 'top' // Any valid Popper.js placement. See https://popper.js.org/popper-documentation.html#Popper.placements
            }
          },{
            target: '.v-step-4',  // We're using document.querySelector() under the hood
            content: `<h2 class="font-bold">Details </h2>This allows you to see what artifact are in the selected version`,
            params: {
              placement: 'top' // Any valid Popper.js placement. See https://popper.js.org/popper-documentation.html#Popper.placements
            }
          },{
            target: '.v-step-5',  // We're using document.querySelector() under the hood
            content: `Delete a version, only authorized user can perform this action <br> If you try , admin and owner/member of the project will be notified ;)`,
            params: {
              placement: 'top' // Any valid Popper.js placement. See https://popper.js.org/popper-documentation.html#Popper.placements
            }
          },{
            target: '.v-step-6',  // We're using document.querySelector() under the hood
            content: `Add you own project to the list ! <br> naming convention <strong>forbid</strong> you to use <strong>_</strong> in the name of the project`
          },{
            target: '.v-step-7',  // We're using document.querySelector() under the hood
            content: `Customize setting of a project by adding more members or to promote member to owners !`
          },{
            target: '.v-step-8',  // We're using document.querySelector() under the hood
            content: `If you project change or a project where you are a member, you'll be notify here !`
          },{
            target: '.v-step-9',  // We're using document.querySelector() under the hood
            content: `Thanks for using the catalog, have fun !`,
             params: {
              highlight: true,
            }
          },
        ],
      version: "",
      columns: [
        {
          label: "App",
          jsonId: ""
        },
        {
          label: "ChartVersion",
          jsonId: ""
        },
        {
          label: "AppVersion",
          jsonId: ""
        }
      ]
    };
  },
  computed: {
    projects() {
      return this.$store.state.datastore.projects;
    },
    config() {
      return this.$store.state.datastore.config;
    }
  },
  mounted: function () {
     if (!localStorage.getItem('projectTour')){
          this.showTour = true
          this.$tours['projectTour'].start()
      }
    }
};
</script>
<style scoped>
.v-tour__target--highlighted {
  box-shadow: 0 0 0 99999px rgba(0,0,0,.4);
}
</style>