const RegistrationForm = {
    data() {
        return {
            addressSameChecked: true,
        }
    },
    props: ["items"],
    template:
        `
        <h3> Registration </h3>
        <hr>
        <form autocomplete="off" class="needs-validation" novalidate>
            <text-input required="required" label="First Name" type="text" name="first-name"></text-input>
            <text-input required="required" label="Last Name" type="text" name="last-name"></text-input>
            <text-input required="required" label="email" type="email" name="email"></text-input>
            <text-input required="required" label="password" type="password" name="password"></text-input>
            <select-input label="Favourite Colour" name="color" :items="items"></select-input>

            <text-input required="required" label="Address" type="email" name="address"></text-input>
            <text-input required="required" label="City/Town" type="password" name="city"></text-input>
            <text-input required="required" label="State/Province" type="email" name="state/province"></text-input>
            <text-input required="required" label="Zip/Postal" type="password" name="zip"></text-input>

            <check-input v-on:click="addressSame" label="Mailing Address Same" checked="true" v-model="addressSameChecked"></check-input>

            <div v-if="addressSameChecked === false">
                <div class="mt-3">
                    <text-input label="Mailing Address" type="email" name="address2"></text-input>
                    <text-input label="Mailing City/Town" type="password" name="city2"></text-input>
                    <text-input label="Mailing State/Province" type="email" name="state2"></text-input>
                    <text-input label="Mailing Zip/Postal" type="password" name="zip2"></text-input>
                </div>
            </div>

            <check-input label="I agree to terms and conditions" required="true"></check-input>
            <hr>
            <input type="submit" class="btn btn-primary", value="Register">
        </form>
    `,
    methods:{
        addressSame() {
            console.log("address same fired");
            if (this.addressSameChecked === true){
                console.log("was checked on click");
                this.addressSameChecked = false;
            } else {
                console.log("was not checked on click");
                this.addressSameChecked = true;
            }

        }
    },
    // Nesting component in existing components
    components: {
        'text-input': TextInput,
        'select-input': SelectInput,
        'check-input': CheckInput
    },
    // mounted lifecycle hooks
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