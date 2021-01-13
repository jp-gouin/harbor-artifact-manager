<template>
    <div>
         <h1 class="text-2xl m-12 font-bold uppercase">Notification Center</h1>
        <input v-model="searchQuery" class="shadow appearance-none border rounded w-4/5 py-2  m-auto text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Search">
        <h2 class="text-xl m-4 font-bold uppercase">Stream Notification</h2>
        <table class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-4/5  m-auto">
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
                <tr v-for="(notification, index) in notifications " :key="index" >
                    <td class="border-0 px-4 py-2">
                        {{notification.date}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.owner}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.type}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.severity}}
                    </td>
                    <td class="border-0 px-4 py-2">
                      <span v-html="notification.payload"></span>
                    </td>
                </tr>
            </tbody>
        </table>
        <h2 class="text-xl m-4 mt-12 font-bold uppercase">Archived Notification</h2>
        <table class="table-auto text-sm border-0 overflow-auto h-8 sm:h-12 md:h-16 lg:h-20 xl:h-24 w-4/5  m-auto">
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
                <tr v-for="(notification, index) in notificationsList2 " :key="index" >
                    <td class="border-0 px-4 py-2">
                        {{notification.date}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.owner}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.type}}
                    </td>
                    <td class="border-0 px-4 py-2">
                        {{notification.severity}}
                    </td>
                    <td class="border-0 px-4 py-2">
                         <span v-html="notification.payload"></span>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    name: "NotificationCenter",
    data(){
        return({
            notificationsList2: [],
            sortKey: '',
            searchQuery: '',
            reverse: false,
            columns: [{
                label:'Date',
                jsonId:'date'
            },{
                label:'Owner',
                jsonId:'owner'
            },{
                label: 'Type',
                jsonId: 'type'
            },{
                label: 'Severity',
                jsonId: 'severity'
            },{
                label: 'Message',
                jsonId: 'payload'
            }]
        })
    },
    mounted: function () {
     this.refresh()
    },
    computed: {
        notifications(){
            let data =  this.$store.state.socketstore.notifications.filter((value)=>{
                var searchRegex = new RegExp(this.searchQuery, 'i')
                return (
                ( searchRegex.test(value.owner) ||
                    searchRegex.test(value.type) ||
                    searchRegex.test(value.severity))
                )
            })
            return this._.orderBy(data, this.sortKey,this.reverse?'desc':'asc')
        }
    },
    methods:{
        sortBy(sortKey) {
            this.reverse = (this.sortKey == sortKey) ? ! this.reverse : false;
            this.sortKey = sortKey;
        },
        refresh(){
          this.$http.get(this.$store.state.datastore.backendUrl+"/v1/getNotifications")
            .then((response) => {
                this.notificationsList2 = response.data
            })
            .catch(() => {
          });
        }
    }
}
</script>