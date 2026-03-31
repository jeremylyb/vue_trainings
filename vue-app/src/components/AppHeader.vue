<!-- 
 
    3. Navigation — AppHeader.vue
    Uses <router-link> instead of <a href> tags. The key difference is that router-link does client-side navigation — 
    no full page reload, it just tells the router to swap the component in <router-view>. 

-->

<template>
    <nav class="navbar navbar-expand-lg mavbar-dark bg-dark">
    <div class="container-fluid">
        <a class="navbar-brand" href="#">Navbar</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
            <li class="nav-item">
            <!-- <a class="nav-link active" aria-current="page" href="#">Home</a> -->

                <router-link class="nav-link active" aria-current="page" to="/">Home</router-link>
            </li>
            <li class="nav-item">
            <!-- <a class="nav-link" href="#">Features</a> -->
                <!-- 
                    - When user click the button, this routes to /login using router-link tag
                    - Based on index.js, this 
                -->
                <router-link class="nav-link" to="/login">Login</router-link>
            </li>
        </ul>
        </div>
    </div>
    </nav>
</template>