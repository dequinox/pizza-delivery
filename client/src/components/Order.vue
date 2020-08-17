<template>
  <div>
      <div v-if="!this.ordered">
    <b-card v-if="!$store.state.isUserLoggedIn && $store.state.orders.length > 0" header="Information Details" style="max-width: 20rem;"
    class="mb-2 d-flex justify-content-center">
        <b-form @submit="orderPizzas">
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
        <b-button type="submit" variant="primary">Order Pizas</b-button>
        </b-form>
    </b-card>
    <order-details class="order-detail" />
    <div v-if="$store.state.orders.length > 0 && $store.state.isUserLoggedIn">
        <b-card
        header="Order Confirmation"
        style="max-width: 20rem;"
        >
            <b-card-body>
                <b-button type="submit" variant="primary" @click=orderPizzas()>Order Pizzas</b-button>
            </b-card-body>
        </b-card>
    </div>

    <div v-if="$store.state.orders.length == 0">
        <b-card
        header="Cart"
        style="max-width: 20rem;"
        >
            <b-card-body>
                <b-card-title>Order List Empty</b-card-title>       
                <b-card-text>Add Items to make an order</b-card-text>
            </b-card-body>
        </b-card>
    </div>
    </div>
    <div v-if="this.ordered">
        <b-card
        header="Order"
        style="max-width: 20rem;"
        >
            <b-card-body>
                <b-card-title>Order Sucessful</b-card-title>       
                <b-card-text v-if="$store.state.isUserLoggedIn" class="text-success">
                    Head to History to see your orders
                </b-card-text>
            </b-card-body>
        </b-card>
    </div>
  </div>
</template>

<script>
import OrderDetails from './OrderDetails'
import {mapState} from 'vuex'
import OrdersService from '@/services/OrdersService'

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
        },
        ordered: false
      }
    },
    computed: {
        ...mapState([
            'isUserLoggedIn',
            'user',
            'orders',
            'cost',
            'currency',
            'costInCurrency'
        ])
    },
    methods: {
      async orderPizzas() {
        if (this.isUserLoggedIn){
            console.log("TRying order with a user")
            var orderDetail = []
            this.orders.forEach(function (arrayItem){
                orderDetail.push({
                    pizzaID: arrayItem.id,
                    amount: arrayItem.amount,
                    title: arrayItem.title
                })
            });
            try {

                const response = await OrdersService.createOrder({
                    accountID: this.user.id,
                    cost: Math.round(this.costInCurrency) + 10,
                    currency: this.currency,
                    orderDetails: orderDetail
                })
                this.ordered = true
                this.$store.dispatch('setOrder', [])
            } catch (error) {
                console.log({
                    accountID: this.user,
                    cost: 0,
                    orderDetails: this.orders
                })
            }
        }
        else {
            this.$store.dispatch('setOrder', [])
            console.log("No User")
            this.ordered = true
        }    
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