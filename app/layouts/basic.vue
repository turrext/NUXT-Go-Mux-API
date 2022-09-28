

<template>
  <!--<nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav-->
  <div>
  <nav class="bg-blue">
    <header class="text-gray-100 bg-blue body-font">
      <div class="container mx-auto flex flex-wrap p-1 flex-col md:flex-row items-center">
        <router-link to="/" class="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="#F64C72" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-10 h-10 orng p-2 bg-blue rounded-full" viewBox="0 0 24 24">
            <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"></path>
          </svg>
          <span class="ml-3 orng text-xl">Shieldhost</span>
        </router-link>
        <nav class="md:ml-auto flex flex-wrap items-center text-base justify-center">
          <router-link to="/about" class="mr-5 text-white hover:text-gray-600">About</router-link>
          <router-link to="/vds" class="mr-5 text-white hover:text-gray-600">VDS</router-link>
          <router-link to="/dedicated" class="mr-5 text-white hover:text-gray-600">VPS</router-link>
          <router-link to="/protection" class="mr-5 text-white hover:text-gray-600">DDoS Protection</router-link>
          <router-link to="/gamehosting" class="mr-5 text-white hover:text-grey-600">Game Hosting</router-link>
        </nav>
        <button style="background-color: #F64C72" href="/login" class="inline-flex items-center border-0 LoginBTN py-1 px-3 focus:outline-none hover:bg-green-400 rounded mt-4 md:mt-0">Login
          <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-4 h-4 ml-1" viewBox="0 0 24 24">
            <path d="M5 12h14M12 5l7 7-7 7"></path>
          </svg>
        </button>
        <button style="background-color: #F64C72" href="/cart" class="inline-flex items-center border-0 LoginBTN py-1 px-3 focus:outline-none hover:bg-green-400 rounded mt-4 md:mt-0">Cart
          <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-4 h-4 ml-1" viewBox="0 0 24 24">
            <path d="M5 12h14M12 5l7 7-7 7"></path>
          </svg>
        </button>
      </div>
    </header>
  </nav>
  <Nuxt />
  <footer class="text-gray-600 body-font">
  <div class="container px-5 py-24 mx-auto flex md:items-center lg:items-start md:flex-row md:flex-nowrap flex-wrap flex-col">
    <div class="w-64 flex-shrink-0 md:mx-0 mx-auto text-center md:text-left">
      <a class="flex title-font font-medium items-center md:justify-start justify-center text-gray-900">
        <!--<svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" class="w-10 h-10 text-white p-2 bg-indigo-500 rounded-full" viewBox="0 0 24 24">
          <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"></path>
        </svg>-->
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mt-2 orng" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01" />
        </svg>
        <span class="ml-3 mb-3 mt-4 orng font-extrabold text-3xl">ShieldHost</span>
      </a>
      <p class="mt-2 font-bold ml-3 text-xl text-gray-500">The Hosting Solution for You!</p>
    </div>
    <div class="flex-grow flex flex-wrap md:pl-20 -mb-10 md:mt-0 mt-10 md:text-left text-center">
      <div class="lg:w-1/4 md:w-1/2 w-full px-4">
        
      </div>
      <div class="lg:w-1/4 md:w-1/2 w-full px-4">
        <h2 class="title-font text-gray-900 tracking-widest text-xl font-bold mb-1">Useful Links</h2>
        <nav class="list-none mb-10">
          <li>
            <a class="text-gray-600 hover:text-gray-800">Terms of Service</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Privacy Policy</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Cancellation and Refund Policy</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Discord</a>
          </li>
        </nav>
      </div>
      <div class="lg:w-1/4 md:w-1/2 w-full px-4">
        <h2 class="title-font text-gray-900 tracking-widest text-xl font-bold mb-3">Products</h2>
        <nav class="list-none mb-10">
          <li>
            <a class="text-gray-600 hover:text-gray-800">VDS</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">VPS</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Dedicated Servers</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Game Servers</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Discord Bot Hosting</a>
          </li>
        </nav>
      </div>
      <div class="lg:w-1/4 md:w-1/2 w-full px-4">
        <h2 class="title-font text-gray-900 tracking-widest text-xl font-bold mb-3">Game Servers</h2>
        <nav class="list-none mb-10">
          <li>
            <a class="text-gray-600 hover:text-gray-800">FiveM</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Minecraft</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">Rust</a>
          </li>
          <li>
            <a class="text-gray-600 hover:text-gray-800">ARK</a>
          </li>
        </nav>
      </div>
    </div>
  </div>
  <div class="bg-gray-100">
    <div class="container mx-auto py-4 px-5 flex flex-wrap flex-col sm:flex-row">
      <p class="text-gray-500 text-sm text-center sm:text-left">© 2022 Shieldhost
        <!--<a href="https://twitter.com/knyttneve" rel="noopener noreferrer" class="text-gray-600 ml-1" target="_blank">@knyttneve</a>-->
      </p>
    </div>
  </div>
</footer>
</div>
</template>


<style>

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #242582;
}
body {
  color: #242582;
}
.home {
  background-color: #242582;
  color: #242582;
}
.blue {
  color: #242582;
}
.bg-blue {
  background-color: #242582;
}
.bg-orng {
  background-color: #F64C72;
}
.orng {
  color: #F64C72;
}
.LoginBtn {
  background-color: #F64C72;
}
nav {
  padding: 30px;
}
nav a {
  font-weight: bold;
  color: #2c3e50;
}
nav a.router-link-exact-active {
  color: #42b983;
}
</style>