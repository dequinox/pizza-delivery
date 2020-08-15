
<template>
  <div>
    <b-card header="Login" style="max-width: 20rem;"
    class="mb-2 d-flex justify-content-center">
        <b-form @submit="onSubmit">
          <p v-if="!authenticationResult" class="text-danger">Wrong Password or Email</p>
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
        },
        authenticationResult : true
      }
    },
    methods: {
      async onSubmit() {
        try {
            const response = await AuthenticationService.login({
                email: this.form.email,
                password: this.form.password
            })
            this.$store.dispatch('setUser', response.data)
            this.$router.push({
                name: 'pizzas'
            })
        } catch (error) {
            this.authenticationResult = false
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