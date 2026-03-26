<template>
    <form 
        @submit.prevent="submit"
        :ref="name"
        :event="event"
        autocomplete="off" 
        :method="method" 
        :action="action" 
        class="needs-validation" 
        novalidate>
        <slot></slot>
    </form>
</template>


<script>
export default {
    name: 'FormTag',
    props: ["method", "action", "name", "event"],
    methods: {
        submit() {
            let myForm = this.$refs[this.$props.name]

            if (myForm.checkValidity()){
                console.log("My event name", this.$props['event']);
                console.log("Name", this.$props.name);
                this.$emit(this.$props['event']);
            }
            myForm.classList.add("was-validated");
        }
    },
    mounted() {
        (function () {
            'use strict'

            // Fetch all the forms we want to apply custom Bootstrap validation styles to
            var forms = document.querySelectorAll('.needs-validation')

            // Loop over them and prevent submission
            Array.prototype.slice.call(forms)
                .forEach(function (form) {
                    form.addEventListener('submit', function (event) {
                        if (!form.checkValidity()) {
                            event.preventDefault()
                            event.stopPropagation()
                        }

                        form.classList.add('was-validated')
                    }, false)
                })
        })()
    }
}

</script>