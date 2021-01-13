<template>
    <div class="w-full max-w-xs m-auto mt-4">
  <div class="bg-white shadow-md rounded px-8 pt-40 pb-8 mb-4">
  <form @submit.prevent="handleSubmit" >
     <div class="login__check w-24 absolute"></div>
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        <svg class="login__icon name svg-icon inline-block w-6 h-6 overflow-visible" viewBox="0 0 20 20">
            <path fill="none" d="M0,20 a10,8 0 0,1 20,0z M10,0 a4,4 0 0,1 0,8 a4,4 0 0,1 0,-8" />
          </svg>
        Username
      </label>
      <input v-model="username" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Username">
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        <svg class="login__icon pass svg-icon inline-block w-6 h-6 overflow-visible" viewBox="0 0 20 20">
            <path fill="none" d="M0,20 20,20 20,8 0,8z M10,13 10,16z M4,8 a6,8 0 0,1 12,0" />
          </svg>
        Password
      </label>
      <input v-model="password" class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="******************">
      <p class="text-red-500 text-xs italic">Please choose a password.</p>
    </div>
    <div class="flex items-center justify-between">
      <button class="bg-transparent hover:bg-blue-700 hover:text-white text-blue-700 font-semibold py-2 px-4 border border-blue-500 hover:border-transparent rounded" :disabled="loggingIn">Login</button>
        <img v-show="loggingIn" src="data:image/gif;base64,R0lGODlhEAAQAPIAAP///wAAAMLCwkJCQgAAAGJiYoKCgpKSkiH/C05FVFNDQVBFMi4wAwEAAAAh/hpDcmVhdGVkIHdpdGggYWpheGxvYWQuaW5mbwAh+QQJCgAAACwAAAAAEAAQAAADMwi63P4wyklrE2MIOggZnAdOmGYJRbExwroUmcG2LmDEwnHQLVsYOd2mBzkYDAdKa+dIAAAh+QQJCgAAACwAAAAAEAAQAAADNAi63P5OjCEgG4QMu7DmikRxQlFUYDEZIGBMRVsaqHwctXXf7WEYB4Ag1xjihkMZsiUkKhIAIfkECQoAAAAsAAAAABAAEAAAAzYIujIjK8pByJDMlFYvBoVjHA70GU7xSUJhmKtwHPAKzLO9HMaoKwJZ7Rf8AYPDDzKpZBqfvwQAIfkECQoAAAAsAAAAABAAEAAAAzMIumIlK8oyhpHsnFZfhYumCYUhDAQxRIdhHBGqRoKw0R8DYlJd8z0fMDgsGo/IpHI5TAAAIfkECQoAAAAsAAAAABAAEAAAAzIIunInK0rnZBTwGPNMgQwmdsNgXGJUlIWEuR5oWUIpz8pAEAMe6TwfwyYsGo/IpFKSAAAh+QQJCgAAACwAAAAAEAAQAAADMwi6IMKQORfjdOe82p4wGccc4CEuQradylesojEMBgsUc2G7sDX3lQGBMLAJibufbSlKAAAh+QQJCgAAACwAAAAAEAAQAAADMgi63P7wCRHZnFVdmgHu2nFwlWCI3WGc3TSWhUFGxTAUkGCbtgENBMJAEJsxgMLWzpEAACH5BAkKAAAALAAAAAAQABAAAAMyCLrc/jDKSatlQtScKdceCAjDII7HcQ4EMTCpyrCuUBjCYRgHVtqlAiB1YhiCnlsRkAAAOwAAAAAAAAAAAA==" />
      <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href="#">
        Forgot Password?
      </a>
    </div>
  </form>
  <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold mt-4 py-2 px-4 rounded inline-flex items-center" :disabled="loggingIn" v-on:click="signIn">
    <svg viewBox="0 0 533.5 544.3" class=" w-6 h-6  mr-2" xmlns="http://www.w3.org/2000/svg"><path d="M533.5 278.4c0-18.5-1.5-37.1-4.7-55.3H272.1v104.8h147c-6.1 33.8-25.7 63.7-54.4 82.7v68h87.7c51.5-47.4 81.1-117.4 81.1-200.2z" fill="#4285f4"/><path d="M272.1 544.3c73.4 0 135.3-24.1 180.4-65.7l-87.7-68c-24.4 16.6-55.9 26-92.6 26-71 0-131.2-47.9-152.8-112.3H28.9v70.1c46.2 91.9 140.3 149.9 243.2 149.9z" fill="#34a853"/><path d="M119.3 324.3c-11.4-33.8-11.4-70.4 0-104.2V150H28.9c-38.6 76.9-38.6 167.5 0 244.4l90.4-70.1z" fill="#fbbc04"/><path d="M272.1 107.7c38.8-.6 76.3 14 104.4 40.8l77.7-77.7C405 24.6 339.7-.8 272.1 0 169.2 0 75.1 58 28.9 150l90.4 70.1c21.5-64.5 81.8-112.4 152.8-112.4z" fill="#ea4335"/></svg>
    Login With Google
  </button>
  </div>
  <p class="text-center text-gray-500 text-xs">
    &copy;Made by JPG.
  </p>
</div>
</template>

<script>
export default {
    name:'LoginPage',
    data () {
        return {
            username: '',
            password: '',
            submitted: false
        }
    },
    computed: {
        loggingIn () {
            return this.$store.state.authentication.status.loggingIn;
        }
    },
    created () {
        // reset login status
        this.$store.dispatch('authentication/logout');
    },
    methods: {
        handleSubmit () {
            this.submitted = true;
            const { username, password } = this;
            if (username && password) {
                this.$store.dispatch('authentication/login', { username, password });
            }
        },
        signIn(){
          this.$gAuth
            .signIn()
            .then((GoogleUser) => {
              //on success do something
              console.log("GoogleUser", GoogleUser);
              console.log("getId", GoogleUser.getId());
              console.log("getBasicProfile", GoogleUser.getBasicProfile());
              console.log("getAuthResponse", GoogleUser.getAuthResponse());
              console.log(
                "getAuthResponse",
                this.$gAuth.GoogleAuth.currentUser.get().getAuthResponse()
              );
              let isSignIn = this.$gAuth.isAuthorized;
             if (isSignIn){
                const profile = GoogleUser.getBasicProfile()
                const authResponse = GoogleUser.getAuthResponse()
                this.$store.dispatch('authentication/signIn', { profile, authResponse});
              }
            })
            .catch(error => {
              //on fail do something
              console.error(error);
            });
        }
    }
}
</script>

<style scoped>
 path {
    stroke: rgba(44, 44, 44, 1);
    stroke-width: 1;
  }
.name path {
      stroke-dasharray: 73.50196075439453;
      stroke-dashoffset: 73.50196075439453;
      animation: animatePath 2s 0.5s forwards;
    }
.pass path {
      stroke-dasharray: 92.10662841796875;
      stroke-dashoffset: 92.10662841796875;
      animation: animatePath 2s 0.5s forwards;
    }
.login__check {
    width: 6rem;
    height: 1.5rem;
    margin-left: 7rem;
    top: 7rem;
    background: rgb(51, 51, 51);
    transform-origin: 0 100%;
    transform: rotate(-45deg);
}
    
    .login__check:before {
      content: "";
      position: absolute;
      left: 0;
      bottom: 100%;
      width: 1.5rem;
      height: 3rem;
      background: rgb(201, 201, 201);
      box-shadow: inset -0.2rem -2rem 2rem rgba(0,0,0,0.5);
    }
@keyframes animatePath {
  to {
    stroke-dashoffset: 0;
  }
}
</style>