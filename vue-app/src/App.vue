
<!-- 

2. Shell — App.vue
This is the persistent layout wrapper. It always shows Header and Footer, and the middle <router-view> acts as a dynamic slot — 
it swaps in whatever page component matches the current URL. So the header/footer never re-render on navigation, only the middle content changes. 

-->

<template>
  <div>
    <Header />
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