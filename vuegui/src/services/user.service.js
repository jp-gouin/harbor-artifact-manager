import { authHeader } from './auth.header';
import { datastore } from '../store/store'
import VueJwtDecode from 'vue-jwt-decode'
//import Axios from 'axios' 
export const userService = {
    login,
    logout,
    getAll
};

function login(username, password) {
    const requestOptions = {
        method: 'POST',
        body: JSON.stringify({ username, password })
    };
    return fetch(`${datastore.state.tokenUrl}`, requestOptions)
        .then(handleResponse)
        .then(user => {
            // login successful if there's a jwt token in the response
            if (user.token) {
                const tk =  VueJwtDecode.decode(user.token);
                user.username = tk.user;
                if(tk.id){
                    user.id = tk.id
                }
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('user', JSON.stringify(user));
            }

            return user;
        });
}


function logout() {
    // remove user from local storage to log user out
    localStorage.removeItem('user');
}

function getAll() {
    const requestOptions = {
        method: 'GET',
        headers: authHeader()
    };

    return fetch(`${datastore.state.backendUrl}/users`, requestOptions).then(handleResponse);
}

function handleResponse(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        if (!response.ok) {
            if (response.status === 401) {
                // auto logout if 401 response returned from api
                logout();
                location.reload(true);
            }

            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}