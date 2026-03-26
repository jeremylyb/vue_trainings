<template>
  <!-- <div class="container">
    <div class="row">
      <div class="col">
        <h1>My App</h1>
          <h2>I am Part of this tag</h2>
      </div>
    </div>
  </div> -->
  <div>
    <Header />
    <!-- <Body /> -->
    <div>
        <router-view></router-view>
    </div>
    <Footer />
  </div>

</template>

<script>
import Header from "./components/AppHeader.vue"
// import Body from "./components/AppBody.vue"
import Footer from "./components/AppFooter.vue"

export default {
  name: 'App',
  components: {
    Header,
    // Body,
    Footer,
  },
}
</script>

<style>

</style>