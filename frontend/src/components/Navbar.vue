<template>
    <b-navbar class="navbar is-fixed-top">
        <template slot="brand">
          <b-navbar-item id="french-ai" tag="router-link" :to="{ name: 'products' }">
            <span id="website-icon">
              -- 🧠 -
            </span>
            <span id="website-name-first">
              Dmitry 
            </span>
            <span id="website-name-sep">
            </span>
            <span id="website-name-last">
              Roitman
            </span>
          </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" :to="{ name: 'home' }">
              Home
            </b-navbar-item>
            <b-navbar-dropdown tag="router-link" :to="{ name: 'products' }" label="About"  class="navbar-item has-dropdown" style="z-index: 100;">
                  <a class="navbar-item" style="z-index: 100;" href="/about/me">
                   About me
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/product/customer">
                   How I can help
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/product/automation">
                   How I work
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/product/features">
                   Resume
                  </a>
            </b-navbar-dropdown>


            <b-navbar-item tag="router-link" :to="{ name: 'pricing' }">
              Pricing
            </b-navbar-item>


              <b-navbar-dropdown label="Solutions" class="navbar-item has-dropdown" style="z-index: 100;">
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/howitworks">
                   How it works?
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/howtochoose">
                   How to choose?
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/sales">
                   Sales
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/support">
                   Support
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/remote">
                   Remote Teams
                  </a>
                  <hr class="navbar-divider" style="z-index: 100;">
                  <a class="navbar-item" style="z-index: 100;" href="/solutions/version">
                    Version 0.9.0
                  </a>
              </b-navbar-dropdown>

              <b-navbar-dropdown label="Company" class="navbar-item has-dropdown" style="z-index: 100;">
                  <a class="navbar-item" style="z-index: 100;" href="/company/about">
                    About Us
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/company/press">
                    Press Releases
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/company/careers">
                    Careers
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/company/customers">
                    Customers
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/company/events">
                    Events
                  </a>
              </b-navbar-dropdown>



              <b-navbar-dropdown label="Resources" class="navbar-item has-dropdown" :to="{ name: 'resources' }" style="z-index: 100;" aria-controls="dropdown-menu2">
                  <a class="navbar-item" style="z-index: 100;" href="/resources/blog">
                          <span class="icon is-small has-text-primary">
                                <i class="fas fa-newspaper"></i>
                          </span>
                          <span class="ml-3">Blog</span>
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/resources/docs">
                  <span class="icon is-small has-text-success">
                      <i class="fas fa-book"></i>
                  </span> 
                  <span class="ml-3">Documentation</span>

                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/resources/faq">
                  <span class="icon is-small has-text-warning">
                      <i class="fas fa-comment"></i>
                  </span>
                  <span class="ml-3">FAQ</span>
                  </a>
                  <a class="navbar-item" style="z-index: 100;" href="/resources/links">
                     <span class="icon is-small has-text-link"> 
                     <i class="fas fa-link"></i>
                     </span> 
                     <span class="ml-3">Links</span>
                  </a>
              </b-navbar-dropdown>


        </template>

        <template slot="end">
                 <b-tooltip
                 type="is-light"
                 label="Sign In"
                 class="mt-1"
                 position="is-bottom"
                 >
                 <b-navbar-item tag="router-link" :to="{ name: 'signin' }">
                 <span class="icon is-small has-text-primary" data-tooltip="Sign In">
                     <i class="fas fa-sign-in-alt"></i>
                 </span>
                 </b-navbar-item>
                 </b-tooltip>

                 <span class="ml-4 mr-4 mt-1">|</span>
              
                 <b-tooltip
                 type="is-light"
                 label="Contact Us"
                 position="is-bottom"
                 class="mt-1"
                 >
                <b-navbar-item tag="router-link" :to="{ name: 'contact' }">
                 <span class="icon is-small has-text-primary" data-tooltip="Contact Us">
                     <i class="fas fa-envelope"></i>
                 </span>
                 </b-navbar-item>
                 </b-tooltip>
                 <span class="ml-6"></span>

        </template>
    </b-navbar>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      open: false,
    };
  },
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Bree+Serif&display=swap');

  #french-ai {
    font-family: 'Bree Serif', serif;
  }

  #website-icon {
    color: #fff;
  }

  #website-name-first {
    color: #1e3799;
  }

  #website-name-sep {
    color: #f3f3f3;
  }

  #website-name-last {
    color: #eb2f06;
  }

</style>
