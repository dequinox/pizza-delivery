<template>
<div>
    <div v-if="$store.state.isUserLoggedIn">
        <b-card
        header="History of Orders"
        style="max-width: 50rem;"
        >
            <b-card-body>
                <b-table striped hover :items="history"></b-table>
            </b-card-body>
        </b-card>
    </div>

    <div v-if="!$store.state.isUserLoggedIn">
        <b-card
        header="History"
        style="max-width: 20rem;"
        >
            <b-card-body>
                <b-card-title>Not Logged In</b-card-title>       
                <b-card-text v-if="!$store.state.isUserLoggedIn" class="text-danger">
                    Log in to see your history of orders
                </b-card-text>
            </b-card-body>
        </b-card>
    </div>
</div>
</template>

<script>
import OrdersService from '@/services/OrdersService'
import {mapState} from 'vuex'

export default {
    data () {
        return {
            history: null
        }
    },
    computed: {
    ...mapState([
      'user'
    ])
  },
    async mounted () {
        var history_data = (await OrdersService.index(this.user)).data
        this.history = []
        history_data.forEach(element => {
            var titles = []
            element.orderDetails.forEach(title => {
                titles.push(title.title)
            })
            this.history.push({
                ID: element.id,
                Cost: element.cost,
                Currency: element.currency,
                OrderDetails: titles

            })
        });
    },
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