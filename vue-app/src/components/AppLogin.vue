<!-- 
 
4. Page — AppLogin.vue
This is one of the routed pages. Its responsibilities are:

Owns the data — email and password live here as reactive state
Uses v-model on the inputs to keep state in sync (two-way binding)
Defines submitHandler — what happens on a valid form submission (currently just a console log, but this is where you'd call your backend)
It delegates form rendering and validation logic to child components 

-->

<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-5">Login</h1>
                <hr>
                <!-- replace this setup by encapsulating the form validations in a formtag component -->
                <!-- <form method="post" action="/login" class="needs-validation" novalidate>
                    <text-input label="Email" type="email" name="email" required="true">
                    </text-input>
                    <text-input label="Password" type="password" name="password" required="true">
                    </text-input>
                    <input type="submit" class="btn btn-primary" value="Login">
                </form> -->

                <!-- 
                    shorthand for v-on 
                    Flow:
                        - Whenever this component hears an email named submitHandler -> going to call the handler below
                        - Call submitHandler in the methods
                        - In methods, the submiHandler will print successful to console and then connect to backend with form data
                        - We have also mounted lifecycle hook -> which is the bootstrap client side validation
                -->
                <!-- <FormTag @myevent=""submitHandler">-->
                <form-tag v-on:myevent="submitHandler" name="myform" event="myevent">
                    <text-input 
                        v-model="email"
                        label="Email" 
                        type="email" 
                        name="email" 
                        required="true">
                    </text-input>
                    <text-input 
                        v-model="password"
                        label="Password" 
                        type="password" 
                        name="password" 
                        required="true">
                    </text-input>
                    <input type="submit" class="btn btn-primary" value="Login">                    
                </form-tag>
                <hr>

            </div>
        </div>
    </div>
</template>

<script>

import TextInput from './forms/TextInput.vue'
import FormTag from './forms/FormTag.vue'

export default {
    name: 'AppLogin',
    components: {
        TextInput,
        FormTag,
    },
    data() {
        return {
            email: "",
            password: "",
        }
    },
    methods: {
        submitHandler() {
            console.log("submitHandler called - success!")

            // integrating with backend
            const payload = {
                email: this.email,
                password: this.password,
            }
            const requestOptions = {
                method: "POST",
                body: JSON.stringify(payload),
            }
            fetch("http://localhost:8081/users/login", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.log("Error:", data.message);
                } else {
                    console.log(data)
                }
            })
        }
    }//,
    // Since we already encapsulate the bootstrap validation in FormTag component, we do not need to do another validation here in AppLogin.vue
    // bootstrap validation
    // mounted() {
    //     (function () {
    //         'use strict'

    //         // Fetch all the forms we want to apply custom Bootstrap validation styles to
    //         var forms = document.querySelectorAll('.needs-validation')

    //         // Loop over them and prevent submission
    //         Array.prototype.slice.call(forms)
    //             .forEach(function (form) {
    //                 form.addEventListener('submit', function (event) {
    //                     if (!form.checkValidity()) {
    //                         event.preventDefault()
    //                         event.stopPropagation()
    //                     }

    //                     form.classList.add('was-validated')
    //                 }, false)
    //             })
    //     })()
    // }
}
</script>