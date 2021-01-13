<template>
  <div>
    <div class="flex flex-col">
      <div class="flex flex-col p-1 bg-gray-700 rounded-lg shadow self-center w-1/2 z-10">
        <div
          v-if="project.editMode"
          class="flex items-center border-b border-b-2 border-teal-500 py-2"
        >
        <span class="text-white">Project_</span>
          <input
            v-model="editName"
            class="text-white appearance-none bg-transparent border-none w-full mr-3 py-1 px-2 leading-tight focus:outline-none"
            type="text"
            placeholder="Name"
            aria-label="Name"
          />
          <button
            class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
            type="button"
            v-on:click="validate(editName)"
          >Validate</button>
          <button
            class="flex-shrink-0 border-transparent border-4 text-teal-500 hover:text-teal-800 text-sm py-1 px-2 rounded"
            type="button"
          >Cancel</button>
        </div>
        <h2
          v-else
          class="text-lg capitalize text-white whitespace-no-wrap overflow-auto"
        >{{project.name}}</h2>
      </div>
      <div class="flex flex-col bg-white rounded-lg p-6 shadow-md -mt-4 relative">
        <button v-on:click="showSettings = true" class="v-step-7 bg-transparent absolute right-0 top-0 m-1">
          <svg class="fill-current w-6 h-6 text-gray-400 " xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" ><path d="M0 0h24v24H0V0z" fill="none"/><path d="M19.43 12.98c.04-.32.07-.64.07-.98 0-.34-.03-.66-.07-.98l2.11-1.65c.19-.15.24-.42.12-.64l-2-3.46c-.09-.16-.26-.25-.44-.25-.06 0-.12.01-.17.03l-2.49 1c-.52-.4-1.08-.73-1.69-.98l-.38-2.65C14.46 2.18 14.25 2 14 2h-4c-.25 0-.46.18-.49.42l-.38 2.65c-.61.25-1.17.59-1.69.98l-2.49-1c-.06-.02-.12-.03-.18-.03-.17 0-.34.09-.43.25l-2 3.46c-.13.22-.07.49.12.64l2.11 1.65c-.04.32-.07.65-.07.98 0 .33.03.66.07.98l-2.11 1.65c-.19.15-.24.42-.12.64l2 3.46c.09.16.26.25.44.25.06 0 .12-.01.17-.03l2.49-1c.52.4 1.08.73 1.69.98l.38 2.65c.03.24.24.42.49.42h4c.25 0 .46-.18.49-.42l.38-2.65c.61-.25 1.17-.59 1.69-.98l2.49 1c.06.02.12.03.18.03.17 0 .34-.09.43-.25l2-3.46c.12-.22.07-.49-.12-.64l-2.11-1.65zm-1.98-1.71c.04.31.05.52.05.73 0 .21-.02.43-.05.73l-.14 1.13.89.7 1.08.84-.7 1.21-1.27-.51-1.04-.42-.9.68c-.43.32-.84.56-1.25.73l-1.06.43-.16 1.13-.2 1.35h-1.4l-.19-1.35-.16-1.13-1.06-.43c-.43-.18-.83-.41-1.23-.71l-.91-.7-1.06.43-1.27.51-.7-1.21 1.08-.84.89-.7-.14-1.13c-.03-.31-.05-.54-.05-.74s.02-.43.05-.73l.14-1.13-.89-.7-1.08-.84.7-1.21 1.27.51 1.04.42.9-.68c.43-.32.84-.56 1.25-.73l1.06-.43.16-1.13.2-1.35h1.39l.19 1.35.16 1.13 1.06.43c.43.18.83.41 1.23.71l.91.7 1.06-.43 1.27-.51.7 1.21-1.07.85-.89.7.14 1.13zM12 8c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4zm0 6c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2z"/></svg>
        </button>
        <div class="flex items-center">
          <img class="h-16 w-16 md:h-24 md:w-24 mx-auto md:mx-0 md:mr-6" :src="icon" />
          <div class="flex flex-col overflow-auto">
            <table
              class="table-auto text-sm border-0 overflow-auto h-auto w-32"
            >
              <thead>
                <tr class="border-b-2 border-gray-500">
                  <th v-for="(column, index) in columns" :key="index" class="px-2 py-1">
                    {{ column.label }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(version, index) in project.versions" :key="index">
                  <td class="border-0 v-step-1">{{version.name}}</td>
                  <td class="border-0 ">
                    <router-link class="self-end m-2 v-step-2" :to="'/quick/'+version.name">
                      <button class="bg-blue-600 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded-full inline-flex items-center">
                        <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="bolt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><path fill="currentColor" d="M296 160H180.6l42.6-129.8C227.2 15 215.7 0 200 0H56C44 0 33.8 8.9 32.2 20.8l-32 240C-1.7 275.2 9.5 288 24 288h118.7L96.6 482.5c-3.6 15.2 8 29.5 23.3 29.5 8.4 0 16.4-4.4 20.8-12l176-304c9.3-15.9-2.2-36-20.7-36z"></path></svg>
                      </button>
                    </router-link>
                  </td>
                  <td class="border-0 ">
                    <button v-on:click="openModal(version)" class="v-step-3 bg-blue-600 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded-full inline-flex items-center">
                        <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="cogs"  role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"><path fill="currentColor" d="M512.1 191l-8.2 14.3c-3 5.3-9.4 7.5-15.1 5.4-11.8-4.4-22.6-10.7-32.1-18.6-4.6-3.8-5.8-10.5-2.8-15.7l8.2-14.3c-6.9-8-12.3-17.3-15.9-27.4h-16.5c-6 0-11.2-4.3-12.2-10.3-2-12-2.1-24.6 0-37.1 1-6 6.2-10.4 12.2-10.4h16.5c3.6-10.1 9-19.4 15.9-27.4l-8.2-14.3c-3-5.2-1.9-11.9 2.8-15.7 9.5-7.9 20.4-14.2 32.1-18.6 5.7-2.1 12.1.1 15.1 5.4l8.2 14.3c10.5-1.9 21.2-1.9 31.7 0L552 6.3c3-5.3 9.4-7.5 15.1-5.4 11.8 4.4 22.6 10.7 32.1 18.6 4.6 3.8 5.8 10.5 2.8 15.7l-8.2 14.3c6.9 8 12.3 17.3 15.9 27.4h16.5c6 0 11.2 4.3 12.2 10.3 2 12 2.1 24.6 0 37.1-1 6-6.2 10.4-12.2 10.4h-16.5c-3.6 10.1-9 19.4-15.9 27.4l8.2 14.3c3 5.2 1.9 11.9-2.8 15.7-9.5 7.9-20.4 14.2-32.1 18.6-5.7 2.1-12.1-.1-15.1-5.4l-8.2-14.3c-10.4 1.9-21.2 1.9-31.7 0zm-10.5-58.8c38.5 29.6 82.4-14.3 52.8-52.8-38.5-29.7-82.4 14.3-52.8 52.8zM386.3 286.1l33.7 16.8c10.1 5.8 14.5 18.1 10.5 29.1-8.9 24.2-26.4 46.4-42.6 65.8-7.4 8.9-20.2 11.1-30.3 5.3l-29.1-16.8c-16 13.7-34.6 24.6-54.9 31.7v33.6c0 11.6-8.3 21.6-19.7 23.6-24.6 4.2-50.4 4.4-75.9 0-11.5-2-20-11.9-20-23.6V418c-20.3-7.2-38.9-18-54.9-31.7L74 403c-10 5.8-22.9 3.6-30.3-5.3-16.2-19.4-33.3-41.6-42.2-65.7-4-10.9.4-23.2 10.5-29.1l33.3-16.8c-3.9-20.9-3.9-42.4 0-63.4L12 205.8c-10.1-5.8-14.6-18.1-10.5-29 8.9-24.2 26-46.4 42.2-65.8 7.4-8.9 20.2-11.1 30.3-5.3l29.1 16.8c16-13.7 34.6-24.6 54.9-31.7V57.1c0-11.5 8.2-21.5 19.6-23.5 24.6-4.2 50.5-4.4 76-.1 11.5 2 20 11.9 20 23.6v33.6c20.3 7.2 38.9 18 54.9 31.7l29.1-16.8c10-5.8 22.9-3.6 30.3 5.3 16.2 19.4 33.2 41.6 42.1 65.8 4 10.9.1 23.2-10 29.1l-33.7 16.8c3.9 21 3.9 42.5 0 63.5zm-117.6 21.1c59.2-77-28.7-164.9-105.7-105.7-59.2 77 28.7 164.9 105.7 105.7zm243.4 182.7l-8.2 14.3c-3 5.3-9.4 7.5-15.1 5.4-11.8-4.4-22.6-10.7-32.1-18.6-4.6-3.8-5.8-10.5-2.8-15.7l8.2-14.3c-6.9-8-12.3-17.3-15.9-27.4h-16.5c-6 0-11.2-4.3-12.2-10.3-2-12-2.1-24.6 0-37.1 1-6 6.2-10.4 12.2-10.4h16.5c3.6-10.1 9-19.4 15.9-27.4l-8.2-14.3c-3-5.2-1.9-11.9 2.8-15.7 9.5-7.9 20.4-14.2 32.1-18.6 5.7-2.1 12.1.1 15.1 5.4l8.2 14.3c10.5-1.9 21.2-1.9 31.7 0l8.2-14.3c3-5.3 9.4-7.5 15.1-5.4 11.8 4.4 22.6 10.7 32.1 18.6 4.6 3.8 5.8 10.5 2.8 15.7l-8.2 14.3c6.9 8 12.3 17.3 15.9 27.4h16.5c6 0 11.2 4.3 12.2 10.3 2 12 2.1 24.6 0 37.1-1 6-6.2 10.4-12.2 10.4h-16.5c-3.6 10.1-9 19.4-15.9 27.4l8.2 14.3c3 5.2 1.9 11.9-2.8 15.7-9.5 7.9-20.4 14.2-32.1 18.6-5.7 2.1-12.1-.1-15.1-5.4l-8.2-14.3c-10.4 1.9-21.2 1.9-31.7 0zM501.6 431c38.5 29.6 82.4-14.3 52.8-52.8-38.5-29.6-82.4 14.3-52.8 52.8z"></path></svg>
                        <Spinner v-if="loading.includes(version.name)" class="ml-1"/>
                    </button>
                  </td>
                  <td class="border-0 ">
                    <button v-on:click="$emit('detail',version.name)" class="v-step-4 bg-blue-600 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded-full inline-flex items-center">
                        <svg  class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path></svg>
                    </button>
                  </td>
                 <td class="border-0 px-4 py-2">
                    <button v-on:click="removeItem(version),$emit('delete',{label:version.label, name:name})" class="v-step-5 bg-red-600 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-full inline-flex items-center">
                      <svg class="fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="trash-alt"  role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M32 464a48 48 0 0 0 48 48h288a48 48 0 0 0 48-48V128H32zm272-256a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zm-96 0a16 16 0 0 1 32 0v224a16 16 0 0 1-32 0zM432 32H312l-9.4-18.7A24 24 0 0 0 281.1 0H166.8a23.72 23.72 0 0 0-21.4 13.3L136 32H16A16 16 0 0 0 0 48v32a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16z"></path></svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
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
        </div>
      </div>
    </div>
    <modal v-if="showModal" @close="showModal = false">
        <h1 slot="header" class="text-xl m-4 font-bold uppercase">Generation</h1>
        <div slot="body" >
          <div class="flex justify-around">
          <!--<button :disabled="selectedVersion.lastLink != ''" -->
            <button
          class="w-2/5 h-48 2xl:h-64 bg-gray-100 hover:bg-gray-200 shadow-md"
           v-on:click="generate(selectedVersion.name, 's3')">
            <svg class="m-auto" xmlns="http://www.w3.org/2000/svg" height="80px" viewBox="0 0 24 24" width="80px"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M19.35 10.04C18.67 6.59 15.64 4 12 4 9.11 4 6.6 5.64 5.35 8.04 2.34 8.36 0 10.91 0 14c0 3.31 2.69 6 6 6h13c2.76 0 5-2.24 5-5 0-2.64-2.05-4.78-4.65-4.96zM19 18H6c-2.21 0-4-1.79-4-4 0-2.05 1.53-3.76 3.56-3.97l1.07-.11.5-.95C8.08 7.14 9.94 6 12 6c2.62 0 4.88 1.86 5.39 4.43l.3 1.5 1.53.11c1.56.1 2.78 1.41 2.78 2.96 0 1.65-1.35 3-3 3zm-5.55-8h-2.9v3H8l4 4 4-4h-2.55z"/></svg>
            <span>Generate the deliveries and upload it on S3</span>
          </button>
          <button class="w-2/5 h-48 2xl:h-64 bg-gray-100 hover:bg-gray-200 shadow-md" v-on:click="generate(selectedVersion.name, 'script')">
            <svg class="m-auto" xmlns="http://www.w3.org/2000/svg" height="80px" viewBox="0 0 24 24" width="80px"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zM6 20V4h7v5h5v11H6z"/></svg>
            <span>Download the convenience script</span>
          </button>
          </div>
          <div v-if="selectedVersion.lastLink && selectedVersion.lastLink != ''">
            <button class="crossRotate m-auto mt-6 flex rounded-full items-center shadow-md text-white uppercase px-2 py-1 text-xs mr-3 bg-gray-600 bg-gradient-to-r from-gray-800 via-gray-600 to-gray-500" v-on:click="generateLink(selectedVersion.name)">
              <span class="mr-2 ">Download is ready </span>
              <span class="flex rounded-full mr-2 color uppercase px-1 py-1 text-xs font-bold shadow-md">
                <svg class="glyphe fill-current w-4 h-4" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="sync-alt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M370.72 133.28C339.458 104.008 298.888 87.962 255.848 88c-77.458.068-144.328 53.178-162.791 126.85-1.344 5.363-6.122 9.15-11.651 9.15H24.103c-7.498 0-13.194-6.807-11.807-14.176C33.933 94.924 134.813 8 256 8c66.448 0 126.791 26.136 171.315 68.685L463.03 40.97C478.149 25.851 504 36.559 504 57.941V192c0 13.255-10.745 24-24 24H345.941c-21.382 0-32.09-25.851-16.971-40.971l41.75-41.749zM32 296h134.059c21.382 0 32.09 25.851 16.971 40.971l-41.75 41.75c31.262 29.273 71.835 45.319 114.876 45.28 77.418-.07 144.315-53.144 162.787-126.849 1.344-5.363 6.122-9.15 11.651-9.15h57.304c7.498 0 13.194 6.807 11.807 14.176C478.067 417.076 377.187 504 256 504c-66.448 0-126.791-26.136-171.315-68.685L48.97 471.03C33.851 486.149 8 475.441 8 454.059V320c0-13.255 10.745-24 24-24z"></path></svg>
              </span>
            </button>
          </div>
          <div v-if="mode == 's3'">
           <div class="flex items-center border-b border-b-2 border-teal-500 py-2" >
              <input disabled class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" :value="'https://storage.cloud.google.com/catalog-deliveries/'+selectedVersion.name+'.tar.gz?authuser=1'" aria-label="Full name">
            </div>
            <div class="flex items-center border-b border-b-2 border-teal-500 py-2">
            <input disabled class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="" :value="'gs://catalog-deliveries/'+selectedVersion.name+'.tar.gz'" aria-label="Full name">
            </div>
          </div>
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

export default {
  name: "ProjectCards",
  components:{ Spinner, Modal},
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
      return this.project.owners.concat(this.project.members);
    }
  },
  data() {
    return {
      editName: "",
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
            data: JSON.stringify({label:version, mode:mode})
          })
          .then(response => {
            if(mode == 'script'){
              this.forceFileDownload(response,version)
            }
            this.loading.splice(this.loading.indexOf(version))
          })
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