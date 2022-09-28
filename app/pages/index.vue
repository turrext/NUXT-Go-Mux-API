  <!--<main>
    <h1>Home page</h1>
    <NuxtLink to="/fun">
      About (internal link that belongs to the Nuxt App)
    </NuxtLink>
    <a href="https://nuxtjs.org">External Link to another page</a>
  </main>-->
<template>
  <div class="bluebg1 py-8 min-h-full">
    <!-- Container for demo purpose -->
    <div class="container my-24 px-6 mx-auto">

      <!-- Section: Design Block -->
      <section class="mb-32 text-white text-center">
        <h2 class="text-3xl font-bold orng mb-20">Why Shieldhost?</h2>

        <div class="grid lg:gap-x-12 lg:grid-cols-3">
          <div class="mb-12 lg:mb-0">
            <div class="rounded-lg shadow-lg h-full block bg-lb">
              <div class="flex justify-center">
                <div class="p-4 bg-blue-600 rounded-full shadow-lg inline-block -mt-8">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 orng" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
              </div>
              <div class="p-6">
                <h3 class="text-2xl font-bold text-blue-600 mb-4">24/7</h3>
                <h5 class="text-lg db text-gray-800 font-bold mb-4">Support</h5>
                <p class="text-gray-500">
                  Our Support team is available around the clock both though Discord and Tickets to resolve your issues!
                </p>
              </div>
            </div>
          </div>

          <div class="mb-12 lg:mb-0">
            <div class="rounded-lg shadow-lg h-full block bg-lb">
              <div class="flex justify-center">
                <div class="p-4 bg-blue-600 rounded-full shadow-lg inline-block -mt-8">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 orng" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M20.618 5.984A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016zM12 9v2m0 4h.01" />
                  </svg>
                </div>
              </div>
              <div class="p-6">
                <h3 class="text-2xl font-bold text-blue-600 mb-4">99%</h3>
                <h5 class="text-lg db text-gray-800 font-bold mb-4">Uptime</h5>
                <p class="text-gray-500">
                  Between our custom DDoS Filters to running the best hardware, we guarentee 99% uptime for your projects!
                </p>
              </div>
            </div>
          </div>

          <div class="">
            <div class="rounded-lg shadow-lg h-full block bg-lb">
              <div class="flex justify-center">
                <div class="p-4 bg-blue-600 rounded-full shadow-lg inline-block -mt-8">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 orng" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01" />
                  </svg>
                </div>
              </div>
              <div class="p-6">
                <h3 class="text-2xl font-bold text-blue-600 mb-4">100%</h3>
                <h5 class="text-lg db text-gray-800 font-bold mb-4">Dedicated Resources</h5>
                <p class="text-gray-500">
                  Whether you buy a dedicated server, game server or a VDS, we give dedicated Ram and CPU Cores!
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!-- Section: Design Block -->

    </div>
    

  <!-- Section: Design Block -->
  <section class="mb-32 text-white text-center w-4/5 mb-16 text-gray-800 mx-auto text-center">
    <h2 class="text-3xl font-bold orng text-center mb-12">Popular Products</h2>

    <div class="grid lg:grid-cols-3 gap-6 lg:gap-x-12">
      <div class="mb-6 lg:mb-0">
        <div class="block rounded-lg shadow-lg bg-lb h-full">
          <div class="p-6 border-b border-gray-300 text-center">
            <p class="uppercase mb-4 blue text-sm">
              <strong>VDS</strong>
            </p>
            <h3 class="text-2xl mb-6">
              <strong>From $6</strong>
              <small class="text-gray-500 text-sm">/month</small>
            </h3>

            <button type="button"
              class="inline-block px-6 py-2.5 bg-orng text-white font-bold text-xs leading-tight uppercase rounded hover:text-blue-700 hover:bg-gray-100 focus:bg-gray-100 focus:outline-none focus:ring-0 active:bg-gray-200 transition duration-150 ease-in-out w-full"
              data-ripple-color="primary">
              Configure
            </button>
          </div>
          <div class="p-6">
            <ol class="list-inside">
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>DDoS Protected
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>France and US Locations
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Dedicated Resources
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>VPS Control Panel
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Variety of Linux Distros and Windows OS Available
              </li>
            </ol>
          </div>
        </div>
      </div>

      <div class="mb-6 lg:mb-0">
        <div class="block rounded-lg shadow-lg bg-lb h-full">
          <div class="p-6 border-b border-gray-300 text-center">
            <p class="uppercase mb-4 blue text-sm">
              <strong>Game Servers</strong>
            </p>
            <h3 class="text-2xl mb-6">
              <strong>From $2</strong>
              <small class="text-gray-500 text-sm">/Monthly</small>
            </h3>
            <button type="button"
              class="inline-block px-6 py-2.5 bg-orng text-white font-bold text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out w-full"
               data-ripple-color="light">
              Continue
            </button>
          </div>
          <div class="p-6">
            <ol class="list-inside">
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>FiveM
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Minecraft
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Rust
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>ARK
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Game Panel
              </li>
            </ol>
          </div>
        </div>
      </div>

      <div class="mb-6 lg:mb-0">
        <div class="block rounded-lg shadow-lg bg-lb h-full">
          <div class="p-6 border-b border-gray-300 text-center">
            <p class="uppercase mb-4 blue text-sm">
              <strong>Dedicated Servers</strong>
            </p>
            <h3 class="text-2xl mb-6">
              <strong>From $80</strong>
              <small class="text-gray-500 text-sm">/Monthly</small>
            </h3>
            <button type="button"
              class="inline-block px-6 py-2.5 bg-orng text-white font-bold text-xs leading-tight uppercase rounded hover:text-blue-700 hover:bg-gray-100 focus:bg-gray-100 focus:outline-none focus:ring-0 active:bg-gray-200 transition duration-150 ease-in-out w-full"  
              data-ripple-color="primary">
              Continue
            </button>
          </div>
          <div class="p-6">
            <ol class="list-inside">
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Additional IPS Available
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Delivery within 3-7 Business Days
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>DDoS Protected
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Guarenteed Network Speeds
              </li>
              <li class="mb-4 flex items-center">
                <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check"
                  class="text-green-600 w-4 h-4 mr-2" role="img" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 512 512">
                  <path fill="currentColor"
                    d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z">
                  </path>
                </svg>Multiple Locations in the US
              </li>
            </ol>
          </div>
        </div>
      </div>
    </div>
  </section>
  <!-- Section: Design Block -->

  <!-- Section: Design Block -->
  <section class="text-gray-600 body-font">
    <div class="container px-5 py-8 mx-auto">
      <h1 class="text-3xl font-bold orng text-center mb-12">Testimonials</h1>
      <div class="flex flex-wrap -m-4">
        <div class="p-4 md:w-1/2 w-full">
          <div class="h-full bg-lb p-8 rounded">
            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="block w-5 h-5 text-gray-400 mb-4" viewBox="0 0 975.036 975.036">
              <path d="M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z"></path>
            </svg>
            <p class="leading-relaxed mb-6">I was with several hosts before, all of them either were too pricey or had isssues with certain scripts that caused disturbances to both me and my clients, Shieldhost is not only afforadble but I've been able to get my work done much faster as their RDP's are much faster than my previous hosts.</p>
            <a class="inline-flex items-center">
              <img alt="testimonial" src="https://dummyimage.com/106x106" class="w-12 h-12 rounded-full flex-shrink-0 object-cover object-center">
              <span class="flex-grow flex flex-col pl-4">
                <span class="title-font font-medium db">THEDEVKILLER</span>
                <span class="text-gray-500 text-sm">Freelancer</span>
              </span>
            </a>
          </div>
        </div>
        <div class="p-4 md:w-1/2 w-full">
          <div class="h-full bg-lb p-8 rounded">
            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="block w-5 h-5 text-gray-400 mb-4" viewBox="0 0 975.036 975.036">
              <path d="M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z"></path>
            </svg>
            <p class="leading-relaxed mb-6">I was facing DDoS Issues with my last host and my server was taking a really long time to start which to also cause it to crash sometimes but since I've Switched to Shieldhost both of those issues are gone! I 100% recommend Shield Host to Fellow Game Server Owners!</p>
            <a class="inline-flex items-center">
              <img alt="testimonial" src="https://dummyimage.com/107x107" class="w-12 h-12 rounded-full flex-shrink-0 object-cover object-center">
              <span class="flex-grow flex flex-col pl-4">
                <span class="title-font font-medium db">Trey Crawf</span>
                <span class="text-gray-500 text-sm">FIVEM SERVER OWNER</span>
              </span>
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>

  </div>
</template>

<script lang="ts">
import Vue from 'vue'


export default {
  layout: 'basic',
  name: 'IndexPage'
}
</script>
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