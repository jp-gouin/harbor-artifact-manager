import Vue from 'vue';
import Vuex from 'vuex';

import { authentication } from './authentication.module';
import { users } from './user.module';
import { datastore } from './store'
import {socketstore} from './socket'

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        datastore,
        authentication,
        users,
        socketstore
    }
});