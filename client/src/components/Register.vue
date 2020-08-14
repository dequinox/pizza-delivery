<template>
  <div>
    <b-card header="Register" style="max-width: 20rem;"
    class="mb-2 d-flex justify-content-center">
        <b-form @submit="onSubmit">
        <b-form-group
            id="input-group-1"
            label="Email address:"
            label-for="input-1"
        >
            <b-form-input
            id="input-1"
            v-model="form.email"
            type="email"
            required
            placeholder="Enter email"
            ></b-form-input>
        </b-form-group>

        <b-form-group id="input-group-2" label="Your Password:" label-for="input-2">
            <b-form-input
            id="input-2"
            v-model="form.password"
            type="password"
            required
            placeholder="Enter Password"
            ></b-form-input>
        </b-form-group>


        <b-form-group id="input-group-3" label="Your Name:" label-for="input-3">
            <b-form-input
            id="input-3"
            v-model="form.name"
            required
            placeholder="Enter Name"
            ></b-form-input>
        </b-form-group>

        <b-form-group id="input-group-4" label="Your Telephone:" label-for="input-4">
            <b-form-input
            id="input-4"
            v-model="form.telephone"
            required
            placeholder="Enter Phone Number"
            ></b-form-input>
        </b-form-group>


        <b-form-group id="input-group-5" label="Your Address:" label-for="input-5">
            <b-form-input
            id="input-5"
            v-model="form.address"
            required
            placeholder="Enter Address"
            ></b-form-input>
        </b-form-group>

        <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      <pre class="m-0">{{ form }}</pre>
    </b-card>
  </div>
</template>


<script>
import AuthenticationService from '@/services/AuthenticationService'

  export default {
    data() {
      return {
        form: {
          email: '',
          password: '',
          name: '',
          telephone: '',
          address: ''
        }
      }
    },
    methods: {
      async onSubmit() {
        try {
            const response = await AuthenticationService.register({
                email: this.form.email,
                password: this.form.password,
                name: this.form.name,
                telephone: this.form.telephone,
                address: this.form.address
            })
            this.$store.dispatch('setUser', response.data)
            this.$router.push({
                name: 'pizzas'
            })
        } catch (error) {
            this.error = error.response.data.error
        }
      },
    }
  }
</script>

<style scoped>
.card {
    margin: 0 auto; /* Added */
    float: none; /* Added */
    margin-bottom: 10px; /* Added */
    margin-top: 10px;
}
</style>