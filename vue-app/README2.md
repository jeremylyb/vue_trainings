File Structure
src/
  main.js
  App.vue
  router/
    index.js
  components/
    AppHeader.vue
    AppFooter.vue
    AppLogin.vue
    forms/
      FormTag.vue
      TextInput.vue

Step 1 — App Boots (main.js)
Browser loads index.html
  → finds <div id="app">
  → main.js runs
  → createApp(App)        ← App.vue is the root
      .use(router)         ← router is registered globally
      .mount('#app')       ← everything renders into #app

Step 2 — Shell Renders (App.vue)
App.vue renders
  → <Header />            ← always visible
  → <router-view />       ← placeholder, swaps based on URL
  → <Footer />            ← always visible
At this point, router-view checks the current URL and renders the matching page component inside itself.

Step 3 — Router Matches URL (router/index.js)
URL: /          → renders AppHome.vue    inside <router-view>
URL: /login     → renders AppLogin.vue   inside <router-view>
The router acts as the traffic controller — it reads the URL and decides which component fills the router-view slot.

Step 4 — User Clicks "Login" in Navbar (AppHeader.vue)
User clicks <router-link to="/login">
  → NO page reload (client-side navigation)
  → Router updates URL to /login
  → Router finds { path: '/login', component: AppLogin }
  → AppLogin.vue swaps into <router-view>

Step 5 — Login Page Renders (AppLogin.vue)
AppLogin.vue renders
  → owns reactive data:
      email: ""
      password: ""

  → renders <form-tag> with:
      v-on:myevent="submitHandler"   ← listening for valid form signal
      name="myform"
      event="myevent"

  → inside form-tag renders:
      <text-input v-model="email" ...>      ← two-way binds to email
      <text-input v-model="password" ...>   ← two-way binds to password
      <input type="submit">

Step 6 — User Types (v-model)
User types in email field
  → v-model updates AppLogin.data.email in real time

User types in password field
  → v-model updates AppLogin.data.password in real time
AppLogin always has the latest values without needing to manually read the DOM.

Step 7 — User Clicks Submit (FormTag.vue)
User clicks submit button
  → @submit.prevent fires
      ← blocks default browser form submission (no page reload)
  
  → FormTag.submit() runs
      → grabs form via this.$refs[name]
      → calls form.checkValidity()

      IF INVALID:
        → adds "was-validated" CSS class
        → Bootstrap shows red error styles on empty/wrong fields
        → stops here, nothing emitted to parent

      IF VALID:
        → adds "was-validated" CSS class
        → this.$emit("myevent")    ← fires event up to AppLogin

Step 8 — Parent Handles Valid Submission (AppLogin.vue)
AppLogin hears "myevent"
  → calls submitHandler()
  → currently: console.log("submitHandler called - success!")
  → future:    POST { email, password } to backend API

Complete Data Flow Diagram
User types
  └→ v-model → AppLogin.data { email, password }

User clicks submit
  └→ FormTag intercepts
      └→ Invalid? → show Bootstrap errors, stop
      └→ Valid?   → emit "myevent"
                      └→ AppLogin.submitHandler() fires
                            └→ use email + password to call backend

Key Concepts Recap
Concept                             Where                       What it does
createApp().use(router).mount()     main.js                     Boots the app
<router-view>                       App.vue                     Renders current page
<router-link>                       AppHeader.vue               Navigates without reload
routes[]                            router/index.js             Maps URLs to components
v-model                             AppLogin.vue                Syncs input → data
@submit.prevent                     FormTag.vue                 Blocks browser default
checkValidity()                     FormTag.vue                 Bootstrap HTML5 validation
$emit("myevent")                    FormTag.vue                 Signals parent on valid submit
submitHandler()                     AppLogin.vue                Handles valid form data



What IS Industry Standard
Vue Router with router-view / router-link
This is exactly how production Vue apps handle navigation. No issues here.
Component-based layout (App.vue with Header/Footer + router-view)
Very standard shell pattern used everywhere.
v-model for form binding
Completely standard, this is the Vue way.
Separating form components (TextInput, FormTag)
The idea is right — reusable form components are industry practice.

What is OUTDATED / Non-Standard
Manual Bootstrap validation in mounted()
javascript// This is the old jQuery-era approach
var forms = document.querySelectorAll('.needs-validation')
In modern Vue, you don't manually query the DOM like this. This pattern fights against Vue's reactivity system. Industry practice today is to use a form validation library instead.
Custom event naming (myevent)
javascriptthis.$emit("myevent")
Generic event names like myevent are a code smell. Real apps use descriptive names like form:submit, on-submit, etc.
$refs for form validity
Directly grabbing DOM elements via $refs for validation is discouraged when better tools exist.

What Industry Actually Uses Today
For Vue specifically:
ConcernThis CodeIndustry StandardForm validationManual Bootstrap + $refsVeeValidate or ZodForm stateManual v-model per fieldVeeValidate / FormKitStylingBootstrapTailwind CSSState managementLocal component dataPiniaAPI calls(not yet implemented)Axios + async/await

What This Code IS Good For
This looks like a learning project, and as that it's actually well structured because:

It teaches you why abstractions exist (FormTag wrapping validation)
It shows you how Vue's emit/props system works
It demonstrates component composition bottom-up

The patterns here are the conceptual foundation of what industry tools like VeeValidate do under the hood — they just do it more robustly. So understanding this code first makes you better at using those tools later.