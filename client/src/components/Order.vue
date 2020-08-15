<template>
  <div>
    <b-card v-if="!$store.state.isUserLoggedIn && $store.state.orders.length > 0" header="Information Details" style="max-width: 20rem;"
    class="mb-2 d-flex justify-content-center">
        <b-form>
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

        </b-form>
      <pre class="m-0">{{ form }}</pre>
    </b-card>
    <order-details class="order-detail" />
    <div v-if="$store.state.orders.length > 0">
        <b-button type="submit" variant="primary" @click=orderPizzas()>Order Pizzas</b-button>
    </div>

    <div v-if="$store.state.orders.length == 0">
        <b-card
        header="Error"
        style="max-width: 20rem;"
        >
            <b-card-body>
                <b-card-title>Order List Empty</b-card-title>       
                <b-card-text>Add Items to make an order</b-card-text>
            </b-card-body>
        </b-card>
    </div>
  </div>
</template>

<script>
import OrderDetails from './OrderDetails'

export default {
    components: {
        OrderDetails
    },
    data() {
      return {
        form: {
          email: '',
          name: '',
          telephone: '',
          address: ''
        }
      }
    },
    methods: {
      async orderPizzas() {
        this.$router.push({
            name: 'pizzas'
        })
      },
    }
}
</script>

<style scoped>
.order-detail {
    width: 20%;
    position: fixed;
    top: 70px;
    right: 30px;
}
.card {
    margin: 0 auto; /* Added */
    float: none; /* Added */
    margin-bottom: 10px; /* Added */
    margin-top: 10px;
}
</style>