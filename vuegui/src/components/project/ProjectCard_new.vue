<template>
  <div>
    <Cardbase  :icon="icon" :title="project.name" :editmode="project.editMode" @title="validate" @cancel="$store.commit('datastore/cancelEditProject')" >
      <div slot="subheader" class="m-1">
          <span class="text-xs text-gray-600">Access : </span>
          <span>{{ownership}}</span>
      </div>
      <div slot="rightheader">
        <button v-on:click="showSettings = true"  class="mr-1 m-auto flex rounded-full items-center  text-xs bg-gray-200 border border-gray-400">
          <span class="flex rounded-full mr-1 color text-gray-800 uppercase px-1 py-1 text-xs font-bold">
            <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="users" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"><path fill="currentColor" d="M96 224c35.3 0 64-28.7 64-64s-28.7-64-64-64-64 28.7-64 64 28.7 64 64 64zm448 0c35.3 0 64-28.7 64-64s-28.7-64-64-64-64 28.7-64 64 28.7 64 64 64zm32 32h-64c-17.6 0-33.5 7.1-45.1 18.6 40.3 22.1 68.9 62 75.1 109.4h66c17.7 0 32-14.3 32-32v-32c0-35.3-28.7-64-64-64zm-256 0c61.9 0 112-50.1 112-112S381.9 32 320 32 208 82.1 208 144s50.1 112 112 112zm76.8 32h-8.3c-20.8 10-43.9 16-68.5 16s-47.6-6-68.5-16h-8.3C179.6 288 128 339.6 128 403.2V432c0 26.5 21.5 48 48 48h288c26.5 0 48-21.5 48-48v-28.8c0-63.6-51.6-115.2-115.2-115.2zm-223.7-13.4C161.5 263.1 145.6 256 128 256H64c-35.3 0-64 28.7-64 64v32c0 17.7 14.3 32 32 32h65.9c6.3-47.4 34.9-87.3 75.2-109.4z"></path></svg>
          </span>
          <span class="mr-2 ">{{userList.length}}</span>
        </button>
      </div>
      <div slot="content" class="flex flex-col flex-wrap mt-2 w-full content-around">
        <div v-for="(version, index) in project.versions" :key="index" class="border w-full bg-gray-200 rounded p-2  m-1  flex flex-row justify-between leading-normal">
          <span class="text-sm self-center break-words w-1/3">{{version.name}}</span>
          <router-link class="self-center v-step-2" :to="'/quick/'+version.name">
            <button class="btn-secondary font-bold py-1 px-2 rounded inline-flex items-center">
              <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="bolt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><path fill="currentColor" d="M296 160H180.6l42.6-129.8C227.2 15 215.7 0 200 0H56C44 0 33.8 8.9 32.2 20.8l-32 240C-1.7 275.2 9.5 288 24 288h118.7L96.6 482.5c-3.6 15.2 8 29.5 23.3 29.5 8.4 0 16.4-4.4 20.8-12l176-304c9.3-15.9-2.2-36-20.7-36z"></path></svg>
            </button>
          </router-link>
          <button v-on:click="openModal(version)" class="v-step-3 self-center btn-secondary font-bold py-1 px-2 rounded inline-flex items-center">
              <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="download" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M216 0h80c13.3 0 24 10.7 24 24v168h87.7c17.8 0 26.7 21.5 14.1 34.1L269.7 378.3c-7.5 7.5-19.8 7.5-27.3 0L90.1 226.1c-12.6-12.6-3.7-34.1 14.1-34.1H192V24c0-13.3 10.7-24 24-24zm296 376v112c0 13.3-10.7 24-24 24H24c-13.3 0-24-10.7-24-24V376c0-13.3 10.7-24 24-24h146.7l49 49c20.1 20.1 52.5 20.1 72.6 0l49-49H488c13.3 0 24 10.7 24 24zm-124 88c0-11-9-20-20-20s-20 9-20 20 9 20 20 20 20-9 20-20zm64 0c0-11-9-20-20-20s-20 9-20 20 9 20 20 20 20-9 20-20z"></path></svg>
              <Spinner v-if="loading.includes(version.name)" class="ml-1"/>
          </button>
           <button v-on:click="$emit('detail',version.name)" class="v-step-4 self-center btn-secondary font-bold py-1 px-2 rounded inline-flex items-center">
              <svg  class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path></svg>
          </button>
            <button v-on:click="removeItem(version),$emit('delete',{label:version.label, name:name})" class="v-step-5 self-center bg-red-600 hover:bg-red-800 text-white font-bold py-2 px-4 rounded inline-flex items-center">
            <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="trash-alt"  role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M32 464a48 48 0 0 0 48 48h288a48 48 0 0 0 48-48V128H32zm272-256a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zM432 32H312l-9.4-18.7A24 24 0 0 0 281.1 0H166.8a23.72 23.72 0 0 0-21.4 13.3L136 32H16A16 16 0 0 0 0 48v32a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16z"></path></svg>
          </button>
        </div>
            <div
              class="flex items-center border-b border-b-2 border-teal-500 py-2 w-3/4 self-center"
            >
            <span >Project_{{project.name}}_Version_</span>
              <input
                v-model="version"
                class="appearance-none bg-transparent border-none w-1/2 mr-3 py-1 px-2 leading-tight focus:outline-none"
                type="number"
                placeholder="Version"
                aria-label="Name"
              />
              <button
                class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
                type="button"
                v-on:click="validate(project.name)">
                  <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"></path></svg>
              </button>
            </div>
      </div>
    </Cardbase>
    <modal v-if="showModal" @close="showModal = false" v-bind:large="true">
        <h1 slot="header" class="text-xl m-4 font-bold uppercase">Generation</h1>
        <div slot="body" >
          <div class="flex justify-around">
          <!--<button :disabled="selectedVersion.lastLink != ''" -->
            <div >
              <button
            class="h-48 w-1/2 bg-gray-100 hover:bg-gray-200 shadow-md"
            v-on:click="generate(selectedVersion.name, 's3')">
              <svg class="m-auto" xmlns="http://www.w3.org/2000/svg" height="80px" viewBox="0 0 24 24" width="80px"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M19.35 10.04C18.67 6.59 15.64 4 12 4 9.11 4 6.6 5.64 5.35 8.04 2.34 8.36 0 10.91 0 14c0 3.31 2.69 6 6 6h13c2.76 0 5-2.24 5-5 0-2.64-2.05-4.78-4.65-4.96zM19 18H6c-2.21 0-4-1.79-4-4 0-2.05 1.53-3.76 3.56-3.97l1.07-.11.5-.95C8.08 7.14 9.94 6 12 6c2.62 0 4.88 1.86 5.39 4.43l.3 1.5 1.53.11c1.56.1 2.78 1.41 2.78 2.96 0 1.65-1.35 3-3 3zm-5.55-8h-2.9v3H8l4 4 4-4h-2.55z"/></svg>
              <span>Generate the deliveries and upload it on S3</span>
            </button>
            <div>
            <input type="checkbox" class="mt-4" id="checkbox" v-model="force">
            <span> Force the download</span>
            </div>
            <div class=" m-2">
              <h2>Quickstart</h2>
              <div class="text-xs flex flex-col text-left m-2">
                <span>1. Download the archive</span>
                <span>2. Copy and untar the datapack</span>
                <span class="bg-gray-900 text-gray-200 rounded-lg p-2 m-1 shadow-md">tar -xzf {{selectedVersion.name}}.tar.gz</span>
                <span>3. Load the datapack</span>
                <span class="bg-gray-900 text-gray-200 rounded-lg p-2 m-1 shadow-md">./{{selectedVersion.name}}.sh --load true -a <span class="italic">harbor_url</span></span>
              </div>
              <div v-if="selectedVersion.lastLink && selectedVersion.lastLink != ''">
                <button class="crossRotate m-auto mt-6 flex rounded-lg items-center shadow-xl text-white uppercase px-2 py-1 text-xs mr-3 bg-gray-600 bg-gradient-to-r from-gray-800 via-gray-600 to-gray-500" v-on:click="generateLink(selectedVersion.name)">
                  <span class="mr-2 ">Download is ready </span>
                  <span class="flex rounded-full mr-2 color uppercase px-1 py-1 text-xs font-bold shadow-md bg-gray-700">
                    <svg class="glyphe fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="sync-alt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M370.72 133.28C339.458 104.008 298.888 87.962 255.848 88c-77.458.068-144.328 53.178-162.791 126.85-1.344 5.363-6.122 9.15-11.651 9.15H24.103c-7.498 0-13.194-6.807-11.807-14.176C33.933 94.924 134.813 8 256 8c66.448 0 126.791 26.136 171.315 68.685L463.03 40.97C478.149 25.851 504 36.559 504 57.941V192c0 13.255-10.745 24-24 24H345.941c-21.382 0-32.09-25.851-16.971-40.971l41.75-41.749zM32 296h134.059c21.382 0 32.09 25.851 16.971 40.971l-41.75 41.75c31.262 29.273 71.835 45.319 114.876 45.28 77.418-.07 144.315-53.144 162.787-126.849 1.344-5.363 6.122-9.15 11.651-9.15h57.304c7.498 0 13.194 6.807 11.807 14.176C478.067 417.076 377.187 504 256 504c-66.448 0-126.791-26.136-171.315-68.685L48.97 471.03C33.851 486.149 8 475.441 8 454.059V320c0-13.255 10.745-24 24-24z"></path></svg>
                  </span>
                </button>
              </div>
            </div>
            </div>
            <div>
              <button class="w-1/2 h-48 bg-gray-100 hover:bg-gray-200 shadow-md" v-on:click="generate(selectedVersion.name, 'script')">
                <svg class="m-auto" xmlns="http://www.w3.org/2000/svg" height="80px" viewBox="0 0 24 24" width="80px"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zM6 20V4h7v5h5v11H6z"/></svg>
                <span>Download the convenience script</span>
              </button>
              <div class=" m-2">
                <h2>Quickstart</h2>
                <div class="text-xs flex flex-col text-left">
                <span>1. Download the script</span>
                <span>2. This will create the datapack</span>
                <span class="bg-gray-900 text-gray-200 rounded-lg p-2 m-1 shadow-md">./{{selectedVersion.name}}.sh </span>
                <span>3. Copy and untar the datapack</span>
                <span class="bg-gray-900 text-gray-200 rounded-lg p-2 m-1 shadow-md">tar -xzf {{selectedVersion.name}}.tar.gz</span>
                <span>4. Load the datapack</span>
                <span class="bg-gray-900 text-gray-200 rounded-lg p-2 m-1 shadow-md">./{{selectedVersion.name}}.sh --load true -a <span class="italic">harbor_url</span></span>
              </div>
              </div>
            </div>
          </div>
          
          <!-- <div v-if="mode == 's3'">
           <div class="flex items-center border-b border-b-2 border-teal-500 py-2" >
              <input disabled class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" :value="'https://storage.cloud.google.com/catalog-deliveries/'+selectedVersion.name+'.tar.gz?authuser=1'" aria-label="Full name">
            </div>
            <div class="flex items-center border-b border-b-2 border-teal-500 py-2">
            <input disabled class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="" :value="'gs://catalog-deliveries/'+selectedVersion.name+'.tar.gz'" aria-label="Full name">
            </div>
          </div> -->
        </div>
      </modal>
      <modal v-if="showSettings" @close="showSettings = false">
        <h1 slot="header" class="text-xl m-4 font-bold uppercase">Settings</h1>
        <div slot="body" >
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2 mt-4 text-left" for="grid-last-name">
              Memberlist
            </label>
            <table
              class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-full"
            >
              <thead>
                <tr class="border-b-2 border-gray-500">
                  <th  class="px-2 py-1">
                    User
                  </th>
                  <th  class="px-2 py-1">
                    Rigth
                  </th>
                  <th  class="px-2 py-1">
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(user, index) in userList" :key="index+user">
                  <td class="border-0 ">{{user}}</td>
                  <td class="border-0 ">
                    <div class="relative">
                      <select v-on:change="changeMembership(user, $event.target.value)"  class="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="grid-state">
                        <option :selected="project.owners.includes(user)">Owner</option>
                        <option :selected="project.members.includes(user)">Member</option>
                      </select>
                      <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/></svg>
                      </div>
                    </div>
                  </td>
                   <td class="border-0 px-4 py-2">
                    <button v-on:click="removeUser(user)" class="bg-red-600 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-full inline-flex items-center">
                      <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="trash-alt"  role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M32 464a48 48 0 0 0 48 48h288a48 48 0 0 0 48-48V128H32zm272-256a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zM432 32H312l-9.4-18.7A24 24 0 0 0 281.1 0H166.8a23.72 23.72 0 0 0-21.4 13.3L136 32H16A16 16 0 0 0 0 48v32a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16z"></path></svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          <label class="block uppercase tracking-wide text-gray-700 text-xs mt-4" for="grid-password">
            Hint : Owner have full right on a project and it's members.
                   Member have only read access to the project and can trigger generation
          </label>
          <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2 mt-4 text-left" for="grid-last-name">
              Add user to the project
            </label>
         <div class="flex items-center border-b border-b-2 border-teal-500 py-2 ">
            <input v-model="newUser" class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="Username" aria-label="Full name">
            <button v-on:click="addUser()" class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded" type="button">
              Add
            </button>
          </div>
          <button v-on:click="validateMembers(); showSettings = false" class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded mt-4" type="button">
              Validate
            </button>
        </div>
      </modal>
  </div>
