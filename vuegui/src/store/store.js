import Vue from 'vue'
//import Axios from 'axios' 
//import { authHeader } from '../services/auth.header';
export const datastore = {
    namespaced: true,
    state: {
        charts: [],
        chartsFull: [],
        config: {},
        projects: [],
        backendUrl: "",
        tokenUrl: "",
        chartData: [],
        annoucements: [],
        starredArtifacts: [],
        userInfo: {}
    },
    mutations: {
        FETCH_CHARTS(state, charts) {
            state.charts = charts
        },
        FETCH_FULL_CHARTS(state, chart){
            state.chartsFull = chart
        },
        FETCH_CONFIG(state, config) {
            state.config = config
            if(config.id){
                localStorage.setItem('configid', config.id);
            }
        },
        FETCH_PROJECTS(state, projects) {
            if(projects){
                state.projects = projects
            }
        },
        FETCH_ANNOUCEMENTS(state, annoucements){
            if(annoucements){
                state.annoucements = annoucements
            }
        },
        SET_PROJECT(state, project){
          let result= state.projects.find((e)=>{
            if (e.name == project.name){
                return true
            }
          })
          if(result){
            result = project;
          }
        },
        FETCH_STA(state, sta){
            state.starredArtifacts = sta
        },
        FETCH_USERINFO(state, userinfo){
            state.userInfo = userinfo
        },
        cancelEditProject(state){
            state.projects.pop();
        },
        addEditProject(state){
            state.projects.push({editMode: true})
        },
        setLocalConfig(state, localConfig) {
           // For dev purpose
            state.backendUrl = localConfig.backendUrl
            state.tokenUrl = localConfig.tokenUrl
           // For production
            state.backendUrl = window.location.origin+'/sapi'
            state.tokenUrl = window.location.origin+'/api/v1/token'
        }
    },
    actions: {
        fetchChartData({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getChartList?filter=&quick=true")
                    .then((response) => {
                        commit("FETCH_CHARTS", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchFullChartData({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getChartList?filter=&quick=false")
                    .then((response) => {
                        commit("FETCH_FULL_CHARTS", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchConfiguration({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getConfig")
                    .then((response) => {
                        commit("FETCH_CONFIG", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchProjects({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getProjects")
                    .then((response) => {
                        commit("FETCH_PROJECTS", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchAnnoucements({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getAnnoucement")
                    .then((response) => {
                        commit("FETCH_ANNOUCEMENTS", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchStarredArtifacts({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getStarredArtifacts")
                    .then((response) => {
                        commit("FETCH_STA", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        },
        fetchUserInfo({ commit }) {
            return new Promise((resolve) => {
                Vue.prototype.$http.get(this.state.datastore.backendUrl+"/v1/getUserInfo")
                    .then((response) => {
                        commit("FETCH_USERINFO", response.data);
                        resolve();
                    })
                    .catch(() => {
                    });
            });
        }
    }
}