</template>
<script>
import Spinner from './Spinner.vue'
import Modal from "./Modal.vue";
import Cardbase from "../dashboard/cards/Cards_base"
export default {
  name: "ProjectCards",
  components:{ Spinner, Modal, Cardbase},
  props: {
    icon: String,
    project: Object,
    value: Object,
  },
  computed:{
    dataTest(){
      return this.project
    },
    userList(){
      if(!this.project.owners){
        return 0;
      }
      return this.project.owners.concat(this.project.members);
    },
    user () {
      return this.$store.state.authentication.user;
    },
    config(){
      return this.$store.state.datastore.config;
    },
    ownership(){
      if (this.user.id && this.config.id){
        return 'Admin'
      }
      var r = this.project.owners.find((e)=>{
        if (e.name == this.user.name){
          return e;
        }
      })
      if(r){
        return 'Owner'
      }else{
        return 'Member'
      }
    }
  },
  data() {
    return {
      editName: "",
      force: false,
      version: "0",
      selectedVersion: Object,
      loading: [],
      files: new Map(),
      showModal: false,
      showSettings: false,
      newUser: "",
      columns: [
            {
              label: "Version",
              jsonId: ""
            },{
              label: "Quick Setup",
              jsonId: ""
            },
            {
              label: "Build Delivery",
              jsonId: ""
            },
            {
              label: "Details",
              jsonId: ""
            },
            {
              label: "Delete",
              jsonId: ""
            }
            
          ]
    }
  },
  methods:{
      openModal(version){
        this.mode = ""
        this.showModal = true;
        this.selectedVersion = version;
      },
      forceFileDownload(response, filename){
        const url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', filename + '_delivery-script-v1.0.sh') //or any other extension
        document.body.appendChild(link)
        link.click()
      },
      generateLink(name){
        this.$http.get(this.$store.state.datastore.backendUrl+"/v1/refreshDownloadLink?project="+name)
          .then((response) => {
              const link = document.createElement('a')
              link.href = response.data
              document.body.appendChild(link)
              link.click()
          })
          .catch(() => {
          });
      },
      removeItem(label){
        this.project.versions.splice(this.project.versions.indexOf(label),1)
      },
      generate(version, mode){
        
        this.mode = mode
        this.showModal = false;
      /*  if(this.mode == 'script'){
            
        }else{
          this.selectedVersion.lastLink = '';
        }*/
        if(this.loading.indexOf(version) === -1 ){
          this.loading.push(version)
          }
          this.$http({
            method: 'post',
            url: this.$store.state.datastore.backendUrl +'/v1/generateDelivery',
            responseType: 'arraybuffer',
            data: JSON.stringify({label:version, mode:mode, force:this.force})
          })
          .then(response => {
            if(mode == 'script'){
              this.forceFileDownload(response,version)
            }
            this.loading.splice(this.loading.indexOf(version))
          }).catch(() => {
            this.loading.splice(this.loading.indexOf(version))
            this.$toast.open({
              message: "Build already in progress",
              type: "error",
            });
        });
      },
      validate(name){
        this.$emit('createProject','Project_'+name.replace('_','')+'_Version_'+this.version)
      },
      changeMembership(user, value){
        if(value === 'Member'){
          this.project.owners.splice(this.project.owners.indexOf(user), 1);
          this.project.members.push(user)
        }else{
          this.project.members.splice(this.project.members.indexOf(user), 1);
          this.project.owners.push(user)
        }
      },
      addUser(){
        this.project.members.push(this.newUser)
      },
      validateMembers(){
        console.log('send change to server')
        console.log(this.project)
        this.$emit('updateMemberlist',this.project)
      },
      removeUser(user){
        if(this.project.owners.indexOf(user) != -1){
          this.project.owners.splice(this.project.owners.indexOf(user), 1);
        }
        if(this.project.members.indexOf(user) != -1){
          this.project.members.splice(this.project.members.indexOf(user), 1);
        }
      }
  }
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
.crossRotate {
  margin: auto;
  margin-top: 1rem;
  background: linear-gradient(90deg, rgba(45,55,72,1) 0%, rgba(74,85,104,1) 70%, rgba(113,128,150,1) 100%);
}
.crossRotate:hover .glyphe {    
      -webkit-transform: rotateZ(1080deg);
      -moz-transform: rotateZ(1080deg);
      transform: rotateZ(1080deg);
  }
  .glyphe{ 
    -webkit-transition: 3s ease-out;
    -moz-transition:  3s ease-out;
    transition:  3s ease-out;
  }
</style>